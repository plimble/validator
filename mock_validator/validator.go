package mock_validator

import "github.com/stretchr/testify/mock"

import "time"

type Validator struct {
	mock.Mock
}

func (_m *Validator) HasError() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
func (_m *Validator) Messages() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}
func (_m *Validator) GetError() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *Validator) RequiredString(val string, name string, err ...error) {
	_m.Called(val, name, err)
}
func (_m *Validator) RequiredBytes(val []byte, name string, err ...error) {
	_m.Called(val, name, err)
}
func (_m *Validator) RequiredInt(val int, name string, err ...error) {
	_m.Called(val, name, err)
}
func (_m *Validator) RequiredInt64(val int64, name string, err ...error) {
	_m.Called(val, name, err)
}
func (_m *Validator) RequiredFloat64(val float64, name string, err ...error) {
	_m.Called(val, name, err)
}
func (_m *Validator) RequiredBool(val bool, name string, err ...error) {
	_m.Called(val, name, err)
}
func (_m *Validator) RequiredEmail(val string, name string, err ...error) {
	_m.Called(val, name, err)
}
func (_m *Validator) NotNil(val interface{}, name string, err ...error) {
	_m.Called(val, name, err)
}
func (_m *Validator) RequiredTime(val time.Time, name string, err ...error) {
	_m.Called(val, name, err)
}
func (_m *Validator) MinInt(val int, n int, name string, err ...error) {
	_m.Called(val, n, name, err)
}
func (_m *Validator) MaxInt(val int, n int, name string, err ...error) {
	_m.Called(val, n, name, err)
}
func (_m *Validator) MinInt64(val int64, n int64, name string, err ...error) {
	_m.Called(val, n, name, err)
}
func (_m *Validator) MaxInt64(val int64, n int64, name string, err ...error) {
	_m.Called(val, n, name, err)
}
func (_m *Validator) MinFloat64(val float64, n float64, name string, err ...error) {
	_m.Called(val, n, name, err)
}
func (_m *Validator) MaxFloat64(val float64, n float64, name string, err ...error) {
	_m.Called(val, n, name, err)
}
func (_m *Validator) MinChar(val string, n int, name string, err ...error) {
	_m.Called(val, n, name, err)
}
func (_m *Validator) MaxChar(val string, n int, name string, err ...error) {
	_m.Called(val, n, name, err)
}
func (_m *Validator) Email(val string, name string, err ...error) {
	_m.Called(val, name, err)
}
func (_m *Validator) Gender(val string, name string, err ...error) {
	_m.Called(val, name, err)
}
func (_m *Validator) Confirm(val string, confirm string, name string, confirmName string, err ...error) {
	_m.Called(val, confirm, name, confirmName, err)
}
func (_m *Validator) ISO8601DataTime(val string, name string, err ...error) {
	_m.Called(val, name, err)
}
func (_m *Validator) InString(val string, in []string, name string, err ...error) {
	_m.Called(val, in, name, err)
}
