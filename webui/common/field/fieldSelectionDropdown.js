function initFieldSelectionDropdown(params) {
	
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
		 	$fieldSelect.append(createFieldSelectionItem(fieldInfo))
		} // for each field	
		
				
	}
	
	loadSortedFieldInfo(params.databaseID,params.fieldTypes,function(sortedFields) {
		populateAddFilterDropdownMenu(sortedFields)
	})
}