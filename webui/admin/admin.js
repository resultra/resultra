// Javascript for admin page

$(document).ready(function() {	
	
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }


	$('#adminPage').layout({
			inset: zeroPaddingInset,
			north: fixedUILayoutPaneParams(40),
			west: {
				size: 250,
				resizable:false,
				slidable: false,
				spacing_open:4,
				spacing_closed:4,
				initClosed:false // panel is initially open	
			}
		})
		
	var tocConfig = {
		databaseID: adminContext.databaseID,
		newItemFormButtonFunc: openSubmitFormDialog
	}
		
	initDatabaseTOC(tocConfig)
		
	initUserDropdownMenu()
		
	// Initialize the different settings panels
	initUserRoleSettings(adminContext.databaseID)
	initUserListSettings(adminContext.databaseID)
	initAdminFormSettings(adminContext.databaseID)
	initAdminListSettings(adminContext.databaseID)
	initAdminFieldSettings(adminContext.databaseID)
	initAdminDashboardSettings(adminContext.databaseID)
	initAdminGeneralProperties(adminContext.databaseID)
	initAdminGlobals(adminContext.databaseID)
	
})