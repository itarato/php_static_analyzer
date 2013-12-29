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
	rule_stack := types.NewStack()
	rule_stack.Push(&rules.PHPFileRule{})

	// Buffer for the latest word to analyze. It's emptied after each analized segment.
	buffer_string := ""

	// var r *rules.IRule{}
	// Iterate through characters.
	for _, char := range file_content {
		buffer_string += string(char)
		// Debug information.
		fmt.Println("Read:\t", types.ReplaceWhitespaces(string(char)), "\tinto:\t", types.ReplaceWhitespaces(buffer_string))

		// Analyze the current position with the latest rule.
		active_rule, err := rule_stack.Top()
		if err != nil {
			log.Fatalln("Error: empty rule stack")
		}
		active_rule.(rules.IRule).Read(char, &buffer_string, rule_stack)
	}

}
