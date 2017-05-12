// Javascript for admin page

$(document).ready(function() {	
	
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }
	$('#trackerMainPage').layout({
			inset: zeroPaddingInset,
			north: fixedUILayoutPaneParams(40),
		})
		
	var tocConfig = {
		databaseID: trackerContext.databaseID,
		newItemFormButtonFunc: openSubmitFormDialog
	}
	initDatabaseTOC(tocConfig)
		
	initUserDropdownMenu()
			
})