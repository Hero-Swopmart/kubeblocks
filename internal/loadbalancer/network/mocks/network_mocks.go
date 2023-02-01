/*
Copyright ApeCloud, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/apecloud/kubeblocks/internal/loadbalancer/network (interfaces: Client)

// Package mock_network is a generated GoMock package.
package mock_network

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	cloud "github.com/apecloud/kubeblocks/internal/loadbalancer/cloud"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CleanNetworkForENI mocks base method.
func (m *MockClient) CleanNetworkForENI(arg0 *cloud.ENIMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanNetworkForENI", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanNetworkForENI indicates an expected call of CleanNetworkForENI.
func (mr *MockClientMockRecorder) CleanNetworkForENI(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanNetworkForENI", reflect.TypeOf((*MockClient)(nil).CleanNetworkForENI), arg0)
}

// CleanNetworkForService mocks base method.
func (m *MockClient) CleanNetworkForService(arg0 string, arg1 *cloud.ENIMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanNetworkForService", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanNetworkForService indicates an expected call of CleanNetworkForService.
func (mr *MockClientMockRecorder) CleanNetworkForService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanNetworkForService", reflect.TypeOf((*MockClient)(nil).CleanNetworkForService), arg0, arg1)
}

// SetupNetworkForENI mocks base method.
func (m *MockClient) SetupNetworkForENI(arg0 *cloud.ENIMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupNetworkForENI", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupNetworkForENI indicates an expected call of SetupNetworkForENI.
func (mr *MockClientMockRecorder) SetupNetworkForENI(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupNetworkForENI", reflect.TypeOf((*MockClient)(nil).SetupNetworkForENI), arg0)
}

// SetupNetworkForService mocks base method.
func (m *MockClient) SetupNetworkForService(arg0 string, arg1 *cloud.ENIMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupNetworkForService", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupNetworkForService indicates an expected call of SetupNetworkForService.
func (mr *MockClientMockRecorder) SetupNetworkForService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupNetworkForService", reflect.TypeOf((*MockClient)(nil).SetupNetworkForService), arg0, arg1)
}
