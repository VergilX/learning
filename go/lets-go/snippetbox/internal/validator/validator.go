package validator

import (
    "slices"
    "strings"
    "unicode/utf8"
)

// contain map of field errors
type Validator struct {
    FieldErrors     map[string]string
}

func (v *Validator) Valid() bool {
    return len(v.FieldErrors) == 0
}

// add field error if it doesn't exist already (key)
func (v *Validator) AddFieldError(key, message string) {
    if v.FieldErrors ==  nil {
        v.FieldErrors = make(map[string]string)
    }

    if _, exists := v.FieldErrors[key]; !exists {
        v.FieldErrors[key] = message
    }
}

// adds error field only if !ok
func (v *Validator) CheckField(ok bool, key, message string) {
    if !ok {
        v.AddFieldError(key, message)
    }
}

func NotBlank(value string) bool {
    return strings.TrimSpace(value) != ""
}

// returns true of no_of_chars < n
func MaxChars(value string, n int) bool {
    return utf8.RuneCountInString(value) <= n
}

// returns true if value in list of permitted values
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
    return slices.Contains(permittedValues, value)
}
func
