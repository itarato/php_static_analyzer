package rules

import (
	"itarato/types/array"
)

type PHPScriptRule Rule

func (rule *PHPScriptRule) Read(char byte, buffer *string, context_state *array.Stack) {
	if *buffer == "<?php" {
		// @todo use proper states
		context_state.Push("php-script")
	}

	// @todo add later closing php tag
}
