package rules

import (
	"custom_validator/validators/rules/methods"
	"reflect"
	"strings"
)

type (
	Rule string

	Rules interface {
		Run() error
	}
)

var mapRules map[string]methods.Methods = map[string]methods.Methods{
	"required": methods.Required{},
}

func Process(val reflect.Value, rules ...string) (err []error) {
	for _, r := range rules {
		met, ok := mapRules[strings.ToLower(r)]
		if ok {
			err = append(err, met.Prepare(val))
		}
	}
	return err
}
