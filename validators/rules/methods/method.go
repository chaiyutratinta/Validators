package methods

import "reflect"

type (
	MethodFunc func(reflect.Value) error

	Methods interface {
		Prepare(reflect.Value) error
	}
	Required struct{}
)

func (Required) Prepare(val reflect.Value) error {
	if val.IsZero() {
		return nil
	}
	return nil
}
