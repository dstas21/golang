package main

func main() {
	a := 1
	b := 10

	if a >= 1 {
		println("a >= 1")
	}
	if a > 0 {
		println("a > 0")
	}
	if a < 1 {
		println("a < 1")
	}
	if a > 0 && b > 5 {
		println("TRUE a > 0 && b > 5")
	}
	if a > 0 && b > 11 {
		println("TRUE a > 0 && b > 11")
	}
	if a > 0 || b > 5 {
		println("TRUE a > 0 || b > 5")
	}
	if a > 0 || b > 11 {
		println("TRUE a > 0 || b > 11")
	}
	if a > 1 || b > 11 {
		println("TRUE a > 0 || b > 5")
	}
	if a != 2 {
		println("TRUE a != 2")
	}
	if b != 10 {
		println("TRUE b != 10")
	}
	if a != 2 && b > 5 || a > 6 {
		println("TRUE a != 2 && b > 5 || a > 6")
	}
	if a != 2 && b > 5 && a > 6 {
		println("TRUE a != 2 && b > 5 && a > 6")
	}
}
