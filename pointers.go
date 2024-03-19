package main

import "fmt"

func main() {
	age := 32 // regular variable

	agePointer := &age

	fmt.Println("Age:", *agePointer)

	editAgeToGetAdultYears(agePointer)
	fmt.Println(age)
}

func editAgeToGetAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}