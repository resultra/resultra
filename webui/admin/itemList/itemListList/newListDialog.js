function openNewListDialog(databaseID) {
	
	var $newListDialogForm = $('#newListDialogForm')
	var $nameInput = $('#newListNameInput')
	var currViewConfig = null
	
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
			itemListViewSelection: { required:true }
		},
		messages: {
			newListNameInput: {
				required: "List name is required"
			},
			itemListViewSelection: { 
				required: "Select a form or table"
			}
		}
	})
	
	function updateViewConfig(viewOptions) {
		console.log("Setting view options for list: " + JSON.stringify(viewOptions))
		currViewConfig = viewOptions
	}
	
	var itemListViewConfig = {
		setViewCallback: updateViewConfig,
		databaseID: databaseID
	}
	initItemListViewSelection(itemListViewConfig)
	

	resetFormValidationFeedback($newListDialogForm)
	$nameInput.val("")
	validator.resetForm()
	
	var selectFormParams = {
		menuSelector: "#newListFormSelection",
		parentDatabaseID: databaseID
	}	
	populateFormSelectionMenu(selectFormParams)
			
	$('#newListDialog').modal('show')
	
	initButtonClickHandler('#newListSaveButton',function() {
		console.log("New list save button clicked")
		if($newListDialogForm.valid() && currViewConfig != null) {				
			var newListParams = { 
				parentDatabaseID: databaseID,
				defaultView: currViewConfig,
				name: $nameInput.val() }
			jsonAPIRequest("itemList/new",newListParams,function(newListInfo) {
				console.log("Created new list: " + JSON.stringify(newListInfo))
				addListToAdminItemListList(newListInfo)
				$('#newListDialog').modal('hide')
				
				var editPropsContentURL = '/admin/itemList/' + newListInfo.listID
				setSettingsPageContent(editPropsContentURL,function() {
					initItemListPropsSettingsPageContent(databaseID,listInfo)
				})
			})
			

		}
	})
	
}