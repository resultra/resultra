


function openNewUserSelectionDialog(databaseID,formID,containerParams) {
	
	function createNewUserSelection($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/userSelection/new",newComponentParams,function(newUserSelectionObjectRef) {
	          console.log("createNewUserSelection: Done getting new ID:response=" + 
						JSON.stringify(newUserSelectionObjectRef));
			  
			  var componentLabel = getFieldRef(newUserSelectionObjectRef.properties.fieldID).name
			  containerParams.containerObj.find('label').text(componentLabel)

	  		  var newComponentSetupParams = {
				  parentFormID: formID,
	  		  	  $container: containerParams.containerObj,
				  componentID: newUserSelectionObjectRef.userSelectionID,
				  componentObjRef: newUserSelectionObjectRef,
				  designFormConfig: userSelectionDesignFormConfig
	  		  }
			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
			  		  			  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
		
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "userSelection_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeUser],
		hideCreateCalcFieldCheckbox: true,
		containerParams: containerParams,
		createNewFormComponent: createNewUserSelection
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
}

function initUserSelectionDialog() {
	
}