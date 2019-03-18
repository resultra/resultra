// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function openNewGlobalDialog(databaseID) {
	
	var $newGlobalDialogForm = $('#adminNewGlobalForm')
	var $newGlobalDialog = $('#adminNewGlobalDialog')
	var $nameInput = $('#adminGlobalNewGlobalNameInput')
	var $refNameInput = $('#adminGlobalNewGlobalReferenceNameInput')
	var $typeSelection = $("#adminGlobalNewGlobalTypeSelection")
	
	var validator = $newGlobalDialogForm.validate({
		rules: {
			adminGlobalNewGlobalNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/global/validateNewName',
					data: {
						databaseID: databaseID,
						globalName: function() { return $nameInput.val(); }
					}
				} // remote
			}, // new
			adminGlobalNewGlobalReferenceNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/global/validateNewReferenceName',
					data: {
						databaseID: databaseID,
						refName: function() { return $refNameInput.val(); }
					}
				} // remote		
			},
			adminGlobalNewGlobalTypeSelection: { optionSelectionRequired:"type" }
		},
		messages: {
			adminGlobalNewGlobalNameInput: {
				required: "Global name is required"
			},
			adminGlobalNewGlobalReferenceNameInput: {
				required: "Reference name is required"
			}
		}
	})

	$nameInput.val("")
	$typeSelection.val("")
	validator.resetForm()
	
	
	$newGlobalDialog.modal('show')

	initButtonClickHandler('#newGlobalDialogSaveGlobalButton',function() {
		console.log("New global save button clicked")
		if($newGlobalDialogForm.valid()) {				
			var newGlobalParams = { 
				parentDatabaseID:databaseID, 
				name: $nameInput.val(),
				refName: $refNameInput.val(),
				type: $typeSelection.val()}
			jsonAPIRequest("global/new",newGlobalParams,function(newGlobalInfo) {
				console.log("Created new global: " + JSON.stringify(newGlobalInfo))
				addGlobalToAdminList(newGlobalInfo)
				$newGlobalDialog.modal('hide')
			})
		}
	})
	
}