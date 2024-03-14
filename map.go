package main

import "fmt"

func main() {
	var money map[string]int = map[string]int{
		"dollar": 1000,
		"euro":   100,
		"ap":     5,
	}
	fmt.Println(money)
	fmt.Println(money["dollar"])

	money2 := map[string]int{
		"dollar": 1000,
		"euro":   100,
		"ap":     5,
	}
	fmt.Println(money2)

	money2["dollar"] = 5000
	fmt.Println(money2)

	delete(money2, "ap")
	fmt.Println(money2)

	fmt.Println(money2["gg"]) //default for int is 0

	val, status := money2["dollar"]
	fmt.Println(val, status)

	val2, status2 := money2["gg"]
	fmt.Println(val2, status2)
}
