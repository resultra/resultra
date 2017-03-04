


function openNewRatingDialog(databaseID,formID,containerParams) {
	
	function createNewRating($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/rating/new",newComponentParams,function(newRatingObjectRef) {
	          console.log("createNewRating: Done getting new ID:response=" + JSON.stringify(newRatingObjectRef));
	    			  			  
	  		  var componentLabel = getFieldRef(newRatingObjectRef.properties.fieldID).name
			  containerParams.containerObj.find('label').text(componentLabel)
			  
  	  		  var newComponentSetupParams = {
  				  parentFormID: formID,
  	  		  	  $container: containerParams.containerObj,
  				  componentID: newRatingObjectRef.ratingID,
  				  componentObjRef: newRatingObjectRef,
  				  designFormConfig: ratingDesignFormConfig
  	  		  }
  			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
			  		  			  
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