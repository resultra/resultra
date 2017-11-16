function openNewTrackerDialog() {
	
	var $newTrackerDialogForm = $('#newTrackerDialogForm')
	var $templateSelection = $('#newTrackerTemplateSelection')
	
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
	
	var getDBListParams = {} // no parameters necessary - gets the tracker list for the currently signed in user
	
	jsonAPIRequest("database/getTemplateList",getDBListParams,function(templateList) {
		
		$templateSelection.empty()
		
		for (var trackerIndex=0; trackerIndex<templateList.length; trackerIndex++) {	
			var trackerInfo = templateList[trackerIndex]
			$templateSelection.append(selectOptionHTML(trackerInfo.databaseID,trackerInfo.databaseName))
		}
			
	})
	
	
	
		
	$('#newTrackerDialog').modal('show')
	
	initButtonClickHandler('#newTrackerSaveButton',function() {
		console.log("New Tracker save button clicked")
		if($newTrackerDialogForm.valid()) {	
			
			var newTrackerParams = {  
				name: $('#newTrackerNameInput').val(),
				templateDatabaseID: $templateSelection.val()
			}
			jsonAPIRequest("database/new",newTrackerParams,function(newTrackerInfo) {
				console.log("Created new tracker: " + JSON.stringify(newTrackerInfo))
				addTrackerListItem(newTrackerInfo)	
				$('#newTrackerDialog').modal('hide')
			})
			

		}
	})
	
}