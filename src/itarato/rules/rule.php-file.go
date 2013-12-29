package rules

import (
	"itarato/types"
)

type PHPFileRule Rule

func (rule *PHPFileRule) Read(char byte, buffer *string, context_state *types.Stack) {
	if *buffer == "<?php" {
		// @todo use proper states
		context_state.Push(&PHPScriptRule{})

		// Clear buffer
		*buffer = ""
	}

	// @todo add later closing php tag
}
