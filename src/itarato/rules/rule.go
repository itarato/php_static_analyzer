package rules

import (
	"itarato/types/array"
	"log"
)

type Rule struct {
}

type IRule interface {
	Read(char byte, buffer *string, context_state *array.Stack)
}

func (rule *Rule) Read(char byte, buffer *string, context_state *array.Stack) {
	log.Fatalln("Usage of abstract type method")
}
