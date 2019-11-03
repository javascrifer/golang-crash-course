package sectionpointers

import "fmt"

type Person struct {
	name string
}

// Mutate value of a primitive by reference.
func changeByReference(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}

// Mutate value of a struct.
func changeName(person *Person, newName string) {
	person.name = newName
	// (*person).name = newName - Same as line above.
}

// Example usage of pointers.
func Example() {
	a := 42

	// & - holds a reference to a memory slot.
	// * - used to dereference memory slot to a value.
	fmt.Println(a)
	fmt.Printf("%v - %T, %v - %T \n", a, a, &a, &a)
	fmt.Printf("%v - %T, %v - %T \n", &a, &a, *&a, *&a)

	// Mutating value by reference for primitive.
	x := 42

	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	changeByReference(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)

	// Mutating value by reference for struct.
	person := Person{name: "James"}
	fmt.Println(person)
	changeName(&person, "John")
	fmt.Println(person)
}
