var formatParams = {
	setOffLabel: function(newLabel) {	
		var labelParams = {
			parentFormID: toggleRef.parentFormID,
			toggleID: toggleRef.toggleID,
			label: newLabel
		}
		jsonAPIRequest("frm/toggle/setOffLabel",labelParams,function(updatedToggleRef) {
			reInitToggleComponentControl($container,updatedToggleRef)
			setContainerComponentInfo($container,updatedToggleRef,updatedToggleRef.toggleID)		
		})
	},
	setOnLabel: function(newLabel) {
		var labelParams = {
			parentFormID: toggleRef.parentFormID,
			toggleID: toggleRef.toggleID,
			label: newLabel
		}
		jsonAPIRequest("frm/toggle/setOnLabel",labelParams,function(updatedToggleRef) {
			reInitToggleComponentControl($container,updatedToggleRef)
			setContainerComponentInfo($container,updatedToggleRef,updatedToggleRef.toggleID)		
		})
	},
	setOffColorScheme: function(newColorScheme) {
		var colorSchemeParams = {
			parentFormID: toggleRef.parentFormID,
			toggleID: toggleRef.toggleID,
			colorScheme: newColorScheme
		}
		console.log("Setting new color scheme: " + JSON.stringify(colorSchemeParams))
	
		jsonAPIRequest("frm/toggle/setOffColorScheme",colorSchemeParams,function(updatedToggleRef) {
			reInitToggleComponentControl($container,updatedToggleRef)
			setContainerComponentInfo($container,updatedToggleRef,updatedToggleRef.toggleID)		
		})
	},
	setOnColorScheme: function(newColorScheme) {
		var colorSchemeParams = {
			parentFormID: toggleRef.parentFormID,
			toggleID: toggleRef.toggleID,
			colorScheme: newColorScheme
		}
		console.log("Setting new color scheme: " + JSON.stringify(colorSchemeParams))
	
		jsonAPIRequest("frm/toggle/setOnColorScheme",colorSchemeParams,function(updatedToggleRef) {
			reInitToggleComponentControl($container,updatedToggleRef)
			setContainerComponentInfo($container,updatedToggleRef,updatedToggleRef.toggleID)	
		})		
	}
}


function initToggleFormatProperties(params) {
	
	var $form = $('#adminToggleComponentFormatForm')
	
	var $offLabel = $form.find("input[name=adminToggleOffComponentLabel]")
	$offLabel.val(toggleRef.properties.offLabel)
	var validationOffLabelParams = {
		url: '/api/generic/stringValidation/validateItemLabel',
		data: { label: function() { return $offLabel.val(); } }
	}
	function saveValidatedOffLabel(newLabel) {
		console.log("Toggle format properties: new off label: " + newLabel)
		params.setOffLabel(newLabel)
	}

	var $onLabel = $form.find("input[name=adminToggleOnComponentLabel]")
	$onLabel.val(toggleRef.properties.onLabel)
	var validationOnLabelParams = {
		url: '/api/generic/stringValidation/validateItemLabel',
		data: { label: function() { return $onLabel.val(); } }
	}
	function saveValidatedOnLabel(newLabel) {
		console.log("Toggle format properties: new on label: " + newLabel)
		params.setOnLabel(newLabel)
	}
	
	
	var validationRules = {
		adminToggleOffComponentLabel: { remote: validationOffLabelParams, required:true },
		adminToggleOnComponentLabel: { remote: validationOnLabelParams, required:true }
	}
	var validationSettings = createInlineFormValidationSettings({ rules: validationRules })	
	var validator = $form.validate(validationSettings)
	
	initInlineInputControlValidationOnBlur(validator,$offLabel,validationOffLabelParams, saveValidatedOffLabel)
	initInlineInputControlValidationOnBlur(validator,$onLabel,validationOnLabelParams, saveValidatedOnLabel)

	
	var $toggleOffColorSchemeSelection = $('#adminToggleOffComponentColorSchemeSelection')
	$toggleOffColorSchemeSelection.val(toggleRef.properties.offColorScheme)
	initSelectControlChangeHandler($toggleOffColorSchemeSelection,function(newColorScheme) {
		params.setOffColorScheme(newColorScheme)
	})
	
	var $toggleOnColorSchemeSelection = $('#adminToggleOnComponentColorSchemeSelection')
	$toggleOnColorSchemeSelection.val(toggleRef.properties.onColorScheme)
	initSelectControlChangeHandler($toggleOnColorSchemeSelection,function(newColorScheme) {
		params.setOnColorScheme(newColorScheme)
	})
	
}
