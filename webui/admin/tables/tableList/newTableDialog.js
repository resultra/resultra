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
	
	var tableCreated = false
	initButtonClickHandler('#newTableSaveButton',function() {
		console.log("New table save button clicked")
		if($newTableDialogForm.valid()) {	
			if (tableCreated === false) {
				tableCreated = true // only allow the creation of one table from the table dialog
				var newTableParams = { 
					parentDatabaseID: pageContext.databaseID, 
					name: $tableNameInput.val() }
				jsonAPIRequest("tableView/new",newTableParams,function(newTableInfo) {
					console.log("Created new table: " + JSON.stringify(newTableInfo))
					$newTableDialog.modal('hide')

					navigateToTablePropsPage(pageContext,newTableInfo)
				})
			}			
			

		}
	})
	
	
}