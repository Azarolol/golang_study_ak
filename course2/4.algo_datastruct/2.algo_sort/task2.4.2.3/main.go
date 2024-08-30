package main

// Структура пользователя
type User struct {
	ID   int
	Name string
	Age  int
}

// Функция слияния двух отсортированных массивов пользователей
func Merge(arr1 []User, arr2 []User) []User {
	out := make([]User, 0, len(arr1)+len(arr2))
	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0].ID <= arr2[0].ID {
			out = append(out, arr1[0])
			arr1 = arr1[1:]
		} else {
			out = append(out, arr2[0])
			arr2 = arr2[1:]
		}
	}
	out = append(out, arr1...)
	out = append(out, arr2...)
	return out
}
