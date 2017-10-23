

function loadNewItemView(pageLayout,databaseID,formLinkID) {
		
	GlobalFormPagePrivs = "edit"
	
	pageLayout.clearCenterContentArea()
	hideSiblingsShowOne("#newItemViewFooterControls")
	hideSiblingsShowOne("#newItemFormPageLayoutCanvas")
	pageLayout.showFooterLayout()
	
	var getNewItemInfoParams = { formLinkID: formLinkID }	

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
	
}