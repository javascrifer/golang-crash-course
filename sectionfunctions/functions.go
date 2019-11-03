package sectionfunctions

import "fmt"

// Void func.
func greet() {
	fmt.Println("Hello!")
}

// Func with parameter and return value.
func getGreeting(name string) string {
	return fmt.Sprint("Hello, ", name, "!")
}

// Func with parameter and multiple values return.
// Note: argument passed to func is passed by value.
func createPerson(name string) (int, bool) {
	if name == "Foo" {
		return 0, false
	}

	return len(name), true
}

// Func with variadic parameters.
func printSlice(slice ...int) {
	fmt.Printf("You passed %v arguments to the function. \n", len(slice))
	fmt.Printf("Type of slice %T. \n", slice)
	fmt.Printf("Value %v. \n", slice)
}

// ExampleFunctions of functions usage.
func ExampleFunctions() {
	// Usage of void function.
	greet()

	// Usage of func with argument.
	greeting := getGreeting("Foo")
	fmt.Println(greeting)

	// Usage of func with argument
	// and multiple values returned.
	id1, isCreated1 := createPerson("Foo")
	id2, isCreated2 := createPerson("Bar")

	fmt.Println(id1, isCreated1)
	fmt.Println(id2, isCreated2)

	// Usage of func with variadic parameters.
	printSlice(1, 2, 3)
}
