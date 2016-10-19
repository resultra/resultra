


function openNewRatingDialog(databaseID,formID,parentTableID,containerParams) {
	
	function createNewRating($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/rating/new",newComponentParams,function(newRatingObjectRef) {
	          console.log("createNewRating: Done getting new ID:response=" + JSON.stringify(newRatingObjectRef));
	  
	  		  var componentLink = newRatingObjectRef.properties.componentLink
	  
			  var componentLabel
			  if(componentLink.linkedValType == linkedComponentValTypeField) {
				  componentLabel = getFieldRef(componentLink.fieldID).name;
			  } else {
			  	componentLabel = "Global Value"
			  }
	  			  
			  var placeholderSelector = '#'+containerParams.containerID
			  
			  $(placeholderSelector).find('label').text(componentLabel)
			  $(placeholderSelector).attr("id",newRatingObjectRef.ratingID)
	  
			  // Set up the newly created checkbox for resize, selection, etc.
			  var componentIDs = { formID: formID, 
				  componentID:newRatingObjectRef.ratingID }
			  initFormComponentDesignBehavior(componentIDs,newRatingObjectRef,ratingDesignFormConfig)
		  
			  // Put a reference to the check box's reference object in the check box's DOM element.
			  // This reference can be retrieved later for property setting, etc.
			  setElemObjectRef(newRatingObjectRef.ratingID,newRatingObjectRef)
		  			  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
		
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "rating_",
		parentTableID: parentTableID,
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeNumber],
		globalTypes: [globalTypeNumber],
		containerParams: containerParams,
		createNewFormComponent: createNewRating
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
}

function initNewRatingDialog() {
	
}