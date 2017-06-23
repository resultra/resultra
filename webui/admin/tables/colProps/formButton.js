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