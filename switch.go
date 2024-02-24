package main

import "fmt"

func main() {
	name := "Andre"

	switch name {
	case "Jordan":
		fmt.Println("Hello Jordan")
	case "Kate":
		fmt.Println("Hello Kate")
	case "John":
		fmt.Println("Hello John")
	default:
		fmt.Println("I don't know you")
	}

	number := 10

	switch {
	case number > 1:
		fmt.Println("Number is grater than 1")
		fallthrough
	case number < 11:
		fmt.Println("Number is less than 11")
	}
}
