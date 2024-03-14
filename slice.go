package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{3, 1, 3, 4, 5, 4, 6}
	fmt.Println(slice)
	slice = append(slice, 10)
	fmt.Println(slice)
	sort.Ints(slice)
	fmt.Println(slice)

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, element := range slice {
		fmt.Printf("%d: %d\n", i, element)
	}
	for element := range slice {
		fmt.Printf("%d\n", element)
	}

	sliceString := []string{"f", "h", "d", "a", "c"}
	sort.Strings(sliceString)
	fmt.Println(sliceString)
}
