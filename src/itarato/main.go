package main

import (
	"fmt"
	"io/ioutil"
	"itarato/types/array"
	"log"
	"os"
)

func main() {

	file_name := os.Args[1]
	log.Println("File name: ", file_name)

	file_content, err := ioutil.ReadFile(file_name)
	log.Println("File length:", len(file_content), "bytes")

	if err != nil {
		log.Fatalln("Error occured during file read")
	}

	state_stack := array.New()
	for _, char := range file_content {
		fmt.Println(string(char))
		state_stack.Push(char)
	}

}
