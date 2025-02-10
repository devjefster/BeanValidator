package main

import (
	"fmt"
	"github.com/devjefster/GoValidator/validator"
)

type User struct {
	Username      string   `validate:"required,minSize=3,maxSize=15"`
	Email         string   `validate:"required,email"`
	Age           int      `validate:"required,positive"`
	Balance       float64  `validate:"positiveOrZero"`
	IsActive      bool     `validate:"isTrue"`
	Tags          []string `validate:"minSize=1,maxSize=5"`
	Birthdate     string   `validate:"date=2006-01-02,pastInclusive=2006-01-02"`
	Subscription  string   `validate:"date=2006-01-02,future=2006-01-02"`
	PhoneNumbers  []string `validate:"size=2"`
	Comment       string   `validate:"maxSize=200"`
	FavoriteItems []int    `validate:"minSize=2"`
}

func main() {
	// Valid user data
	user := User{
		Username:      "johndoe",
		Email:         "johndoe@example.com",
		Age:           25,
		Balance:       100.50,
		IsActive:      true,
		Tags:          []string{"Go", "Golang"},
		Birthdate:     "1995-06-15",
		Subscription:  "2025-01-01",
		PhoneNumbers:  []string{"123-456-7890", "987-654-3210"},
		Comment:       "This is a sample comment.",
		FavoriteItems: []int{1, 2},
	}

	// Validate user struct
	errors := validator.Validate(user)

	if errors.HasErrors() {
		fmt.Println("Validation failed:")
		for _, err := range errors {
			fmt.Println("-", err)
		}
	} else {
		fmt.Println("Validation passed ✅")
	}
}
