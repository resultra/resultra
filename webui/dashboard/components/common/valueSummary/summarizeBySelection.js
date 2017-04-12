

function populateSummarizeBySelection(selectionSelector, fieldType) {
	$(selectionSelector).empty()
	$(selectionSelector).append(defaultSelectOptionPromptHTML("Choose how to summarize values"))
	if(fieldType === fieldTypeNumber) {
		$(selectionSelector).append(selectOptionHTML("count","Count of values"))
		$(selectionSelector).append(selectOptionHTML("sum","Sum of values"))
		$(selectionSelector).append(selectOptionHTML("average","Average of values"))
	} else if (fieldType === fieldTypeText) {
		$(selectionSelector).append(selectOptionHTML("count","Count of values"))
	} else if (fieldType === fieldTypeBool) {
		$(selectionSelector).append(selectOptionHTML("count","Count of values"))
		$(selectionSelector).append(selectOptionHTML("countTrue","Count of true values"))
		$(selectionSelector).append(selectOptionHTML("countFalse","Count of false values"))
		$(selectionSelector).append(selectOptionHTML("percFalse","Percentage false values"))
		$(selectionSelector).append(selectOptionHTML("percTrue","Percentage true values"))	
	} else {
		console.log("unrecocognized field type: " + fieldType)
	}
}
