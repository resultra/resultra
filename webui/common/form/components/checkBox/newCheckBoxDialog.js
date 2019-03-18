// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.



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