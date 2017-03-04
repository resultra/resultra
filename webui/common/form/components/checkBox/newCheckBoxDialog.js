


function openNewCheckboxDialog(databaseID,formID,containerParams) {
	
	function createNewCheckbox($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/checkBox/new",newComponentParams,function(newCheckBoxObjectRef) {
	          console.log("createNewCheckbox: Done getting new ID:response=" + JSON.stringify(newCheckBoxObjectRef));
	  	  			  
			  var componentLabel = getFieldRef(newCheckBoxObjectRef.properties.fieldID).name			  
			  containerParams.containerObj.find('span').text(componentLabel)

	  		  var newComponentSetupParams = {
				  parentFormID: formID,
	  		  	  $container: containerParams.containerObj,
				  componentID: newCheckBoxObjectRef.checkBoxID,
				  componentObjRef: newCheckBoxObjectRef,
				  designFormConfig: checkBoxDesignFormConfig
	  		  }
			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
		  			  
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