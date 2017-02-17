

function openNewSelectionDialog(databaseID,formID,containerParams)
{
		
	function createNewSelection($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/selection/new",newComponentParams,function(newSelectionObjectRef) {
	          console.log("createNewSelection: Done getting new ID:response=" + JSON.stringify(newSelectionObjectRef));
		  
			  // The new text box has been saved on the server, but only a placeholder of the text box 
			  // is currently shown in the layout. The following code is needed to update and finalize the placeholder
			  // as a complete and fully-functional text box.

			  var placeholderSelector = '#'+containerParams.containerID
			  
			  var fieldName = getFieldRef(newSelectionObjectRef.properties.fieldID).name
			  $(placeholderSelector).find('label').text(fieldName)			  	
				  			  
			  $(placeholderSelector).attr("id",newSelectionObjectRef.selectionID)
		  
			  // Set up the newly created checkbox for resize, selection, etc.
			  var componentIDs = { formID: formID, componentID: newSelectionObjectRef.selectionID}
			  initFormComponentDesignBehavior(containerParams.containerObj,componentIDs,newSelectionObjectRef,selectionDesignFormConfig)

			  // Put a reference to the check box's reference object in the check box's DOM element.
			  // This reference can be retrieved later for property setting, etc.
			  setElemObjectRef(newSelectionObjectRef.selectionID,newSelectionObjectRef)
		
			  $parentDialog.modal("hide")
			  
			  // TODO -  Now that the text box has been finalized, the layout containing the text box needs to be saved as well.
			  containerParams.finalizeLayoutIncludingNewComponentFunc()

	       }) // newLayoutContainer API request
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "selection_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeText,fieldTypeNumber],
		containerParams: containerParams,
		createNewFormComponent: createNewSelection
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
			
} // newLayoutContainer

function initNewSelectionDialog() {
}


