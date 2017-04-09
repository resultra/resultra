function openNewDashboardDialog(databaseID) {
	
	var $newDashboardDialogForm = $('#newDashboardDialogForm')
	var $nameInput = $('#newDashboardNameInput')
	var $newDashboardDialog = $('#newDashboardDialog')
	
	var validator = $newDashboardDialogForm.validate({
		rules: {
			newDashboardNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/dashboard/validateNewDashboardName',
					data: {
						databaseID: databaseID,
						dashboardName: function() { return $nameInput.val(); }
					}
				} // remote
			}, // newFormNameInput
		},
		messages: {
			newDashboardNameInput: {
				required: "Dashboard name is required"
			}
		}
	})

	resetFormValidationFeedback($newDashboardDialogForm)
	$nameInput.val("")
	validator.resetForm()
		
	$newDashboardDialog.modal('show')
	
	initButtonClickHandler('#newDashboardSaveButton',function() {
		console.log("New dashboard save button clicked")
		if($newDashboardDialogForm.valid()) {	
			
			var newDashboardParams = { 
				databaseID: databaseID,
				name: $nameInput.val()}
				
			jsonAPIRequest("dashboard/new",newDashboardParams,function(newDashboardInfo) {
				console.log("Created new dashboard: " + JSON.stringify(newDashboardInfo))
				navigateToURL('/admin/dashboard/' + newDashboardInfo.dashboardID)
				$newDashboardDialog.modal('hide')
			})
			

		}
	})
	
}