package rules

func GetRules() (rules []IRule) {
	// Main rule is the php-file rule.
	return []IRule{&PHPFileRule{}}
}
