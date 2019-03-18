// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.



function openNewRatingDialog(databaseID,formID,containerParams) {
	
	function createNewRating($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/rating/new",newComponentParams,function(newRatingObjectRef) {
	          console.log("createNewRating: Done getting new ID:response=" + JSON.stringify(newRatingObjectRef));
	    			  			  
			  var $ratingContainer 
			  
			  initRatingFormComponentContainer(containerParams.containerObj,
				  	newRatingObjectRef)			  		  
			  
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