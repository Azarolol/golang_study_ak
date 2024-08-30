package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Node struct {
	data *Commit
	prev *Node
	next *Node
}

type DoubleLinkedList struct {
	head *Node // начальный элемент в списке
	tail *Node // последний элемент в списке
	curr *Node // текущий элемент меняется при использовании методов next, prev
	len  int   // количество элементов в списке
}

type LinkedLister interface {
	LoadData(path string) error
	Init(c []Commit)
	Len() int
	SetCurrent(n int) error
	Current() *Node
	Next() *Node
	Prev() *Node
	Insert(n int, c Commit) error
	Push(c Commit) error
	Delete(n int) error
	DeleteCurrent() error
	Index() (int, error)
	GetByIndex(n int) (*Node, error)
	Pop() *Node
	Shift() *Node
	SearchUUID(uuID string) *Node
	Search(message string) *Node
	Reverse() *DoubleLinkedList
}

// LoadData loads data from a JSON file at the given path into the list.
func (d *DoubleLinkedList) LoadData(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var commits []Commit
	if err := json.Unmarshal(bytes, &commits); err != nil {
		return err
	}
	QuickSort(commits)
	d.Init(commits)
	return nil
}

func (d *DoubleLinkedList) Init(cs []Commit) {
	for _, c := range cs {
		d.Push(c)
	}
}

// Len получения длины списка
func (d *DoubleLinkedList) Len() int {
	return d.len
}

// Current получение текущего элемента
func (d *DoubleLinkedList) Current() *Node {
	return d.curr
}

// Next получение следующего элемента
func (d *DoubleLinkedList) Next() *Node {
	d.curr = d.curr.next
	return d.curr
}

func (d *DoubleLinkedList) Prev() *Node {
	d.curr = d.curr.prev
	return d.curr
}

// Insert вставка элемента после n элемента
func (d *DoubleLinkedList) Insert(n int, c Commit) error {
	if n < 0 || n > d.len {
		return errors.New("index out of bounds")
	}

	newNode := &Node{data: &c}
	if n == 0 {
		if d.head == nil {
			d.head = newNode
			d.tail = newNode
			d.curr = newNode
		} else {
			newNode.next = d.head
			d.head.prev = newNode
			d.head = newNode
		}
	} else if n == d.len {
		d.tail.next = newNode
		newNode.prev = d.tail
		d.tail = newNode
	} else {
		current := d.head
		for i := 0; i < n; i++ {
			current = current.next
		}
		newNode.next = current
		newNode.prev = current.prev
		current.prev.next = newNode
		current.prev = newNode
	}
	d.len++
	return nil
}

func (d *DoubleLinkedList) GetByIndex(n int) (*Node, error) {
	if n < 0 || n > d.len {
		return nil, errors.New("index out of bounds")
	}

	curr := d.head
	for i := 0; i < n; i++ {
		curr = curr.next
	}
	return curr, nil
}

func (d *DoubleLinkedList) SetCurrent(n int) error {
	if n < 0 || n > d.len {
		return errors.New("index out of bounds")
	}
	curr := d.head
	for i := 0; i < n; i++ {
		curr = curr.next
	}
	d.curr = curr
	return nil
}

func (d *DoubleLinkedList) Push(c Commit) error {
	newNode := &Node{data: &c}
	if d.tail == nil {
		d.head = newNode
		d.tail = newNode
		d.curr = newNode
	} else {
		d.tail.next = newNode
		newNode.prev = d.tail
		d.tail = newNode
	}
	d.len++
	return nil
}

// Delete удаление n элемента
func (d *DoubleLinkedList) Delete(n int) error {
	if n < 0 || n > d.len {
		return errors.New("index out of bounds")
	}

	if n == 0 {
		d.head.next.prev = nil
		d.head = d.head.next
	} else if n == d.len {
		d.tail.prev.next = nil
		d.tail = d.tail.prev
	} else {
		current := d.head
		for i := 0; i < n; i++ {
			current = current.next
		}
		current.prev.next = current.next
		current.next.prev = current.prev
	}
	d.len--
	return nil
}

// DeleteCurrent удаление текущего элемента
func (d *DoubleLinkedList) DeleteCurrent() error {
	if d.curr == nil {
		return errors.New("current is nil")
	}
	if d.curr.next != nil && d.curr.prev != nil {
		d.curr.prev.next = d.curr.next
		d.curr.next.prev = d.curr.prev
		d.curr = d.curr.next
	} else if d.curr.next == nil {
		d.curr.prev.next = nil
		d.tail = d.curr.prev
		d.curr = d.curr.prev
	} else if d.curr.prev == nil {
		d.curr.next.prev = nil
		d.head = d.curr.next
		d.curr = d.curr.next
	} else {
		d.head = nil
		d.tail = nil
		d.curr = nil
	}
	d.len--
	return nil
}

// Index получение индекса текущего элемента
func (d *DoubleLinkedList) Index() (int, error) {
	if d.curr == nil {
		return 0, errors.New("current is nil")
	}
	curr := d.curr
	i := 0
	for curr.prev != nil {
		i++
		curr = curr.prev
	}
	return i, nil
}

// Pop Операция Pop
func (d *DoubleLinkedList) Pop() *Node {
	if d.tail == nil {
		return nil
	}

	out := d.tail
	if d.tail.prev != nil {
		d.tail.prev.next = nil
	}
	d.tail = d.tail.prev
	d.len--

	return out
}

// Shift операция shift
func (d *DoubleLinkedList) Shift() *Node {
	if d.head == nil {
		return nil
	}

	out := d.head
	if d.head.next != nil {
		d.head.next.prev = nil
	}
	d.head = d.head.next
	d.len--

	return out
}

// SearchUUID поиск коммита по uuid
func (d *DoubleLinkedList) SearchUUID(uuID string) *Node {
	curr := d.head
	for curr != nil {
		if curr.data.UUID == uuID {
			return curr
		}
		curr = curr.next
	}
	return nil
}

func (d *DoubleLinkedList) Search(message string) *Node {
	curr := d.head
	for curr != nil {
		if curr.data.Message == message {
			return curr
		}
		curr = curr.next
	}
	return nil
}

// Reverse возвращает перевернутый список
func (d *DoubleLinkedList) Reverse() *DoubleLinkedList {
	newD := &DoubleLinkedList{}
	curr := d.tail
	for curr != nil {
		newD.Push(*curr.data)
		curr = curr.prev
	}
	return newD
}

type Commit struct {
	Message string `json:"message"`
	UUID    string `json:"uuid"`
	Date    string `json:"date"`
}

func QuickSort(commits []Commit) {
	quicksort(commits, 0, len(commits)-1)
}

func quicksort(arr []Commit, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quicksort(arr, low, pi-1)
		quicksort(arr, pi+1, high)
	}
}

func partition(arr []Commit, low, high int) int {
	pivot := arr[high].Date
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j].Date > pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func GenerateData(numCommits int) []Commit {
	var commits []Commit
	gofakeit.Seed(0)

	for i := 0; i < numCommits; i++ {
		commit := Commit{
			Message: gofakeit.Sentence(5),
			UUID:    gofakeit.UUID(),
			Date:    gofakeit.DateRange(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)).Format("2006-01-02"),
		}
		commits = append(commits, commit)
	}

	return commits
}

func main() {
	var d LinkedLister = &DoubleLinkedList{}
	name := "data.json"
	err := createJSON(name)
	if err != nil {
		log.Fatal(err)
	}
	d.LoadData(name)
	for curr := d.Current(); curr != nil; curr = d.Next() {
		fmt.Println(*curr.data)
	}
}

func createJSON(name string) error {
	commits := GenerateData(10)
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(commits)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
