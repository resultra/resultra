function initDashboardPropertiesDashboardName(dashboardInfo) {
	
	$('#dashboardPropsDashboardNameInput').val(dashboardInfo.name)
	
	var $dashboardNameForm = $('#dashboardNamePropertyPanelForm')
	
	var remoteValidationParams = {
		url: '/api/dashboard/validateDashboardName',
		data: {
			dashboardID: function() { return dashboardInfo.dashboardID },
			dashboardName: function() { return $('#dashboardPropsDashboardNameInput').val(); }
		}
	}
	
	var validationSettings = createInlineFormValidationSettings({
		rules: {
			dashboardPropsDashboardNameInput: {
				minlength: 3,
				required: true,
				remote:  remoteValidationParams
			}
		},
	})	
	var validator = $dashboardNameForm.validate(validationSettings)
	
	
	initInlineInputValidationOnBlur(validator,'#dashboardPropsDashboardNameInput',
			remoteValidationParams, function(validatedName) {
				jsonAPIRequest("dashboard/setName",{dashboardID:dashboardInfo.dashboardID,
							newName:validatedName},function(formInfo) {
					console.log("Done changing dashboard name: " + validatedName)
				})		
	})
		

	validator.resetForm()
	
}