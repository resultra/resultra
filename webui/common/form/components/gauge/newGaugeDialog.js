// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function openNewGaugeDialog(databaseID,formID,containerParams) {
	
	function createNewGaugeComponent($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/gauge/new",newComponentParams,function(newGaugeObjectRef) {
	          console.log("openNewGaugeDialog: Done getting new gauge component:response=" 
						+ JSON.stringify(newGaugeObjectRef));
						
			setGaugeComponentLabel(containerParams.containerObj, newGaugeObjectRef)
	  	  						 
  	  		  var newComponentSetupParams = {
  				  parentFormID: formID,
  	  		  	  $container: containerParams.containerObj,
  				  componentID: newGaugeObjectRef.gaugeID,
  				  componentObjRef: newGaugeObjectRef,
  				  designFormConfig: gaugeDesignFormConfig
  	  		  }
  			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
	  		  			  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
		
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "gauge_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeNumber],
		containerParams: containerParams,
		createNewFormComponent: createNewGaugeComponent
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
}