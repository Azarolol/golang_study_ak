package main

import (
	"container/list"
	"fmt"
	"hash"
	"hash/crc32"
	"time"
)

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type Item struct {
	key   string
	value interface{}
}

type MapSlice struct {
	m    []Item
	hash hash.Hash
}

type MapList struct {
	m    *list.List
	hash hash.Hash
}

func (m *MapSlice) Set(key string, value interface{}) {
	m.hash.Reset()
	hashKey := string(m.hash.Sum([]byte(key)))
	m.m = append(m.m, Item{hashKey, value})
}

func (m *MapSlice) Get(key string) (interface{}, bool) {
	m.hash.Reset()
	hashKey := string(m.hash.Sum([]byte(key)))
	for _, i := range m.m {
		if i.key == hashKey {
			return i.value, true
		}
	}
	return nil, false
}

func (m *MapList) Set(key string, value interface{}) {
	m.hash.Reset()
	hashKey := string(m.hash.Sum([]byte(key)))
	m.m.PushBack(Item{hashKey, value})
}

func (m *MapList) Get(key string) (interface{}, bool) {
	m.hash.Reset()
	hashKey := string(m.hash.Sum([]byte(key)))
	cur := m.m.Front()
	for cur.Value != nil {
		if cur.Value.(Item).key == hashKey {
			return cur.Value.(Item).value, true
		}
		cur = cur.Next()
	}
	return nil, false
}

func NewHashMapSlice(count int, options ...func(*MapSlice)) HashMaper {
	out := &MapSlice{m: make([]Item, 0, count)}
	if len(options) == 0 {
		out.hash = crc32.NewIEEE()
	} else {
		for _, opt := range options {
			opt(out)
		}
	}
	return out
}

func NewHashMapList(count int, options ...func(*MapList)) HashMaper {
	out := &MapList{m: list.New()}
	if len(options) == 0 {
		out.hash = crc32.NewIEEE()
	} else {
		for _, opt := range options {
			opt(out)
		}
	}
	return out
}

func MeasureTime(f func()) time.Duration {
	start := time.Now()
	f()
	since := time.Since(start)
	return since
}

func main() {
	time := MeasureTime(TestSlice16)
	fmt.Println(time)
	time = MeasureTime(TestSlice1000)
	fmt.Println(time)

	time = MeasureTime(TestList16)
	fmt.Println(time)
	time = MeasureTime(TestList1000)
	fmt.Println(time)

}

func TestList16() {
	m := NewHashMapList(16)
	for i := 0; i < 16; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}

	for i := 0; i < 16; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}

func TestList1000() {
	m := NewHashMapList(1000)
	for i := 0; i < 1000; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}

	for i := 0; i < 1000; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}

func TestSlice16() {
	m := NewHashMapSlice(16)
	for i := 0; i < 16; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}

	for i := 0; i < 16; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}

func TestSlice1000() {
	m := NewHashMapSlice(1000)
	for i := 0; i < 1000; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}

	for i := 0; i < 1000; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}
