


function openNewUserSelectionDialog(databaseID,formID,parentTableID,containerParams) {
	
	function createNewUserSelection($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/userSelection/new",newComponentParams,function(newUserSelectionObjectRef) {
	          console.log("createNewUserSelection: Done getting new ID:response=" + 
						JSON.stringify(newUserSelectionObjectRef));
	  
	  		  var componentLink = newUserSelectionObjectRef.properties.componentLink
	  
			  var componentLabel
			  if(componentLink.linkedValType == linkedComponentValTypeField) {
				  componentLabel = getFieldRef(componentLink.fieldID).name;
			  } else {
			  	componentLabel = "Global Value"
			  }
	  			  
			  var placeholderSelector = '#'+containerParams.containerID
			  
			  $(placeholderSelector).find('label').text(componentLabel)
			  $(placeholderSelector).attr("id",newUserSelectionObjectRef.userSelectionID)
	  
			  // Set up the newly created checkbox for resize, selection, etc.
			  var componentIDs = { formID: formID, 
				  componentID:newUserSelectionObjectRef.userSelectionID }
			  initFormComponentDesignBehavior(componentIDs,newUserSelectionObjectRef,userSelectionDesignFormConfig)
		  
			  // Put a reference to the check box's reference object in the check box's DOM element.
			  // This reference can be retrieved later for property setting, etc.
			  setElemObjectRef(newUserSelectionObjectRef.userSelectionID,newUserSelectionObjectRef)
		  			  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
		
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "userSelection_",
		parentTableID: parentTableID,
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeUser],
		globalTypes: [globalTypeUser],
		containerParams: containerParams,
		createNewFormComponent: createNewUserSelection
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
}

function initUserSelectionDialog() {
	
}