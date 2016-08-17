
function initDatabaseNameProperties(databaseInfo) {
	
	$('#databasePropsNameInput').val(databaseInfo.name)
	
	var $databaseNameForm = $('#databaseNamePropertyForm')
	
	
	var nameValidationParams = {
		databaseID: function() { return databaseInfo.databaseID },
		databaseName: function() { return $('#databasePropsNameInput').val() }
	}
	
	var validationSettings = createInlineFormValidationSettings({
		rules: {
			databasePropsNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/database/validateDatabaseName',
					data: nameValidationParams
				} // remote
			} // newRoleNameInput
		}
	})	
	
	
	var validator = $databaseNameForm.validate(validationSettings)
	
	$('#databasePropsNameInput').unbind("blur")
	$('#databasePropsNameInput').blur(function() {
		if(validator.element('#databasePropsNameInput')) {
			
			var newName = $('#databasePropsNameInput').val()
			
			console.log("Starting database name change (pending remote validation): " + newName)
			
			doubleCheckRemoteFormValidation('/api/database/validateDatabaseName',nameValidationParams, 
							function(validationResult) {
				
				if(validationResult == true) {
					console.log("Changing database name: " + newName)
					var setNameParams = {
						databaseID:databaseInfo.databaseID,
						newName:newName
					}
					jsonAPIRequest("database/setName",setNameParams,function(dbInfo) {
						console.log("Done changing database name: " + newName)
					})
				} else {
					console.log("Remote validation failed, aborting database name change: " + newName)
				}
			})
			
		}
	})	
	

	validator.resetForm()
	
}

function initAdminGeneralProperties(databaseID) {
	var getDBInfoParams = { databaseID: databaseID }
	jsonAPIRequest("database/getInfo",getDBInfoParams,function(dbInfo) {
		initDatabaseNameProperties(dbInfo.databaseInfo)
	})
	
}