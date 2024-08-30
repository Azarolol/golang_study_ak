package main

import "fmt"

type CircuitRinger interface {
	Add(val int)
	Get() (int, bool)
}

type RingBuffer struct {
	buf  []int
	head int
	tail int
	len  int
	cap  int
}

func NewRingBuffer(size int) CircuitRinger {
	return &RingBuffer{
		buf:  make([]int, size),
		len:  0,
		cap:  size,
		head: 0,
		tail: 0,
	}
}

func (b *RingBuffer) Add(val int) {
	b.buf[b.tail] = val
	if b.len < b.cap {
		b.len++
	} else {
		b.head++
		if b.head == b.cap {
			b.head = 0
		}
	}
	b.tail++
	if b.tail == b.cap {
		b.tail = 0
	}
}

func (b *RingBuffer) Get() (int, bool) {
	if b.len == 0 {
		return 0, false
	}
	out := b.buf[b.head]
	b.head++
	if b.head == b.cap {
		b.head = 0
	}
	b.len--
	return out, true
}

func main() {
	rb := NewRingBuffer(3)
	rb.Add(1)
	rb.Add(2)
	rb.Add(3)
	rb.Add(4) // Перезаписывает значение 1

	for val, ok := rb.Get(); ok; val, ok = rb.Get() {
		fmt.Println(val) // Выводит: 2, 3, 4
	}

	if _, ok := rb.Get(); !ok {
		fmt.Println("Buffer is empty") // Выводит: Buffer is empty
	}
}
