package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

type Tabler interface {
	TableName() string
}

func (*User) TableName() string {
	return "users"
}

// Интерфейс для генерации SQL-запросов
type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}

// Интерфейс для генерации фейковых данных
type FakeDataGenerator interface {
	GenerateFakeUser() User
}

type SQLiteGenerator struct{}

func (SQLiteGenerator) CreateTableSQL(table Tabler) string {
	out := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", table.TableName())
	types := reflect.TypeOf(table).Elem()
	for i := 0; i < types.NumField(); i++ {
		field := types.FieldByIndex([]int{i})
		out += fmt.Sprintf("%v ", field.Tag.Get("db_field"))
		out += fmt.Sprintf("%v, ", field.Tag.Get("db_type"))
	}
	return strings.TrimRight(out, ", ") + ");"
}

func (SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	out := fmt.Sprintf("INSERT INTO %s (", model.TableName())
	values := reflect.ValueOf(model).Elem()
	types := reflect.TypeOf(model).Elem()
	for i := 0; i < types.NumField(); i++ {
		field := types.FieldByIndex([]int{i})
		out += fmt.Sprintf("%v, ", field.Tag.Get("db_field"))
	}
	out = strings.TrimRight(out, ", ") + ") VALUES ("
	for i := 0; i < values.NumField(); i++ {
		field := values.Field(i)
		if field.Kind() == reflect.String {
			out += fmt.Sprintf("'%v', ", field)
		} else {
			out += fmt.Sprintf("%v, ", field)
		}
	}
	return strings.TrimRight(out, ", ") + ");"
}

type GoFakeitGenerator struct{}

func (*GoFakeitGenerator) GenerateFakeUser() User {
	return User{
		ID:        gofakeit.IntRange(1, 9999),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
}

func main() {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}

	user := User{}
	sql := sqlGenerator.CreateTableSQL(&user)
	fmt.Println(sql)

	for i := 0; i < 23; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		fmt.Println(query)
	}
}
