package datatype

import "errors"

type DataType interface {
	Validate(interface{}) error
}

type IntDataType struct {
	MinValue int
	MaxValue int
}

// create new IntDataType object
func Int(min int, max int) DataType {
	return IntDataType{
		MinValue: min,
		MaxValue: max,
	}
}

func (t IntDataType) Validate(val interface{}) error {
	num, ok := val.(int)
	if !ok {
		return errors.New("invalid column data type")
	}
	if num < t.MinValue || num > t.MaxValue {
		return errors.New("column INT data type validation failed")
	}
	return nil
}

type StringDataType struct {
	MaxLength int
}

// create new StringDataType object
func String(maxLen int) DataType {
	return StringDataType{
		MaxLength: maxLen,
	}
}

func (t StringDataType) Validate(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return errors.New("invalid column data type")
	}
	if len(str) > t.MaxLength {
		return errors.New("column String data type validation failed")
	}
	return nil
}
