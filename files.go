package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("Some text...")
	e := os.WriteFile("text.txt", data, 0600)

	if e != nil {
		fmt.Println("Файл не создался\n", e)
	}

	fileDate, err := os.ReadFile("text.txt")

	if err != nil {
		fmt.Println("Файл не прочитан\n", err)
	}
	fmt.Println(fileDate)         //байтовый срез
	fmt.Println(string(fileDate)) //нормальная строка

	os.Remove("text.txt")
}
