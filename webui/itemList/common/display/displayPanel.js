function initItemListDisplayConfigPanel(listInfo,changeDisplayPageSizeCallback,changeFormCallback) {
	
	var $pageSizeSelection = $('#viewListPageSizeSelection')
	
	$pageSizeSelection.val(listInfo.properties.defaultPageSize)
	initNumberSelectionChangeHandler($pageSizeSelection, function(newPageSize){
		console.log("Change page size: " + newPageSize)
		changeDisplayPageSizeCallback(newPageSize)
	})
		
	var $displayFormSelection = $('#viewItemListFormSelection')
	
	var limitSelectionToFormIDs = listInfo.properties.alternateForms.slice(0)
	limitSelectionToFormIDs.push(listInfo.formID)
	var selectFormParams = {
		menuSelector: '#viewItemListFormSelection',
		parentDatabaseID: listInfo.parentDatabaseID,
		initialFormID: listInfo.formID,
		limitToFormIDs:limitSelectionToFormIDs
	}	
	populateFormSelectionMenu(selectFormParams)
	
	initSelectControlChangeHandler($displayFormSelection, function(newFormID) {
		changeFormCallback(newFormID)
	})
	
	
}