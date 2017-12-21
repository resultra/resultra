function initFieldSelectionDropdown(clientParams) {
	
	var defaultParams = {
		includeCalcFields: true
	}
	var params = $.extend({},defaultParams,clientParams)
	
	
	function populateAddFilterDropdownMenu(sortedFields) {
		
		var fieldSelectionSelector = createPrefixedSelector(params.elemPrefix,'FieldSelectionDropdownMenu')
		var $fieldSelect = $(fieldSelectionSelector)
		$fieldSelect.empty()
		
		var fieldIDAttrName = "data-fieldID"
		
		function createFieldSelectionItem(fieldInfo) {
			
			
			var $fieldLink = $('<a href="#"></a>')
			$fieldLink.attr(fieldIDAttrName,fieldInfo.fieldID)
			$fieldLink.text(fieldInfo.name)
			
			var $fieldListItem = $('<li></li>')
			$fieldListItem.append($fieldLink)
			$fieldListItem.attr(fieldIDAttrName,fieldInfo.fieldID)
			
			$fieldLink.click(function(e) {
				params.fieldSelectionCallback(fieldInfo)
		
				e.preventDefault();// prevent the default anchor functionality				
			})
			
			return $fieldListItem
		}
			
		// Populate the selection menu for selecting the field to filter on
		for (var fieldIndex in sortedFields) {
			var fieldInfo = sortedFields[fieldIndex]
			
			var includeInSelection = true
			if (fieldInfo.isCalcField && !params.includeCalcFields) {
				includeInSelection = false
			}
			
			if (includeInSelection) {
				$fieldSelect.append(createFieldSelectionItem(fieldInfo))
			}
		} // for each field	
		
				
	}
	
	loadSortedFieldInfo(params.databaseID,params.fieldTypes,function(sortedFields) {
		
		if((params.limitToFieldList !== undefined) && (params.limitToFieldList.length>0)) {
			var sortedLimitedFields = []
			var fieldLookup = new IDLookupTable(params.limitToFieldList)
			for(var fieldIndex in sortedFields) {
				var currField = sortedFields[fieldIndex]
				if (fieldLookup.hasID(currField.fieldID)) {
					sortedLimitedFields.push(currField)
				}
			}
			populateAddFilterDropdownMenu(sortedLimitedFields)
		} else {
			populateAddFilterDropdownMenu(sortedFields)
			
		}
		
	})
}