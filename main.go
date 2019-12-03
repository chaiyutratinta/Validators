package main

import (
	"Validators/errors"
	"fmt"
)

type something struct {
	a string `json:"a",validate:"Required,Email",condition:"Length<=10"`
}

func main() {
	err := errors.New()
	for i := 0; i <= 10; i++ {
		err.Field("Field").ErrMsg("Fuck").Rule("len").Set()
	}

	// err := validate.Struct(stru)
	fmt.Println(err.GetError())
}
