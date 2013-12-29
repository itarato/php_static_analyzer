package rules

import (
	"itarato/types"
	"log"
)

type Rule struct {
}

type IRule interface {
	Read(char byte, buffer *string, context_state *types.Stack)
}

func (rule *Rule) Read(char byte, buffer *string, context_state *types.Stack) {
	log.Fatalln("Usage of abstract type method")
}
