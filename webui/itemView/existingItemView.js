// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function loadExistingItemViewPageContent(viewItemConfig) {
		
	GlobalFormPagePrivs = "edit"
	
	theMainWindowLayout.disableRHSSidebar()
	
	var contentConfig = {
		mainContentURL: "/itemView/existingItemContentLayout",
		offPageContentURL: "/itemView/existingItemOffPageContent"
	}
	setMainWindowPageContent(contentConfig,function() {
		var contentLayout = new ExistingItemContentLayout()
		contentLayout.setCenterContentHeader(viewItemConfig.title)
		getRecordRefAndChangeSetID(viewItemConfig,initRecordFormView)
	})
	
	
}