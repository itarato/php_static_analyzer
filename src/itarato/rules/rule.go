package rules

import (
	"itarato/types"
	"log"
)

// Main rule type.
type Rule struct {
}

// Rule interface.
type IRule interface {
	Read(char byte, buffer *string, context_state *types.Stack)
}

/**
 * Interface function for a Rule to read the next character from the file
 * and make decision about the structure.
 *
 * First deciding if the Rule can be applied: for example PHPScript is applicable
 * after "<?php".
 *
 * Then push the rule to the context.
 *
 * -- not finished --
 */
func (rule *Rule) Read(char byte, buffer *string, context_state *types.Stack) {
	log.Fatalln("Usage of abstract type method")
}
