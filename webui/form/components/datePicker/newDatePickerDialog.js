
function openNewDatePickerDialog(databaseID,formID,containerParams)
{
		
	function createNewDatePicker($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/datePicker/new",newComponentParams,function(newDatePickerObjectRef) {
	          console.log("saveNewDatePicker: Done getting new ID:response=" + JSON.stringify(newDatePickerObjectRef));
		  

			  var componentLink = newDatePickerObjectRef.properties.componentLink
		  
		  		var componentLabel
			  if(componentLink.linkedValType == linkedComponentValTypeField) {
		  	  	 componentLabel = getFieldRef(componentLink.fieldID).name
			  } else {
			  	 componentLabel = "Global Field"
			  }
			  
			  var placeholderSelector = '#'+containerParams.containerID
	
			  $(placeholderSelector).find('label').text(componentLabel)
			  $(placeholderSelector).attr("id",newDatePickerObjectRef.datePickerID)
		  
			  // Set up the newly created checkbox for resize, selection, etc.
			  var componentIDs = { formID: formID, componentID:newDatePickerObjectRef.datePickerID }
			  initFormComponentDesignBehavior(componentIDs,newDatePickerObjectRef,datePickerDesignFormConfig)
			  
			  // Put a reference to the check box's reference object in the check box's DOM element.
			  // This reference can be retrieved later for property setting, etc.
			  setElemObjectRef(newDatePickerObjectRef.datePickerID,newDatePickerObjectRef)
			  				  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "datePicker_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeTime],
		globalTypes: [globalTypeTime],
		containerParams: containerParams,
		createNewFormComponent: createNewDatePicker
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
			
} // newLayoutContainer

function initNewDatePickerDialog() {
}