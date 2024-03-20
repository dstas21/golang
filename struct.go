package main

import "fmt"

type User struct {
	name     string
	age      int
	password string
	score    []int
}

func (user User) getName() string {
	return user.name
}

func (user *User) setName(name string) {
	user.name = name
}

func (user User) isElder() bool {
	isTrue := false
	if user.age >= 18 {
		isTrue = true
	}
	return isTrue
}

func (user User) getHighScore() int {
	var high int
	for _, score := range user.score {
		if high < score {
			high = score
		}
	}
	return high
}

func main() {
	var user User = User{name: "John", age: 18, password: "GG"}
	fmt.Println(user)

	u := User{"Ben", 23, "pas", []int{1, 23, 3, 4, 5}}
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

	fmt.Println(u.getName())
	u.setName("Tom")
	fmt.Println(u.getName())
	fmt.Println(u.isElder())
	fmt.Println(u.getHighScore())
}

func change(u User) {
	u.name = "Kate"
}

func changeTrue(u *User) {
	u.name = "Kate"
	u.age = 15
	u.password = "pp"
}
