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


function initFilterRuleSelection(fieldsByID) {
		
	resetFormControl("#filterRecordsSelectFilterRuleDropdown")
	resetFormControl("#filterRecordsFilterRuleParamInputField")
	
	$('#filterRecordsSelectFilterRuleDropdownMenu').change( function() {
		var fieldID = $("#filterRecordsSelectFieldDropdownMenu").val()
		if(fieldID.length > 0) {
			
			removeFormControlError("#filterRecordsSelectFilterRuleDropdown")
			
			var ruleID = $("#filterRecordsSelectFilterRuleDropdownMenu").val()
			if(ruleID.length > 0) {
				var ruleDef = getFilterRecordsRuleDef(fieldsByID,fieldID,ruleID)
				var fieldInfo = fieldsByID[fieldID]
		
				console.log("filterRecords: rule selection changed: ruleID=" + ruleID)
				console.log("Select filter rule: rule selection changed: filtering rule =" + JSON.stringify(ruleDef))
		
				if(ruleDef.hasParam) {
					if(fieldInfo.type == "text") {
						// TODO Set validation rules based upon text field (non-empty)
						//setFilterRuleFormValidationRules(filterRuleFormValidationTextParam())
					} else {
						// TODO Set validation rules based upon text field (non-empty)
					}
					$("#filterRecordsFilterRuleParamInputField").slideDown()
				} else {
					console.log("setting no filter param rules for filter param validation")
					configureAddFilterDialogNoParam()
				}
				
			} // if ruleID length > 0
			else {
				configureAddFilterDialogNoParam()
			} // ruleID is reset to empty item => also reset the validation rules and hide the parameter input box
		
			
		} // if fieldID length > 0
				
	});
	
}

function initFilterRecordsFieldSelectionMenu(fieldsByID) {
	
	
	resetFormControl("#filterRecordsSelectFieldDropdown")
	$("#filterRecordsSelectFieldDropdownMenu").empty()	
	 $("#filterRecordsSelectFieldDropdownMenu").append(defaultSelectOptionPromptHTML("Select a field"))
	
	// Populate the selection menu for selecting the field to filter on
	for (var fieldID in fieldsByID) {
		fieldInfo = fieldsByID[fieldID]
	 	var selectFieldHTML = selectOptionHTML(fieldID, fieldInfo.name)
	 	$("#filterRecordsSelectFieldDropdownMenu").append(selectFieldHTML)
	} // for each field
	
	
	$('#filterRecordsSelectFieldDropdownMenu').change( function() {
		var fieldID = $("#filterRecordsSelectFieldDropdownMenu").val()
		if(fieldID.length > 0) {
			removeFormControlError("#filterRecordsSelectFieldDropdown")
			console.log("filterRecords: selection changed: fieldID=" + fieldID)
			addRecordFilterChangeField(fieldID,fieldsByID)				
		}
		else {
			configureAddFilterDialogNoParam()
		}
	});
	
	
}

function configureAddFilterDialogNoParam()
{
	$("#filterRecordsFilterRuleParamInputField").hide()
	$("#filterRecordsFilterRuleParamInputField").val("")
	// TODO - Re-integrate with Bootstrap
	return
	setFilterRuleFormValidationRules(filterRuleFormValidationNoParam)
}

function addRecordFilterChangeField(fieldID, fieldsByID)
{
	var fieldInfo = fieldsByID[fieldID]
	console.log("filterRecords: selection changed: fieldID=" + fieldID + " type=" + fieldInfo.type)

	// Change the list of selectable rules to match the type of field
	$("#filterRecordsSelectFilterRuleDropdownMenu").empty()
	$("#filterRecordsSelectFilterRuleDropdownMenu").append(defaultSelectOptionPromptHTML("Select a field"))				
	var rulesByType = filterRulesByType[fieldInfo.type]
	for(var ruleID in rulesByType) {
	 	var selectRuleHTML = selectOptionHTML(ruleID, rulesByType[ruleID].label)
	 	$("#filterRecordsSelectFilterRuleDropdownMenu").append(selectRuleHTML)				
	}
	
	configureAddFilterDialogNoParam()
	
}


function recordFilterValidateAddFilterForm() {
	
	var fieldID = $("#filterRecordsSelectFieldDropdownMenu").val()
	var foundError = false
	if((fieldID == null) || (fieldID.length <= 0)) {
		addFormControlError("#filterRecordsSelectFieldDropdown")
		foundError = true
	}
	
	var ruleID = $("#filterRecordsSelectFilterRuleDropdownMenu").val()
	if((ruleID == null) || (ruleID.length <= 0)) {
		addFormControlError("#filterRecordsSelectFilterRuleDropdown")
		foundError = true
	}
	
	if(foundError) {
		return null
	}
	
	var ruleDef = getFilterRecordsRuleDef(addFilterDialogContext.fieldsByID,fieldID,ruleID)
	var fieldInfo = addFilterDialogContext.fieldsByID[fieldID]
	
	var newFilterRuleParams = {
		fieldID: fieldID,
		ruleID: ruleID,
	}
	
	if(ruleDef.hasParam) {
		
		var paramInput = $("#filterRecordsFilterRuleParamInput").val()
		
		if(fieldInfo.type == "text") {
			newFilterRuleParams.textRuleParam = paramInput
		} else {
			var numberParam = Number(paramInput)
			if(isNaN(numberParam)) {
				foundError = true
				addFormControlError("#filterRecordsFilterRuleParamInputField")
				console.log("Unexpected value in filter paramter field: expecting a number, but got: " + paramInput)
				return;
			} else {
				newFilterRuleParams.numberRuleParam = numberParam	
			}
		}
	}
	
	if(foundError) {
		return null
	}
	
	return newFilterRuleParams
}

function validateThenAddFilterRule(fieldsByID) {
	
	var newFilterRuleParams = recordFilterValidateAddFilterForm()
	if(newFilterRuleParams != null) {
		console.log("filtering rule validated")
		addFilterRule(newFilterRuleParams)
		$( "#filterRecordsAddFilterDialog" ).dialog("close")
	} else {
		console.log("filtering rule not validated")
	}
}

var addFilterDialogContext = {}

function openAddFilterDialog(parentTableID)
{		
	configureAddFilterDialogNoParam()

	loadFieldInfo(parentTableID,[fieldTypeAll],function(fieldsByID) {
		
		addFilterDialogContext.fieldsByID = fieldsByID
		
		initFilterRecordsFieldSelectionMenu(fieldsByID)
		initFilterRuleSelection(fieldsByID)
		
		var filterButton = $('#filterRecordsAddFilterButton')
		$("#filterRecordsAddFilterDialog").dialog({
			autoOpen: false,
			height: 450, width: 300,
			resizable: false,
			modal: false,
			position: { my: "right top", at: "left-10 top", of: filterButton },
			buttons: { 
				"Add Filtering Rule": function() { validateThenAddFilterRule(fieldsByID) },
	  			"Cancel" : function() { $(this).dialog('close'); },
	 		 },	
		 });
	  
		 $("#filterRecordsAddFilterDialog").dialog("open")
	})

// TODO - Reintegrate with Bootstrap	
//	$( "#filterRecordsAddFilterDialog" ).form('clear');
	
}
