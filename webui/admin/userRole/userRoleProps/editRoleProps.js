$(document).ready(function() {
	
	
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
	
	
	
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }


	$('#editRolePropsPage').layout({
			inset: zeroPaddingInset,
			north: fixedUILayoutPaneParams(40),
			west: {
				size: 250,
				resizable:false,
				slidable: false,
				spacing_open:4,
				spacing_closed:4,
				initClosed:false // panel is initially open	
			}
		})
		
	initAdminSettingsTOC(rolePropsContext.databaseID)
		
	initUserDropdownMenu()
	initAlertHeader(rolePropsContext.databaseID)
		
	var getRoleParams = { roleID: rolePropsContext.roleID }
	jsonAPIRequest("userRole/get",getRoleParams,function(roleInfo) {
		initUserRoleNameProperties(roleInfo)		
	}) // set record's number field value
	
	initRoleListPrivProperties(rolePropsContext.roleID)
	initRoleDashboardPrivProperties(rolePropsContext.roleID)
	initRoleNewItemPrivs(rolePropsContext.roleID)
	initRoleAlertPrivs(rolePropsContext.roleID)
	
})