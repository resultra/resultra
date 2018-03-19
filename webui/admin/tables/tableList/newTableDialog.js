function openNewTableDialog(pageContext) {
	
	var $newTableDialogForm = $('#newTableDialogForm')
	var $newTableDialog = $('#newTableDialog')	
	var $tableNameInput = $newTableDialogForm.find("input[name=newTableNameInput]")
	
	var validator = $newTableDialogForm.validate({
		rules: {
			newTableNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/tableView/validateNewTableName',
					data: {
						databaseID: pageContext.databaseID,
						tableName: function() { return $tableNameInput.val(); }
					}
				} // remote
			}, // newFormNameInput
		},
		messages: {
			newTableNameInput: {
				required: "Table name is required"
			}
		}
	})

	resetFormValidationFeedback($newTableDialogForm)
	$tableNameInput.val("")
	validator.resetForm()
	
	$newTableDialog.modal("show")
	
	initButtonClickHandler('#newTableSaveButton',function() {
		console.log("New table save button clicked")
		if($newTableDialogForm.valid()) {				
			var newTableParams = { 
				parentDatabaseID: pageContext.databaseID, 
				name: $tableNameInput.val() }
			jsonAPIRequest("tableView/new",newTableParams,function(newTableInfo) {
				console.log("Created new table: " + JSON.stringify(newTableInfo))
				$newTableDialog.modal('hide')

				navigateToTablePropsPage(pageContext,newTableInfo)
			})
			

		}
	})
	
	
}