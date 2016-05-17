package validationtester

import (
	"fmt"
	"github.com/go-validator/validator"
)

func Test() {
	type Data struct {
		Name string `validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
		Age  int    `validate:"min=18,max=150"`
	}

	d := &Data{
		Name: "osman",
		Age:  250,
	}

	errs := validator.Validate(d)

	fmt.Println(errs)
}
