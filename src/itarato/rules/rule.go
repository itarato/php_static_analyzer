package rules

import (
	"itarato/types"
)

// Main rule type.
type Rule struct {
}

// Rule interface.
type IRule interface {
	/**
	 * Try the rule if it's applicable.
	 */
	TryOn(char byte, buffer *string, context_state *types.Stack) bool

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
	Read(char byte, buffer *string, context_state *types.Stack)

	/**
	 * Get a human readable name of the rule - informational.
	 */
	GetName() string
}
