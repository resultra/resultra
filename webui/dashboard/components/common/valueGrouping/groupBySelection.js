function populateDashboardValueGroupingSelection($selection,fieldType) {
	$selection.empty()
	$selection.append(defaultSelectOptionPromptHTML("Select a grouping"))
	if(fieldType == fieldTypeNumber) {
		$selection.append(selectOptionHTML("none","Don't group values"))
		$selection.append(selectOptionHTML("bucket","Bucket values"))
	}
	else if (fieldType === fieldTypeText) {
		$selection.append(selectOptionHTML("none","Don't group values"))
	}
	else if (fieldType === fieldTypeBool) {
		$selection.append(selectOptionHTML("none","Group into true and false"))
	}
	else if (fieldType === fieldTypeTime) {
		$selection.append(selectOptionHTML("none","Don't group values"))
		$selection.append(selectOptionHTML("day","Day"))	
		$selection.append(selectOptionHTML("week","Week"))	
		$selection.append(selectOptionHTML("monthYear","Month and year"))	
	}
	else {
		console.log("unrecocognized field type: " + fieldType)
	}
}
