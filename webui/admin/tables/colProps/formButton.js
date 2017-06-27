function initFormButtonColPropertiesImpl(formButtonCol) {
	
//	setColPropsHeader(formButtonCol)
	
	var elemPrefix = "button_"
	hideSiblingsShowOne("#formButtonColProps")
	
	function initColorSchemeProperties() {
		var $schemeSelection = $('#adminButtonComponentColorSchemeSelection')
		$schemeSelection.val(formButtonCol.properties.colorScheme)
		initSelectControlChangeHandler($schemeSelection,function(newScheme) {
		
			var sizeParams = {
				parentTableID: formButtonCol.parentTableID,
				buttonID: formButtonCol.buttonID,
				colorScheme: newScheme
			}
			jsonAPIRequest("tableView/formButton/setColorScheme",sizeParams,function(updatedButton) {
			})
		
		})
		
	}
	initColorSchemeProperties()	
	
	
	function initIconProperties() {
		var $iconSelection = $('#adminButtonComponentIconSelection')
		$iconSelection.val(formButtonCol.properties.icon)
		initSelectControlChangeHandler($iconSelection,function(newIcon) {
		
			var iconParams = {
				parentTableID: formButtonCol.parentTableID,
				buttonID: formButtonCol.buttonID,
				icon: newIcon
			}
			jsonAPIRequest("tableView/formButton/setIcon",iconParams,function(updatedButton) {
			})
		
		})
		
	}
	initIconProperties()
	
	function initButtonSizeProperties() {
		var $sizeSelection = $('#adminButtonComponentSizeSelection')
		$sizeSelection.val(formButtonCol.properties.size)
		initSelectControlChangeHandler($sizeSelection,function(newSize) {
		
			var sizeParams = {
				parentTableID: formButtonCol.parentTableID,
				buttonID: formButtonCol.buttonID,
				size: newSize
			}
			jsonAPIRequest("tableView/formButton/setSize",sizeParams,function(updatedButton) {
			})		
		})
		
	}
	initButtonSizeProperties()
	
	function saveBehaviorProperties(popupBehavior) {
		var setPopupBehaviorParams = { 
				parentTableID: formButtonCol.parentTableID,
				buttonID: formButtonCol.buttonID,
			popupBehavior: popupBehavior
		}
		jsonAPIRequest("tableView/formButton/setPopupBehavior",setPopupBehaviorParams,function(updatedButtonRef) {
		})
	}
	initFormButtonBehaviorProperties(formButtonCol,saveBehaviorProperties)
	
	
	
	var elemPrefix = "button_"
	
	var defaultValPropParams = {
		databaseID: colPropsAdminContext.databaseID,
		elemPrefix: elemPrefix,
		defaultDefaultValues: formButtonCol.properties.defaultValues,
		updateDefaultValues: function(updatedDefaultVals) {
			console.log("Updating default values for form button: " + JSON.stringify(updatedDefaultVals))
			
			var setDefaultValsParams = {
				parentTableID: formButtonCol.parentTableID,
				buttonID: formButtonCol.buttonID,
				defaultValues: updatedDefaultVals }
			
			jsonAPIRequest("tableView/formButton/setDefaultVals",setDefaultValsParams,function(updatedButtonRef) {
			})
		}
	}
	initDefaultValuesPropertyPanel(defaultValPropParams)
	
	
}


function initFormButtonColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		buttonID: columnID
	}
	jsonAPIRequest("tableView/formButton/get", getColParams, function(formButtonCol) { 
		initFormButtonColPropertiesImpl(formButtonCol)
	})
}