package main

func main() {
	a := 5
	pointer := &a
	println(pointer)  //ячейка в памяти
	println(*pointer) //значение в ячейке * - указатель разъименования

	s := "S"
	println(s)
	change(s)
	println(s)
	changeTrue(&s)
	println(s)

}

func change(str string) { //копирует переменную в новую ячейку памяти
	str = "GG"
}
func changeTrue(str *string) { //работает с переданной переменной и ее ячейкой в памяти
	*str = "GG"
}
