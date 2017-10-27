function loadExistingItemView(pageLayout,databaseID,viewItemConfig) {
		
	GlobalFormPagePrivs = "edit"
	
	pageLayout.clearCenterContentArea()
	hideSiblingsShowOne("#existingItemViewFooterControls")
	hideSiblingsShowOne("#viewFormPageLayoutCanvas")
	pageLayout.showFooterLayout()
	pageLayout.disablePropertySidebar()
	
	pageLayout.setCenterContentHeader(viewItemConfig.title)
	
	getRecordRefAndChangeSetID(viewItemConfig,initRecordFormView)
	
	
}