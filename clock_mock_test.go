// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go

// Package clock is a generated GoMock package.
package clock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockClock is a mock of Clock interface
type MockClock struct {
	ctrl     *gomock.Controller
	recorder *MockClockMockRecorder
}

// MockClockMockRecorder is the mock recorder for MockClock
type MockClockMockRecorder struct {
	mock *MockClock
}

// NewMockClock creates a new mock instance
func NewMockClock(ctrl *gomock.Controller) *MockClock {
	mock := &MockClock{ctrl: ctrl}
	mock.recorder = &MockClockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClock) EXPECT() *MockClockMockRecorder {
	return m.recorder
}

// After mocks base method
func (m *MockClock) After(d time.Duration) <-chan time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "After", d)
	ret0, _ := ret[0].(<-chan time.Time)
	return ret0
}

// After indicates an expected call of After
func (mr *MockClockMockRecorder) After(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "After", reflect.TypeOf((*MockClock)(nil).After), d)
}

// AfterFunc mocks base method
func (m *MockClock) AfterFunc(d time.Duration, f func()) *Timer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterFunc", d, f)
	ret0, _ := ret[0].(*Timer)
	return ret0
}

// AfterFunc indicates an expected call of AfterFunc
func (mr *MockClockMockRecorder) AfterFunc(d, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterFunc", reflect.TypeOf((*MockClock)(nil).AfterFunc), d, f)
}

// NewTicker mocks base method
func (m *MockClock) NewTicker(d time.Duration) *Ticker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTicker", d)
	ret0, _ := ret[0].(*Ticker)
	return ret0
}

// NewTicker indicates an expected call of NewTicker
func (mr *MockClockMockRecorder) NewTicker(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTicker", reflect.TypeOf((*MockClock)(nil).NewTicker), d)
}

// NewTimer mocks base method
func (m *MockClock) NewTimer(d time.Duration) *Timer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTimer", d)
	ret0, _ := ret[0].(*Timer)
	return ret0
}

// NewTimer indicates an expected call of NewTimer
func (mr *MockClockMockRecorder) NewTimer(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTimer", reflect.TypeOf((*MockClock)(nil).NewTimer), d)
}

// Now mocks base method
func (m *MockClock) Now() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now
func (mr *MockClockMockRecorder) Now() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockClock)(nil).Now))
}

// Since mocks base method
func (m *MockClock) Since(t time.Time) time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Since", t)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Since indicates an expected call of Since
func (mr *MockClockMockRecorder) Since(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Since", reflect.TypeOf((*MockClock)(nil).Since), t)
}

// Sleep mocks base method
func (m *MockClock) Sleep(d time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Sleep", d)
}

// Sleep indicates an expected call of Sleep
func (mr *MockClockMockRecorder) Sleep(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sleep", reflect.TypeOf((*MockClock)(nil).Sleep), d)
}

// Tick mocks base method
func (m *MockClock) Tick(d time.Duration) <-chan time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tick", d)
	ret0, _ := ret[0].(<-chan time.Time)
	return ret0
}

// Tick indicates an expected call of Tick
func (mr *MockClockMockRecorder) Tick(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tick", reflect.TypeOf((*MockClock)(nil).Tick), d)
}

// Until mocks base method
func (m *MockClock) Until(t time.Time) time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Until", t)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Until indicates an expected call of Until
func (mr *MockClockMockRecorder) Until(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Until", reflect.TypeOf((*MockClock)(nil).Until), t)
}

// ContextWithDeadline mocks base method
func (m *MockClock) ContextWithDeadline(parent context.Context, d time.Time) (context.Context, context.CancelFunc) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContextWithDeadline", parent, d)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(context.CancelFunc)
	return ret0, ret1
}

// ContextWithDeadline indicates an expected call of ContextWithDeadline
func (mr *MockClockMockRecorder) ContextWithDeadline(parent, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContextWithDeadline", reflect.TypeOf((*MockClock)(nil).ContextWithDeadline), parent, d)
}

// ContextWithTimeout mocks base method
func (m *MockClock) ContextWithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContextWithTimeout", parent, timeout)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(context.CancelFunc)
	return ret0, ret1
}

// ContextWithTimeout indicates an expected call of ContextWithTimeout
func (mr *MockClockMockRecorder) ContextWithTimeout(parent, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContextWithTimeout", reflect.TypeOf((*MockClock)(nil).ContextWithTimeout), parent, timeout)
}
