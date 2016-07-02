


function openNewCheckboxDialog(formID,parentTableID,containerParams) {
	
	function createNewCheckbox($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/checkBox/new",newComponentParams,function(newCheckBoxObjectRef) {
	          console.log("createNewCheckbox: Done getting new ID:response=" + JSON.stringify(newCheckBoxObjectRef));
	  
		  	  var fieldName = getFieldRef(newCheckBoxObjectRef.properties.fieldID).name;
			  
			  var placeholderSelector = '#'+containerParams.containerID
			  
			  $(placeholderSelector).find('label').text(fieldName)
			  $(placeholderSelector).attr("id",newCheckBoxObjectRef.checkBoxID)
	  
			  // Set up the newly created checkbox for resize, selection, etc.
			  var componentIDs = { formID: formID, componentID:newCheckBoxObjectRef.checkBoxID }
			  initFormComponentDesignBehavior(componentIDs,newCheckBoxObjectRef,checkBoxDesignFormConfig)
		  
			  // Put a reference to the check box's reference object in the check box's DOM element.
			  // This reference can be retrieved later for property setting, etc.
			  setElemObjectRef(newCheckBoxObjectRef.checkBoxID,newCheckBoxObjectRef)
		  			  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
		
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "checkbox_",
		parentTableID: parentTableID,
		formID: formID,
		fieldTypes: [fieldTypeBool],
		containerParams: containerParams,
		createNewFormComponent: createNewCheckbox
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
}

function initNewCheckBoxDialog() {
	
}