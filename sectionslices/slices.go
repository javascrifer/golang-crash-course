package sectionslices

import "fmt"

// Example usage of slices.
func Example() {
	slice := []int{0, 1, 2}

	// To create arrays with dynamic length
	// (aka slices) Go uses following concept:
	// 1. Allocates memory for X items.
	// 2. If limit of items is reached allocates
	// memory for X + Y items and transfers items
	// to the new memory slot.
	fmt.Println(slice)
	fmt.Printf("Capasity: %v\n", cap(slice))
	fmt.Printf("Length: %v\n\n", len(slice))

	// In this case we already reached our items limit.
	// So Go runtime will allocate space for Y more items.
	// Number of additional memory allocation will be figured
	// out on the go. In this case it will be 3.
	slice = append(slice, 3)

	fmt.Println(slice)
	fmt.Printf("Capasity: %v\n", cap(slice))
	fmt.Printf("Length: %v\n\n", len(slice))

	// Looping through slice.
	fmt.Println("Looping through slice:")
	for k, v := range slice {
		fmt.Println(k, v)
	}
	fmt.Println()

	// Slicing a slice. Returns a new slice (fresh reference).
	// It can be used to delete a certain element or get certain
	// elements in rage.

	// Taking values from X (included) to Y (not included) index.
	fmt.Println(slice[1:4])

	// Taking values from X (included) to the last index.
	fmt.Println(slice[2:])

	// Taking values from first to the Y (not included) index.
	fmt.Println(slice[:2])
}
