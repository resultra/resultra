function openNewListDialog(databaseID) {
	
	var $newListDialogForm = $('#newListDialogForm')
	
	var validator = $newListDialogForm.validate({
		rules: {
			newListNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/itemList/validateNewListName',
					data: {
						databaseID: databaseID,
						listName: function() { return $('#newListNameInput').val(); }
					}
				} // remote
			}, // newListNameInput
			newListFormSelection: { required:true }
		},
		messages: {
			newListNameInput: {
				required: "List name is required"
			}
		}
	})

	validator.resetForm()
	
	var selectFormParams = {
		menuSelector: "#newListFormSelection",
		parentDatabaseID: databaseID
	}	
	populateFormSelectionMenu(selectFormParams)
			
	$('#newListDialog').modal('show')
	
	initButtonClickHandler('#newListSaveButton',function() {
		console.log("New list save button clicked")
		if($newListDialogForm.valid()) {	
			console.log("table selection: " + $('#newListTableSelection').val() )
			
			var newListParams = { 
				parentDatabaseID: databaseID,
				formID: $('#newListFormSelection').val(),
				name: $('#newListNameInput').val() }
			jsonAPIRequest("itemList/new",newListParams,function(newListInfo) {
				console.log("Created new list: " + JSON.stringify(newListInfo))
				addListToAdminItemListList(newListInfo)
				$('#newListDialog').modal('hide')
			})
			

		}
	})
	
}