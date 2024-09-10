package constraint

import "errors"

type Constraint interface {
	Validate(interface{}) error
}

type NotEmpty struct {
}

func NotEmptyConstraint() Constraint {
	return NotEmpty{}
}

func (n NotEmpty) Validate(val interface{}) error {
	str := val.(string)
	if str == "" {
		return errors.New("empty string is not valid")
	}

	return nil
}
