// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function openNewAttachmentDialog(databaseID,formID,containerParams)
{
			
	function createNewImageComponent($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/attachment/new",newComponentParams,function(newImageObjectRef) {
	          console.log("saveNewImage: Done getting new ID:response=" + JSON.stringify(newImageObjectRef));
			  
			  setAttachmentComponentLabel(containerParams.containerObj,newImageObjectRef)
			  			   
	  		  var newComponentSetupParams = {
				  parentFormID: formID,
	  		  	  $container: containerParams.containerObj,
				  componentID: newImageObjectRef.imageID,
				  componentObjRef: newImageObjectRef,
				  designFormConfig: attachmentDesignFormConfig
	  		  }
			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
			  			  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "attachment_",
		databaseID: databaseID,
		formID: formID,
		hideCreateCalcFieldCheckbox: true,
		fieldTypes: [fieldTypeAttachment],
		containerParams: containerParams,
		createNewFormComponent: createNewImageComponent
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
		
} // newLayoutContainer

function initNewAttachmentDialog() {
}