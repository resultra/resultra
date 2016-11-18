var filterRulesText = {
	"isBlank": {
		label: "Text not set (blank)",
		hasParam: false,
	},
	"notBlank": {
		label: "Text is set (not blank)",
		hasParam: false,
	},
	"contains": {
		label: "Text contains",
		hasParam: true,
		paramLabel: "Filter if contains"
	},
	"startsWith": {
		label: "Text starts with",
		hasParam: true,
		paramLabel: "Filter if starts with"
	},
}

var filterRulesNumber = {
	"any": {
		label: "Any value (no filtering)",
		hasParam: false,
	},
	"isBlank": {
		label: "Value not set (blank)",
		hasParam: false,
	},
	"notBlank": {
		label: "Value is set (not blank)",
		hasParam: false,
	},
	"greater": {
		label: "Value greater than",
		hasParam: true,
		paramLabel: "Filter if greater than"
	},
	"less": {
		label: "Value less than",
		hasParam: true,
		paramLabel: "Filter if less than"
	}
}

var filterRulesByType = {
	"text":filterRulesText,
	"number":filterRulesNumber,
}

function getFilterRecordsRuleDef(fieldID, ruleID) {
	var fieldInfo = getFieldRef(fieldID)
	var typeRules = filterRulesByType[fieldInfo.type]
	var ruleDef = typeRules[ruleID]
	return ruleDef
}