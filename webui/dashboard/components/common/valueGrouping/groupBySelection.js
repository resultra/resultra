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
	else if (fieldType === fieldTypeTime) {
		$selection.append(selectOptionHTML("none","Don't group values"))
		$selection.append(selectOptionHTML("monthAndYear","Month and year"))	
	}
	else {
		console.log("unrecocognized field type: " + fieldType)
	}
}
