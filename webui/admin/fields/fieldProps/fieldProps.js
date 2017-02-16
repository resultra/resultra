$(document).ready(function() {
	
	
	function initFieldNameProperties(fieldInfo) {
	
		var $nameInput = $('#fieldPropsNameInput')
	
		var $nameForm = $('#fieldNamePropertyForm')
		
		$nameInput.val(fieldInfo.name)
		
		var remoteValidationParams = {
			url: '/api/field/validateExistingFieldName',
			data: {
				fieldID: function() { return fieldInfo.fieldID },
				fieldName: function() { return $nameInput.val() }
			}	
		}
	
		var validationSettings = createInlineFormValidationSettings({
			rules: {
				fieldPropsNameInput: {
					minlength: 3,
					required: true,
					remote: remoteValidationParams
				} // newRoleNameInput
			}
		})	
	
	
		var validator = $nameForm.validate(validationSettings)
	
		initInlineInputValidationOnBlur(validator,'#fieldPropsNameInput',
			remoteValidationParams, function(validatedName) {		
				var setNameParams = {
					fieldID:fieldInfo.fieldID,
					newFieldName:validatedName
				}
				jsonAPIRequest("field/setName",setNameParams,function(updatedFieldInfo) {
					console.log("Done changing field name: " + validatedName)
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
		initFieldNameProperties(fieldInfo)		
	}) // set record's number field value
	
})