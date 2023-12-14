package validator

import (
	"strconv"
	"strings"
)

type Validator struct {
	FieldErrors map[string]string
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}
	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

func (v *Validator) NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func (v *Validator) CorrectLength(value string) bool {
	return len(strings.TrimSpace(value)) == 16
}

func (v *Validator) IsNumeric(value string) bool {
	_, err := strconv.ParseInt(value, 10, 0)
	return err == nil
}
