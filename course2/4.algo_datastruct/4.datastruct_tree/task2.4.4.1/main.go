package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type Node struct {
	index int
	data  *User
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (t *BinaryTree) insert(user *User) *BinaryTree {
	if t.root == nil {
		t.root = &Node{index: user.ID, data: user}
		return t
	}

	curr := t.root
	for {
		if curr.index < user.ID {
			if curr.right == nil {
				curr.right = &Node{index: user.ID, data: user}
				return t
			} else {
				curr = curr.right
			}
		} else if curr.index > user.ID {
			if curr.left == nil {
				curr.left = &Node{index: user.ID, data: user}
				return t
			} else {
				curr = curr.left
			}
		} else {
			return t
		}
	}
}

func (t *BinaryTree) search(key int) *User {
	curr := t.root
	for {
		if curr.index < key {
			if curr.right == nil {
				return nil
			} else {
				curr = curr.right
			}
		} else if curr.index > key {
			if curr.left == nil {
				return nil
			} else {
				curr = curr.left
			}
		} else {
			return curr.data
		}
	}
}

func generateData(n int) *BinaryTree {
	rand.Seed(time.Now().UnixNano())
	bt := &BinaryTree{}
	for i := 0; i < n; i++ {
		val := rand.Intn(100)
		bt.insert(&User{
			ID:   val,
			Name: fmt.Sprintf("User%d", val),
			Age:  rand.Intn(50) + 20,
		})
	}
	return bt
}

func main() {
	bt := generateData(50)
	user := bt.search(30)
	if user != nil {
		fmt.Printf("Найден пользователь: %+v\n", user)
	} else {
		fmt.Println("Пользователь не найден")
	}
}
