// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function openNewEmailAddrDialog(databaseID,formID,containerParams)
{
		
	function createNewEmailAddr($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/emailAddr/new",newComponentParams,function(newEmailAddrObjectRef) {
	          console.log("saveNewEmailAddr: Done getting new ID:response=" + JSON.stringify(newEmailAddrObjectRef));
		  		  
			  // The new text box has been saved on the server, but only a placeholder of the text box 
			  // is currently shown in the layout. The following code is needed to update and finalize the placeholder
			  // as a complete and fully-functional text box.
			  
			  var fieldName = getFieldRef(newEmailAddrObjectRef.properties.fieldID).name
			  containerParams.containerObj.find('label').text(fieldName)			  	
			  	  
  	  		  var newComponentSetupParams = {
  				  parentFormID: formID,
  	  		  	  $container: containerParams.containerObj,
  				  componentID: newEmailAddrObjectRef.emailAddrID,
  				  componentObjRef: newEmailAddrObjectRef,
  				  designFormConfig: emailAddrDesignFormConfig
  	  		  }
  			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
		
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "emailAddr_",
		databaseID: databaseID,
		formID: formID,
		hideCreateCalcFieldCheckbox: true,
		fieldTypes: [fieldTypeEmail],
		containerParams: containerParams,
		createNewFormComponent: createNewEmailAddr
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
			
} // newLayoutContainer

function initNewEmailAddrDialog() {
}


