


function openNewUserSelectionDialog(databaseID,formID,containerParams) {
	
	function createNewUserSelection($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/userSelection/new",newComponentParams,function(newUserSelectionObjectRef) {
	          console.log("createNewUserSelection: Done getting new ID:response=" + 
						JSON.stringify(newUserSelectionObjectRef));
	  	  			  
			  var placeholderSelector = '#'+containerParams.containerID
			  
			  var componentLabel = getFieldRef(newUserSelectionObjectRef.properties.fieldID).name
			  $(placeholderSelector).find('label').text(componentLabel)
			  $(placeholderSelector).attr("id",newUserSelectionObjectRef.userSelectionID)
	  
			  // Set up the newly created checkbox for resize, selection, etc.
			  var componentIDs = { formID: formID, 
				  componentID:newUserSelectionObjectRef.userSelectionID }
			  initFormComponentDesignBehavior(containerParams.containerObj,componentIDs,newUserSelectionObjectRef,userSelectionDesignFormConfig)
		  
			  // Put a reference to the check box's reference object in the check box's DOM element.
			  // This reference can be retrieved later for property setting, etc.
			  setContainerComponentInfo(containerParams.containerObj,newUserSelectionObjectRef,newUserSelectionObjectRef.userSelectionID)
		  			  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
		
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "userSelection_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeUser],
		containerParams: containerParams,
		createNewFormComponent: createNewUserSelection
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
}

function initUserSelectionDialog() {
	
}