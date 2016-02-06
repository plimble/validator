package validator

//go:generate mockery -name Validator -output mock_validator

import (
	"github.com/plimble/errors"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	emailPatern       = regexp.MustCompile(".+@.+\\..+")
	dateiso8601Patern = regexp.MustCompile("^(\\d{4})-(\\d{2})-(\\d{2})T(\\d{2}):(\\d{2}):(\\d{2})(Z|(\\+|-)\\d{2}(:?\\d{2})?)$")
)

type ValidateError struct {
	Name string
	Err  error
}

type Validator interface {
	HasError() bool
	Messages() map[string]string
	GetError() errors.Error
	AddError(name string, err error)
	AddErrorMsg(name, format string, args ...interface{})
	RequiredString(val string, name string, err ...error)
	RequiredBytes(val []byte, name string, err ...error)
	RequiredInt(val int, name string, err ...error)
	RequiredInt64(val int64, name string, err ...error)
	RequiredFloat64(val float64, name string, err ...error)
	RequiredBool(val bool, name string, err ...error)
	RequiredEmail(val string, name string, err ...error)
	NotNil(val interface{}, name string, err ...error)
	RequiredTime(val time.Time, name string, err ...error)
	MinInt(val int, n int, name string, err ...error)
	MaxInt(val int, n int, name string, err ...error)
	MinInt64(val int64, n int64, name string, err ...error)
	MaxInt64(val int64, n int64, name string, err ...error)
	MinFloat64(val float64, n float64, name string, err ...error)
	MaxFloat64(val float64, n float64, name string, err ...error)
	MinChar(val string, n int, name string, err ...error)
	MaxChar(val string, n int, name string, err ...error)
	Email(val string, name string, err ...error)
	Gender(val string, name string, err ...error)
	Confirm(val, confirm string, name string, confirmName string, err ...error)
	ISO8601DataTime(val string, name string, err ...error)
	InString(val string, in []string, name string, err ...error)
}

type validator struct {
	errs []ValidateError
}

func New() Validator {
	return &validator{
		errs: []ValidateError{},
	}
}

func (v *validator) HasError() bool {
	if len(v.errs) > 0 {
		return true
	}

	return false
}

func (v *validator) Messages() map[string]string {
	msgs := make(map[string]string)
	for i := 0; i < len(v.errs); i++ {
		if v.errs[i].Name != "" {
			msgs[v.errs[i].Name] = v.errs[i].Err.Error()
		}
	}

	return msgs
}

func (v *validator) GetError() errors.Error {
	if len(v.errs) > 0 {
		return errors.BadRequest(v.errs[0].Err.Error())
	}

	return nil
}

func (v *validator) add(name string, err error, errs []error) {
	if len(errs) > 0 {
		err = errs[0]
	}

	v.errs = append(v.errs, ValidateError{name, err})
}

func (v *validator) AddError(name string, err error) {
	v.errs = append(v.errs, ValidateError{name, err})
}

func (v *validator) AddErrorMsg(name, format string, args ...interface{}) {
	v.errs = append(v.errs, ValidateError{name, errors.BadRequestf(format, args...)})
}

func (v *validator) RequiredString(val string, name string, err ...error) {
	if len(strings.TrimSpace(val)) == 0 {
		defaultErr := errors.BadRequestf("%s is required", name)
		v.add(name, defaultErr, err)
	}
}

func (v *validator) RequiredBytes(val []byte, name string, err ...error) {
	if len(val) == 0 {
		defaultErr := errors.BadRequestf("%s is required", name)
		v.add(name, defaultErr, err)
	}
}

func (v *validator) RequiredInt(val int, name string, err ...error) {
	if val == 0 {
		defaultErr := errors.BadRequestf("%s is required", name)
		v.add(name, defaultErr, err)
	}
}

func (v *validator) RequiredInt64(val int64, name string, err ...error) {
	if val == 0 {
		defaultErr := errors.BadRequestf("%s is required", name)
		v.add(name, defaultErr, err)
	}
}

func (v *validator) RequiredFloat64(val float64, name string, err ...error) {
	if val == 0 {
		defaultErr := errors.BadRequestf("%s is required", name)
		v.add(name, defaultErr, err)
	}
}

func (v *validator) RequiredBool(val bool, name string, err ...error) {
	if !val {
		defaultErr := errors.BadRequestf("%s is required", name)
		v.add(name, defaultErr, err)
	}
}

func (v *validator) RequiredEmail(val string, name string, err ...error) {
	if val == "" {
		defaultErr := errors.BadRequestf("%s is required", name)
		v.add(name, defaultErr, err)
	}

	v.Email(val, name, err...)
}

func (v *validator) NotNil(val interface{}, name string, err ...error) {
	if val == nil {
		defaultErr := errors.BadRequestf("%s is required", name)
		v.add(name, defaultErr, err)
	}
}

func (v *validator) RequiredTime(val time.Time, name string, err ...error) {
	if val.IsZero() {
		defaultErr := errors.BadRequestf("%s is required", name)
		v.add(name, defaultErr, err)
	}
}

func (v *validator) MinInt(val int, n int, name string, err ...error) {
	if val > n {
		return
	}

	defaultErr := errors.BadRequestf("%s should be atleast %d", name, n)
	v.add(name, defaultErr, err)
}

func (v *validator) MaxInt(val int, n int, name string, err ...error) {
	if val > n {
		defaultErr := errors.BadRequestf("%s should not greater than %d", name, n)
		v.add(name, defaultErr, err)
	}
}

func (v *validator) MinInt64(val int64, n int64, name string, err ...error) {
	if val > n {
		return
	}

	defaultErr := errors.BadRequestf("%s should be atleast %d", name, n)
	v.add(name, defaultErr, err)
}

func (v *validator) MaxInt64(val int64, n int64, name string, err ...error) {
	if val > n {
		defaultErr := errors.BadRequestf("%s should not greater than %d", name, n)
		v.add(name, defaultErr, err)
	}
}

func (v *validator) MinFloat64(val float64, n float64, name string, err ...error) {
	if val < n {
		return
	}

	defaultErr := errors.BadRequestf("%s should be atleast %d", name, n)
	v.add(name, defaultErr, err)
}

func (v *validator) MaxFloat64(val float64, n float64, name string, err ...error) {
	if val > n {
		return
	}
	defaultErr := errors.BadRequestf("%s should not greater than %d", name, n)
	v.add(name, defaultErr, err)
}

func (v *validator) MinChar(val string, n int, name string, err ...error) {
	if utf8.RuneCountInString(val) < n {
		defaultErr := errors.BadRequestf("%s should be atleast %d character", name, n)
		v.add(name, defaultErr, err)
	}
}

func (v *validator) MaxChar(val string, n int, name string, err ...error) {
	if utf8.RuneCountInString(val) > n {
		defaultErr := errors.BadRequestf("%s should not greater than %d character", name, n)
		v.add(name, defaultErr, err)
	}
}

func (v *validator) Email(val string, name string, err ...error) {
	if val == "" {
		return
	}
	if !emailPatern.MatchString(val) {
		defaultErr := errors.BadRequestf("%s invalid email format", name)
		v.add(name, defaultErr, err)
	}
}

func (v *validator) Gender(val string, name string, err ...error) {
	if val != `male` && val != `female` {
		defaultErr := errors.BadRequestf("%s should be male or female", name)
		v.add(name, defaultErr, err)
	}
}

func (v *validator) Confirm(val, confirm string, name string, confirmName string, err ...error) {
	if val != confirm {
		defaultErr := errors.BadRequestf("%s is not matched %s", name, confirmName)
		v.add(name, defaultErr, err)
	}
}

func (v *validator) ISO8601DataTime(val string, name string, err ...error) {
	if val == "" {
		return
	}
	if !dateiso8601Patern.MatchString(val) {
		defaultErr := errors.BadRequestf("%s is invalid date format", name)
		v.add(name, defaultErr, err)
	}
}

func (v *validator) InString(val string, in []string, name string, err ...error) {
	for _, k := range in {
		if k == val {
			return
		}
	}

	defaultErr := errors.BadRequestf("%s is not in", strings.Join(in, ","))
	v.add(name, defaultErr, err)
}
