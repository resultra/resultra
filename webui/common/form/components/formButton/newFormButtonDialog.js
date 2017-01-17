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
					  $(placeholderSelector).find('.formButton').text(buttonLabel)
					  $(placeholderSelector).attr("id",newButtonObjectRef.buttonID)

					  // Set up the newly created checkbox for resize, selection, etc.
					  var componentIDs = { formID: formID, componentID:newButtonObjectRef.buttonID }
	  
					  initFormComponentDesignBehavior(componentIDs,newButtonObjectRef,formButtonDesignFormConfig)
  
					  // Put a reference to the check box's reference object in the check box's DOM element.
					  // This reference can be retrieved later for property setting, etc.
					  setElemObjectRef(newButtonObjectRef.buttonID,newButtonObjectRef)
					  
					  componentCreated = true
					  $newFormButtonDialog.modal('hide')
  			  
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