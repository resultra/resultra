


function openNewRatingDialog(databaseID,formID,containerParams) {
	
	function createNewRating($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/rating/new",newComponentParams,function(newRatingObjectRef) {
	          console.log("createNewRating: Done getting new ID:response=" + JSON.stringify(newRatingObjectRef));
	    			  
			  var placeholderSelector = '#'+containerParams.containerID
			  
	  		  var componentLabel = getFieldRef(newRatingObjectRef.properties.fieldID).name
			  $(placeholderSelector).find('label').text(componentLabel)
			  $(placeholderSelector).attr("id",newRatingObjectRef.ratingID)
	  
			  // Set up the newly created checkbox for resize, selection, etc.
			  var componentIDs = { formID: formID, 
				  componentID:newRatingObjectRef.ratingID }
			  initFormComponentDesignBehavior(containerParams.containerObj,componentIDs,newRatingObjectRef,ratingDesignFormConfig)
		  
			  // Put a reference to the check box's reference object in the check box's DOM element.
			  // This reference can be retrieved later for property setting, etc.
			  setElemObjectRef(newRatingObjectRef.ratingID,newRatingObjectRef)
		  			  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
		
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "rating_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeNumber],
		containerParams: containerParams,
		createNewFormComponent: createNewRating
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
}

function initNewRatingDialog() {
	
}