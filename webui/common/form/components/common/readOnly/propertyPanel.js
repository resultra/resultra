function initFormComponentReadOnlyPropertyPanel(params) {
	
	var checkboxSelector = '#' + params.elemPrefix + "FormComponentReadOnlyProperty"
	
	initCheckboxChangeHandler(checkboxSelector, params.initialVal, params.readOnlyPropertyChangedCallback)
	
}