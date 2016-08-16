function openNewDashboardDialog(databaseID) {
	
	var $newDashboardDialogForm = $('#newDashboardDialogForm')
	
	var validator = $newDashboardDialogForm.validate({
		rules: {
			newDashboardNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/dashboard/validateNewDashboardName',
					data: {
						databaseID: databaseID,
						dashboardName: function() { return $('#newDashboardNameInput').val(); }
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

	validator.resetForm()
		
	$('#newDashboardDialog').modal('show')
	
	initButtonClickHandler('#newDashboardSaveButton',function() {
		console.log("New dashboard save button clicked")
		if($newDashboardDialogForm.valid()) {	
			
			var newDashboardParams = { 
				databaseID: databaseID,
				name: $('#newDashboardNameInput').val()}
				
			jsonAPIRequest("dashboard/new",newDashboardParams,function(newDashboardInfo) {
				console.log("Created new dashboard: " + JSON.stringify(newDashboardInfo))
				$('#newDashboardDialog').modal('hide')
			})
			

		}
	})
	
}