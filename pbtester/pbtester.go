package pbtester

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func Test() {
	fname := ".\\pbtester\\osman.dat"

	book := &AddressBook{}

	// proto2
	// p := &Person{
	// 	Id:    proto.Int32(1234),
	// 	Name:  proto.String("John Doe"),
	// 	Email: proto.String("jdoe@example.com"),
	// 	Phones: []*Person_PhoneNumber{
	// 		{Number: proto.String("555-4321"), Type: &[]Person_PhoneType{Person_HOME}[0]},
	// 	},
	// }

	// proto3
	p := &Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*Person_PhoneNumber{
			{Number: "555-4321", Type: Person_HOME},
		},
	}

	book.People = append(book.People, p)

	out, err := proto.Marshal(book)

	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	fmt.Println("written to " + fname)

	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book = &AddressBook{}

	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Println("read from " + fname)
	fmt.Println(*book)
}
