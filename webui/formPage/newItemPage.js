$(document).ready(function() {	
	 
	initSubmitFormUILayoutPanes()
				
	initUserDropdownMenu()
	
	var submitFormParams = {
		databaseID: submitFormPageContext.databaseID,
		$parentFormCanvas: $('#newItemFormPageLayoutCanvas'),
		formLinkID: submitFormPageContext.formLinkID,
		formID: submitFormPageContext.formID
	}
	initFormPageSubmitForm(submitFormParams)
					
}); // document ready