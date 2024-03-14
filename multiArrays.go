package main

import (
	"fmt"
)

func main() {
	multiArray := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(multiArray)
	fmt.Println(multiArray[1][2])
	fmt.Println(multiArray[2][3])
}
