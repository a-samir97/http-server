package models

import (
	"fmt"
)

type User struct {
	ID        int
	firstName string
	lastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("user is not found")
}

func UpdateUser(u User) (User, error) {
	for i, uu := range users {
		if uu.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("user is not found here")
}

func RemoveUser(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user is not found")
}
