package jsontester

import (
	"encoding/json"
	"fmt"
)

func Test() {
	type datatype struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	d := datatype{Id: 1, Name: "osman"}

	j, _ := json.Marshal(d)

	fmt.Println(string(j))
}
