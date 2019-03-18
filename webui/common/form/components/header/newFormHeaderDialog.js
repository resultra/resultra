// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function openNewFormHeaderDialog(databaseID,formID,containerParams) {
	console.log("New form header dialog")
	
	var newHeaderParams = {
		parentFormID: formID,
		geometry: containerParams.geometry,
		label: "New Header"}
	
	jsonAPIRequest("frm/header/new",newHeaderParams,function(newHeaderObjectRef) {
          console.log("create new form header: Done getting new ID:response=" + JSON.stringify(newHeaderObjectRef));
		  
		  containerParams.containerObj.find('.formHeader').text(newHeaderObjectRef.properties.label)
  
  		  var newComponentSetupParams = {
			  parentFormID: formID,
  		  	  $container: containerParams.containerObj,
			  componentID: newHeaderObjectRef.headerID,
			  componentObjRef: newHeaderObjectRef,
			  designFormConfig: formHeaderDesignFormConfig
  		  }
		  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
  
	  			  
       }) // newLayoutContainer API request
	
}