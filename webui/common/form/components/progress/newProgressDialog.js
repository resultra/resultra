function openNewProgressDialog(databaseID,formID,containerParams) {
	
	function createNewProgressComponent($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/progress/new",newComponentParams,function(newProgressObjectRef) {
	          console.log("createNewProgressComponent: Done getting new progress component:response=" 
						+ JSON.stringify(newProgressObjectRef));
	  	  
				var componentLabel = getFieldRef(newProgressObjectRef.properties.fieldID).name		
				containerParams.containerObj.find('label').text(componentLabel)
						
	  
			  // Set up the newly created checkbox for resize, selection, etc.
			  var componentIDs = { formID: formID, componentID:newProgressObjectRef.progressID }
			  initFormComponentDesignBehavior(containerParams.containerObj,componentIDs,
				  		newProgressObjectRef,progressDesignFormConfig)
		  
			  // Put a reference to the check box's reference object in the check box's DOM element.
			  // This reference can be retrieved later for property setting, etc.
			  setContainerComponentInfo(containerParams.containerObj,newProgressObjectRef,newProgressObjectRef.progressID)
		  			  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
		
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "progress_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeNumber],
		containerParams: containerParams,
		createNewFormComponent: createNewProgressComponent
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
}