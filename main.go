package main

import (
	"Validators/constants"
	"Validators/errors"
	"fmt"
)

type something struct {
	a string `json:"a",validate:"Required,Email",condition:"Length<=10"`
}

func main() {
	err := errors.New()
	for i := 0; i <= 10; i++ {
		err.Field("Name").Rule(constants.Req).Build()
	}

	// err := validate.Struct(stru)
	fmt.Println(err.GetError())
}
