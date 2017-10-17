function openSaveTemplateDialog(databaseID) {
	
	var $form = $('#saveTemplateDialogForm')
	var $templateNameInput = $('#saveTemplateNameInput')
	
	var validator = $form.validate({
		rules: {
			saveTemplateNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/database/validateNewTrackerName',
					data: {
						trackerName: function() { return $templateNameInput.val(); }
					}
				} // remote
			}, // newFormNameInput
		},
		messages: {
			saveTemplateNameInput: {
				required: "Template name is required"
			}
		}
	})

	resetFormValidationFeedback($form)
	$templateNameInput.val("")
	validator.resetForm()
	
	var $dialog = $('#saveTemplateDialog')
		
	$dialog.modal('show')
	
	initButtonClickHandler('#saveTemplateSaveButton',function() {
		if($form.valid()) {	
			
			var saveTemplateParams = { 
				sourceDatabaseID: databaseID, 
				newTemplateName: $templateNameInput.val() }
			jsonAPIRequest("database/saveAsTemplate",saveTemplateParams,function(saveTemplResp) {
				console.log("Created new template: " + JSON.stringify(saveTemplResp))
				$dialog.modal('hide')
			})

		}
	})
	
}


function initSaveTemplateProperties(databaseID) {
	
	initButtonClickHandler('#adminSaveTemplateButton',function() {
		console.log("Save template button clicked")
		openSaveTemplateDialog(databaseID)
	})
	
}