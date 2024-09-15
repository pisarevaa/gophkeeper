package model

import (
	"strings"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=ContextKeyEnum -linecomment -output contextkey_enum_string.go
type ContextKeyEnum int

const (
	ContextKeyUnknown ContextKeyEnum = iota // unknown
	ContextUserID     ContextKeyEnum = iota // text
)

func (em *ContextKeyEnum) SetValue(value string) error {
	contextKey := ContextKeyFromString(value)
	if contextKey == ContextKeyUnknown {
		return &ContextKeyInvalidValueError{
			Value: value,
		}
	}

	*em = contextKey

	return nil
}

func (em *ContextKeyEnum) UnmarshalText(text []byte) error {
	return em.SetValue(string(text))
}

func (em ContextKeyEnum) MarshalText() ([]byte, error) {
	if em == ContextKeyUnknown {
		return nil, &ContextKeyInvalidValueError{
			Value: ContextKeyUnknown.String(),
		}
	}

	return []byte(em.String()), nil
}

func ContextKeyFromString(value string) ContextKeyEnum {
	switch strings.ToLower(value) {
	case "userID":
		return ContextUserID
	default:
		return ContextKeyUnknown
	}
}

type ContextKeyInvalidValueError struct {
	Value string
}

func (e *ContextKeyInvalidValueError) Error() string {
	return "context key invalid value: " + e.Value
}
