package main

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func PrepareQuery(operation string, table string, user User) (string, []interface{}, error) {
	switch operation {
	case "insert":
		return sq.Insert(table).Columns("name", "age").Values(user.Name, user.Age).ToSql()
	case "update":
		return sq.Update(table).Set("age", user.Age).Where(sq.Eq{"name": user.Name}).ToSql()
	case "delete":
		return sq.Delete(table).Where(sq.Eq{"id": user.ID}).ToSql()
	case "select":
		return sq.Select("id", "name", "age").From(table).Where(sq.Eq{"id": user.ID}).ToSql()
	}
	return "", nil, fmt.Errorf("unknown operation")
}

func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

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
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := PrepareQuery("insert", "users", user)
	if err != nil {
		return err
	}

	if _, err := db.Exec(query, args...); err != nil {
		return err
	}
	return nil
}

func SelectUser(id int) (User, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	query, args, err := PrepareQuery("insert", "users", User{ID: id})
	if err != nil {
		return User{}, err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return User{}, err
	}

	rows.Next()
	var user User
	err = rows.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func UpdateUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := PrepareQuery("update", "users", user)
	if err != nil {
		return err
	}

	if _, err := db.Exec(query, args...); err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := PrepareQuery("delete", "users", User{ID: id})
	if err != nil {
		return err
	}

	if _, err := db.Exec(query, args...); err != nil {
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
