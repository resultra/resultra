
var datePickerDialogSelector = "#newDatePicker"

function openNewDatePickerDialog(formID,parentTableID,containerParams)
{
	
	// Must be the same as designForm.go - this is the common prefix on all DOM element IDs to distinguish
	// checkbox related elements from other form elements.
	var datePickerElemPrefix = "datePicker_"
	var dialogProgressSelector = '#' + datePickerElemPrefix + 'NewFormElemDialogProgress'
	
	newDatePickerParams = {
		containerParams: containerParams,
		containerCreated: false,
		placeholderID: containerParams.containerID,
		dialogBox: $( datePickerDialogSelector )
	}
		
	function saveNewDatePicker(newDatePickerDialog) {
		console.log("New date picker: done in dialog")
		
		var newOrExistingFormInfo = getFormFormInfoByPanelID(newDatePickerDialog,createNewOrExistingFieldDialogPanelID)
		var doCreateNewField = $(newOrExistingFormInfo.newFieldRadioSelector).prop("checked")
		
		
		if(doCreateNewField) {
			console.log("saveNewDatePicker: New field selected - not implemented yet")
			$(newDatePickerDialog).dialog('close');
		} else {
			console.log("saveNewDatePicker: Existing field selected")
			console.log("saveNewDatePicker: getting field id from field = " + newOrExistingFormInfo.existingFieldSelection)
			var fieldID = $(newOrExistingFormInfo.existingFieldSelectionSelector).val()			
			console.log("saveNewDatePicker: Selected field ID: " + fieldID)
			
			var newDatePickerAPIParams = {
				parentID: newDatePickerParams.containerParams.parentFormID,
				geometry: newDatePickerParams.containerParams.geometry,
				fieldID: fieldID
			}
			console.log("saveNewDatePicker: API params: " + JSON.stringify(newDatePickerAPIParams))
			
			jsonAPIRequest("frm/datePicker/new",newDatePickerAPIParams,function(newDatePickerObjectRef) {
		          console.log("saveNewDatePicker: Done getting new ID:response=" + JSON.stringify(newDatePickerObjectRef));
			  
			  	  var fieldName = getFieldRef(newDatePickerObjectRef.fieldID).name
				  $('#'+newDatePickerParams.placeholderID).find('label').text(fieldName)
				  $('#'+newDatePickerParams.placeholderID).attr("id",newDatePickerObjectRef.datePickerID)
			  
				  // Set up the newly created checkbox for resize, selection, etc.
				  var componentIDs = { formID: formID, componentID:newDatePickerObjectRef.datePickerID }
				  initFormComponentDesignBehavior(componentIDs,newDatePickerObjectRef,datePickerDesignFormConfig)
				  
				  // Put a reference to the check box's reference object in the check box's DOM element.
				  // This reference can be retrieved later for property setting, etc.
				  setElemObjectRef(newDatePickerObjectRef.datePickerID,newDatePickerObjectRef)
				  
			
				  newDatePickerParams.containerCreated = true				  
					  
				  newDatePickerParams.dialogBox.dialog("close")

		       }) // newLayoutContainer API request
			
			
		} // Create check box with existing field
		
	} // saveNewCheckbox()
		
	var newOrExistingFieldPanel = createNewOrExistingFieldPanelConfig({
		parentTableID: parentTableID,
		elemPrefix:datePickerElemPrefix,
		fieldTypes: [fieldTypeTime],
		doneIfSelectExistingField:true,
		doneFunc:saveNewDatePicker})
	var newFieldPanel = createNewFieldDialogPanelConfig(datePickerElemPrefix)
	
	openWizardDialog({
		closeFunc: function() {
			console.log("Close dialog")
			if(!newDatePickerParams.containerCreated)
			{
			  // If the the text box creation is not complete, remove the placeholder
			  // from the canvas.
				$('#'+newDatePickerParams.placeholderID).remove()
			}
      	},
		width: 500, height: 500,
		dialogDivID: datePickerDialogSelector,
		panels: [newOrExistingFieldPanel,newFieldPanel],
		progressDivID: dialogProgressSelector,
	})
		
} // newLayoutContainer

function initNewDatePickerDialog() {
	// Initialize the newTextBox dialog with the minimum parameters. This is necessary
	// to hide the dialog from view when the document is initially loaded. The
	// dialog is fully re-initialized just prior to it being opened.
	initWizardDialog(datePickerDialogSelector)
}