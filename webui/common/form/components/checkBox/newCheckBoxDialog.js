


function openNewCheckboxDialog(databaseID,formID,containerParams) {
	
	function createNewCheckbox($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/checkBox/new",newComponentParams,function(newCheckBoxObjectRef) {
	          console.log("createNewCheckbox: Done getting new ID:response=" + JSON.stringify(newCheckBoxObjectRef));
	  	  			  
	  		  var newComponentSetupParams = {
				  parentFormID: formID,
	  		  	  $container: containerParams.containerObj,
				  componentID: newCheckBoxObjectRef.checkBoxID,
				  componentObjRef: newCheckBoxObjectRef,
				  designFormConfig: checkBoxDesignFormConfig
	  		  }
			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
			  
			  initCheckboxComponentFormContainer(containerParams.containerObj,newCheckBoxObjectRef)
			    			  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
		
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "checkbox_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeBool],
		containerParams: containerParams,
		createNewFormComponent: createNewCheckbox
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
}

function initNewCheckBoxDialog() {
	
}