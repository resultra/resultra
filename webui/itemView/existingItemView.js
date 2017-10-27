function loadExistingItemView(pageLayout,databaseID,itemParams) {
		
	GlobalFormPagePrivs = "edit"
	
	pageLayout.clearCenterContentArea()
	hideSiblingsShowOne("#existingItemViewFooterControls")
	hideSiblingsShowOne("#existingItemFormPageLayoutCanvas")
	pageLayout.showFooterLayout()
	pageLayout.disablePropertySidebar()
	
	$('#existingItemFormPageLayoutCanvas').text("Content TBD")
	pageLayout.setCenterContentHeader("Header TBD for Item View")
	
/*	var getNewItemInfoParams = { formLinkID: formLinkID }	

	jsonAPIRequest("formLink/getNewItemInfo",getNewItemInfoParams,function(newItemInfo) {
		
		pageLayout.setCenterContentHeader(newItemInfo.linkName)		
			
		var submitFormParams = {
			databaseID: databaseID,
			$parentFormCanvas: $('#newItemFormPageLayoutCanvas'),
			formLinkID: formLinkID,
			formID: newItemInfo.formID
		}
	
		var $addAnotherButton = $('#newItemPageAddAnotherButton')
		initButtonControlClickHandler($addAnotherButton, function() {
			initFormPageSubmitForm(submitFormParams)
		})
	
		initFormPageSubmitForm(submitFormParams)
	}) 
*/
	
}