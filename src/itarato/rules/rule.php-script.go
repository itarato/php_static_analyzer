package rules

import (
	"itarato/types"
)

type PHPScriptRule Rule

func (rule *PHPScriptRule) Read(char byte, buffer *string, context_state *types.Stack) {
	if *buffer == "<?php" {
		// @todo use proper states
		context_state.Push("php-script")

		// Clear buffer
		*buffer = ""
	}

	// @todo add later closing php tag
}
