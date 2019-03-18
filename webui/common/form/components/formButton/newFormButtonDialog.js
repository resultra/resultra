// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function openNewFormButtonDialog(databaseID,formID,containerParams) {
	console.log("New form button dialog")
	
		var $newFormButtonDialogForm = $('#newFormButtonDialogForm')
	
		var validator = $newFormButtonDialogForm.validate({
			rules: {
				newFormButtonFormLinkSelection: {
					required: true,
				}, // newFormButtonFormLinkSelection
			},
			messages: {
				newFormButtonFormLinkSelection: {
					required: "Selection of a popup form is required"
				}
			}
		})

		validator.resetForm()
		
		var selectFormParams = {
			menuSelector: "#newFormButtonFormLinkSelection",
			parentDatabaseID: databaseID
		}	
		populateFormSelectionMenu(selectFormParams)
		
		var $newFormButtonDialog = $('#newFormButtonDialog')
		var placeholderSelector = '#'+containerParams.containerID
		
		var componentCreated = false
		$newFormButtonDialog.modal('show')
	
		initButtonClickHandler('#newFormButtonSaveButton',function() {
			console.log("New form button save button clicked")
			if($newFormButtonDialogForm.valid()) {	
				
				var newButtonParams = {
					parentFormID: formID,
					geometry: containerParams.geometry,
					linkedFormID: $("#newFormButtonFormLinkSelection").val() }
				
				jsonAPIRequest("frm/formButton/new",newButtonParams,function(newButtonObjectRef) {
	
			          console.log("create new form button: Done getting new ID:response=" + JSON.stringify(newButtonObjectRef));
			  
					  var buttonLabel = "TBD - Button Label"
					  
					  jsonAPIRequest("frm/getFormInfo", { formID: newButtonObjectRef.properties.linkedFormID }, function(formInfo) {
						  containerParams.containerObj.find(".formButton").text(formInfo.form.name)		

	  
				  		  var newComponentSetupParams = {
							  parentFormID: formID,
				  		  	  $container: containerParams.containerObj,
							  componentID: newButtonObjectRef.buttonID,
							  componentObjRef: newButtonObjectRef,
							  designFormConfig: formButtonDesignFormConfig
				  		  }
						  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)

					  
						  componentCreated = true
						  $newFormButtonDialog.modal('hide')
					  })
					    
  			  
			    }) // newLayoutContainer API request
			
			}
		})
		
		$newFormButtonDialog.unbind("hidden.bs.modal")
		$newFormButtonDialog.on("hidden.bs.modal", function () {
		    // put your default event here
			if(!componentCreated) {
				console.log("Cancel new button component creation: removing placholder component = " 
									+ placeholderSelector)
				$(placeholderSelector).remove()				
			}
		});
		
		
	
	
	
}