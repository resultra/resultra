// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function openNewTextBoxDialog(databaseID,formID,containerParams)
{
		
	function createNewTextBox($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/textBox/new",newComponentParams,function(newTextBoxObjectRef) {
	          console.log("saveNewTextBox: Done getting new ID:response=" + JSON.stringify(newTextBoxObjectRef));
		  		  
			  // The new text box has been saved on the server, but only a placeholder of the text box 
			  // is currently shown in the layout. The following code is needed to update and finalize the placeholder
			  // as a complete and fully-functional text box.
			  			  	  
  	  		  var newComponentSetupParams = {
  				  parentFormID: formID,
  	  		  	  $container: containerParams.containerObj,
  				  componentID: newTextBoxObjectRef.textBoxID,
  				  componentObjRef: newTextBoxObjectRef,
  				  designFormConfig: textBoxDesignFormConfig
  	  		  }
  			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
		
			  initTextBoxFormComponentContainer(containerParams.containerObj,newTextBoxObjectRef)
		
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "textBox_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeText],
		containerParams: containerParams,
		createNewFormComponent: createNewTextBox
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
			
} // newLayoutContainer

function initNewTextBoxDialog() {
}


