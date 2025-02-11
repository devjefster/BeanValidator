# GoValidator 🚀
A powerful and lightweight **struct validation library** for Go, inspired by Java's Bean Validation.

---

## **📌 Features**
✅ **Simple and powerful validation rules**  
✅ **Supports string, number, boolean, date, and collection validation**  
✅ **Custom error messages**  
✅ **Easy-to-use struct-based validation**

---

## **📌 Installation**
To use **GoValidator**, install it as a Go module:
```sh
  go get github.com/devjefster/GoValidator@latest
```

## 📌 Usage Example

Create a struct with validation tags and run the validator.
```go
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

````
## 📌 Supported Validation Rules

| Rule                   | Description                                            | Example                               |
|------------------------|--------------------------------------------------------|---------------------------------------|
| required               | Ensures the field is not empty or nil.                 | validate:"required"                   |
| email                  | Ensures the field contains a valid email.              | validate:"email"                      |
| isTrue                 | Ensures the field is true.                             | validate:"isTrue"                     |
| positive               | Ensures the field is greater than 0.                   | validate:"positive"                   |
| negative               | Ensures the field is less than 0.                      | validate:"negative"                   |
| positiveOrZero         | Ensures the field is >= 0.                             | validate:"positiveOrZero"             |
| negativeOrZero         | Ensures the field is <= 0.                             | validate:"negativeOrZero"             |
| size=n                 | Ensures the collection has exactly n elements.         | validate:"size=2"                     |
| minSize=n              | Ensures the collection has at least n elements.        | validate:"minSize=2"                  |
| maxSize=n              | Ensures the collection has at most n elements.         | validate:"maxSize=5"                  |
| date=format            | Ensures the field is a valid date in the given format. | validate:"date=2006-01-02"            |
| past=format            | Ensures the field is a past date.                      | validate:"past=2006-01-02"            |
| future=format          | Ensures the field is a future date.                    | validate:"future=2006-01-02"          |
| pastInclusive=format   | Ensures the field is a past or present date.           | validate:"pastInclusive=2006-01-02"   |
| futureInclusive=format | Ensures the field is a future or present date.         | validate:"futureInclusive=2006-01-02" |

## 📌 Running Tests

To run all unit tests:
````shell
  go test -v ./...
````
## 📌 Contributions

Feel free to fork, contribute, and open issues to improve this library! 🚀

---

### **✅ Summary**
✅ **Detailed documentation with examples**  
✅ **List of all supported validation rules**  
✅ **Installation and test instructions**  
✅ **Encourages contributions and improvements**

