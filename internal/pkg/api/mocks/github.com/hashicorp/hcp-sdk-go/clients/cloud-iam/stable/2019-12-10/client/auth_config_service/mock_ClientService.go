// Code generated by mockery. DO NOT EDIT.

package mock_auth_config_service

import (
	auth_config_service "github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/client/auth_config_service"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// MockClientService is an autogenerated mock type for the ClientService type
type MockClientService struct {
	mock.Mock
}

type MockClientService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClientService) EXPECT() *MockClientService_Expecter {
	return &MockClientService_Expecter{mock: &_m.Mock}
}

// AuthConfigServiceCreateAuthConnection provides a mock function with given fields: params, authInfo, opts
func (_m *MockClientService) AuthConfigServiceCreateAuthConnection(params *auth_config_service.AuthConfigServiceCreateAuthConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...auth_config_service.ClientOption) (*auth_config_service.AuthConfigServiceCreateAuthConnectionOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AuthConfigServiceCreateAuthConnection")
	}

	var r0 *auth_config_service.AuthConfigServiceCreateAuthConnectionOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*auth_config_service.AuthConfigServiceCreateAuthConnectionParams, runtime.ClientAuthInfoWriter, ...auth_config_service.ClientOption) (*auth_config_service.AuthConfigServiceCreateAuthConnectionOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*auth_config_service.AuthConfigServiceCreateAuthConnectionParams, runtime.ClientAuthInfoWriter, ...auth_config_service.ClientOption) *auth_config_service.AuthConfigServiceCreateAuthConnectionOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth_config_service.AuthConfigServiceCreateAuthConnectionOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*auth_config_service.AuthConfigServiceCreateAuthConnectionParams, runtime.ClientAuthInfoWriter, ...auth_config_service.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientService_AuthConfigServiceCreateAuthConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthConfigServiceCreateAuthConnection'
type MockClientService_AuthConfigServiceCreateAuthConnection_Call struct {
	*mock.Call
}

// AuthConfigServiceCreateAuthConnection is a helper method to define mock.On call
//   - params *auth_config_service.AuthConfigServiceCreateAuthConnectionParams
//   - authInfo runtime.ClientAuthInfoWriter
//   - opts ...auth_config_service.ClientOption
func (_e *MockClientService_Expecter) AuthConfigServiceCreateAuthConnection(params interface{}, authInfo interface{}, opts ...interface{}) *MockClientService_AuthConfigServiceCreateAuthConnection_Call {
	return &MockClientService_AuthConfigServiceCreateAuthConnection_Call{Call: _e.mock.On("AuthConfigServiceCreateAuthConnection",
		append([]interface{}{params, authInfo}, opts...)...)}
}

func (_c *MockClientService_AuthConfigServiceCreateAuthConnection_Call) Run(run func(params *auth_config_service.AuthConfigServiceCreateAuthConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...auth_config_service.ClientOption)) *MockClientService_AuthConfigServiceCreateAuthConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]auth_config_service.ClientOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(auth_config_service.ClientOption)
			}
		}
		run(args[0].(*auth_config_service.AuthConfigServiceCreateAuthConnectionParams), args[1].(runtime.ClientAuthInfoWriter), variadicArgs...)
	})
	return _c
}

func (_c *MockClientService_AuthConfigServiceCreateAuthConnection_Call) Return(_a0 *auth_config_service.AuthConfigServiceCreateAuthConnectionOK, _a1 error) *MockClientService_AuthConfigServiceCreateAuthConnection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClientService_AuthConfigServiceCreateAuthConnection_Call) RunAndReturn(run func(*auth_config_service.AuthConfigServiceCreateAuthConnectionParams, runtime.ClientAuthInfoWriter, ...auth_config_service.ClientOption) (*auth_config_service.AuthConfigServiceCreateAuthConnectionOK, error)) *MockClientService_AuthConfigServiceCreateAuthConnection_Call {
	_c.Call.Return(run)
	return _c
}

// AuthConfigServiceDeleteAuthConnectionFromOrganization provides a mock function with given fields: params, authInfo, opts
func (_m *MockClientService) AuthConfigServiceDeleteAuthConnectionFromOrganization(params *auth_config_service.AuthConfigServiceDeleteAuthConnectionFromOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...auth_config_service.ClientOption) (*auth_config_service.AuthConfigServiceDeleteAuthConnectionFromOrganizationOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AuthConfigServiceDeleteAuthConnectionFromOrganization")
	}

	var r0 *auth_config_service.AuthConfigServiceDeleteAuthConnectionFromOrganizationOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*auth_config_service.AuthConfigServiceDeleteAuthConnectionFromOrganizationParams, runtime.ClientAuthInfoWriter, ...auth_config_service.ClientOption) (*auth_config_service.AuthConfigServiceDeleteAuthConnectionFromOrganizationOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*auth_config_service.AuthConfigServiceDeleteAuthConnectionFromOrganizationParams, runtime.ClientAuthInfoWriter, ...auth_config_service.ClientOption) *auth_config_service.AuthConfigServiceDeleteAuthConnectionFromOrganizationOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth_config_service.AuthConfigServiceDeleteAuthConnectionFromOrganizationOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*auth_config_service.AuthConfigServiceDeleteAuthConnectionFromOrganizationParams, runtime.ClientAuthInfoWriter, ...auth_config_service.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientService_AuthConfigServiceDeleteAuthConnectionFromOrganization_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthConfigServiceDeleteAuthConnectionFromOrganization'
type MockClientService_AuthConfigServiceDeleteAuthConnectionFromOrganization_Call struct {
	*mock.Call
}

// AuthConfigServiceDeleteAuthConnectionFromOrganization is a helper method to define mock.On call
//   - params *auth_config_service.AuthConfigServiceDeleteAuthConnectionFromOrganizationParams
//   - authInfo runtime.ClientAuthInfoWriter
//   - opts ...auth_config_service.ClientOption
func (_e *MockClientService_Expecter) AuthConfigServiceDeleteAuthConnectionFromOrganization(params interface{}, authInfo interface{}, opts ...interface{}) *MockClientService_AuthConfigServiceDeleteAuthConnectionFromOrganization_Call {
	return &MockClientService_AuthConfigServiceDeleteAuthConnectionFromOrganization_Call{Call: _e.mock.On("AuthConfigServiceDeleteAuthConnectionFromOrganization",
		append([]interface{}{params, authInfo}, opts...)...)}
}

func (_c *MockClientService_AuthConfigServiceDeleteAuthConnectionFromOrganization_Call) Run(run func(params *auth_config_service.AuthConfigServiceDeleteAuthConnectionFromOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...auth_config_service.ClientOption)) *MockClientService_AuthConfigServiceDeleteAuthConnectionFromOrganization_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]auth_config_service.ClientOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(auth_config_service.ClientOption)
			}
		}
		run(args[0].(*auth_config_service.AuthConfigServiceDeleteAuthConnectionFromOrganizationParams), args[1].(runtime.ClientAuthInfoWriter), variadicArgs...)
	})
	return _c
}

func (_c *MockClientService_AuthConfigServiceDeleteAuthConnectionFromOrganization_Call) Return(_a0 *auth_config_service.AuthConfigServiceDeleteAuthConnectionFromOrganizationOK, _a1 error) *MockClientService_AuthConfigServiceDeleteAuthConnectionFromOrganization_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClientService_AuthConfigServiceDeleteAuthConnectionFromOrganization_Call) RunAndReturn(run func(*auth_config_service.AuthConfigServiceDeleteAuthConnectionFromOrganizationParams, runtime.ClientAuthInfoWriter, ...auth_config_service.ClientOption) (*auth_config_service.AuthConfigServiceDeleteAuthConnectionFromOrganizationOK, error)) *MockClientService_AuthConfigServiceDeleteAuthConnectionFromOrganization_Call {
	_c.Call.Return(run)
	return _c
}

// AuthConfigServiceEditAuthConnection provides a mock function with given fields: params, authInfo, opts
func (_m *MockClientService) AuthConfigServiceEditAuthConnection(params *auth_config_service.AuthConfigServiceEditAuthConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...auth_config_service.ClientOption) (*auth_config_service.AuthConfigServiceEditAuthConnectionOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AuthConfigServiceEditAuthConnection")
	}

	var r0 *auth_config_service.AuthConfigServiceEditAuthConnectionOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*auth_config_service.AuthConfigServiceEditAuthConnectionParams, runtime.ClientAuthInfoWriter, ...auth_config_service.ClientOption) (*auth_config_service.AuthConfigServiceEditAuthConnectionOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*auth_config_service.AuthConfigServiceEditAuthConnectionParams, runtime.ClientAuthInfoWriter, ...auth_config_service.ClientOption) *auth_config_service.AuthConfigServiceEditAuthConnectionOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth_config_service.AuthConfigServiceEditAuthConnectionOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*auth_config_service.AuthConfigServiceEditAuthConnectionParams, runtime.ClientAuthInfoWriter, ...auth_config_service.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientService_AuthConfigServiceEditAuthConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthConfigServiceEditAuthConnection'
type MockClientService_AuthConfigServiceEditAuthConnection_Call struct {
	*mock.Call
}

// AuthConfigServiceEditAuthConnection is a helper method to define mock.On call
//   - params *auth_config_service.AuthConfigServiceEditAuthConnectionParams
//   - authInfo runtime.ClientAuthInfoWriter
//   - opts ...auth_config_service.ClientOption
func (_e *MockClientService_Expecter) AuthConfigServiceEditAuthConnection(params interface{}, authInfo interface{}, opts ...interface{}) *MockClientService_AuthConfigServiceEditAuthConnection_Call {
	return &MockClientService_AuthConfigServiceEditAuthConnection_Call{Call: _e.mock.On("AuthConfigServiceEditAuthConnection",
		append([]interface{}{params, authInfo}, opts...)...)}
}

func (_c *MockClientService_AuthConfigServiceEditAuthConnection_Call) Run(run func(params *auth_config_service.AuthConfigServiceEditAuthConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...auth_config_service.ClientOption)) *MockClientService_AuthConfigServiceEditAuthConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]auth_config_service.ClientOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(auth_config_service.ClientOption)
			}
		}
		run(args[0].(*auth_config_service.AuthConfigServiceEditAuthConnectionParams), args[1].(runtime.ClientAuthInfoWriter), variadicArgs...)
	})
	return _c
}

func (_c *MockClientService_AuthConfigServiceEditAuthConnection_Call) Return(_a0 *auth_config_service.AuthConfigServiceEditAuthConnectionOK, _a1 error) *MockClientService_AuthConfigServiceEditAuthConnection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClientService_AuthConfigServiceEditAuthConnection_Call) RunAndReturn(run func(*auth_config_service.AuthConfigServiceEditAuthConnectionParams, runtime.ClientAuthInfoWriter, ...auth_config_service.ClientOption) (*auth_config_service.AuthConfigServiceEditAuthConnectionOK, error)) *MockClientService_AuthConfigServiceEditAuthConnection_Call {
	_c.Call.Return(run)
	return _c
}

// AuthConfigServiceGetAuthConnections provides a mock function with given fields: params, authInfo, opts
func (_m *MockClientService) AuthConfigServiceGetAuthConnections(params *auth_config_service.AuthConfigServiceGetAuthConnectionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...auth_config_service.ClientOption) (*auth_config_service.AuthConfigServiceGetAuthConnectionsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AuthConfigServiceGetAuthConnections")
	}

	var r0 *auth_config_service.AuthConfigServiceGetAuthConnectionsOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*auth_config_service.AuthConfigServiceGetAuthConnectionsParams, runtime.ClientAuthInfoWriter, ...auth_config_service.ClientOption) (*auth_config_service.AuthConfigServiceGetAuthConnectionsOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*auth_config_service.AuthConfigServiceGetAuthConnectionsParams, runtime.ClientAuthInfoWriter, ...auth_config_service.ClientOption) *auth_config_service.AuthConfigServiceGetAuthConnectionsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth_config_service.AuthConfigServiceGetAuthConnectionsOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*auth_config_service.AuthConfigServiceGetAuthConnectionsParams, runtime.ClientAuthInfoWriter, ...auth_config_service.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientService_AuthConfigServiceGetAuthConnections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthConfigServiceGetAuthConnections'
type MockClientService_AuthConfigServiceGetAuthConnections_Call struct {
	*mock.Call
}

// AuthConfigServiceGetAuthConnections is a helper method to define mock.On call
//   - params *auth_config_service.AuthConfigServiceGetAuthConnectionsParams
//   - authInfo runtime.ClientAuthInfoWriter
//   - opts ...auth_config_service.ClientOption
func (_e *MockClientService_Expecter) AuthConfigServiceGetAuthConnections(params interface{}, authInfo interface{}, opts ...interface{}) *MockClientService_AuthConfigServiceGetAuthConnections_Call {
	return &MockClientService_AuthConfigServiceGetAuthConnections_Call{Call: _e.mock.On("AuthConfigServiceGetAuthConnections",
		append([]interface{}{params, authInfo}, opts...)...)}
}

func (_c *MockClientService_AuthConfigServiceGetAuthConnections_Call) Run(run func(params *auth_config_service.AuthConfigServiceGetAuthConnectionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...auth_config_service.ClientOption)) *MockClientService_AuthConfigServiceGetAuthConnections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]auth_config_service.ClientOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(auth_config_service.ClientOption)
			}
		}
		run(args[0].(*auth_config_service.AuthConfigServiceGetAuthConnectionsParams), args[1].(runtime.ClientAuthInfoWriter), variadicArgs...)
	})
	return _c
}

func (_c *MockClientService_AuthConfigServiceGetAuthConnections_Call) Return(_a0 *auth_config_service.AuthConfigServiceGetAuthConnectionsOK, _a1 error) *MockClientService_AuthConfigServiceGetAuthConnections_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClientService_AuthConfigServiceGetAuthConnections_Call) RunAndReturn(run func(*auth_config_service.AuthConfigServiceGetAuthConnectionsParams, runtime.ClientAuthInfoWriter, ...auth_config_service.ClientOption) (*auth_config_service.AuthConfigServiceGetAuthConnectionsOK, error)) *MockClientService_AuthConfigServiceGetAuthConnections_Call {
	_c.Call.Return(run)
	return _c
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// MockClientService_SetTransport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTransport'
type MockClientService_SetTransport_Call struct {
	*mock.Call
}

// SetTransport is a helper method to define mock.On call
//   - transport runtime.ClientTransport
func (_e *MockClientService_Expecter) SetTransport(transport interface{}) *MockClientService_SetTransport_Call {
	return &MockClientService_SetTransport_Call{Call: _e.mock.On("SetTransport", transport)}
}

func (_c *MockClientService_SetTransport_Call) Run(run func(transport runtime.ClientTransport)) *MockClientService_SetTransport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(runtime.ClientTransport))
	})
	return _c
}

func (_c *MockClientService_SetTransport_Call) Return() *MockClientService_SetTransport_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockClientService_SetTransport_Call) RunAndReturn(run func(runtime.ClientTransport)) *MockClientService_SetTransport_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClientService creates a new instance of MockClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClientService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClientService {
	mock := &MockClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
