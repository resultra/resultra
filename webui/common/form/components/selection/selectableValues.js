function selectComponentValueInputID(elemPrefix)
{
	return elemPrefix + "_valueInput"
}

function selectComponentLabelInputID(elemPrefix) {
	return elemPrefix + "_labelInput"	
}

function getSelectComponentSelectableValues() {
	var selectableVals = []
	$(".seletableValuesListItem").each(function() {
		var elemPrefix = $(this).attr('id')
		
		var selectVal = $('#'+selectComponentValueInputID(elemPrefix)).val()
		var selectLabel = $('#'+selectComponentLabelInputID(elemPrefix)).val()
			
		if((selectVal != null && selectVal.length > 0) &&
			(selectLabel != null && selectLabel.length > 0)) {
			selectableVals.push({
				val: selectVal,
				label: selectLabel
			})
			
		}
	})
	
	console.log("Selectable values: " + JSON.stringify(selectableVals))
	
	return selectableVals;
}

function saveSelectComponentSelectableValues($selection,selectionRef) {
	var selectableVals = getSelectComponentSelectableValues()
	
	var saveSelectableValueParams = {
		parentFormID: selectionRef.parentFormID,
		selectionID: selectionRef.selectionID,
		selectableVals: selectableVals
	}
	
	jsonAPIRequest("frm/selection/setSelectableVals",saveSelectableValueParams,function(updatedSelection) {
		setContainerComponentInfo($selection,updatedSelection,updatedSelection.selectionID)			
	})		
	
}


function selectableValuesListItem($selection,elemPrefix,selectionRef,selectableVal) {
		
	var listItemHTML = 	'<div class="list-group-item seletableValuesListItem" id="'+elemPrefix+'"></div>';
	var $listItem = $(listItemHTML)
	
	var $selectableValForm = $('<form></form>')
	$listItem.append($selectableValForm)
	
	var $valueInputFormGroup = $('<div class="form-group marginBottom5"></div>')
	var valueInputID = selectComponentValueInputID(elemPrefix)	
	var $valueInput = $('<input type="text" class="form-control" placeholder="Value when selected"' +
									' id="'+ valueInputID + '" name="'+ valueInputID + '">')
	$valueInput.val(selectableVal.val)
	$valueInputFormGroup.append($valueInput)
		
	var $labelInputFormGroup = $('<div class="form-group marginBottom0"></div>')
	var labelInputID = selectComponentLabelInputID(elemPrefix)
	var $labelInput = $('<input type="text" class="form-control" placeholder="Label for selection"' +
									' id="'+ labelInputID + '" name="'+ labelInputID + '">')
	$labelInput.val(selectableVal.label)
	$labelInputFormGroup.append($labelInput)
	
	$selectableValForm.append($valueInputFormGroup)
	$selectableValForm.append($labelInputFormGroup)
	
	var validationRules = {}
	validationRules[valueInputID] = { required: true}
	validationRules[labelInputID] = { required: true}
	
	var validationSettings = createInlineFormValidationSettings({ 
		rules: validationRules,
		onkeyup: false
	})	
	var validator = $selectableValForm.validate(validationSettings)
	
	$valueInput.blur(function() {
		console.log("validating value input")
		saveSelectComponentSelectableValues($selection,selectionRef)
	})
	$labelInput.blur(function() {
		console.log("validating value input")
		saveSelectComponentSelectableValues($selection,selectionRef)
	})
	
	return $listItem
	
	
}

// Generate unique element IDs for the different sort rule list items.
var currSelectableValueElemPrefixNum = 0;
function generateSelectableValuePrefix() {
	currSelectableValueElemPrefixNum++;
	return "selectableVal" + currSelectableValueElemPrefixNum + "_"
}




function initSelectableValuesProperties($selection,selectionRef) {
	
	var $selectableValuesList = $('#selectableValuesList')
	
	$selectableValuesList.empty()
	
	if(selectionRef.properties.selectableVals !== null) {
		for(var selValIndex = 0; selValIndex < selectionRef.properties.selectableVals.length; selValIndex++) {
			var selectableVal = selectionRef.properties.selectableVals[selValIndex]	
			var $listItem = selectableValuesListItem($selection,generateSelectableValuePrefix(),selectionRef,selectableVal)
			$selectableValuesList.append($listItem)
		}		
	}
		
	initButtonClickHandler('#selectionAddSelectableValueButton',function(e) {
		console.log("add selectable button clicked")
		var initialEmptySelectableVal = { val:"", label:""}
		var $listItem = selectableValuesListItem($selection,generateSelectableValuePrefix(),selectionRef,initialEmptySelectableVal) 
		$selectableValuesList.append($listItem)
	})
	
	
}