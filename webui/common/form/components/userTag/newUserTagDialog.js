// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.



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