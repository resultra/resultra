function openNewButtonTableColDialog(pageContext,tableRef) {
	
	var $newButtonDialogForm = $('#newButtonTableColForm')
	var $newButtonColDialog = $('#newButtonTableColDialog')	

	var validator = $newButtonDialogForm.validate({
		rules: {
			newButtonColFormLinkSelection: {
				required: true,
			}, // newFormButtonFormLinkSelection
		},
		messages: {
			newButtonColFormLinkSelection: {
				required: "Selection of a popup form is required"
			}
		}
	})
	
	var $formSelection = $('#newButtonColFormLinkSelection')

	$formSelection.val("")
	validator.resetForm()
	
	var selectFormParams = {
		menuSelector: "#newButtonColFormLinkSelection",
		parentDatabaseID: tableRef.parentDatabaseID
	}	
	populateFormSelectionMenu(selectFormParams)
	
	initButtonClickHandler('#newButtonColSaveButton',function() {
		console.log("New form button save button clicked")
		if($newButtonDialogForm.valid()) {	
			
			var newButtonParams = {
				parentTableID: tableRef.tableID,
				linkedFormID: $formSelection.val() 
			}
			
			jsonAPIRequest("tableView/formButton/new",newButtonParams,function(newButtonObjectRef) {
				navigateToNewColumnSettingsPage(pageContext,newButtonObjectRef)			
				$newButtonColDialog.modal('hide')
			})		
		}
	})
	
	
	
	$newButtonColDialog.modal("show")
}