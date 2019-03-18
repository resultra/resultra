// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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