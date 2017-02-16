$(document).ready(function() {
	
	
	function initFieldNameProperties(fieldInfo) {
	
		var $nameInput = $('#fieldPropsNameInput')
	
		var $nameForm = $('#fieldNamePropertyForm')
		
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
	
		initInlineInputValidationOnBlur(validator,'#fieldPropsNameInput',
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
	$('#editFieldPropsPage').layout({
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
		
	var tocConfig = {
		databaseID: fieldPropsContext.databaseID,
		newItemFormButtonFunc: openSubmitFormDialog
	}
	initDatabaseTOC(tocConfig)
		
	initUserDropdownMenu()
		
	var getFieldParams = { fieldID: fieldPropsContext.fieldID }
	jsonAPIRequest("field/get",getFieldParams,function(fieldInfo) {
//		initUserRoleNameProperties(roleInfo)		
	}) // set record's number field value
	
})