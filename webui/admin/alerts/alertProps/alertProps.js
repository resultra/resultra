$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#alertPropsPage'))	
	
	initUserDropdownMenu()
	initAlertHeader(alertPropsContext.databaseID)
	
	initAdminSettingsTOC(alertPropsContext.databaseID)
	
	function initAlertFormProperties(alertInfo) {
		var selectFormParams = {
			menuSelector: "#alertPropFormSelection",
			parentDatabaseID: alertPropsContext.databaseID,
			initialFormID: alertInfo.properties.formID
		}	
		populateFormSelectionMenu(selectFormParams)
		var $formSelection = $("#alertPropFormSelection")
		initSelectControlChangeHandler($formSelection, function(selectedFormID) {

			var setFormParams = {
				alertID: alertPropsContext.alertID,
				formID: selectedFormID
			}	

			jsonAPIRequest("alert/setForm",setFormParams,function(setFormParams) {
				console.log("Done setting form for alert")
			}) 
			
		})
		
	} // initItemListFormProperties
	
	function initAlertFieldProperties(alertInfo) {
		var $fieldSelection = $('#alertPropFieldSelection')
		
		loadSortedFieldInfo(alertPropsContext.databaseID,[fieldTypeText],function(sortedFields) {
			
			populateSortedFieldSelectionMenu($fieldSelection,sortedFields)
			$fieldSelection.val(alertInfo.properties.summaryFieldID)
			
			initSelectControlChangeHandler($fieldSelection, function(selectedFieldID) {

				var setFieldParams = {
					alertID: alertInfo.alertID,
					summaryFieldID: selectedFieldID
				}	

				jsonAPIRequest("alert/setSummaryField",setFieldParams,function(setFieldParams) {
					console.log("Done setting field for alert")
				}) 
			
			})
			
		})
		
		
	}
	
	function initTriggerConditionProps(alertInfo) {
		var alertConditionPrefix = "alertTriggerCondition_"
		var alertConditionPropertyPanelParams = {
			elemPrefix: alertConditionPrefix,
			databaseID: alertPropsContext.databaseID,
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
					alertID: alertPropsContext.alertID,
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
		
		var getDecodedCaptionParams = { alertID: alertInfo.alertID }
		jsonAPIRequest("alert/getDecodedCaptionMessage",getDecodedCaptionParams,function(decodedMsg) {
			editor.setValue(decodedMsg)
		})
		
		editor.setHighlightActiveLine(false);
		
		
		function populateFieldReferenceDropdown() {
			
			var $fieldSelection = $("#alertFieldRefSelection")
	 	    $fieldSelection.empty()
			$fieldSelection.append('<option value="" disabled selected>Insert Field Reference</option>')
			
			loadSortedFieldInfo(alertInfo.parentDatabaseID, [fieldTypeAll],function(sortedFields) {
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
				alertID: alertPropsContext.alertID,
				captionMessage:newCaptionMsg
			}
			jsonAPIRequest("alert/setCaptionMessage",setCaptionParams,function(updatedAlertInfo) {
				console.log("Done changing alert caption: " + updatedAlertInfo)
			})
		})
		
	}
	
	
	var getAlertParams = { 
		alertID: alertPropsContext.alertID
	}
	jsonAPIRequest("alert/get",getAlertParams,function(alertInfo) {
		initAlertFormProperties(alertInfo)
		initAlertFieldProperties(alertInfo)
		initAlertRecipientProps(alertInfo)
		initTriggerConditionProps(alertInfo)
		initNameProperties(alertInfo)
		initCaptionMessageProperty(alertInfo)
	}) 
	
	var conditionPropsParams = {
		databaseID: alertPropsContext.databaseID,
		alertID: alertPropsContext.alertID
	}
	initAlertConditionProps(conditionPropsParams)
})