function initUserRolePropsAdminSettingsPageContent(pageContext,roleInfo) {
	
	
	function initUserRoleNameProperties(roleInfo) {
	
		var $nameInput = $('#rolePropsNameInput')
	
		var $listNameForm = $('#roleNamePropertyForm')
		
		$nameInput.val(roleInfo.roleName)
		
		var remoteValidationParams = {
			url: '/api/userRole/validateRoleName',
			data: {
				roleID: function() { return roleInfo.roleID },
				roleName: function() { return $nameInput.val() }
			}	
		}
	
		var validationSettings = createInlineFormValidationSettings({
			rules: {
				itemListPropsNameInput: {
					minlength: 3,
					required: true,
					remote: remoteValidationParams
				} // newRoleNameInput
			}
		})	
	
	
		var validator = $listNameForm.validate(validationSettings)
	
		initInlineInputValidationOnBlur(validator,'#rolePropsNameInput',
			remoteValidationParams, function(validatedName) {		
				var setNameParams = {
					roleID:roleInfo.roleID,
					newRoleName:validatedName
				}
				jsonAPIRequest("userRole/setName",setNameParams,function(listInfo) {
					console.log("Done changing list name: " + validatedName)
				})
		})	

		validator.resetForm()
	
	} // initItemListNameProperties				
		
	var getRoleParams = { roleID: roleInfo.roleID }
	jsonAPIRequest("userRole/get",getRoleParams,function(roleInfo) {
		initUserRoleNameProperties(roleInfo)		
	}) // set record's number field value
	
	initRoleListPrivProperties(roleInfo.roleID)
	initRoleDashboardPrivProperties(roleInfo.roleID)
	initRoleNewItemPrivs(roleInfo.roleID)
	initRoleAlertPrivs(roleInfo.roleID)
	initRoleCollaborators(pageContext.databaseID,roleInfo.roleID)
	
	initSettingsPageButtonLink('#rolePropsBackToRoleListLink',"roles")
	
	
}