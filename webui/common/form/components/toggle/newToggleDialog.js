// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.



function openNewToggleDialog(databaseID,formID,containerParams) {
	
	function createNewToggle($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/toggle/new",newComponentParams,function(newToggleObjectRef) {
	          console.log("createNewToggle: Done getting new ID:response=" + JSON.stringify(newToggleObjectRef));
	  	  			  
	  		  var newComponentSetupParams = {
				  parentFormID: formID,
	  		  	  $container: containerParams.containerObj,
				  componentID: newToggleObjectRef.toggleID,
				  componentObjRef: newToggleObjectRef,
				  designFormConfig: toggleDesignFormConfig
	  		  }
			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
			  
			  initToggleComponentFormComponentContainer(containerParams.containerObj,newToggleObjectRef)
			  reInitToggleComponentControl(containerParams.containerObj,newToggleObjectRef)
		  			  
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