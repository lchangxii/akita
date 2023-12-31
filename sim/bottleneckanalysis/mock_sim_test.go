// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/akita/akita/v3/sim (interfaces: TimeTeller)

package bottleneckanalysis

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sim "gitlab.com/akita/akita/v3/sim"
)

// MockTimeTeller is a mock of TimeTeller interface.
type MockTimeTeller struct {
	ctrl     *gomock.Controller
	recorder *MockTimeTellerMockRecorder
}

// MockTimeTellerMockRecorder is the mock recorder for MockTimeTeller.
type MockTimeTellerMockRecorder struct {
	mock *MockTimeTeller
}

// NewMockTimeTeller creates a new mock instance.
func NewMockTimeTeller(ctrl *gomock.Controller) *MockTimeTeller {
	mock := &MockTimeTeller{ctrl: ctrl}
	mock.recorder = &MockTimeTellerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTimeTeller) EXPECT() *MockTimeTellerMockRecorder {
	return m.recorder
}

// CurrentTime mocks base method.
func (m *MockTimeTeller) CurrentTime() sim.VTimeInSec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentTime")
	ret0, _ := ret[0].(sim.VTimeInSec)
	return ret0
}

// CurrentTime indicates an expected call of CurrentTime.
func (mr *MockTimeTellerMockRecorder) CurrentTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentTime", reflect.TypeOf((*MockTimeTeller)(nil).CurrentTime))
}
