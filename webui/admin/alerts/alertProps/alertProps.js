// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initAlertSettingsAdminPageContent(databaseID,alertInfo) {
		
	function initAlertFormProperties(alertInfo) {
		var selectFormParams = {
			menuSelector: "#alertPropFormSelection",
			parentDatabaseID: databaseID,
			initialFormID: alertInfo.properties.formID
		}	
		populateFormSelectionMenu(selectFormParams)
		var $formSelection = $("#alertPropFormSelection")
		initSelectControlChangeHandler($formSelection, function(selectedFormID) {

			var setFormParams = {
				alertID: alertInfo.alertID,
				formID: selectedFormID
			}	

			jsonAPIRequest("alert/setForm",setFormParams,function(setFormParams) {
				console.log("Done setting form for alert")
			}) 
			
		})
		
	} // initItemListFormProperties
	
	
	function initTriggerConditionProps(alertInfo) {
		var alertConditionPrefix = "alertTriggerCondition_"
		var alertConditionPropertyPanelParams = {
			elemPrefix: alertConditionPrefix,
			databaseID: databaseID,
			defaultFilterRules: alertInfo.properties.triggerConditions,
			initDone: function () {},
			updateFilterRules: function (updatedFilterRules) {
				var setConditionParams = {
					alertID: alertInfo.alertID,
					triggerConditions: updatedFilterRules
				}
				jsonAPIRequest("alert/setTriggerConditions",setConditionParams,function(updatedAlert) {
					console.log("Trigger conditions updated")
				}) // set record's number field value
			}
		}
		initFilterPropertyPanel(alertConditionPropertyPanelParams)
		
	}
	
	
	function initNameProperties(alertInfo) {
	
		var $nameInput = $('#alertPropsNameInput')
	
		var $nameForm = $('#alertPropsNameForm')
		
		$nameInput.val(alertInfo.name)
		
		
		var remoteValidationParams = {
			url: '/api/generic/stringValidation/validateItemLabel',
			data: {
				label: function() { return $nameInput.val(); }
			}
		}
	
		var validationSettings = createInlineFormValidationSettings({
			rules: {
				itemListPropsNameInput: {
					minlength: 3,
					required: true,
					remote: remoteValidationParams
				}
			}
		})	
	
	
		var validator = $nameForm.validate(validationSettings)
	
		initInlineInputValidationOnBlur(validator,'#alertPropsNameInput',
			remoteValidationParams, function(validatedName) {		
				var setNameParams = {
					alertID: alertInfo.alertID,
					alertName:validatedName
				}
				jsonAPIRequest("alert/setName",setNameParams,function(updatedAlertInfo) {
					console.log("Done changing alert name: " + validatedName)
				})
		})	

		validator.resetForm()
	
	
	} // initFormLinkNameProperties
	
	function initCaptionMessageProperty(alertInfo) {
		var editor = ace.edit("alertCaptionMessageEditor")
			
		// Address a console warning message on scrolling
		editor.$blockScrolling = Infinity;
		editor.setTheme("ace/theme/tomorrow_night")
		editor.setShowPrintMargin(false);
		editor.getSession().setOption('indentedSoftWrap', false); // disable indent on subsequent lines
		editor.getSession().setUseWrapMode(true); // enable line wrapping.
		editor.renderer.setShowGutter(false);
		
		var getDecodedCaptionParams = { alertID: alertInfo.alertID }
		jsonAPIRequest("alert/getDecodedCaptionMessage",getDecodedCaptionParams,function(decodedMsg) {
			editor.setValue(decodedMsg)
		})
		
		editor.setHighlightActiveLine(false);
		
		
		function populateFieldReferenceDropdown() {
			
			var $fieldSelection = $("#alertFieldRefSelection")
	 	    $fieldSelection.empty()
			$fieldSelection.append('<option value="" disabled selected>Insert Field Reference</option>')
			
			var supportedFieldReferenceTypes = [fieldTypeNumber,fieldTypeBool,fieldTypeText,fieldTypeTime]
			loadSortedFieldInfo(alertInfo.parentDatabaseID, supportedFieldReferenceTypes,function(sortedFields) {
				for (var fieldIndex in sortedFields) {
	
					var fieldInfo = sortedFields[fieldIndex]		
	
			 	   var menuItemHTML = '<option value="' + fieldInfo.refName + 
						'">' + fieldInfo.name + '</option>'
		
					console.log("Adding selection to insert formula menu:" + menuItemHTML)
			
				 	$fieldSelection.append(menuItemHTML)			

				} // for each  field
			})
	
			$fieldSelection.on('change',function() {
				var fieldRefName = $(this).find("option:selected").val();
				if(fieldRefName.length > 0) {
					editor.insert("[" + fieldRefName + "]")	
					$fieldSelection.prop('selectedIndex',0);
				}
			})
			
		}
		populateFieldReferenceDropdown()
		
		editor.on("blur",function() {
			var newCaptionMsg = editor.getValue()
			console.log("Caption message changed: " + newCaptionMsg)
			var setCaptionParams = {
				alertID: alertInfo.alertID,
				captionMessage:newCaptionMsg
			}
			jsonAPIRequest("alert/setCaptionMessage",setCaptionParams,function(updatedAlertInfo) {
				console.log("Done changing alert caption: " + updatedAlertInfo)
			})
		})
		
	}
	
	var getAlertParams = { 
		alertID: alertInfo.alertID
	}
	jsonAPIRequest("alert/get",getAlertParams,function(alertInfo) {
		initAlertFormProperties(alertInfo)
		initAlertRecipientProps(alertInfo)
		initTriggerConditionProps(alertInfo)
		initNameProperties(alertInfo)
		initCaptionMessageProperty(alertInfo)
	}) 
	
	var conditionPropsParams = {
		databaseID: databaseID,
		alertID: alertInfo.alertID
	}
	initAlertConditionProps(conditionPropsParams)
	
	
	initSettingsPageButtonLink('#alertPropsBackToAlertListLink',"alerts")
	
	
	
	
}