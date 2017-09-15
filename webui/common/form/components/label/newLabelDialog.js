function openNewLabelDialog(databaseID,formID,containerParams) {
	
	function createNewLabel($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/label/new",newComponentParams,function(newLabelObjectRef) {
	          console.log("createNewLabel: Done getting new ID:response=" + 
						JSON.stringify(newLabelObjectRef));
			  
			  var componentLabel = getFieldRef(newLabelObjectRef.properties.fieldID).name
			  containerParams.containerObj.find('label').text(componentLabel)

	  		  var newComponentSetupParams = {
				  parentFormID: formID,
	  		  	  $container: containerParams.containerObj,
				  componentID: newLabelObjectRef.labelID,
				  componentObjRef: newLabelObjectRef,
				  designFormConfig: labelDesignFormConfig
	  		  }
			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
			  		  			  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
		
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "label_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeTag],
		containerParams: containerParams,
		createNewFormComponent: createNewLabel
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
}

function initLabelDialog() {
	
}