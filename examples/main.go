package main

import (
	"bean-validator/validator"
	"fmt"
)

type User struct {
	Name    string   `validate:"required,non-blank"`
	Age     int      `validate:"min=18,max=100"`
	Email   string   `validate:"email"`
	Hobbies []string `validate:"non-empty"`
}

func main() {
	user := User{
		Name:    "   ",
		Age:     16,
		Email:   "invalid-email",
		Hobbies: []string{},
	}

	errs := validator.Validate(user)
	if len(errs) > 0 {
		fmt.Println("Validation Errors:")
		for _, err := range errs {
			fmt.Printf("- %s\n", err.Message)
		}
	} else {
		fmt.Println("Validation Passed!")
	}
}
