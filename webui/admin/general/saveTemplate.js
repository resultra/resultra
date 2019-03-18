// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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