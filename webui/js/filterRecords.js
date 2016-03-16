

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

var filterRuleFormValidationNoParam = {
    filterRecordsSelectRule: {
      rules: [
        {
          type   : 'empty',
          prompt : 'Please enter a filtering rule'
        }
      ]
    }, // newFieldName validation
    filterRecordsSelectField: {
      rules: [
        {
          type   : 'empty',
          prompt : 'Please select a field'
        }
      ]
    }, // newFieldValTypeSelection validation
}

function filterRuleFormValidationTextParam() {
	var textValidation = filterRuleFormValidationNoParam
	
	textValidation.filterRecordsFilterRuleParamInput = {
        rules: [
          {
            type   : 'empty',
            prompt : 'Enter text to match'
          }
        ]		
	}
	
	return textValidation
}

function filterRuleFormValidationNumberParam() {
	var numValidation = filterRuleFormValidationNoParam
	
	numValidation.filterRecordsFilterRuleParamInput = {
        rules: [
          {
            type   : 'number',
            prompt : 'Enter a number'
          },
          {
            type   : 'empty',
            prompt : 'Enter a number'
          }
        ]		
	}
	
	return numValidation
}


function setFilterRuleFormValidationRules(validationRules) {
	$( "#filterRecordsAddFilterDialog" ).form({
		fields:validationRules,
		inline:true,
	});
	
}

function getFilterRecordsRuleDef(fieldsByID, fieldID, ruleID) {
	var fieldInfo = fieldsByID[fieldID]
	var typeRules = filterRulesByType[fieldInfo.type]
	var ruleDef = typeRules[ruleID]
	return ruleDef
}

function populateFilterPanelWithOneFilterRule(filterRuleRef)
{
	var fieldName = filterRuleRef.fieldRef.fieldInfo.name
	var ruleLabel = filterRuleRef.filterRuleDef.label
	
	// TODO - Filter rule items need better formatting & CSS style
	var filterRecordRuleItem = itemDivHTML(
		contentHTML(headerWithBodyHTML(fieldName,ruleLabel)) +
		contentHTML('<button class="ui compact icon button" style="padding:3px"><i class="remove icon"></i></button>')
	)
			
	$('#filterRecordsRuleList').append(filterRecordRuleItem)
	
}

function addFilterRule(newFilterRuleParams)
{
	console.log("Adding new filter rule: params = " + JSON.stringify(newFilterRuleParams))
	
	jsonAPIRequest("newRecordFilterRule",newFilterRuleParams,function(newFilterRuleRef) {
		populateFilterPanelWithOneFilterRule(newFilterRuleRef)
		// TODO - Also need to invoke a callback function to trigger an update to the view
		// (dashboard or form) which has a filter. The records shown in these views will 
		// change.
	}) // set record's number field value
}

function populateFilterPanel()
{
	var getFilterRulesParams = {} // Params are initially empty. TODO - Add parameters for which rules to retrieve
	jsonAPIRequest("getRecordFilterRules",getFilterRulesParams,function(filterRuleRefs) {
		for (ruleIter in filterRuleRefs) {
			filterRuleRef = filterRuleRefs[ruleIter]
			populateFilterPanelWithOneFilterRule(filterRuleRef)
		}
	}) // set record's number field value
	
}


function validateThenAddFilterRule(fieldsByID) {
	if($('#filterRecordsAddFilterDialog').form('validate form')) {
		console.log("filtering rule validated")
		
		var fieldID = $("#filterRecordsSelectFieldDropdown").dropdown("get value")
		var ruleID = $("#filterRecordsSelectFilterRuleDropdown").dropdown("get value")
		var ruleDef = getFilterRecordsRuleDef(fieldsByID,fieldID,ruleID)
		var fieldInfo = fieldsByID[fieldID]
		
		var newFilterRuleParams = {
			fieldID: fieldID,
			ruleID: ruleID,
		}
		
		if(ruleDef.hasParam) {
			
			var paramInput = $("#filterRecordsFilterRuleParamInput").dropdown("get value")
			
			if(fieldInfo.type == "text") {
				newFilterRuleParams.textRuleParam = paramInput
			} else {
				var numberParam = Number(paramInput)
				if(isNaN(numberParam)) {
					console.log("Unexpected value in filter paramter field: expecting a number, but got: " + paramInput)
					return;
				}
				newFilterRuleParams.numberRuleParam = numberParam
			}
		}
		
		addFilterRule(newFilterRuleParams)
		
		$( "#filterRecordsAddFilterDialog" ).dialog("close")
	} else {
		console.log("filtering rule not validated")
	}
}

function openAddFilterDialog(fieldsByID)
{
	var filterButton = $('#filterRecordsAddFilterButton')
		
	resetValidationRulesNoParam()	
	
	$( "#filterRecordsAddFilterDialog" ).form('clear');
	
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
}

function resetValidationRulesNoParam()
{
	setFilterRuleFormValidationRules(filterRuleFormValidationNoParam)
	$("#filterRecordsFilterRuleParamInputField").form('clear')
	$("#filterRecordsFilterRuleParamInputField").hide()
}

function addRecordFilterChangeField(fieldID, fieldsByID)
{
	var fieldInfo = fieldsByID[fieldID]
	console.log("filterRecords: selection changed: fieldID=" + fieldID + " type=" + fieldInfo.type)

	// Change the list of selectable rules to match the type of field
	$("#filterRecordsSelectFilterRuleDropdownMenu").empty()
	var rulesByType = filterRulesByType[fieldInfo.type]
	for(var ruleID in rulesByType) {
	 	var selectRuleHTML = dropdownSelectItemHTML(ruleID, rulesByType[ruleID].label)
	 	$("#filterRecordsSelectFilterRuleDropdownMenu").append(selectRuleHTML)				
	}
	
	$('#filterRecordsSelectFilterRuleDropdown').dropdown("clear")
	
	resetValidationRulesNoParam()
	
}


function initFilterRuleSelection(fieldsByID) {
	
//	$('#filterRecordsSelectFilterRuleDropdown').dropdown({placeholder: "Select a rule::"});
	$("#filterRecordsSelectFilterRuleDropdownMenu").empty()
	
	$('#filterRecordsSelectFilterRuleDropdown').dropdown({
	 	onChange: function() {
			var fieldID = $("#filterRecordsSelectFieldDropdown").dropdown("get value")
			if(fieldID.length > 0) {
				var ruleID = $("#filterRecordsSelectFilterRuleDropdown").dropdown("get value")
				if(ruleID.length > 0) {
					var ruleDef = getFilterRecordsRuleDef(fieldsByID,fieldID,ruleID)
					var fieldInfo = fieldsByID[fieldID]
			
					console.log("filterRecords: rule selection changed: ruleID=" + ruleID)
					console.log("Select filter rule: rule selection changed: filtering rule =" + JSON.stringify(ruleDef))
			
					if(ruleDef.hasParam) {
						if(fieldInfo.type == "text") {
							setFilterRuleFormValidationRules(filterRuleFormValidationTextParam())
						} else {
							console.log("setting number rules for filter param validation")
							setFilterRuleFormValidationRules(filterRuleFormValidationNumberParam())
						}
						$("#filterRecordsFilterRuleParamInputField").show()
					} else {
						console.log("setting no filter param rules for filter param validation")
						resetValidationRulesNoParam()
					}
					
				} // if ruleID length > 0
				else {
					resetValidationRulesNoParam()
				} // ruleID is reset to empty item => also reset the validation rules and hide the parameter input box
			
				
			} // if fieldID length > 0
				
		}
	});
	
}

function initFilterRecordsFieldSelectionMenu(fieldsByID) {
	
	$('#filterRecordsSelectFieldDropdown').dropdown();
	
	$("#filterRecordsSelectFieldDropdownMenu").empty()
	
	for (var fieldID in fieldsByID) {
		
		fieldInfo = fieldsByID[fieldID]
		
	 	var selectFieldHTML = dropdownSelectItemHTML(fieldID, fieldInfo.name)
		
	 	$("#filterRecordsSelectFieldDropdownMenu").append(selectFieldHTML)			

	} // for each field
	
	
	$('#filterRecordsSelectFieldDropdown').dropdown({
	 	onChange: function() {
			var fieldID = $("#filterRecordsSelectFieldDropdown").dropdown("get value")
			if(fieldID.length > 0) {
				console.log("filterRecords: selection changed: fieldID=" + fieldID)
				addRecordFilterChangeField(fieldID,fieldsByID)				
			}
			else {
				resetValidationRulesNoParam()
			}
		}
	});
	
	
}

function initFilterRecordsElems(fieldsByID) {
	
	$('#filterRecordsAddFilterButton').click(function(e){
		e.preventDefault();
		console.log("add filter button clicked")
		openAddFilterDialog(fieldsByID)
	})
	
	$( "#filterRecordsAddFilterDialog" ).dialog({ autoOpen: false })
	
	initFilterRecordsFieldSelectionMenu(fieldsByID)
	initFilterRuleSelection(fieldsByID)
	
	// Populate the filter panel using a JSON call to retrieve the list of filtering rules. This will
	// get more elaborate once record filtering is actually implemented, but this suffices for 
	// prototyping.
	populateFilterPanel()
	

}