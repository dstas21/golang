package main

import "fmt"

type User struct {
	name     string
	age      int
	password string
}

func main() {
	var user User = User{name: "John", age: 18, password: "GG"}
	fmt.Println(user)

	u := User{"Ben", 23, "pas"}
	fmt.Println(u)
	fmt.Println(u.name)
	fmt.Println(u.age)
	fmt.Println(u.password)
	u.name = "IO"
	fmt.Println(u)

	change(u)
	fmt.Println(u)

	changeTrue(&u)
	fmt.Println(u)
}

func change(u User) {
	u.name = "Kate"
}

func changeTrue(u *User) {
	u.name = "Kate"
	u.age = 35
	u.password = "pp"
}
