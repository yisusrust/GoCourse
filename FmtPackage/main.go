package main

import "fmt"

func main() {
	// Print adding end of line at the end
	fmt.Println("Hello")
	fmt.Println("World")

	// Print including vars parsed
	name := "Lore"
	age := 28
	fmt.Printf("%s has %d yearls old\n", name, age)
	fmt.Printf("%v has %v yearls old\n", name, age)

	// Sprintf save a string into a var
	message := fmt.Sprintf("%v has %v yearls old", name, age)
	fmt.Println(message)

	//look data time
	fmt.Printf("Name: %T", name)
}
