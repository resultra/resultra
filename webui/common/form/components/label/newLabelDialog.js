// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
		elemPrefix: "tag_",
		databaseID: databaseID,
		formID: formID,
		hideCreateCalcFieldCheckbox: true,
		fieldTypes: [fieldTypeTag],
		containerParams: containerParams,
		createNewFormComponent: createNewLabel
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
}

function initLabelDialog() {
	
}