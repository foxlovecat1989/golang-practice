package main

import "fmt"

type book struct {
	name string
}

func (b *book) String() string {
	return fmt.Sprintf("This book's name is %s", b.name)
}

type number int

func (n *number) String() string {
	return fmt.Sprintf("This number is %d", *n)
}

func main() {
	b := book{
		"hello world",
	}
	fmt.Println(&b)

	var c number = 1
	fmt.Println(&c)
}
