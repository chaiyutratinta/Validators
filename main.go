package main

import (
	"custom_validator/validators"
	"fmt"
)

type something struct {
	A string `json:"a" rules:"required" condition:"Length<=10" validate:"required"`
	B string `json:"b" rules:"required"`
	C int    `json:"c" validate:"required"`
}

type something2 struct {
	productName string
	amount      int
	price       int
}

func main() {
	stru := &something{
		A: "a",
		B: "B",
		C: 0,
	}
	// stru2 := something2{}

	validators.Validate(stru)

	// validators.Validate(stru2)

	// validate := validator.New()
	// err := validate.Struct(stru)
	// fmt.Println(err)
	fmt.Println("RUN MAIN")
}
