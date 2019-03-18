// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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