package validator

//go:generate mockery -name Validator

import (
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
	GetError() error
	RequiredString(val string, err error, name ...string)
	RequiredBytes(val []byte, err error, name ...string)
	RequiredInt(val int, err error, name ...string)
	RequiredFloat64(val float64, err error, name ...string)
	RequiredBool(val bool, err error, name ...string)
	RequiredEmail(val string, err error, name ...string)
	NotNil(val interface{}, err error, name ...string)
	RequiredTime(val time.Time, err error, name ...string)
	MinInt(val int, n int, err error, name ...string)
	MaxInt(val int, n int, err error, name ...string)
	MinFloat64(val float64, n float64, err error, name ...string)
	MaxFloat64(val float64, n float64, err error, name ...string)
	MinChar(val string, n int, err error, name ...string)
	MaxChar(val string, n int, err error, name ...string)
	Email(val string, err error, name ...string)
	Gender(val string, err error, name ...string)
	Confirm(val, confirm string, err error, name ...string)
	ISO8601DataTime(val string, err error, name ...string)
	Length(val int, atleast int, err error, name ...string)
	InString(val string, in []string, err error, name ...string)
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

func (v *validator) GetError() error {
	if len(v.errs) > 0 {
		return v.errs[0].Err
	}

	return nil
}

func (v *validator) addError(err error, name []string) {
	n := ""
	if len(name) > 0 {
		n = name[0]
	}

	v.errs = append(v.errs, ValidateError{n, err})
}

func (v *validator) RequiredString(val string, err error, name ...string) {
	if len(strings.TrimSpace(val)) == 0 {
		v.addError(err, name)
	}
}

func (v *validator) RequiredBytes(val []byte, err error, name ...string) {
	if len(val) == 0 {
		v.addError(err, name)
	}
}

func (v *validator) RequiredInt(val int, err error, name ...string) {
	if val == 0 {
		v.addError(err, name)
	}
}

func (v *validator) RequiredFloat64(val float64, err error, name ...string) {
	if val == 0 {
		v.addError(err, name)
	}
}

func (v *validator) RequiredBool(val bool, err error, name ...string) {
	if !val {
		v.addError(err, name)
	}
}

func (v *validator) RequiredEmail(val string, err error, name ...string) {
	if val == "" {
		v.addError(err, name)
	}

	v.Email(val, err, name...)
}

func (v *validator) NotNil(val interface{}, err error, name ...string) {
	if val == nil {
		v.addError(err, name)
	}
}

func (v *validator) RequiredTime(val time.Time, err error, name ...string) {
	if val.IsZero() {
		v.addError(err, name)
	}
}

func (v *validator) MinInt(val int, n int, err error, name ...string) {
	if val > n {
		return
	}

	v.addError(err, name)
}

func (v *validator) MaxInt(val int, n int, err error, name ...string) {
	if val > n {
		v.addError(err, name)
	}
}

func (v *validator) MinFloat64(val float64, n float64, err error, name ...string) {
	if val < n {
		return
	}

	v.addError(err, name)
}

func (v *validator) MaxFloat64(val float64, n float64, err error, name ...string) {
	if val > n {
		return
	}
	v.addError(err, name)
}

func (v *validator) MinChar(val string, n int, err error, name ...string) {
	if utf8.RuneCountInString(val) < n {
		v.addError(err, name)
	}
}

func (v *validator) MaxChar(val string, n int, err error, name ...string) {
	if utf8.RuneCountInString(val) > n {
		v.addError(err, name)
	}
}

func (v *validator) Email(val string, err error, name ...string) {
	if val == "" {
		return
	}
	if !emailPatern.MatchString(val) {
		v.addError(err, name)
	}
}

func (v *validator) Gender(val string, err error, name ...string) {
	if val != `male` && val != `female` {
		v.addError(err, name)
	}
}

func (v *validator) Confirm(val, confirm string, err error, name ...string) {
	if val != confirm {
		v.addError(err, name)
	}
}

func (v *validator) ISO8601DataTime(val string, err error, name ...string) {
	if val == "" {
		return
	}
	if !dateiso8601Patern.MatchString(val) {
		v.addError(err, name)
	}
}

func (v *validator) Length(val int, atleast int, err error, name ...string) {
	if val <= atleast {
		v.addError(err, name)
	}
}

func (v *validator) InString(val string, in []string, err error, name ...string) {
	for _, k := range in {
		if k == val {
			return
		}
	}

	v.addError(err, name)
}
