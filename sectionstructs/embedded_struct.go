package sectionstructs

import "fmt"

// Person generic struct.
type Person struct {
	firstName string
	lastName  string
}

// SecretAgent embedded struct with Person inside of it.
type SecretAgent struct {
	Person
	tools []string
}

// Example embedded struct usage.
func Example() {
	secretAgent := SecretAgent{
		Person: Person{
			firstName: "Foo",
			lastName:  "Bar",
		},
		tools: []string{"goggles", "microphone"},
	}

	// Struct is kind a same thing as a TypeScript interface.
	// It allows you to design a blueprint for an entity.
	// In this case we are using anonymous t
	fmt.Printf("Hi I'm %v %v!\n", secretAgent.firstName, secretAgent.lastName)
	fmt.Println("My main tools are:")

	for _, tool := range secretAgent.tools {
		fmt.Printf("\t%v\n", tool)
	}
}
