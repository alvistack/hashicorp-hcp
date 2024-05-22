// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuildArtifactResponseV1 Metadata for a single build within a product release
//
// swagger:model build_artifact_response_v1
type BuildArtifactResponseV1 struct {

	// The target architecture for this build
	// Example: amd64
	// Enum: [386 aarch64 all amd64 arm arm5 arm6 arm64 arm7 armelv5 armhf armhfv6 i386 i686 mips mips64 mipsle ppc64le s390x ui x86_64]
	Arch string `json:"arch,omitempty"`

	// The target operating system for this build
	// Example: darwin
	// Enum: [archlinux centos darwin debian dragonfly freebsd linux netbsd openbsd plan9 python solaris terraform web windows]
	Os string `json:"os,omitempty"`

	// True if this build is not supported by HashiCorp.  Some os/arch combinations are built
	// by HashiCorp for customer convenience but not officially supported.
	//
	// Example: false
	Unsupported bool `json:"unsupported,omitempty"`

	// The URL where this build can be downloaded
	// Example: https://releases.hashicorp.com/consul/1.10.0+ent/consul_1.10.0+ent_darwin_amd64.zip
	URL string `json:"url,omitempty"`
}

// Validate validates this build artifact response v1
func (m *BuildArtifactResponseV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var buildArtifactResponseV1TypeArchPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["386","aarch64","all","amd64","arm","arm5","arm6","arm64","arm7","armelv5","armhf","armhfv6","i386","i686","mips","mips64","mipsle","ppc64le","s390x","ui","x86_64"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buildArtifactResponseV1TypeArchPropEnum = append(buildArtifactResponseV1TypeArchPropEnum, v)
	}
}

const (

	// BuildArtifactResponseV1ArchNr386 captures enum value "386"
	BuildArtifactResponseV1ArchNr386 string = "386"

	// BuildArtifactResponseV1ArchAarch64 captures enum value "aarch64"
	BuildArtifactResponseV1ArchAarch64 string = "aarch64"

	// BuildArtifactResponseV1ArchAll captures enum value "all"
	BuildArtifactResponseV1ArchAll string = "all"

	// BuildArtifactResponseV1ArchAmd64 captures enum value "amd64"
	BuildArtifactResponseV1ArchAmd64 string = "amd64"

	// BuildArtifactResponseV1ArchArm captures enum value "arm"
	BuildArtifactResponseV1ArchArm string = "arm"

	// BuildArtifactResponseV1ArchArm5 captures enum value "arm5"
	BuildArtifactResponseV1ArchArm5 string = "arm5"

	// BuildArtifactResponseV1ArchArm6 captures enum value "arm6"
	BuildArtifactResponseV1ArchArm6 string = "arm6"

	// BuildArtifactResponseV1ArchArm64 captures enum value "arm64"
	BuildArtifactResponseV1ArchArm64 string = "arm64"

	// BuildArtifactResponseV1ArchArm7 captures enum value "arm7"
	BuildArtifactResponseV1ArchArm7 string = "arm7"

	// BuildArtifactResponseV1ArchArmelv5 captures enum value "armelv5"
	BuildArtifactResponseV1ArchArmelv5 string = "armelv5"

	// BuildArtifactResponseV1ArchArmhf captures enum value "armhf"
	BuildArtifactResponseV1ArchArmhf string = "armhf"

	// BuildArtifactResponseV1ArchArmhfv6 captures enum value "armhfv6"
	BuildArtifactResponseV1ArchArmhfv6 string = "armhfv6"

	// BuildArtifactResponseV1ArchI386 captures enum value "i386"
	BuildArtifactResponseV1ArchI386 string = "i386"

	// BuildArtifactResponseV1ArchI686 captures enum value "i686"
	BuildArtifactResponseV1ArchI686 string = "i686"

	// BuildArtifactResponseV1ArchMips captures enum value "mips"
	BuildArtifactResponseV1ArchMips string = "mips"

	// BuildArtifactResponseV1ArchMips64 captures enum value "mips64"
	BuildArtifactResponseV1ArchMips64 string = "mips64"

	// BuildArtifactResponseV1ArchMipsle captures enum value "mipsle"
	BuildArtifactResponseV1ArchMipsle string = "mipsle"

	// BuildArtifactResponseV1ArchPpc64le captures enum value "ppc64le"
	BuildArtifactResponseV1ArchPpc64le string = "ppc64le"

	// BuildArtifactResponseV1ArchS390x captures enum value "s390x"
	BuildArtifactResponseV1ArchS390x string = "s390x"

	// BuildArtifactResponseV1ArchUI captures enum value "ui"
	BuildArtifactResponseV1ArchUI string = "ui"

	// BuildArtifactResponseV1ArchX8664 captures enum value "x86_64"
	BuildArtifactResponseV1ArchX8664 string = "x86_64"
)

// prop value enum
func (m *BuildArtifactResponseV1) validateArchEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, buildArtifactResponseV1TypeArchPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BuildArtifactResponseV1) validateArch(formats strfmt.Registry) error {
	if swag.IsZero(m.Arch) { // not required
		return nil
	}

	// value enum
	if err := m.validateArchEnum("arch", "body", m.Arch); err != nil {
		return err
	}

	return nil
}

var buildArtifactResponseV1TypeOsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["archlinux","centos","darwin","debian","dragonfly","freebsd","linux","netbsd","openbsd","plan9","python","solaris","terraform","web","windows"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buildArtifactResponseV1TypeOsPropEnum = append(buildArtifactResponseV1TypeOsPropEnum, v)
	}
}

const (

	// BuildArtifactResponseV1OsArchlinux captures enum value "archlinux"
	BuildArtifactResponseV1OsArchlinux string = "archlinux"

	// BuildArtifactResponseV1OsCentos captures enum value "centos"
	BuildArtifactResponseV1OsCentos string = "centos"

	// BuildArtifactResponseV1OsDarwin captures enum value "darwin"
	BuildArtifactResponseV1OsDarwin string = "darwin"

	// BuildArtifactResponseV1OsDebian captures enum value "debian"
	BuildArtifactResponseV1OsDebian string = "debian"

	// BuildArtifactResponseV1OsDragonfly captures enum value "dragonfly"
	BuildArtifactResponseV1OsDragonfly string = "dragonfly"

	// BuildArtifactResponseV1OsFreebsd captures enum value "freebsd"
	BuildArtifactResponseV1OsFreebsd string = "freebsd"

	// BuildArtifactResponseV1OsLinux captures enum value "linux"
	BuildArtifactResponseV1OsLinux string = "linux"

	// BuildArtifactResponseV1OsNetbsd captures enum value "netbsd"
	BuildArtifactResponseV1OsNetbsd string = "netbsd"

	// BuildArtifactResponseV1OsOpenbsd captures enum value "openbsd"
	BuildArtifactResponseV1OsOpenbsd string = "openbsd"

	// BuildArtifactResponseV1OsPlan9 captures enum value "plan9"
	BuildArtifactResponseV1OsPlan9 string = "plan9"

	// BuildArtifactResponseV1OsPython captures enum value "python"
	BuildArtifactResponseV1OsPython string = "python"

	// BuildArtifactResponseV1OsSolaris captures enum value "solaris"
	BuildArtifactResponseV1OsSolaris string = "solaris"

	// BuildArtifactResponseV1OsTerraform captures enum value "terraform"
	BuildArtifactResponseV1OsTerraform string = "terraform"

	// BuildArtifactResponseV1OsWeb captures enum value "web"
	BuildArtifactResponseV1OsWeb string = "web"

	// BuildArtifactResponseV1OsWindows captures enum value "windows"
	BuildArtifactResponseV1OsWindows string = "windows"
)

// prop value enum
func (m *BuildArtifactResponseV1) validateOsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, buildArtifactResponseV1TypeOsPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BuildArtifactResponseV1) validateOs(formats strfmt.Registry) error {
	if swag.IsZero(m.Os) { // not required
		return nil
	}

	// value enum
	if err := m.validateOsEnum("os", "body", m.Os); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this build artifact response v1 based on context it is used
func (m *BuildArtifactResponseV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BuildArtifactResponseV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildArtifactResponseV1) UnmarshalBinary(b []byte) error {
	var res BuildArtifactResponseV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
