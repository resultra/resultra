function initAlertPageLayout()
{
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }
	
	// Initialize the page layout
	var mainLayout = $('#layoutPage').layout({
		inset: zeroPaddingInset,
		north: fixedUILayoutPaneParams(40),
		west: {
			size: 250,
			resizable:false,
			slidable: false,
			spacing_open:4,
			spacing_closed:4,
			initClosed:true // panel is initially closed	
		}
	})
	
	var formLayout = $('#alertsPane').layout({
		north: fixedUILayoutPaneAutoSizeToFitContentsParams(),
		north__showOverflowOnHover:	true,
	})
		
	initButtonClickHandler("#viewTableOfContentsMenuButton", function() {
		console.log("TOC button clicked")
		mainLayout.toggle("west")
	})
	
}


$(document).ready(function() {	
	 
	initAlertPageLayout()
				
	initUserDropdownMenu()
	
	var tocConfig = {
		databaseID: alertPageContext.databaseID,
		newItemFormButtonFunc: openSubmitFormDialog
	}	
	initDatabaseTOC(tocConfig)
	
						
}); // document ready