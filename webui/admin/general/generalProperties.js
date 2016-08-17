
function initDashboardNameProperties(databaseInfo) {
	
	$('#databasePropsNameInput').val(databaseInfo.name)
	
	var $databaseNameForm = $('#databaseNamePropertyForm')
	
	var validator = $databaseNameForm.validate({
		rules: {
			databasePropsNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/database/validateDatabaseName',
					data: {
						databaseID: databaseInfo.databaseID,
						databaseName: $('#databasePropsNameInput').val()
					}
				} // remote
			} // newRoleNameInput
		},
		// Since there is already a value in place, the validation can be made "eager" by 
		// always triggering the validation when the key goes up. The default behavior is
		// to only trigger the first validation when the input focus is lost (blur), then
		// become more eager when an error is detected.
		onkeyup: function(element) { $(element).valid() }
	})
	
	$('#databasePropsNameInput').unbind("blur")
	$('#databasePropsNameInput').blur(function() {
		if(validator.element('#databasePropsNameInput')) {
			
			var newName = $('#databasePropsNameInput').val()
			
			console.log("Starting database name change (pending remote validation): " + newName)
			
			var validationParams = { databaseID: databaseInfo.databaseID,
						databaseName: newName }
			doubleCheckRemoteFormValidation('/api/database/validateDatabaseName',validationParams, function(validationResult) {
				
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
		initDashboardNameProperties(dbInfo.databaseInfo)
	})
	
}