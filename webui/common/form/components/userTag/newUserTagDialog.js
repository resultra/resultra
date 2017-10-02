


function openNewUserTagDialog(databaseID,formID,containerParams) {
	
	function createNewUserTag($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/userTag/new",newComponentParams,function(newUserTagObjectRef) {
	          console.log("createNewUserTag: Done getting new ID:response=" + 
						JSON.stringify(newUserTagObjectRef));
			  
			  var componentLabel = getFieldRef(newUserTagObjectRef.properties.fieldID).name
			  containerParams.containerObj.find('label').text(componentLabel)

	  		  var newComponentSetupParams = {
				  parentFormID: formID,
	  		  	  $container: containerParams.containerObj,
				  componentID: newUserTagObjectRef.userTagID,
				  componentObjRef: newUserTagObjectRef,
				  designFormConfig: userTagDesignFormConfig
	  		  }
			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
			  		  			  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
		
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "userTag_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeUsers],
		hideCreateCalcFieldCheckbox: true,
		containerParams: containerParams,
		createNewFormComponent: createNewUserTag
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
}

function initUserTagDialog() {
	
}