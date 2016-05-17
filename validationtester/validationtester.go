package validationtester

import (
	"fmt"
	"gopkg.in/validator.v2"
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
