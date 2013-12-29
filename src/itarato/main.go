package main

import (
	"fmt"
	"io/ioutil"
	"itarato/rules"
	"itarato/types"
	"log"
	"os"
)

/**
 * Main entry point.
 */
func main() {
	fmt.Println("PHP parser V0.1")

	// File of subject.
	file_name := os.Args[1]
	log.Println("File name:", file_name)

	// Load file content.
	file_content, err := ioutil.ReadFile(file_name)
	log.Println("File length:", len(file_content), "bytes")

	if err != nil {
		log.Fatalln("Error occured during file read")
	}

	// Stack for the actual context state.
	context_state_stack := types.NewStack()
	buffer_string := ""
	rules := rules.GetRules()

	for _, char := range file_content {
		buffer_string += string(char)
		fmt.Println("Read:\t", types.ReplaceWhitespaces(string(char)), "\tinto:\t", types.ReplaceWhitespaces(buffer_string))

		for _, rule := range rules {
			rule.Read(char, &buffer_string, context_state_stack)
		}
	}

}
