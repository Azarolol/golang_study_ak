package main

import "fmt"

type BrowserHistory struct {
	stack []string
}

func (h *BrowserHistory) Visit(url string) {
	fmt.Println("Посещение", url)
	h.stack = append(h.stack, url)
}

func (h *BrowserHistory) Back() {
	if len(h.stack) == 0 {
		fmt.Println("Нет больше истории для возврата")
		return
	}
	h.stack = h.stack[:len(h.stack)-1]
	fmt.Println("Возвращение к", h.stack[len(h.stack)-1])
}

func (h *BrowserHistory) PrintHistory() {
	fmt.Println("История браузера:")
	for i := len(h.stack) - 1; i >= 0; i-- {
		fmt.Println(h.stack[i])
	}
}

func main() {
	history := &BrowserHistory{}
	history.Visit("www.google.com")
	history.Visit("www.github.com")
	history.Visit("www.openai.com")
	history.Back()
	history.PrintHistory()
}
