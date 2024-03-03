package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("./os/write/message.txt")
	if err != nil {
		log.Fatalln("err: ", err)
	}
	defer f.Close()

	_, err = f.Write([]byte("Hello World2"))
	if err != nil {
		log.Fatalln("err: ", err)
	}
}
