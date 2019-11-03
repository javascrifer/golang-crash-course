package sectionfunctions

import "fmt"

// Human is interface which defines human behaviour.
type Human interface {
	speak()
}

// Person is a demo struct.
type Person struct {
	firstName string
	lastName  string
}

// Speak is a method of a Person struct.
// It receives values from a struct via receiver.
func (person Person) speak() {
	fmt.Printf("Hi, I'm %v %v! \n", person.firstName, person.lastName)
}

type NanoRobot struct {
	name string
}

func (nanoRobot NanoRobot) speak() {
	fmt.Printf("Hi, I'm nano robot %v!\n", nanoRobot.name)
}

// Using Polymorphism we can pass both Person and NanoRobot structs
// to use implemented speak method of the Human interface.
// Also, you can assert type of struct and perform actions depended on the type.
func tryToSpeak(human Human) {
	switch human.(type) {
	case Person:
		fmt.Println("Polymorphism fount that it is human is person")
		fmt.Println(human.(Person).firstName, human.(Person).lastName)
	case NanoRobot:
		fmt.Println("Polymorphism fount that it is human is nano robot")
		fmt.Println(human.(NanoRobot).name)
	}

	human.speak()
}

// Immediately invoked function expression example.
func iife() {
	func() {
		fmt.Println("IIFE called")
	}()
}

// Example of callback as a parameter.
func evenSum(callback func(sum int), numbers ...int) {
	var sum = 0

	for _, number := range numbers {
		if (number % 2) == 0 {
			sum = sum + number
		}
	}

	callback(sum)
}

// Counter closure example.
func Counter(initialValue int) func() int {
	value := initialValue

	f := func() int {
		value++
		return value
	}

	return f
}

// Factorial recursive function example.
func factorial(value int) int {
	if value == 0 {
		return 1
	}

	return value * factorial(value-1)
}

// ExampleMethods of methods usage.
func ExampleMethods() {
	// Polymorphism in action.
	person := Person{
		firstName: "Foo",
		lastName:  "Bar",
	}

	nanoRobot := NanoRobot{name: "Robodo"}

	tryToSpeak(person)
	tryToSpeak(nanoRobot)

	// IIFE in action.
	iife()

	// Closure example.
	counter := Counter(0)
	count := counter()
	fmt.Println(count)

	counter()
	counter()
	count = counter()
	fmt.Println(count)

	// Callback example.
	evenSumHandler := func(sum int) {
		fmt.Printf("Even number sum is %v\n", sum)
	}
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	evenSum(evenSumHandler, numbers...)

	// Recursion example.
	fmt.Printf("Factorial of %v is %v\n", 5, factorial(5))
}
