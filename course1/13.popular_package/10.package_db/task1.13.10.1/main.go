package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *sql.DB

func CreateUserTable() error {
	schemaSQL := `CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	age INTEGER,
);`
	if _, err := db.Exec(schemaSQL); err != nil {
		return err
	}
	return nil
}

func InsertUser(user User) error {
	insertSQL := "INSERT INTO users (name, age) VALUES (?, ?)"
	if _, err := db.Exec(insertSQL, user.Name, user.Age); err != nil {
		return err
	}
	return nil
}

func SelectUser(id int) (User, error) {
	selectSQL := `SELECT * FROM users WHERE id = ?`

	rows, err := db.Query(selectSQL, id)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	rows.Next()
	var user User
	err = rows.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func UpdateUser(user User) error {
	updateSQL := "UPDATE users SET age = ? WHERE name = ?"
	if _, err := db.Exec(updateSQL, user.Age, user.Name); err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	deleteSQL := `DELETE FROM users WHERE id = ?`
	if _, err := db.Exec(deleteSQL, id); err != nil {
		return err
	}
	return nil
}

func main() {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Can't connect to database")
		return
	}
	defer db.Close()

	err = CreateUserTable()
	if err != nil {
		fmt.Println("Can't create user table")
		return
	}

	user := User{1, "Azamat", 22}
	err = InsertUser(user)
	if err != nil {
		fmt.Println("Can't insert user")
		return
	}

	selectedUser, err := SelectUser(1)
	if err != nil {
		fmt.Println("Can't select user")
		return
	}
	fmt.Println(selectedUser.Name, selectedUser.Age)

	user.Age = 23
	err = UpdateUser(user)
	if err != nil {
		fmt.Println("Can't update user")
		return
	}

	selectedUser, err = SelectUser(1)
	if err != nil {
		fmt.Println("Can't select user")
		return
	}
	fmt.Println(selectedUser.Name, selectedUser.Age)

	err = DeleteUser(1)
	if err != nil {
		fmt.Println("Can't delete user")
		return
	}

	_, err = SelectUser(1)
	if err == nil {
		fmt.Println("User was not deleted")
		return
	}
}
