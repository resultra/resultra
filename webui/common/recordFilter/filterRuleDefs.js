// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
	"greaterEqual": {
		label: "Value greater or equal to",
		hasParam: true,
		paramLabel: "Enter value"
	},
	"less": {
		label: "Value less than",
		hasParam: true,
		paramLabel: "Filter if less than"
	},
	"lessEqual": {
		label: "Value less than or equal",
		hasParam: true,
		paramLabel: "Enter value"
	},
	"equal": {
		label: "Value equal",
		hasParam: true,
		paramLabel: "Enter value"
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