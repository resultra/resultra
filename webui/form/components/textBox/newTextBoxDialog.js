

function openNewTextBoxDialog(formID,parentTableID,containerParams)
{
		
	function createNewTextBox($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/textBox/new",newComponentParams,function(newTextBoxObjectRef) {
	          console.log("saveNewTextBox: Done getting new ID:response=" + JSON.stringify(newTextBoxObjectRef));
		  
			  // The new text box has been saved on the server, but only a placeholder of the text box 
			  // is currently shown in the layout. The following code is needed to update and finalize the placeholder
			  // as a complete and fully-functional text box.
			  var fieldName = getFieldRef(newTextBoxObjectRef.properties.fieldID).name

			  var placeholderSelector = '#'+containerParams.containerID

			  $(placeholderSelector).find('label').text(fieldName)
			  $(placeholderSelector).attr("id",newTextBoxObjectRef.textBoxID)
		  
			  // Set up the newly created checkbox for resize, selection, etc.
			  var componentIDs = { formID: formID, componentID: newTextBoxObjectRef.textBoxID}
			  initFormComponentDesignBehavior(componentIDs,newTextBoxObjectRef,textBoxDesignFormConfig)

			  // Put a reference to the check box's reference object in the check box's DOM element.
			  // This reference can be retrieved later for property setting, etc.
			  setElemObjectRef(newTextBoxObjectRef.textBoxID,newTextBoxObjectRef)
		
			  $parentDialog.modal("hide")
			  
			  // TODO -  Now that the text box has been finalized, the layout containing the text box needs to be saved as well.
			  containerParams.finalizeLayoutIncludingNewComponentFunc()

	       }) // newLayoutContainer API request
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "textBox_",
		parentTableID: parentTableID,
		formID: formID,
		fieldTypes: [fieldTypeText,fieldTypeNumber],
		containerParams: containerParams,
		createNewFormComponent: createNewTextBox
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
			
} // newLayoutContainer

function initNewTextBoxDialog() {
}


