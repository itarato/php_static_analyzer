package main

import (
	"fmt"
	"io/ioutil"
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

	for _, char := range file_content {
		fmt.Println(string(char))
	}

}
