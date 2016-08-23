function initDashboardComponentTitlePropertyPanel(elemPrefix, panelParams) {
	
	var titleElemInfo = createPrefixedTemplElemInfo(elemPrefix,'DashboardPropsComponentTitleInput')
	var formSelector = createPrefixedSelector(elemPrefix,'DashboardComponentTitlePropertyPanelForm')
	
	$(titleElemInfo.selector).val(panelParams.title)
	
	var $componentTitleForm = $(formSelector)

	var remoteValidationParams = {
		url: '/api/dashboard/validateComponentTitle',
		data: {
			title: function() { return $(titleElemInfo.selector).val(); }
		}
	}
	
	var validationRules = {}
	validationRules[titleElemInfo.id] = {
				minlength: 1,
				required: true,
				remote:  remoteValidationParams
			}
	
	var validationSettings = createInlineFormValidationSettings({ rules: validationRules })	
	var validator = $componentTitleForm.validate(validationSettings)
	
	
	initInlineInputValidationOnBlur(validator,titleElemInfo.selector,
			remoteValidationParams, function(validatedTitle) {
				
				console.log("Component title validated: " + validatedTitle)
				panelParams.setTitleFunc(validatedTitle)
				/*
				jsonAPIRequest("dashboard/setName",{dashboardID:dashboardInfo.dashboardID,
							newName:validatedName},function(formInfo) {
					console.log("Done changing dashboard name: " + validatedName)
				}) */		
	})
		

	validator.resetForm()
	
}