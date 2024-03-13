package main

import (
	"fmt"
	"math"
)

func main() {
	println("Калькулятор")
	print("Введите действие (+, -, *, /, sqrt, round) :")

	var action string
	fmt.Scan(&action)

	print("Введите число a: ")
	var a float64
	fmt.Scan(&a)

	var b float64
	if action != "sqrt" {
		if action != "round" {
			print("Введите число b: ")
			fmt.Scan(&b)
		}
	}

	var result float64
	switch {
	case action == "+":
		result = a + b
	case action == "-":
		result = a - b
	case action == "*":
		result = a * b
	case action == "/":
		result = a / b
	case action == "sqrt":
		result = math.Sqrt(a)
	case action == "round":
		result = math.Round(a)
	}

	print("Результат")
	print(result)

}
