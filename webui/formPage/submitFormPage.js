

$(document).ready(function() {	
	 
	initSubmitFormUILayoutPanes()
				
	initUserDropdownMenu()
	
	var submitFormParams = {
		databaseID: submitFormPageContext.databaseID,
		$parentFormCanvas: $('#submitFormPageLayoutCanvas'),
		formLinkID: submitFormPageContext.formLinkID,
		formID: submitFormPageContext.formID
	}
	
	var $addAnotherButton = $('#newItemPageAddAnotherButton')
	initButtonControlClickHandler($addAnotherButton, function() {
		initFormPageSubmitForm(submitFormParams)
	})
	
	
	initFormPageSubmitForm(submitFormParams)
					
}); // document ready