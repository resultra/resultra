// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function openNewProgressDialog(databaseID,formID,containerParams) {
	
	function createNewProgressComponent($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/progress/new",newComponentParams,function(newProgressObjectRef) {
	          console.log("createNewProgressComponent: Done getting new progress component:response=" 
						+ JSON.stringify(newProgressObjectRef));
	  	  
				var componentLabel = getFieldRef(newProgressObjectRef.properties.fieldID).name		
				containerParams.containerObj.find('label').text(componentLabel)
						 
  	  		  var newComponentSetupParams = {
  				  parentFormID: formID,
  	  		  	  $container: containerParams.containerObj,
  				  componentID: newProgressObjectRef.progressID,
  				  componentObjRef: newProgressObjectRef,
  				  designFormConfig: progressDesignFormConfig
  	  		  }
  			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
	  		  			  
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