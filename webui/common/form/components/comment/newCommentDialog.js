// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function openNewCommentComponentDialog(databaseID,formID,containerParams)
{
		
	function createNewCommentComponent($parentDialog, newComponentParams) {
		
		var newCommentBoxParams = {
			fieldID: newComponentParams.fieldID,
			geometry: newComponentParams.geometry,
			parentFormID: newComponentParams.parentFormID
		}
		
		jsonAPIRequest("frm/comment/new",newCommentBoxParams,function(newCommentObjectRef) {
	          console.log("createNewComment: Done getting new ID:response=" + JSON.stringify(newCommentObjectRef));
		  			  
			  var fieldID = newCommentObjectRef.properties.fieldID
			  var componentLabel = getFieldRef(fieldID).name
			  containerParams.containerObj.find('label').text(componentLabel)
			  
			  
	  		  var newComponentSetupParams = {
				  parentFormID: formID,
	  		  	  $container: containerParams.containerObj,
				  componentID: newCommentObjectRef.commentID,
				  componentObjRef: newCommentObjectRef,
				  designFormConfig: commentDesignFormConfig
	  		  }
			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
			  			  				  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "comment_",
		databaseID: databaseID,
		formID: formID,
		hideCreateCalcFieldCheckbox: true,
		fieldTypes: [fieldTypeComment],
		containerParams: containerParams,
		createNewFormComponent: createNewCommentComponent
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
			
} // newLayoutContainer

function initNewCommentComponentDialog() {
}