function initFieldSelectionDropdown(params) {
	
	function populateAddFilterDropdownMenu(fieldsByID) {
		
		var fieldSelectionSelector = createPrefixedSelector(params.elemPrefix,'FieldSelectionDropdownMenu')
		var $fieldSelect = $(fieldSelectionSelector)
	
		$fieldSelect.empty()
		
		var fieldIDAttrName = "data-fieldID"
	
		// Populate the selection menu for selecting the field to filter on
		for (var fieldID in fieldsByID) {
			var fieldInfo = fieldsByID[fieldID]
			
			var $fieldListItem = $('<li><a data-fieldID="'+fieldID+'" href="#">' + fieldInfo.name + '</a></li>')
		 	$fieldSelect.append($fieldListItem)
			$fieldListItem.attr(fieldIDAttrName,fieldID)
		} // for each field	
		
		$fieldSelect.find("a").click(function(e) {
			
			var fieldID = $(this).attr(fieldIDAttrName)
			
			console.log("click on field: " + fieldID)
			var fieldInfo = fieldsByID[fieldID]
			params.fieldSelectionCallback(fieldInfo)
		
			e.preventDefault();// prevent the default anchor functionality
		})
				
	}
	
	
	loadFieldInfo(params.tableID,params.fieldTypes,function(fieldsByID) {
		populateAddFilterDropdownMenu(fieldsByID)
	})
}