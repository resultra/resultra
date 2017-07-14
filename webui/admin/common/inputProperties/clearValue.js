function initClearValueProps(params) {
	
	initCheckboxChangeHandler('#adminClearValueSupported', 
				params.initialVal, function (newVal) {
			
		params.setClearValueSupported(newVal)
			
	})
	
}