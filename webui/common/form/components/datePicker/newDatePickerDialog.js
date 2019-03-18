// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function openNewDatePickerDialog(databaseID,formID,containerParams)
{
		
	function createNewDatePicker($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/datePicker/new",newComponentParams,function(newDatePickerObjectRef) {
	          console.log("saveNewDatePicker: Done getting new ID:response=" + JSON.stringify(newDatePickerObjectRef));
		  
			  var componentLabel = getFieldRef(newDatePickerObjectRef.properties.fieldID).name
			  containerParams.containerObj.find('label').text(componentLabel)
		  
	  		  var newComponentSetupParams = {
				  parentFormID: formID,
	  		  	  $container: containerParams.containerObj,
				  componentID: newDatePickerObjectRef.datePickerID,
				  componentObjRef: newDatePickerObjectRef,
				  designFormConfig: datePickerDesignFormConfig
	  		  }
			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
			  				  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "datePicker_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeTime],
		containerParams: containerParams,
		createNewFormComponent: createNewDatePicker
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
			
} // newLayoutContainer

function initNewDatePickerDialog() {
}