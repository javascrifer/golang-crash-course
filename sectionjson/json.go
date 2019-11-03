package sectionjson

import (
	"encoding/json"
	"fmt"
)

// Something to keep in mind:
// 1. In Go public variables starts with capital case.
// 2. In Go private variables starts with lower case.
// 3. Private variables will not be included to the JSON.
// 4. Private variables can not be accessed outside the package.

// Human is demo struct with 3 public and 1 private variable
type Human struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	gender    string
}

// Encoding struct to a JSON.
func encodeJson() {
	fmt.Println("Encoding")

	human := Human{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}

	b, err := json.Marshal(human)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Bytes:", b)
	fmt.Println("String:", string(b))
}

func decodeJson() {
	var human Human

	fmt.Println("Decoding")

	s := `{"first_name":"John","last_name":"Doe","age":25}`
	bs := []byte(s)
	err := json.Unmarshal(bs, &human)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(human)
}

func Example() {
	encodeJson()
	decodeJson()
}
