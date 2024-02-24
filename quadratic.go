package main

import (
	"fmt"
	"math"
)

func quadratic() {
	var a, b, c float64
	fmt.Println("Квадратное уравнение")
	fmt.Println("Введите a")
	fmt.Scan(&a)
	fmt.Println("Введите b")
	fmt.Scan(&b)
	fmt.Println("Введите c")
	fmt.Scan(&c)

	D := b*b - 4*(a*c)

	if D > 0 {
		var x1, x2 float64

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Ваше квадратное уравнение имеет два корня\nD = " + fmt.Sprint(D) +
			"\nx1 = " + fmt.Sprint(x1) +
			"\nx2 = " + fmt.Sprint(x2))
	} else if D == 0 {
		var x float64

		x = (-b) / (2 * a)

		fmt.Println("Ваше квадратное уравнение имеет один корен\nD = " + fmt.Sprint(D) +
			"\nx = " + fmt.Sprint(x))
	} else if D < 0 {

		fmt.Println("Ваше квадратное уравнение не имеет корней\nD = " + fmt.Sprint(D))
	}

	fmt.Scan(&a)
}
