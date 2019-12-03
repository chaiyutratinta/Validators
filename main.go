package main

import "fmt"

type something struct {
	a string `json:"a",validate:"Required,Email",condition:"Length<=10"`
}

func main() {
	stru := something{}

	// validate := validator.Validate()
	// err := validate.Struct(stru)
	fmt.Println("RUN MAIN")
}
