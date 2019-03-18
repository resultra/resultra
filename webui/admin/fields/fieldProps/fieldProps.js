// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initFieldPropsSettingsPageContent(fieldID) {
	
	
	function initFieldNameProperties(fieldInfo) {
	
		var $nameInput = $('#fieldPropsNameInput')
		
		$nameInput.blur() // don't auto focus
	
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
	
	function initFieldRefNameProperties(fieldInfo) {
	
		var $nameInput = $('#fieldPropsRefNameInput')
	
		var $nameForm = $('#fieldRefNamePropertyForm')
		
		$nameInput.val(fieldInfo.refName)
		
		var remoteValidationParams = {
			url: '/api/field/validateExistingFieldRefName',
			data: {
				fieldID: function() { return fieldInfo.fieldID },
				fieldRefName: function() { return $nameInput.val() }
			}	
		}
	
		var validationSettings = createInlineFormValidationSettings({
			rules: {
				fieldPropsRefNameInput: {
					minlength: 3,
					required: true,
					remote: remoteValidationParams
				} // newRoleNameInput
			}
		})	
	
	
		var validator = $nameForm.validate(validationSettings)
	
		initInlineInputValidationOnBlur(validator,'#fieldPropsRefNameInput',
			remoteValidationParams, function(validatedName) {		
				var setNameParams = {
					fieldID:fieldID,
					newFieldRefName:validatedName
				}
				jsonAPIRequest("field/setRefName",setNameParams,function(updatedFieldInfo) {
					console.log("Done changing field name: " + validatedName)
				})
		})	

		validator.resetForm()
	
	}
	
	
		
	var getFieldParams = { fieldID: fieldID }
	jsonAPIRequest("field/get",getFieldParams,function(fieldInfo) {
		initFieldNameProperties(fieldInfo)
		initFieldRefNameProperties(fieldInfo)
		
		
		var $formulaEditorProperty = $('#adminCalcFieldFormulaProperty')
		if (fieldInfo.isCalcField) {
			$formulaEditorProperty.show()
			function showFormulaEditPane() { /* no-op */ }
			function hideFormulaEditPanel() { /* no-op */ }
			var formulaEditorParams = {
				databaseID: fieldInfo.parentDatabaseID,
				showEditorFunc:showFormulaEditPane,
				hideEditorFunc:hideFormulaEditPanel
			}
			initFormulaEditor(formulaEditorParams)
			openFormulaEditor(fieldInfo)
			
		} else {
			$formulaEditorProperty.hide()
		}
			
	}) // set record's number field value
	
	var $valueListLink = $('#fieldPropsBackToFieldListLink')
	$valueListLink.click(function(e) {
		e.preventDefault()
		$valueListLink.blur()
		navigateToSettingsPageContent("fields")	
	})	
		
		
		
	
}