package main

import "fmt"

func vars() {
	fmt.Println("Для оптимизации exe файла необходимо прописать go build -ldflags '-s -w' main.go")

	var age int

	var number float64 = 4.343
	fmt.Println(number)

	age2 := 16
	fmt.Println(age2)

	var name string

	what := "What is your name?"
	fmt.Println(what)
	fmt.Scan(&name)
	fmt.Println("Hello " + name + "!")
	fmt.Println("How old are you?")
	fmt.Scan(&age)
	fmt.Println("You are " + fmt.Sprint(age))

	var h int = 2
	var g float64 = float64(h)
	fmt.Println(g)
}
