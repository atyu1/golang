package main

import (
	"fmt"
	"time"
)

type User struct {
	Username  string
	Email     string
	LastLogin time.Time
}

func (u *User) PrintDebug() {
	fmt.Println("### DEBUG: ")
	fmt.Println("# Username: ", u.Username)
	fmt.Println("# Email: ", u.Email)
	fmt.Println("# LastLogin: ", u.LastLogin)
}

func main() {
	var user1 User
	user1.Username = "john1"
	user1.Email = user1.Username + "@gmail.com"
	user1.LastLogin = time.Now()

	user2 := User{
		Username:  "katy2",
		Email:     "Kate@gmail.com",
		LastLogin: time.Now(),
	}

	user1.PrintDebug()
	user2.PrintDebug()
}
