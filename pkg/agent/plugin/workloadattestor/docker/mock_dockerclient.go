// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/pkg/agent/plugin/workloadattestor/docker (interfaces: DockerClient)

// Package docker is a generated GoMock package.
package docker

import (
	context "context"
	reflect "reflect"

	types "github.com/docker/docker/api/types"
	gomock "github.com/golang/mock/gomock"
)

// MockDockerClient is a mock of DockerClient interface
type MockDockerClient struct {
	ctrl     *gomock.Controller
	recorder *MockDockerClientMockRecorder
}

// MockDockerClientMockRecorder is the mock recorder for MockDockerClient
type MockDockerClientMockRecorder struct {
	mock *MockDockerClient
}

// NewMockDockerClient creates a new mock instance
func NewMockDockerClient(ctrl *gomock.Controller) *MockDockerClient {
	mock := &MockDockerClient{ctrl: ctrl}
	mock.recorder = &MockDockerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDockerClient) EXPECT() *MockDockerClientMockRecorder {
	return m.recorder
}

// ContainerInspect mocks base method
func (m *MockDockerClient) ContainerInspect(arg0 context.Context, arg1 string) (types.ContainerJSON, error) {
	ret := m.ctrl.Call(m, "ContainerInspect", arg0, arg1)
	ret0, _ := ret[0].(types.ContainerJSON)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerInspect indicates an expected call of ContainerInspect
func (mr *MockDockerClientMockRecorder) ContainerInspect(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerInspect", reflect.TypeOf((*MockDockerClient)(nil).ContainerInspect), arg0, arg1)
}