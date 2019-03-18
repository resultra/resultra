// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function initDatabaseNameProperties(databaseInfo) {
	
	$('#databasePropsNameInput').val(databaseInfo.name)
	
	var $databaseNameForm = $('#databaseNamePropertyForm')
		
	var remoteValidationParams = {
		url: '/api/database/validateDatabaseName',
		data: {
			databaseID: function() { return databaseInfo.databaseID },
			databaseName: function() { return $('#databasePropsNameInput').val() }
		}	
	}
	
	var validationSettings = createInlineFormValidationSettings({
		rules: {
			databasePropsNameInput: {
				minlength: 3,
				required: true,
				remote: remoteValidationParams
			} // newRoleNameInput
		}
	})	
	
	
	var validator = $databaseNameForm.validate(validationSettings)
	
	
	$('#databasePropsNameInput').blur() // Prevent auto-focus
	initInlineInputValidationOnBlur(validator,'#databasePropsNameInput',
		remoteValidationParams, function(validatedName) {		
			var setNameParams = {
				databaseID:databaseInfo.databaseID,
				newName:validatedName
			}
			jsonAPIRequest("database/setName",setNameParams,function(dbInfo) {
				console.log("Done changing database name: " + validatedName)
			})
	})	

	validator.resetForm()
	
}

function initAdminGeneralProperties(databaseID) {
	var getDBInfoParams = { databaseID: databaseID }
	jsonAPIRequest("database/getInfo",getDBInfoParams,function(dbInfo) {
		initDatabaseNameProperties(dbInfo.databaseInfo)
		initTrackerDescriptionPropertyPanel(dbInfo.databaseInfo)
		initActiveTrackerPropertyPanel(dbInfo.databaseInfo)
	})
	initSaveTemplateProperties(databaseID)
	
}