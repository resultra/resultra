// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initValueListSettingsPageContent(valueListInfo) {
		
	function initValueListNameProperties(valueListInfo) {
	
		var $nameInput = $('#valueListPropsNameInput')
		$nameInput.blur()
	
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
	
	
	initValueListNameProperties(valueListInfo)
	initValueListValueListProperties(valueListInfo)
	
	var $valueListLink = $('#valueListPropsBackToValueListLink')
	$valueListLink.click(function(e) {
		e.preventDefault()
		$valueListLink.blur()
		navigateToSettingsPageContent("valueLists")	
	})	
	
}