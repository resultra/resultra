
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