function initNewItemPageLayout()
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
		
	var formLayout = $('#newItemFormPane').layout({
		north: fixedUILayoutPaneAutoSizeToFitContentsParams(),
		north__showOverflowOnHover:	true
	})
	
	initButtonClickHandler("#viewTableOfContentsMenuButton", function() {
		mainLayout.toggle("west")
	})
		
}



$(document).ready(function() {	
	 
	initNewItemPageLayout()
				
	initUserDropdownMenu()
	
	var tocConfig = {
		databaseID: submitFormPageContext.databaseID,
		newItemFormButtonFunc: openSubmitFormDialog
	}	
	initDatabaseTOC(tocConfig)
	
	
	var submitFormParams = {
		databaseID: submitFormPageContext.databaseID,
		$parentFormCanvas: $('#newItemFormPageLayoutCanvas'),
		formLinkID: submitFormPageContext.formLinkID,
		formID: submitFormPageContext.formID
	}
	initFormPageSubmitForm(submitFormParams)
					
}); // document ready