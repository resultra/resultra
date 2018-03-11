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