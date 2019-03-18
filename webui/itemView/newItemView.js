// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


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