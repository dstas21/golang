package main

import "fmt"

type Numbers struct {
	num1 int
	num2 int
}

type NumbersInterface interface {
	sum() int
	multipl() int
	div() float64
	substract() int
}

func (n Numbers) sum() int {
	return n.num1 + n.num2
}

func (n Numbers) multipl() int {
	return n.num1 * n.num2
}

func (n Numbers) div() float64 {
	return float64(n.num1) / float64(n.num2)
}

func (n Numbers) substract() int {
	return n.num1 - n.num2
}

func main() {
	var i NumbersInterface
	numb := Numbers{9, 4}
	i = numb
	fmt.Println(i.sum())
	fmt.Println(i.substract())
	fmt.Println(i.multipl())
	fmt.Println(i.div())

}
