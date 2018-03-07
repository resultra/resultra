

function loadNewItemView(params) {
		
	GlobalFormPagePrivs = "edit"
	
//	hideSiblingsShowOne("#newItemViewFooterControls")
//	params.pageLayout.disablePropertySidebar()
	
	var getNewItemInfoParams = { formLinkID: params.formLinkID }	

	jsonAPIRequest("formLink/getNewItemInfo",getNewItemInfoParams,function(newItemInfo) {
		
		params.pageLayout.setCenterContentHeader(newItemInfo.linkName)		
			
		var submitFormParams = {
			databaseID: params.databaseID,
			$parentFormCanvas: $('#newItemFormPageLayoutCanvas'),
			formLinkID: params.formLinkID,
			formID: newItemInfo.formID,
			loadLastViewCallback: params.loadLastViewCallback
		}
	
		var $addAnotherButton = $('#newItemPageAddAnotherButton')
		initButtonControlClickHandler($addAnotherButton, function() {
			initFormPageSubmitForm(submitFormParams)
		})
	
		initFormPageSubmitForm(submitFormParams)
	}) 
	
	
}