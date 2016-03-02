

var filterRulesText = {
	"isBlank": {
		label: "Value not set (blank)",
		hasParam: false,
	},
	"notBlank": {
		label: "Value is set (not blank)",
		hasParam: false,
	},
	"contains": {
		label: "Value contains",
		hasParam: true,
		paramLabel: "Filter if contains"
	},
	"startsWith": {
		label: "Starts with",
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
		label: "Greater than",
		hasParam: true,
		paramLabel: "Filter if greater than"
	},
	"less": {
		label: "Less than",
		hasParam: true,
		paramLabel: "Filter if less than"
	}
}

var filterRulesByType = {
	"text":filterRulesText,
	"number":filterRulesNumber,
}

function openAddFilterDialog()
{
	var filterButton = $('#filterRecordsAddFilterButton')
	
	$("#filterRecordsAddFilterDialog").dialog({
		autoOpen: false,
		height: 450, width: 300,
		modal: false,
		position: { my: "right top", at: "left-10 top", of: filterButton },
		buttons: { 
			"Add Filter": function() { $(this).dialog('close'); },
  			"Cancel" : function() { $(this).dialog('close'); },
 		 },	
	 });
	  
	 $("#filterRecordsAddFilterDialog").dialog("open")
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
}

function initFilterRuleSelection(fieldsByID) {
	
	$('#filterRecordsSelectFilterRuleDropdown').dropdown();
	$("#filterRecordsSelectFilterRuleDropdownMenu").empty()
	
	$('#filterRecordsSelectFilterRuleDropdown').dropdown({
	 	onChange: function() {
			var fieldID = $("#filterRecordsSelectFieldDropdown").dropdown("get value")
			var ruleID = $("#filterRecordsSelectFilterRuleDropdown").dropdown("get value")
			
			var fieldInfo = fieldsByID[fieldID]
			var typeRules = filterRulesByType[fieldInfo.type]
			var ruleDef = typeRules[ruleID]
			
			console.log("filterRecords: rule selection changed: ruleID=" + ruleID)
			console.log("Select filter rule: rule selection changed: filtering rule =" + JSON.stringify(ruleDef))
			
			if(ruleDef.hasParam) {
				$("#filterRecordsFilterRuleParamInputField").show()
			} else {
				$("#filterRecordsFilterRuleParamInputField").hide()
			}
				
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
			console.log("filterRecords: selection changed: fieldID=" + fieldID)
			addRecordFilterChangeField(fieldID,fieldsByID)
		}
	});
	
	
}

function initFilterRecordsElems(fieldsByID) {
	
	
	$('#filterRecordsAddFilterButton').click(function(e){
		e.preventDefault();
		console.log("add filter button clicked")
		openAddFilterDialog()
	})
	
	$( "#filterRecordsAddFilterDialog" ).dialog({ autoOpen: false })
	
	initFilterRecordsFieldSelectionMenu(fieldsByID)
	initFilterRuleSelection(fieldsByID)
}