


function openNewCheckboxDialog(databaseID,formID,containerParams) {
	
	function createNewCheckbox($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/checkBox/new",newComponentParams,function(newCheckBoxObjectRef) {
	          console.log("createNewCheckbox: Done getting new ID:response=" + JSON.stringify(newCheckBoxObjectRef));
	  
	  		  var componentLink = newCheckBoxObjectRef.properties.componentLink
	  
			  var componentLabel
			  if(componentLink.linkedValType == linkedComponentValTypeField) {
				  componentLabel = getFieldRef(componentLink.fieldID).name;
			  } else {
			  	componentLabel = "Global Value"
			  }
	  			  
			  var placeholderSelector = '#'+containerParams.containerID
			  
			  $(placeholderSelector).find('label').text(componentLabel)
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
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeBool],
		globalTypes: [globalTypeBool],
		containerParams: containerParams,
		createNewFormComponent: createNewCheckbox
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
}

function initNewCheckBoxDialog() {
	
}