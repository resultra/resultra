


function openNewToggleDialog(databaseID,formID,containerParams) {
	
	function createNewToggle($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/toggle/new",newComponentParams,function(newToggleObjectRef) {
	          console.log("createNewToggle: Done getting new ID:response=" + JSON.stringify(newToggleObjectRef));
	  	  			  
			  var componentLabel = getFieldRef(newToggleObjectRef.properties.fieldID).name			  
			  containerParams.containerObj.find('span').text(componentLabel)

	  		  var newComponentSetupParams = {
				  parentFormID: formID,
	  		  	  $container: containerParams.containerObj,
				  componentID: newToggleObjectRef.toggleID,
				  componentObjRef: newToggleObjectRef,
				  designFormConfig: toggleDesignFormConfig
	  		  }
			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
		  			  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
		
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "toggle_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeBool],
		containerParams: containerParams,
		createNewFormComponent: createNewToggle
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
}

function initNewToggleDialog() {
	
}