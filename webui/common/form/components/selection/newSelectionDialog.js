// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function openNewSelectionDialog(databaseID,formID,containerParams)
{
		
	function createNewSelection($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/selection/new",newComponentParams,function(newSelectionObjectRef) {
	          console.log("createNewSelection: Done getting new ID:response=" + JSON.stringify(newSelectionObjectRef));
		  
			  // The new text box has been saved on the server, but only a placeholder of the text box 
			  // is currently shown in the layout. The following code is needed to update and finalize the placeholder
			  // as a complete and fully-functional text box.
			  
			  var fieldName = getFieldRef(newSelectionObjectRef.properties.fieldID).name
			  containerParams.containerObj.find('label').text(fieldName)
			  
  	  		  var newComponentSetupParams = {
  				  parentFormID: formID,
  	  		  	  $container: containerParams.containerObj,
  				  componentID: newSelectionObjectRef.selectionID,
  				  componentObjRef: newSelectionObjectRef,
  				  designFormConfig: selectionDesignFormConfig
  	  		  }
  			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
			  			  			
			  $parentDialog.modal("hide")
			  
	       }) // newLayoutContainer API request
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "selection_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeText,fieldTypeNumber],
		containerParams: containerParams,
		createNewFormComponent: createNewSelection
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
			
} // newLayoutContainer

function initNewSelectionDialog() {
}


