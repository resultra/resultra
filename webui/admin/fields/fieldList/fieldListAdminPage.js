function initFieldsSettingsPageContent(pageContext) {
		
	initAdminFieldSettings(pageContext.databaseID)
		
}

function navigateToFieldListSettingsPage(pageContext) {
	
	const contentURL = '/admin/fields/mainContent/' + pageContext.databaseID
	setSettingsPageContent(contentURL,function() {
		initFieldsSettingsPageContent(pageContext)
	})
	
	const offPageContentURL = '/admin/fields/offPageContent'
	setSettingsPageOffPageContent(offPageContentURL,function() {
	})
	
	
	
}