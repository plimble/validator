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
func (m *MockValidator) RequiredString(val string, name string, err error) {
	m.Called(val, name, err)
}
func (m *MockValidator) RequiredBytes(val []byte, name string, err error) {
	m.Called(val, name, err)
}
func (m *MockValidator) RequiredInt(val int, name string, err error) {
	m.Called(val, name, err)
}
func (m *MockValidator) RequiredFloat64(val float64, name string, err error) {
	m.Called(val, name, err)
}
func (m *MockValidator) RequiredBool(val bool, name string, err error) {
	m.Called(val, name, err)
}
func (m *MockValidator) RequiredEmail(val string, name string, err error) {
	m.Called(val, name, err)
}
func (m *MockValidator) NotNil(val interface{}, name string, err error) {
	m.Called(val, name, err)
}
func (m *MockValidator) RequiredTime(val time.Time, name string, err error) {
	m.Called(val, name, err)
}
func (m *MockValidator) MinInt(val int, n int, name string, err error) {
	m.Called(val, n, name, err)
}
func (m *MockValidator) MaxInt(val int, n int, name string, err error) {
	m.Called(val, n, name, err)
}
func (m *MockValidator) MinFloat64(val float64, n float64, name string, err error) {
	m.Called(val, n, name, err)
}
func (m *MockValidator) MaxFloat64(val float64, n float64, name string, err error) {
	m.Called(val, n, name, err)
}
func (m *MockValidator) MinChar(val string, n int, name string, err error) {
	m.Called(val, n, name, err)
}
func (m *MockValidator) MaxChar(val string, n int, name string, err error) {
	m.Called(val, n, name, err)
}
func (m *MockValidator) Email(val string, name string, err error) {
	m.Called(val, name, err)
}
func (m *MockValidator) Gender(val string, name string, err error) {
	m.Called(val, name, err)
}
func (m *MockValidator) Confirm(val string, confirm string, name string, confirmName string, err error) {
	m.Called(val, confirm, name, confirmName, err)
}
func (m *MockValidator) ISO8601DataTime(val string, name string, err error) {
	m.Called(val, name, err)
}
func (m *MockValidator) InString(val string, in []string, name string, err error) {
	m.Called(val, in, name, err)
}
