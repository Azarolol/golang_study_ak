package main

import "fmt"

type User struct {
	ID       int
	Username string
	Email    string
	Role     string
}

type UserOption func(*User)

func NewUser(ID int, uops ...UserOption) *User {
	user := &User{ID: ID}
	for _, uop := range uops {
		uop(user)
	}
	return user
}

func WithUsername(username string) UserOption {
	return func(u *User) {
		u.Username = username
	}
}

func WithEmail(email string) UserOption {
	return func(u *User) {
		u.Email = email
	}
}

func WithRole(role string) UserOption {
	return func(u *User) {
		u.Role = role
	}
}

func main() {
	user := NewUser(1, WithUsername("testuser"), WithEmail("testuser@example.com"), WithRole("admin"))
	fmt.Printf("User: %+v\n", user)
}
