$(document).ready(function() {
		
	function initValueListNameProperties(valueListInfo) {
	
		var $nameInput = $('#valueListPropsNameInput')
	
		var $nameForm = $('#valueListNamePropertyForm')
		
		$nameInput.val(valueListInfo.name)
		
		var remoteValidationParams = {
			url: '/api/generic/stringValidation/validateItemLabel',
			data: {
				label: function() { return $nameInput.val(); }
			}
		}
	
		var validationSettings = createInlineFormValidationSettings({
			rules: {
				valueListPropsNameInput: {
					minlength: 3,
					required: true,
					remote: remoteValidationParams
				}
			}
		})	
	
	
		var validator = $nameForm.validate(validationSettings)
	
		initInlineInputValidationOnBlur(validator,'#valueListPropsNameInput',
			remoteValidationParams, function(validatedName) {		
				var setNameParams = {
					valueListID:valueListInfo.valueListID,
					newName:validatedName
				}
				jsonAPIRequest("valueList/setName",setNameParams,function(updatedLinkInfo) {
					console.log("Done changing value list name: " + validatedName)
				})
		})	

		validator.resetForm()
	
	
	} // initFormLinkNameProperties
	
	
	
	
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }

	$('#editValueListPropsPage').layout({
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
		
	initAdminSettingsTOC(valueListPropsContext.databaseID)
		
	initUserDropdownMenu()
		
		var formLinkElemPrefix = "valueList_"
		
		var getValueListParams = { valueListID: valueListPropsContext.valueListID }
		jsonAPIRequest("valueList/get",getValueListParams,function(valueListInfo) {
			
			initValueListNameProperties(valueListInfo)
			initValueListValueListProperties(valueListInfo)
	
		})
	
})