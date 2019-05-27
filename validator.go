package validator

//go:generate mockery -name Validator -output mock_validator

import (
    "regexp"
    "strings"
    "time"
    "unicode/utf8"

    "github.com/onedaycat/errors"

    "fmt"
)

var (
    emptyStr          = ""
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
    GetMsg() string
    Wrap(errors.Error) errors.Error
    AddError(name string, err errors.Error)
    AddErrorMsg(name, format string, args ...interface{})
    NotNil(val interface{}, name string, err ...errors.Error)
    Email(val string, name string, err ...errors.Error)
    Gender(val string, name string, err ...errors.Error)
    Confirm(val, confirm string, name string, confirmName string, err ...errors.Error)
    ISO8601DataTime(val string, name string, err ...errors.Error)
    InString(val string, in []string, name string, err ...errors.Error)
    RequiredString(val string, name string, err ...errors.Error)
    RequiredBytes(val []byte, name string, err ...errors.Error)
    RequiredInt(val int, name string, err ...errors.Error)
    RequiredInt32(val int32, name string, err ...errors.Error)
    RequiredInt64(val int64, name string, err ...errors.Error)
    RequiredFloat32(val float32, name string, err ...errors.Error)
    RequiredFloat64(val float64, name string, err ...errors.Error)
    RequiredBool(val bool, name string, err ...errors.Error)
    RequiredEmail(val string, name string, err ...errors.Error)
    RequiredTime(val time.Time, name string, err ...errors.Error)
    RequiredArrayString(val []string, name string, err ...errors.Error)
    MinChar(val string, n int, name string, err ...errors.Error)
    MinInt(val int, n int, name string, err ...errors.Error)
    MinInt32(val int32, n int32, name string, err ...errors.Error)
    MinInt64(val int64, n int64, name string, err ...errors.Error)
    MinFloat32(val float32, n float32, name string, err ...errors.Error)
    MinFloat64(val float64, n float64, name string, err ...errors.Error)
    MaxChar(val string, n int, name string, err ...errors.Error)
    MaxInt(val int, n int, name string, err ...errors.Error)
    MaxInt32(val int32, n int32, name string, err ...errors.Error)
    MaxInt64(val int64, n int64, name string, err ...errors.Error)
    MaxFloat32(val float32, n float32, name string, err ...errors.Error)
    MaxFloat64(val float64, n float64, name string, err ...errors.Error)
    RangeInt(val, min, max int, name string, err ...errors.Error)
    RangeInt32(val, min, max int32, name string, err ...errors.Error)
    RangeInt64(val, min, max int64, name string, err ...errors.Error)
    RangeFloat32(val, min, max float32, name string, err ...errors.Error)
    RangeFloat64(val, min, max float64, name string, err ...errors.Error)
    LenArrayString(val []string, nLen int, name string, err ...errors.Error)
    LenArrayInt(val []int, nLen int, name string, err ...errors.Error)
    LenArrayInt32(val []int32, nLen int, name string, err ...errors.Error)
    LenArrayInt64(val []int64, nLen int, name string, err ...errors.Error)
    LenArrayFloat32(val []float32, nLen int, name string, err ...errors.Error)
    LenArrayFloat64(val []float64, nLen int, name string, err ...errors.Error)
    PointerRequiredEmail(val *string, name string, err ...errors.Error)
    PointerRequiredString(val *string, name string, err ...errors.Error)
    PointerRequiredInt(val *int, name string, err ...errors.Error)
    PointerRequiredInt32(val *int32, name string, err ...errors.Error)
    PointerRequiredInt64(val *int64, name string, err ...errors.Error)
    PointerRequiredFloat32(val *float32, name string, err ...errors.Error)
    PointerRequiredFloat64(val *float64, name string, err ...errors.Error)
    PointerRequiredBool(val *bool, name string, err ...errors.Error)
    PointerRangeInt(val *int, min, max int, name string, err ...errors.Error)
    PointerRangeInt32(val *int32, min, max int32, name string, err ...errors.Error)
    PointerRangeInt64(val *int64, min, max int64, name string, err ...errors.Error)
    PointerRangeFloat32(val *float32, min, max float32, name string, err ...errors.Error)
    PointerRangeFloat64(val *float64, min, max float64, name string, err ...errors.Error)
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
        return errors.BadRequest(emptyStr, v.errs[0].Err.Error())
    }

    return nil
}
func (v *validator) GetMsg() string {
    if len(v.errs) > 0 {
        return v.errs[0].Err.Error()
    }

    return emptyStr
}
func (v *validator) Wrap(err errors.Error) errors.Error {
    if len(v.errs) > 0 {
        return err.WithMessage(v.errs[0].Err.Error())
    }

    return nil
}
func (v *validator) add(name string, err error, errs []errors.Error) {
    if len(errs) > 0 {
        err = errs[0]
    }

    v.errs = append(v.errs, ValidateError{name, err})
}
func (v *validator) AddError(name string, err errors.Error) {
    v.errs = append(v.errs, ValidateError{name, err})
}
func (v *validator) AddErrorMsg(name, format string, args ...interface{}) {
    v.errs = append(v.errs, ValidateError{name, fmt.Errorf(format, args...)})
}

func (v *validator) NotNil(val interface{}, name string, err ...errors.Error) {
    if val == nil {
        defaultErr := fmt.Errorf("%s is required", name)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) Email(val string, name string, err ...errors.Error) {
    if val == "" {
        return
    }
    if !emailPatern.MatchString(val) {
        defaultErr := fmt.Errorf("%s invalid email format", name)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) Gender(val string, name string, err ...errors.Error) {
    if val != `male` && val != `female` {
        defaultErr := fmt.Errorf("%s should be male or female", name)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) Confirm(val, confirm string, name string, confirmName string, err ...errors.Error) {
    if val != confirm {
        defaultErr := fmt.Errorf("%s is not matched %s", name, confirmName)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) ISO8601DataTime(val string, name string, err ...errors.Error) {
    if val == "" {
        return
    }
    if !dateiso8601Patern.MatchString(val) {
        defaultErr := fmt.Errorf("%s is invalid date format", name)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) InString(val string, in []string, name string, err ...errors.Error) {
    for _, k := range in {
        if k == val {
            return
        }
    }

    defaultErr := fmt.Errorf("%s is not in", strings.Join(in, ","))
    v.add(name, defaultErr, err)
}

func (v *validator) RequiredString(val string, name string, err ...errors.Error) {
    if len(strings.TrimSpace(val)) == 0 {
        defaultErr := fmt.Errorf("%s is required", name)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) RequiredBytes(val []byte, name string, err ...errors.Error) {
    if len(val) == 0 {
        defaultErr := fmt.Errorf("%s is required", name)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) RequiredInt(val int, name string, err ...errors.Error) {
    if val == 0 {
        defaultErr := fmt.Errorf("%s is required", name)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) RequiredInt32(val int32, name string, err ...errors.Error) {
    if val == 0 {
        defaultErr := fmt.Errorf("%s is required", name)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) RequiredInt64(val int64, name string, err ...errors.Error) {
    if val == 0 {
        defaultErr := fmt.Errorf("%s is required", name)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) RequiredFloat32(val float32, name string, err ...errors.Error) {
    if val == 0 {
        defaultErr := fmt.Errorf("%s is required", name)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) RequiredFloat64(val float64, name string, err ...errors.Error) {
    if val == 0 {
        defaultErr := fmt.Errorf("%s is required", name)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) RequiredBool(val bool, name string, err ...errors.Error) {
    if !val {
        defaultErr := fmt.Errorf("%s is required", name)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) RequiredEmail(val string, name string, err ...errors.Error) {
    if val == "" {
        defaultErr := fmt.Errorf("%s is required", name)
        v.add(name, defaultErr, err)
    }

    v.Email(val, name, err...)
}
func (v *validator) RequiredTime(val time.Time, name string, err ...errors.Error) {
    if val.IsZero() {
        defaultErr := fmt.Errorf("%s is required", name)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) RequiredArrayString(val []string, name string, err ...errors.Error) {
    if len(val) == 0 {
        defaultErr := fmt.Errorf("%s is required", name)
        v.add(name, defaultErr, err)
    }
}

func (v *validator) MinChar(val string, n int, name string, err ...errors.Error) {
    if utf8.RuneCountInString(val) < n {
        defaultErr := fmt.Errorf("%s should be atleast %d character", name, n)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) MinInt(val int, n int, name string, err ...errors.Error) {
    if val > n {
        return
    }

    defaultErr := fmt.Errorf("%s should be atleast %d", name, n)
    v.add(name, defaultErr, err)
}
func (v *validator) MinInt32(val int32, n int32, name string, err ...errors.Error) {
    if val > n {
        return
    }

    defaultErr := fmt.Errorf("%s should be atleast %d", name, n)
    v.add(name, defaultErr, err)
}
func (v *validator) MinInt64(val int64, n int64, name string, err ...errors.Error) {
    if val > n {
        return
    }

    defaultErr := fmt.Errorf("%s should be atleast %d", name, n)
    v.add(name, defaultErr, err)
}
func (v *validator) MinFloat32(val float32, n float32, name string, err ...errors.Error) {
    if val < n {
        return
    }

    defaultErr := fmt.Errorf("%s should be atleast %v", name, n)
    v.add(name, defaultErr, err)
}
func (v *validator) MinFloat64(val float64, n float64, name string, err ...errors.Error) {
    if val < n {
        return
    }

    defaultErr := fmt.Errorf("%s should be atleast %v", name, n)
    v.add(name, defaultErr, err)
}

func (v *validator) MaxChar(val string, n int, name string, err ...errors.Error) {
    if utf8.RuneCountInString(val) > n {
        defaultErr := fmt.Errorf("%s should not greater than %d character", name, n)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) MaxInt(val int, n int, name string, err ...errors.Error) {
    if val > n {
        defaultErr := fmt.Errorf("%s should not greater than %d", name, n)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) MaxInt32(val int32, n int32, name string, err ...errors.Error) {
    if val > n {
        defaultErr := fmt.Errorf("%s should not greater than %d", name, n)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) MaxInt64(val int64, n int64, name string, err ...errors.Error) {
    if val > n {
        defaultErr := fmt.Errorf("%s should not greater than %d", name, n)
        v.add(name, defaultErr, err)
    }
}
func (v *validator) MaxFloat32(val float32, n float32, name string, err ...errors.Error) {
    if val > n {
        return
    }
    defaultErr := fmt.Errorf("%s should not greater than %v", name, n)
    v.add(name, defaultErr, err)
}
func (v *validator) MaxFloat64(val float64, n float64, name string, err ...errors.Error) {
    if val > n {
        return
    }
    defaultErr := fmt.Errorf("%s should not greater than %v", name, n)
    v.add(name, defaultErr, err)
}

func (v *validator) RangeInt(val, min, max int, name string, err ...errors.Error) {
    if val >= min && val <= max {
        return
    }

    defaultErr := fmt.Errorf("%s is out of range", name)
    v.add(name, defaultErr, err)
}
func (v *validator) RangeInt32(val, min, max int32, name string, err ...errors.Error) {
    if val >= min && val <= max {
        return
    }

    defaultErr := fmt.Errorf("%s is out of range", name)
    v.add(name, defaultErr, err)
}
func (v *validator) RangeInt64(val, min, max int64, name string, err ...errors.Error) {
    if val >= min && val <= max {
        return
    }

    defaultErr := fmt.Errorf("%s is out of range", name)
    v.add(name, defaultErr, err)
}
func (v *validator) RangeFloat32(val, min, max float32, name string, err ...errors.Error) {
    if val >= min && val <= max {
        return
    }

    defaultErr := fmt.Errorf("%s is out of range", name)
    v.add(name, defaultErr, err)
}
func (v *validator) RangeFloat64(val, min, max float64, name string, err ...errors.Error) {
    if val >= min && val <= max {
        return
    }

    defaultErr := fmt.Errorf("%s is out of range", name)
    v.add(name, defaultErr, err)
}

func (v *validator) LenArrayString(val []string, nLen int, name string, err ...errors.Error) {
    if len(val) == nLen {
        return
    }

    defaultErr := fmt.Errorf("%s must len equal to %d", name, nLen)
    v.add(name, defaultErr, err)
}
func (v *validator) LenArrayInt(val []int, nLen int, name string, err ...errors.Error) {
    if len(val) == nLen {
        return
    }

    defaultErr := fmt.Errorf("%s must len equal to %d", name, nLen)
    v.add(name, defaultErr, err)
}
func (v *validator) LenArrayInt32(val []int32, nLen int, name string, err ...errors.Error) {
    if len(val) == nLen {
        return
    }

    defaultErr := fmt.Errorf("%s must len equal to %d", name, nLen)
    v.add(name, defaultErr, err)
}
func (v *validator) LenArrayInt64(val []int64, nLen int, name string, err ...errors.Error) {
    if len(val) == nLen {
        return
    }

    defaultErr := fmt.Errorf("%s must len equal to %d", name, nLen)
    v.add(name, defaultErr, err)
}
func (v *validator) LenArrayFloat32(val []float32, nLen int, name string, err ...errors.Error) {
    if len(val) == nLen {
        return
    }

    defaultErr := fmt.Errorf("%s must len equal to %d", name, nLen)
    v.add(name, defaultErr, err)
}
func (v *validator) LenArrayFloat64(val []float64, nLen int, name string, err ...errors.Error) {
    if len(val) == nLen {
        return
    }

    defaultErr := fmt.Errorf("%s must len equal to %d", name, nLen)
    v.add(name, defaultErr, err)
}

func (v *validator) PointerRequiredString(val *string, name string, err ...errors.Error) {
    if val != nil {
        v.RequiredString(*val, name)
    }
}
func (v *validator) PointerRequiredEmail(val *string, name string, err ...errors.Error) {
    if val != nil {
        v.RequiredString(*val, name)
    }
}
func (v *validator) PointerRequiredInt(val *int, name string, err ...errors.Error) {
    if val != nil {
        v.RequiredInt(*val, name)
    }
}
func (v *validator) PointerRequiredInt32(val *int32, name string, err ...errors.Error) {
    if val != nil {
        v.RequiredInt32(*val, name)
    }
}
func (v *validator) PointerRequiredInt64(val *int64, name string, err ...errors.Error) {
    if val != nil {
        v.RequiredInt64(*val, name)
    }
}
func (v *validator) PointerRequiredFloat64(val *float64, name string, err ...errors.Error) {
    if val != nil {
        v.RequiredFloat64(*val, name)
    }
}
func (v *validator) PointerRequiredFloat32(val *float32, name string, err ...errors.Error) {
    if val != nil {
        v.RequiredFloat32(*val, name)
    }
}
func (v *validator) PointerRequiredBool(val *bool, name string, err ...errors.Error) {
    if val != nil {
        v.RequiredBool(*val, name)
    }
}
func (v *validator) PointerRangeInt(val *int, min, max int, name string, err ...errors.Error) {
    if val == nil || (*val >= min && *val <= max) {
        return
    }

    defaultErr := fmt.Errorf("%s is out of range", name)
    v.add(name, defaultErr, err)
}
func (v *validator) PointerRangeInt32(val *int32, min, max int32, name string, err ...errors.Error) {
    if val == nil || *val >= min && *val <= max {
        return
    }

    defaultErr := fmt.Errorf("%s is out of range", name)
    v.add(name, defaultErr, err)
}
func (v *validator) PointerRangeInt64(val *int64, min, max int64, name string, err ...errors.Error) {
    if val == nil || *val >= min && *val <= max {
        return
    }

    defaultErr := fmt.Errorf("%s is out of range", name)
    v.add(name, defaultErr, err)
}
func (v *validator) PointerRangeFloat32(val *float32, min, max float32, name string, err ...errors.Error) {
    if val == nil || *val >= min && *val <= max {
        return
    }

    defaultErr := fmt.Errorf("%s is out of range", name)
    v.add(name, defaultErr, err)
}
func (v *validator) PointerRangeFloat64(val *float64, min, max float64, name string, err ...errors.Error) {
    if val == nil || *val >= min && *val <= max {
        return
    }

    defaultErr := fmt.Errorf("%s is out of range", name)
    v.add(name, defaultErr, err)
}