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
// Source: github.com/apecloud/kubeblocks/internal/loadbalancer/netlink (interfaces: NetLink)

// Package mock_netlink is a generated GoMock package.
package mock_netlink

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	netlink "github.com/vishvananda/netlink"
)

// MockNetLink is a mock of NetLink interface.
type MockNetLink struct {
	ctrl     *gomock.Controller
	recorder *MockNetLinkMockRecorder
}

// MockNetLinkMockRecorder is the mock recorder for MockNetLink.
type MockNetLinkMockRecorder struct {
	mock *MockNetLink
}

// NewMockNetLink creates a new mock instance.
func NewMockNetLink(ctrl *gomock.Controller) *MockNetLink {
	mock := &MockNetLink{ctrl: ctrl}
	mock.recorder = &MockNetLinkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetLink) EXPECT() *MockNetLinkMockRecorder {
	return m.recorder
}

// AddrAdd mocks base method.
func (m *MockNetLink) AddrAdd(arg0 netlink.Link, arg1 *netlink.Addr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddrAdd", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddrAdd indicates an expected call of AddrAdd.
func (mr *MockNetLinkMockRecorder) AddrAdd(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrAdd", reflect.TypeOf((*MockNetLink)(nil).AddrAdd), arg0, arg1)
}

// AddrDel mocks base method.
func (m *MockNetLink) AddrDel(arg0 netlink.Link, arg1 *netlink.Addr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddrDel", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddrDel indicates an expected call of AddrDel.
func (mr *MockNetLinkMockRecorder) AddrDel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrDel", reflect.TypeOf((*MockNetLink)(nil).AddrDel), arg0, arg1)
}

// AddrList mocks base method.
func (m *MockNetLink) AddrList(arg0 netlink.Link, arg1 int) ([]netlink.Addr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddrList", arg0, arg1)
	ret0, _ := ret[0].([]netlink.Addr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddrList indicates an expected call of AddrList.
func (mr *MockNetLinkMockRecorder) AddrList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrList", reflect.TypeOf((*MockNetLink)(nil).AddrList), arg0, arg1)
}

// LinkList mocks base method.
func (m *MockNetLink) LinkList() ([]netlink.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkList")
	ret0, _ := ret[0].([]netlink.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LinkList indicates an expected call of LinkList.
func (mr *MockNetLinkMockRecorder) LinkList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkList", reflect.TypeOf((*MockNetLink)(nil).LinkList))
}

// LinkSetMTU mocks base method.
func (m *MockNetLink) LinkSetMTU(arg0 netlink.Link, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkSetMTU", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetMTU indicates an expected call of LinkSetMTU.
func (mr *MockNetLinkMockRecorder) LinkSetMTU(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetMTU", reflect.TypeOf((*MockNetLink)(nil).LinkSetMTU), arg0, arg1)
}

// LinkSetUp mocks base method.
func (m *MockNetLink) LinkSetUp(arg0 netlink.Link) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkSetUp", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetUp indicates an expected call of LinkSetUp.
func (mr *MockNetLinkMockRecorder) LinkSetUp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetUp", reflect.TypeOf((*MockNetLink)(nil).LinkSetUp), arg0)
}

// NewRule mocks base method.
func (m *MockNetLink) NewRule() *netlink.Rule {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRule")
	ret0, _ := ret[0].(*netlink.Rule)
	return ret0
}

// NewRule indicates an expected call of NewRule.
func (mr *MockNetLinkMockRecorder) NewRule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRule", reflect.TypeOf((*MockNetLink)(nil).NewRule))
}

// RouteDel mocks base method.
func (m *MockNetLink) RouteDel(arg0 *netlink.Route) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteDel", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RouteDel indicates an expected call of RouteDel.
func (mr *MockNetLinkMockRecorder) RouteDel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteDel", reflect.TypeOf((*MockNetLink)(nil).RouteDel), arg0)
}

// RouteList mocks base method.
func (m *MockNetLink) RouteList(arg0 netlink.Link, arg1 int) ([]netlink.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteList", arg0, arg1)
	ret0, _ := ret[0].([]netlink.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteList indicates an expected call of RouteList.
func (mr *MockNetLinkMockRecorder) RouteList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteList", reflect.TypeOf((*MockNetLink)(nil).RouteList), arg0, arg1)
}

// RouteReplace mocks base method.
func (m *MockNetLink) RouteReplace(arg0 *netlink.Route) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteReplace", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RouteReplace indicates an expected call of RouteReplace.
func (mr *MockNetLinkMockRecorder) RouteReplace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteReplace", reflect.TypeOf((*MockNetLink)(nil).RouteReplace), arg0)
}

// RuleAdd mocks base method.
func (m *MockNetLink) RuleAdd(arg0 *netlink.Rule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RuleAdd", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RuleAdd indicates an expected call of RuleAdd.
func (mr *MockNetLinkMockRecorder) RuleAdd(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuleAdd", reflect.TypeOf((*MockNetLink)(nil).RuleAdd), arg0)
}

// RuleDel mocks base method.
func (m *MockNetLink) RuleDel(arg0 *netlink.Rule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RuleDel", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RuleDel indicates an expected call of RuleDel.
func (mr *MockNetLinkMockRecorder) RuleDel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuleDel", reflect.TypeOf((*MockNetLink)(nil).RuleDel), arg0)
}

// RuleList mocks base method.
func (m *MockNetLink) RuleList(arg0 int) ([]netlink.Rule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RuleList", arg0)
	ret0, _ := ret[0].([]netlink.Rule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RuleList indicates an expected call of RuleList.
func (mr *MockNetLinkMockRecorder) RuleList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuleList", reflect.TypeOf((*MockNetLink)(nil).RuleList), arg0)
}
