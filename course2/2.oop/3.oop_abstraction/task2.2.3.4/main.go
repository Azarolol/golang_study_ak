package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

func (u *User) TableName() string {
	return "users"
}

type Tabler interface {
	TableName() string
}

type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}

type SQLiteGenerator struct{}

func (g SQLiteGenerator) CreateTableSQL(table Tabler) string {
	out := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", table.TableName())
	types := reflect.TypeOf(table).Elem()
	for i := 0; i < types.NumField(); i++ {
		field := types.FieldByIndex([]int{i})
		out += fmt.Sprintf("%v ", field.Tag.Get("db_field"))
		out += fmt.Sprintf("%v, ", field.Tag.Get("db_type"))
	}
	return strings.TrimRight(out, ", ") + ");"
}

func (g SQLiteGenerator) CreateInsertSQL(model Tabler) string {
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

type Migrator struct {
	db           *sql.DB
	sqlGenerator SQLGenerator
}

func NewMigrator(db *sql.DB, sqlGenerator SQLGenerator) *Migrator {
	return &Migrator{
		db:           db,
		sqlGenerator: sqlGenerator,
	}
}

func (m Migrator) Migrate(models ...Tabler) error {
	for _, model := range models {
		_, err := m.db.Exec(m.sqlGenerator.CreateTableSQL(model))
		if err != nil {
			return fmt.Errorf("failed to create table for model %v: %v", model, err)
		}
	}
	return nil
}

func main() {
	// Подключение к SQLite БД
	db, err := sql.Open("sqlite3", "file:my_database.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Создание мигратора с использованием вашего SQLGenerator
	migrator := NewMigrator(db, SQLiteGenerator{})

	// Миграция таблицы User
	if err := migrator.Migrate(&User{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}
