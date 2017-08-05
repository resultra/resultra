
function loadDatePickerProperties($container,datePickerRef) {
	console.log("Loading date picker properties")
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#datePickerProps')
	
	var formatParams = {
		initialFormat: datePickerRef.properties.dateFormat,
		setFormat: function(newFormat) {
			var formatParams = {
				parentFormID: datePickerRef.parentFormID,
				datePickerID: datePickerRef.datePickerID,
				dateFormat: newFormat
			}
			jsonAPIRequest("frm/datePicker/setFormat",formatParams,function(updatedDatePicker) {
				setContainerComponentInfo($container,updatedDatePicker,updatedDatePicker.datePickerID)	
				initDatePickerFormComponentInput($container,updatedDatePicker)
				var sampleDate = moment("2013-02-08 09:30")
				setDatePickerFormComponentDate($container, updatedDatePicker, sampleDate)
			})
		}
	}
	initDateFormatProperties(formatParams)
	
	
	var validationParams = {
		setValidationConfig: function(validationConfig) {
			var validationParams = {
				parentFormID: datePickerRef.parentFormID,
				datePickerID: datePickerRef.datePickerID,
				validation: validationConfig
			}
			jsonAPIRequest("frm/datePicker/setValidation", validationParams, function(updatedDatePicker) {
				setContainerComponentInfo($container,updatedDatePicker,updatedDatePicker.datePickerID)
			})			
		}
	
	}
	initDateValidationProperties(validationParams)
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
				parentFormID: datePickerRef.parentFormID,
				datePickerID: datePickerRef.datePickerID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/datePicker/setLabelFormat", formatParams, function(updatedDatePicker) {
			setDatePickerComponentLabel($container,updatedDatePicker)
			setContainerComponentInfo($container,updatedDatePicker,updatedDatePicker.datePickerID)	
		})	
	}
	
	var elemPrefix = "datePicker_"
	
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: datePickerRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: datePickerRef.parentFormID,
			datePickerID: datePickerRef.datePickerID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/datePicker/setVisibility",params,function(updatedDatePicker) {
			setContainerComponentInfo($container,updatedDatePicker,updatedDatePicker.datePickerID)	
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: datePickerRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: datePickerRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: datePickerRef.parentFormID,
				datePickerID: datePickerRef.datePickerID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/datePicker/setPermissions",params,function(updatedDatePicker) {
				initDatePickerAddonControls($container,updatedDatePicker)	
				setContainerComponentInfo($container,updatedDatePicker,updatedDatePicker.datePickerID)	
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	
	var clearValueParams = {
		initialVal: datePickerRef.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentFormID: datePickerRef.parentFormID,
				datePickerID: datePickerRef.datePickerID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("frm/datePicker/setClearValueSupported",formatParams,function(updatedDatePicker) {
				setContainerComponentInfo($container,updatedDatePicker,updatedDatePicker.datePickerID)
				initDatePickerAddonControls($container,updatedDatePicker)	
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: datePickerRef.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentFormID: datePickerRef.parentFormID,
				datePickerID: datePickerRef.datePickerID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("frm/datePicker/setHelpPopupMsg",params,function(updatedDatePicker) {
				setContainerComponentInfo($container,updatedDatePicker,updatedDatePicker.datePickerID)
				updateComponentHelpPopupMsg($container, updatedDatePicker)
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)


	var conditionalFormatParams = {
		initialFormats: datePickerRef.properties.conditionalFormats,
		setConditionalFormats: function(formats) {
			var params = {
				parentFormID: datePickerRef.parentFormID,
				datePickerID: datePickerRef.datePickerID,
				conditionalFormats: formats
			}
			jsonAPIRequest("frm/datePicker/setConditionalFormats",params,function(updatedDatePicker) {
				setContainerComponentInfo($container,updatedDatePicker,updatedDatePicker.datePickerID)
			})	
		}
	}
	initDateConditionalFormatPropertyPanel(conditionalFormatParams)

	
	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: datePickerRef.parentFormID,
		componentID: datePickerRef.datePickerID,
		componentLabel: 'date picker',
		$componentContainer: $container
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	toggleFormulaEditorForField(datePickerRef.properties.fieldID)
	
}