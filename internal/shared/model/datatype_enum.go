package model

import (
	"strings"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=DataTypeEnum -linecomment -output datatype_enum_string.go
type DataTypeEnum int

const (
	TypeUnknown DataTypeEnum = iota // unknown
	TextType    DataTypeEnum = iota // text
	BinaryType  DataTypeEnum = iota // binary
)

func (em *DataTypeEnum) SetValue(value string) error {
	dataType := MethodFromString(value)
	if dataType == TypeUnknown {
		return &DataTypeInvalidValueError{
			Value: value,
		}
	}

	*em = dataType

	return nil
}

func (em *DataTypeEnum) UnmarshalText(text []byte) error {
	return em.SetValue(string(text))
}

func (em DataTypeEnum) MarshalText() ([]byte, error) {
	if em == TypeUnknown {
		return nil, &DataTypeInvalidValueError{
			Value: TypeUnknown.String(),
		}
	}

	return []byte(em.String()), nil
}

func MethodFromString(value string) DataTypeEnum {
	switch strings.ToLower(value) {
	case "text":
		return TextType
	case "binary":
		return BinaryType
	default:
		return TypeUnknown
	}
}

type DataTypeInvalidValueError struct {
	Value string
}

func (e *DataTypeInvalidValueError) Error() string {
	return "data type invalid value: " + e.Value
}
