package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

type Animal struct {
	Type string
	Name string
	Age  int
}

func getAnimals() []Animal {
	animals := make([]Animal, 3)
	for i := 0; i < 3; i++ {
		animals[i] = Animal{gofakeit.Animal(), gofakeit.FirstName(), gofakeit.Number(1, 10)}
	}
	return animals
}

func preparePrint(animals []Animal) string {
	output := ""
	for _, animal := range animals {
		output += fmt.Sprintf("Тип: %s, Имя: %s, Возраст: %d\n", animal.Type, animal.Name, animal.Age)
	}
	return output
}

func main() {
	// Пример использования функции
	animals := getAnimals() // Получаем срез 3 животных
	result := preparePrint(animals)
	fmt.Println(result)
}
