package main

func test(someFunction func(int) int) {
	println(someFunction(25))
}

func test2(x string) func() {
	return func() {
		println(x)
	}
}

func main() {
	first()
	first()
	sum(5, 6)
	value := sumWithResult(3, 4)
	println(value)
	sum, sub, multip, div := mathFunc(8, 2)
	println(sum, sub, multip, div)

	a := func() {
		println("a")
	}
	a()
	b := func(a int, b int) {
		println(a + b*a)
	}
	b(2, 3)

	c := func(a int, b int) int {
		return a + b*a
	}
	println(c(5, 10))

	sqrt := func(x int) int {
		return x * x
	}

	test(sqrt)
	test2("test2")()
}

func first() {
	println("Hello from first")
}

func sum(a int, b int) {
	println(a + b)
}

func sumWithResult(a int, b int) int {
	return a + b
}

func mathFunc(a int, b int) (int, int, int, int) {
	sum := a + b
	sub := a - b
	multip := a * b
	div := a / b
	return sum, sub, multip, div
}

func mathFunc2(a int, b int) (sum int, sub int, multip int, div int) {
	sum = a + b
	sub = a - b
	multip = a * b
	div = a / b
	return
}
