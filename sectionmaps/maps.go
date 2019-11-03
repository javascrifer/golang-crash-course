package sectionmaps

import "fmt"

// Example maps usage.
func Example() {
	// Key = string, value = int.
	ageMap := map[string]int{
		"Foo": 32,
		"Bar": 22,
	}

	// Adding values to a map.
	ageMap["Foobar"] = 40
	fmt.Println("Foobar is added to the map")

	// Accessing values from a map return value and boolean
	// field which identifies if a value is present in a map.
	// Usually, you should name it as ok. If value is not
	// present in the map return value is default value for
	// that type. In this case it would be 0 (default int).
	if _, ok := ageMap["Bar"]; ok {
		delete(ageMap, "Bar")
		fmt.Println("Bar is deleted from a map")
	}

	// Looping through map with range k, v -> name, age.
	for name, age := range ageMap {
		fmt.Println(name, age)
	}
}
