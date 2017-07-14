function initClearValueProps(params) {
	
	var checkboxSelector = '#' + params.elemPrefix + 'adminClearValueSupported'
	
	initCheckboxChangeHandler(checkboxSelector, 
				params.initialVal, function (newVal) {
			
		params.setClearValueSupported(newVal)
			
	})
	
}