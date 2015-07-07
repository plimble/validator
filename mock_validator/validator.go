package mock_validator

import "github.com/stretchr/testify/mock"

import "time"

type MockValidator struct {
	mock.Mock
}

func NewMockValidator() *MockValidator {
	return &MockValidator{}
}

func (m *MockValidator) HasError() bool {
	ret := m.Called()

	r0 := ret.Get(0).(bool)

	return r0
}
func (m *MockValidator) Messages() map[string]string {
	ret := m.Called()

	var r0 map[string]string
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(map[string]string)
	}

	return r0
}
func (m *MockValidator) GetError() error {
	ret := m.Called()

	r0 := ret.Error(0)

	return r0
}
func (m *MockValidator) RequiredString(val string, err error, name ...string) {
	m.Called(val, err, name)
}
func (m *MockValidator) RequiredBytes(val []byte, err error, name ...string) {
	m.Called(val, err, name)
}
func (m *MockValidator) RequiredInt(val int, err error, name ...string) {
	m.Called(val, err, name)
}
func (m *MockValidator) RequiredFloat64(val float64, err error, name ...string) {
	m.Called(val, err, name)
}
func (m *MockValidator) RequiredBool(val bool, err error, name ...string) {
	m.Called(val, err, name)
}
func (m *MockValidator) RequiredEmail(val string, err error, name ...string) {
	m.Called(val, err, name)
}
func (m *MockValidator) NotNil(val interface{}, err error, name ...string) {
	m.Called(val, err, name)
}
func (m *MockValidator) RequiredTime(val time.Time, err error, name ...string) {
	m.Called(val, err, name)
}
func (m *MockValidator) MinInt(val int, n int, err error, name ...string) {
	m.Called(val, n, err, name)
}
func (m *MockValidator) MaxInt(val int, n int, err error, name ...string) {
	m.Called(val, n, err, name)
}
func (m *MockValidator) MinFloat64(val float64, n float64, err error, name ...string) {
	m.Called(val, n, err, name)
}
func (m *MockValidator) MaxFloat64(val float64, n float64, err error, name ...string) {
	m.Called(val, n, err, name)
}
func (m *MockValidator) MinChar(val string, n int, err error, name ...string) {
	m.Called(val, n, err, name)
}
func (m *MockValidator) MaxChar(val string, n int, err error, name ...string) {
	m.Called(val, n, err, name)
}
func (m *MockValidator) Email(val string, err error, name ...string) {
	m.Called(val, err, name)
}
func (m *MockValidator) Gender(val string, err error, name ...string) {
	m.Called(val, err, name)
}
func (m *MockValidator) Confirm(val string, confirm string, err error, name ...string) {
	m.Called(val, confirm, err, name)
}
func (m *MockValidator) ISO8601DataTime(val string, err error, name ...string) {
	m.Called(val, err, name)
}
func (m *MockValidator) Length(val int, atleast int, err error, name ...string) {
	m.Called(val, atleast, err, name)
}
func (m *MockValidator) InString(val string, in []string, err error, name ...string) {
	m.Called(val, in, err, name)
}
