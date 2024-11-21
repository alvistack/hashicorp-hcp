// Code generated by mockery. DO NOT EDIT.

package mock_authorization_service

import (
	authorization_service "github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/stable/2019-12-10/client/authorization_service"
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

// AuthorizationServiceBatchTestIamPermissions provides a mock function with given fields: params, authInfo, opts
func (_m *MockClientService) AuthorizationServiceBatchTestIamPermissions(params *authorization_service.AuthorizationServiceBatchTestIamPermissionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...authorization_service.ClientOption) (*authorization_service.AuthorizationServiceBatchTestIamPermissionsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AuthorizationServiceBatchTestIamPermissions")
	}

	var r0 *authorization_service.AuthorizationServiceBatchTestIamPermissionsOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*authorization_service.AuthorizationServiceBatchTestIamPermissionsParams, runtime.ClientAuthInfoWriter, ...authorization_service.ClientOption) (*authorization_service.AuthorizationServiceBatchTestIamPermissionsOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*authorization_service.AuthorizationServiceBatchTestIamPermissionsParams, runtime.ClientAuthInfoWriter, ...authorization_service.ClientOption) *authorization_service.AuthorizationServiceBatchTestIamPermissionsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authorization_service.AuthorizationServiceBatchTestIamPermissionsOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*authorization_service.AuthorizationServiceBatchTestIamPermissionsParams, runtime.ClientAuthInfoWriter, ...authorization_service.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientService_AuthorizationServiceBatchTestIamPermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthorizationServiceBatchTestIamPermissions'
type MockClientService_AuthorizationServiceBatchTestIamPermissions_Call struct {
	*mock.Call
}

// AuthorizationServiceBatchTestIamPermissions is a helper method to define mock.On call
//   - params *authorization_service.AuthorizationServiceBatchTestIamPermissionsParams
//   - authInfo runtime.ClientAuthInfoWriter
//   - opts ...authorization_service.ClientOption
func (_e *MockClientService_Expecter) AuthorizationServiceBatchTestIamPermissions(params interface{}, authInfo interface{}, opts ...interface{}) *MockClientService_AuthorizationServiceBatchTestIamPermissions_Call {
	return &MockClientService_AuthorizationServiceBatchTestIamPermissions_Call{Call: _e.mock.On("AuthorizationServiceBatchTestIamPermissions",
		append([]interface{}{params, authInfo}, opts...)...)}
}

func (_c *MockClientService_AuthorizationServiceBatchTestIamPermissions_Call) Run(run func(params *authorization_service.AuthorizationServiceBatchTestIamPermissionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...authorization_service.ClientOption)) *MockClientService_AuthorizationServiceBatchTestIamPermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]authorization_service.ClientOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(authorization_service.ClientOption)
			}
		}
		run(args[0].(*authorization_service.AuthorizationServiceBatchTestIamPermissionsParams), args[1].(runtime.ClientAuthInfoWriter), variadicArgs...)
	})
	return _c
}

func (_c *MockClientService_AuthorizationServiceBatchTestIamPermissions_Call) Return(_a0 *authorization_service.AuthorizationServiceBatchTestIamPermissionsOK, _a1 error) *MockClientService_AuthorizationServiceBatchTestIamPermissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClientService_AuthorizationServiceBatchTestIamPermissions_Call) RunAndReturn(run func(*authorization_service.AuthorizationServiceBatchTestIamPermissionsParams, runtime.ClientAuthInfoWriter, ...authorization_service.ClientOption) (*authorization_service.AuthorizationServiceBatchTestIamPermissionsOK, error)) *MockClientService_AuthorizationServiceBatchTestIamPermissions_Call {
	_c.Call.Return(run)
	return _c
}

// AuthorizationServiceTestIamPermissions provides a mock function with given fields: params, authInfo, opts
func (_m *MockClientService) AuthorizationServiceTestIamPermissions(params *authorization_service.AuthorizationServiceTestIamPermissionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...authorization_service.ClientOption) (*authorization_service.AuthorizationServiceTestIamPermissionsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AuthorizationServiceTestIamPermissions")
	}

	var r0 *authorization_service.AuthorizationServiceTestIamPermissionsOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*authorization_service.AuthorizationServiceTestIamPermissionsParams, runtime.ClientAuthInfoWriter, ...authorization_service.ClientOption) (*authorization_service.AuthorizationServiceTestIamPermissionsOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*authorization_service.AuthorizationServiceTestIamPermissionsParams, runtime.ClientAuthInfoWriter, ...authorization_service.ClientOption) *authorization_service.AuthorizationServiceTestIamPermissionsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authorization_service.AuthorizationServiceTestIamPermissionsOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*authorization_service.AuthorizationServiceTestIamPermissionsParams, runtime.ClientAuthInfoWriter, ...authorization_service.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientService_AuthorizationServiceTestIamPermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthorizationServiceTestIamPermissions'
type MockClientService_AuthorizationServiceTestIamPermissions_Call struct {
	*mock.Call
}

// AuthorizationServiceTestIamPermissions is a helper method to define mock.On call
//   - params *authorization_service.AuthorizationServiceTestIamPermissionsParams
//   - authInfo runtime.ClientAuthInfoWriter
//   - opts ...authorization_service.ClientOption
func (_e *MockClientService_Expecter) AuthorizationServiceTestIamPermissions(params interface{}, authInfo interface{}, opts ...interface{}) *MockClientService_AuthorizationServiceTestIamPermissions_Call {
	return &MockClientService_AuthorizationServiceTestIamPermissions_Call{Call: _e.mock.On("AuthorizationServiceTestIamPermissions",
		append([]interface{}{params, authInfo}, opts...)...)}
}

func (_c *MockClientService_AuthorizationServiceTestIamPermissions_Call) Run(run func(params *authorization_service.AuthorizationServiceTestIamPermissionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...authorization_service.ClientOption)) *MockClientService_AuthorizationServiceTestIamPermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]authorization_service.ClientOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(authorization_service.ClientOption)
			}
		}
		run(args[0].(*authorization_service.AuthorizationServiceTestIamPermissionsParams), args[1].(runtime.ClientAuthInfoWriter), variadicArgs...)
	})
	return _c
}

func (_c *MockClientService_AuthorizationServiceTestIamPermissions_Call) Return(_a0 *authorization_service.AuthorizationServiceTestIamPermissionsOK, _a1 error) *MockClientService_AuthorizationServiceTestIamPermissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClientService_AuthorizationServiceTestIamPermissions_Call) RunAndReturn(run func(*authorization_service.AuthorizationServiceTestIamPermissionsParams, runtime.ClientAuthInfoWriter, ...authorization_service.ClientOption) (*authorization_service.AuthorizationServiceTestIamPermissionsOK, error)) *MockClientService_AuthorizationServiceTestIamPermissions_Call {
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
