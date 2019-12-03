package validators

import (
	"custom_validator/constants"
	"custom_validator/validators/rules"
	"fmt"
	"reflect"
	"strings"
)

func Validate(stru interface{}) error {
	rVal := reflect.ValueOf(stru).Elem()
	rType := rVal.Type()

	for i := 0; i < rVal.NumField(); i++ {
		fieldType := rType.Field(i)
		fieldValue := rVal.Field(i)
		fieldTag := fieldType.Tag

		fieldRules := strings.Split(fieldTag.Get(constants.Rules), ",")

		rules.Process(fieldValue, fieldRules...)

		fmt.Println(fieldType.Name, fieldType.Type, fieldValue.String(), fieldType.Tag)
	}

	return nil
}
