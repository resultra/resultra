


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