// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.



function openNewSocialButtonDialog(databaseID,formID,containerParams) {
	
	function createNewSocialButton($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/socialButton/new",newComponentParams,function(newSocialButtonObjectRef) {
	          console.log("createNewSocialButton: Done getting new ID:response=" + JSON.stringify(newSocialButtonObjectRef));
	    			  			  
			  var $socialButtonContainer 
			  
			  initRatingFormComponentContainer(containerParams.containerObj,
				  	newSocialButtonObjectRef)			  		  
			  
  	  		  var newComponentSetupParams = {
  				  parentFormID: formID,
  	  		  	  $container: containerParams.containerObj,
  				  componentID: newSocialButtonObjectRef.socialButtonID,
  				  componentObjRef: newSocialButtonObjectRef,
  				  designFormConfig: socialButtonDesignFormConfig
  	  		  }
  			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
			  		  			  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
		
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "socialButton_",
		databaseID: databaseID,
		formID: formID,
		hideCreateCalcFieldCheckbox: true,
		fieldTypes: [fieldTypeUser],
		containerParams: containerParams,
		createNewFormComponent: createNewSocialButton
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
}

function initNewSocialButtonDialog() {
	
}