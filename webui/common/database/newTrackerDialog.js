function openNewTrackerDialog() {
	
	var $newTrackerDialogForm = $('#newTrackerDialogForm')
	
	var validator = $newTrackerDialogForm.validate({
		rules: {
			newTrackerNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/database/validateNewTrackerName',
					data: {
						trackerName: function() { return $('#newTrackerNameInput').val(); }
					}
				} // remote
			}, // newFormNameInput
		},
		messages: {
			newTrackerNameInput: {
				required: "Tracker name is required"
			}
		}
	})

	validator.resetForm()
		
	$('#newTrackerDialog').modal('show')
	
	initButtonClickHandler('#newTrackerSaveButton',function() {
		console.log("New Tracker save button clicked")
		if($newTrackerDialogForm.valid()) {	
			
			var newTrackerParams = {  name: $('#newTrackerNameInput').val() }
			jsonAPIRequest("database/new",newTrackerParams,function(newTrackerInfo) {
				console.log("Created new tracker: " + JSON.stringify(newTrackerInfo))
				addTrackerListItem(newTrackerInfo)	
				$('#newTrackerDialog').modal('hide')
			})
			

		}
	})
	
}