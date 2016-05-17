package validationtester

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"gopkg.in/validator.v2"
)

func Test() {
	test1()
	test2()
}

func test1() {
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

func test2() {
	type Echo struct {
		Echo string `json:"echo" valid:"stringlength(10|255),required"`
	}

	request := Echo{}
	request.Echo = "invalid"

	errors, err := govalidator.ValidateStruct(request)

	fmt.Println(errors)
	fmt.Println(err)
}
