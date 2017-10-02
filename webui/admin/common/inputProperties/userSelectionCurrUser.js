

function initSelectionCurrUserProperties(params) {
	
		initCheckboxChangeHandler('#' + params.elemPrefix + 'AdminUserSelectionCurrUserSelectable', 
					params.currUserSelectable, function (newVal) {
		
			params.setCurrUserSelectable(newVal)		
		
		})
}