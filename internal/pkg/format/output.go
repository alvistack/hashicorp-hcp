package format

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"text/tabwriter"
	"text/template"

	"github.com/hashicorp/hcp/internal/pkg/iostreams"
)

// Displayer is the interface for displaying a given payload. By implementing
// this interface, the payload can be outputted in any of the given Formats.
type Displayer interface {
	// DefaultFormat returns the Format in which to display the payload if the
	// user does not specify a Format override.
	DefaultFormat() Format

	// Payload is the object to display. Payload may return a single object or a
	// slice of objects.
	Payload() any

	// FieldTemplates returns a slice of Fields. Each Field represents an field
	// based on the payload to display to the user. It is common that the Field
	// is simply a specific field of the payload struct being outputted.
	FieldTemplates() []Field
}

// NewDisplayer creates a new Displayer with the given payload, default format,
// and fields.
func NewDisplayer[T any](payload T, defaultFormat Format, fields []Field) Displayer {
	return &internalDisplayer[T]{
		payload:       payload,
		fields:        fields,
		defaultFormat: defaultFormat,
	}
}

// DisplayFields displays the given fields about the given payload. If no fields are
// provided, then all fields are displayed. The fields can be specified using
// either the direct struct field name or the json tag name.
func DisplayFields[T any](payload T, format Format, fields ...string) Displayer {
	return NewDisplayer[T](payload, format, inferFields(payload, fields))
}

func inferFields[T any](payload T, columns []string) []Field {
	rv := reflect.ValueOf(payload)

	for rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}

	if rv.Kind() == reflect.Slice {
		rv = rv.Index(0)

		for rv.Kind() == reflect.Pointer {
			rv = rv.Elem()
		}
	}

	if rv.Kind() != reflect.Struct {
		return []Field{NewField("Value", "{{ . }}")}
	}

	toField := map[string]int{}

	for i, col := range columns {
		toField[col] = i
	}

	st := rv.Type()

	var ret []Field

	all := len(toField) == 0

	if !all {
		ret = make([]Field, len(toField))
	}

	for i := 0; i < st.NumField(); i++ {
		f := st.Field(i)
		if !f.IsExported() {
			continue
		}

		df := NewField(f.Name, fmt.Sprintf("{{ .%s }}", f.Name))

		if all {
			ret = append(ret, df)
		} else if idx, ok := toField[f.Name]; ok {
			ret[idx] = df
		}
	}

	return ret
}

type internalDisplayer[T any] struct {
	payload       T
	fields        []Field
	defaultFormat Format
}

func (i *internalDisplayer[T]) DefaultFormat() Format   { return i.defaultFormat }
func (i *internalDisplayer[T]) FieldTemplates() []Field { return i.fields }
func (i *internalDisplayer[T]) Payload() any            { return i.payload }

// TemplatedPayload allows a Displayer to return a different payload if the
// output will be templated using the field templates. This can be useful when
// raw output (e.g. JSON) requires a specific payload but the templated output
// (e.g. table/pretty) would benefit from a different payload type.
type TemplatedPayload interface {
	TemplatedPayload() any
}

// Field represents a field to output.
type Field struct {
	// Name is the displayed name of the field to the user.
	Name string

	// ValueFormat is a text/template that controls how the value will be
	// displayed. If the payload is a struct with the following structure:
	//
	//   type Cluster struct {
	//     Name          string
	//     Description   string
	//     CloudProvider string
	//     Region        string
	//     CreatedAt     time.Time
	//  }
	//
	// Example ValueFormat's would be:
	//   '{{ . Name }}' -> "Example"
	//   '{{ .CloudProvider }}/{{ . Region }}' -> "aws/us-east-1"
	//
	// A more advanced example would be using the text/template to invoke a
	// function. This can be done by implementing a function on the Payload type
	// that can be invoked. Function definitions will shadow fields in the
	// returned Payload.
	//
	// func (c *Cluster) CreatedAt() string {
	//   return humanize.Time(d.cluster.CreatedAt)
	// }
	//
	// A ValueFormat of '{{ .CreatedAt }}' will now invoke this function. If the
	// cluster was recently created an output may display "4s ago".
	ValueFormat string
}

// NewField creates a new Field with the given name and value format string. See
// the Field struct for more information.
func NewField(name, valueFormat string) Field {
	return Field{Name: name, ValueFormat: valueFormat}
}

// Outputter is used to output data to users in a consistent manner. The
// outputter supports outputting data in a number of Formats.
//
// To output data, the Display function should be called with a Displayer. A
// Displayer has a default format. The outputter will use this format unless a
// format has previously been set which overrides the default.
type Outputter struct {
	// io is the iostream to output to.
	io iostreams.IOStreams

	// forcedFormat is the format to output with regardless of the DefaultFormat
	// of the passed Displayer.
	forcedFormat Format
}

// New returns an new outputter that will write to the provided IOStreams.
func New(io iostreams.IOStreams) *Outputter {
	if io == nil {
		panic("io stream must be specified")
	}

	return &Outputter{
		io: io,
	}
}

// SetFormat sets the format to output with regardless of the DefaultFormat
// returned by the displayer.
func (o *Outputter) SetFormat(f Format) {
	o.forcedFormat = f
}

// GetFormat returns the format if set.
func (o *Outputter) GetFormat() Format {
	return o.forcedFormat
}

// Display displays the passed Displayer. The format used is the DefaultFormat
// unless the outputter has had a Format set which overrides the default.
func (o *Outputter) Display(d Displayer) error {
	// Determine what format to use
	format := d.DefaultFormat()
	if o.forcedFormat != Unset {
		format = o.forcedFormat
	}

	// Display the payload based on the selected format.
	switch format {
	case Pretty:
		return o.outputPretty(d)
	case Table:
		return o.outputTable(d)
	case JSON:
		return o.outputJSON(d)
	}

	return fmt.Errorf("invalid output format")
}

// Show outputs the given val using the DisplayFields function.
// If fields are specified, only those fields are shown, otherwise
// all fields are shown. The json tag of a field is honored and can be used to
// specified a field.
//
// This is a simplified version of using .Display, which should be used for all more
// advanced cases that require formatting fields differently.
//
// This function can accept a slice of values as well and formats them correctly.
// If the value being considered (directly or within in a slice) is not a struct,
// it is displayed as is under the field named 'Value'.
func (o *Outputter) Show(val any, format Format, fields ...string) error {
	return o.Display(DisplayFields(val, format, fields...))
}

// outputJSON outputs the payload in JSON.
func (o *Outputter) outputJSON(d Displayer) error {
	data, err := json.MarshalIndent(d.Payload(), "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshall result to JSON: %w", err)
	}

	fmt.Fprintln(o.io.Out(), string(data))
	return nil
}

// outputPretty outputs the payload using a key/value format where each field
// occupies a single row.
func (o *Outputter) outputPretty(d Displayer) error {
	var p any
	if tp, ok := d.(TemplatedPayload); ok {
		p = tp.TemplatedPayload()
	} else {
		p = d.Payload()
	}

	tmpl, err := template.New("hcp").Parse(prettyPrintTemplate(d))
	if err != nil {
		return err
	}

	rv := reflect.ValueOf(p)
	if rv.Kind() == reflect.Slice {
		if rv.Len() == 0 {
			fmt.Fprintln(o.io.Out(), "Listed 0 items.")
			return nil
		}

		for i := 0; i < rv.Len(); i++ {
			vf := rv.Index(i)
			if err := tmpl.Execute(o.io.Out(), vf.Interface()); err != nil {
				return err
			}

			fmt.Fprintln(o.io.Out())
			if i != rv.Len()-1 {
				fmt.Fprintln(o.io.Out(), "---")
			}
		}
	} else {
		if err := tmpl.Execute(o.io.Out(), p); err != nil {
			return err
		}

		fmt.Fprintln(o.io.Out())
	}

	return nil
}

// prettyPrintTemplate returns a text/template string for pretty printing the
// given payload. The template will align the values so they are easily scannable.
func prettyPrintTemplate(d Displayer) string {
	// Write to the buffer using a tabwriter. The Tabwriter will ensure that
	// each key/value is aligned.
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', 0)

	// Go through each field and output a new line
	fields := d.FieldTemplates()
	for i, f := range fields {
		fmt.Fprintf(w, "%s:\t%s", f.Name, f.ValueFormat)
		if i != len(fields)-1 {
			fmt.Fprintln(w)
		}
	}

	// Ignore the error
	_ = w.Flush()
	return buf.String()
}
