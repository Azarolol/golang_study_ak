package main

import (
	"testing"
)

func TestUser_TableName(t *testing.T) {
	tests := []struct {
		name string
		u    *User
		want string
	}{
		{"Test user table name", &User{}, "users"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.TableName(); got != tt.want {
				t.Errorf("User.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLiteGenerator_CreateTableSQL(t *testing.T) {
	type args struct {
		table Tabler
	}
	tests := []struct {
		name string
		g    SQLiteGenerator
		args args
		want string
	}{
		{"Test SQLite generator create table", SQLiteGenerator{}, args{&User{}}, `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		email VARCHAR(100) UNIQUE
	);`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.CreateTableSQL(tt.args.table); got != tt.want {
				t.Errorf("SQLiteGenerator.CreateTableSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLiteGenerator_CreateInsertSQL(t *testing.T) {
	type args struct {
		model Tabler
	}
	tests := []struct {
		name string
		g    SQLiteGenerator
		args args
		want string
	}{
		{"Test SQLite generator create insert", SQLiteGenerator{}, args{&User{ID: 1, FirstName: "Azamat", LastName: "Gimazov", Email: "gimazovazamat@mail.ru"}}, "INSERT INTO users (id, first_name, last_name, email) VALUES (1, 'Azamat', 'Gimazov', 'gimazovazamat@mail.ru');"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.CreateInsertSQL(tt.args.model); got != tt.want {
				t.Errorf("SQLiteGenerator.CreateInsertSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test main"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
