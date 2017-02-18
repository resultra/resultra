function initItemListDisplayConfigPanel(listInfo,changeDisplayPageSizeCallback) {
	
	var $pageSizeSelection = $('#viewListPageSizeSelection')
	
	$pageSizeSelection.val(listInfo.properties.defaultPageSize)
	initNumberSelectionChangeHandler($pageSizeSelection, function(newPageSize){
		console.log("Change page size: " + newPageSize)
		changeDisplayPageSizeCallback(newPageSize)
	})
		
	
}