package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Структура для хранения уникального идентификатора и имени пользователя
type User struct {
	ID   int
	Name string
}

// Структура для хранения словаря для поиска пользователей по ключу и mutex, которые используется для обеспечения потокобезопасной работы со словарем
type Cache struct {
	dict  map[string]*User
	mutex sync.RWMutex
}

// Функция для создания новой структуры Cache
func NewCache() *Cache {
	return &Cache{make(map[string]*User), sync.RWMutex{}}
}

// Функция для потокобезопасного добавления новой пары ключ-значение в словарь
func (c *Cache) Set(key string, user *User) {
	c.mutex.Lock()
	c.dict[key] = user
	c.mutex.Unlock()
}

// Функция для потокобезопасного извлечения структуры пользователя из словаря по ключу
func (c *Cache) Get(key string) *User {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.dict[key]
}

// Функция для создания ключей из множества значений
func keyBuilder(keys ...string) string {
	if len(keys) == 0 {
		return ""
	}
	out := keys[0]
	for i := 1; i < len(keys); i++ {
		out += ":" + keys[i]
	}
	return out
}

func main() {
	cache := NewCache()

	for i := 0; i < 100; i++ {
		go cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{
			ID:   i,
			Name: fmt.Sprint("user-", i),
		})
	}

	time.Sleep(1 * time.Second)

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(cache.Get(keyBuilder("user", strconv.Itoa(i))))
		}(i)
	}

	time.Sleep(1 * time.Second)
}
