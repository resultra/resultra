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
	
	function initNewTrackerDescriptionEditor(initialDescription) {
		var $descInput = $props.find(".adminGeneralTrackerDescriptionInput")
	
		function setTrackerDescription(description) {
			var setDescParams = {
				databaseID:trackerDatabaseInfo.databaseID,
				description:description
			}
			jsonAPIRequest("database/setDescription",setDescParams,function(dbInfo) {
			})
		
		}
	
		$descInput.html(trackerDatabaseInfo.description)
	
		$descInput.dblclick(function() {
			if (!inlineCKEditorEnabled($descInput)) {
			
				var editor = enableInlineCKEditor($descInput)
				$descInput.focus()
			
				editor.on('blur', function(event) {
					var popupMsg = editor.getData();
				
					setTrackerDescription(popupMsg)
							
					disableInlineCKEditor($descInput,editor)
				})
			
			}
		})
		
	}
	
	jsonAPIRequest("database/getTemplateList",getDBListParams,function(templateList) {
		
		$templateSelection.empty()
		
		var templateTrackerInfoByID = {}
		
		for (var trackerIndex=0; trackerIndex<templateList.length; trackerIndex++) {	
			var trackerInfo = templateList[trackerIndex]
			$templateSelection.append(selectOptionHTML(trackerInfo.databaseID,trackerInfo.databaseName))
			templateTrackerInfoByID[trackerInfo.databaseID] = trackerInfo
		}
		
		initSelectControlChangeHandler($templateSelection,function(selectedDatabaseID) {
			var trackerInfo = templateTrackerInfoByID[selectedDatabaseID]
			console.log("new template tracker database selected: " + JSON.stringify(trackerInfo))
			
			var $descGroup = $('#newTrackerTemplateDescriptionGroup') 
			if(trackerInfo.description.length > 0) {
				var $templateDesc = $('#newTrackerTemplateDescription')
				$templateDesc.html(trackerInfo.description)
				$descGroup.show()
			} else {
				$descGroup.hide()
			}
			
		})
			
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