// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function openNewImageDialog(databaseID,formID,containerParams)
{
		
	function createNewImage($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/image/new",newComponentParams,function(newImageObjectRef) {
	          console.log("saveNewImage: Done getting new ID:response=" + JSON.stringify(newImageObjectRef));
		  		  
			  // The new text box has been saved on the server, but only a placeholder of the text box 
			  // is currently shown in the layout. The following code is needed to update and finalize the placeholder
			  // as a complete and fully-functional text box.
			  
			  var fieldName = getFieldRef(newImageObjectRef.properties.fieldID).name
			  containerParams.containerObj.find('label').text(fieldName)			  	
			  	  
  	  		  var newComponentSetupParams = {
  				  parentFormID: formID,
  	  		  	  $container: containerParams.containerObj,
  				  componentID: newImageObjectRef.imageID,
  				  componentObjRef: newImageObjectRef,
  				  designFormConfig: imageDesignFormConfig
  	  		  }
  			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
		
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "image_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeImage],
		hideCreateCalcFieldCheckbox: true,
		containerParams: containerParams,
		createNewFormComponent: createNewImage
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
			
} // newLayoutContainer

function initNewImageDialog() {
}


