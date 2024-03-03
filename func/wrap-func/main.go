package main

import "fmt"

type Student struct{}

func main() {
	s := Student{}
	logInfo(&s)
}

func logInfo(s fmt.Stringer) {
	fmt.Println("message: ", s)
}

func (s *Student) String() string {
	return fmt.Sprint("I am a Student")
}
