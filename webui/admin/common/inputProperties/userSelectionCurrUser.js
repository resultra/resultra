

function initSelectionCurrUserProperties(params) {
	
		initCheckboxChangeHandler('#adminUserSelectionCurrUserSelectable', 
					params.currUserSelectable, function (newVal) {
		
			params.setCurrUserSelectable(newVal)		
		
		})
}