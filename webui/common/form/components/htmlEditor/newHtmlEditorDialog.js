// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function openNewHtmlEditorDialog(databaseID,formID,containerParams)
{
				
	function createNewHtmlEditor($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/htmlEditor/new",newComponentParams,function(newHtmlEditorObjectRef) {
	          console.log("saveNewHtmlEditor: Done getting new ID:response=" + JSON.stringify(newHtmlEditorObjectRef));
		  	  
			  
	  		  var componentLabel = getFieldRef(newHtmlEditorObjectRef.properties.fieldID).name
			  containerParams.containerObj.find('label').text(componentLabel)
			  	  
	  		  var newComponentSetupParams = {
				  parentFormID: formID,
	  		  	  $container: containerParams.containerObj,
				  componentID: newHtmlEditorObjectRef.htmlEditorID,
				  componentObjRef: newHtmlEditorObjectRef,
				  designFormConfig: htmlEditorDesignFormConfig
	  		  }
			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
			  		  			  	
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
		
	var newFormComponentDialogParams = {
		elemPrefix: "htmlEditor_",
		databaseID: databaseID,
		formID: formID,
		hideCreateCalcFieldCheckbox: true,
		fieldTypes: [fieldTypeLongText],
		containerParams: containerParams,
		createNewFormComponent: createNewHtmlEditor
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)

		
} // newLayoutContainer

function initNewHtmlEditorDialog() {
}