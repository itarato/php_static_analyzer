package rules

func GetRules() (rules []IRule) {
	phprule := &PHPScriptRule{}
	return []IRule{phprule}
}
