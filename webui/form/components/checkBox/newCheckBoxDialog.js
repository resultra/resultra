
var checkboxDialogSelector = "#newCheckbox"

function openNewCheckboxDialog(formID,parentTableID,containerParams)
{
	
	// Must be the same as designForm.go - this is the common prefix on all DOM element IDs to distinguish
	// checkbox related elements from other form elements.
	var checkboxElemPrefix = "checkbox_"
	var dialogProgressSelector = '#' + checkboxElemPrefix + 'NewFormElemDialogProgress'
	
	newCheckBoxParams = {
		containerParams: containerParams,
		containerCreated: false,
		placeholderID: containerParams.containerID,
		dialogBox: $( checkboxDialogSelector )
	}
		
	function saveNewCheckbox(newCheckboxDialog) {
		console.log("New checkbox: done in dialog")
		
		var newOrExistingFormInfo = getFormFormInfoByPanelID(newCheckboxDialog,createNewOrExistingFieldDialogPanelID)
		var doCreateNewField = $(newOrExistingFormInfo.newFieldRadioSelector).prop("checked")
		
		if(doCreateNewField) {
			console.log("saveNewCheckbox: New field selected - not implemented yet")
			$(newCheckboxDialog).dialog('close');
		} else {
			console.log("saveNewCheckbox: Existing field selected")
			console.log("saveNewCheckbox: getting field id from field = " + newOrExistingFormInfo.existingFieldSelection)
			
			
			var fieldID = $(newOrExistingFormInfo.existingFieldSelectionSelector).val()
			console.log("saveNewCheckbox: Selected field ID: " + fieldID)
			
			var newCheckBoxAPIParams = {
				fieldParentTableID: designFormContext.tableID,
				parentFormID: newCheckBoxParams.containerParams.parentFormID,
				geometry: newCheckBoxParams.containerParams.geometry,
				fieldID: fieldID
			}
			console.log("saveNewCheckbox: API params: " + JSON.stringify(newCheckBoxAPIParams))
			
			jsonAPIRequest("frm/checkBox/new",newCheckBoxAPIParams,function(newCheckBoxObjectRef) {
		          console.log("saveNewCheckbox: Done getting new ID:response=" + JSON.stringify(newCheckBoxObjectRef));
			  
			  	  var fieldName = getFieldRef(newCheckBoxObjectRef.properties.fieldID).name;
				  $('#'+newCheckBoxParams.placeholderID).find('label').text(fieldName)
				  $('#'+newCheckBoxParams.placeholderID).attr("id",newCheckBoxObjectRef.checkBoxID)
			  
				  // Set up the newly created checkbox for resize, selection, etc.
				  var componentIDs = { formID: formID, componentID:newCheckBoxObjectRef.checkBoxID }
				  initFormComponentDesignBehavior(componentIDs,newCheckBoxObjectRef,checkBoxDesignFormConfig)
				  
				  // Put a reference to the check box's reference object in the check box's DOM element.
				  // This reference can be retrieved later for property setting, etc.
				  setElemObjectRef(newCheckBoxObjectRef.checkBoxID,newCheckBoxObjectRef)
				  
			
				  newCheckBoxParams.containerCreated = true				  
					  
				  newCheckBoxParams.dialogBox.dialog("close")

		       }) // newLayoutContainer API request
			
			
		} // Create check box with existing field
		
	} // saveNewCheckbox()
	
	
	var newOrExistingFieldPanel = createNewOrExistingFieldPanelConfig({
		parentTableID: parentTableID,
		elemPrefix:checkboxElemPrefix,
		fieldTypes: [fieldTypeBool],
		doneIfSelectExistingField:true,
		doneFunc:saveNewCheckbox})
	var newFieldPanel = createNewFieldDialogPanelConfig(checkboxElemPrefix)
	
	openWizardDialog({
		closeFunc: function() {
			console.log("Close dialog")
			if(!newCheckBoxParams.containerCreated)
			{
			  // If the the text box creation is not complete, remove the placeholder
			  // from the canvas.
				$('#'+newCheckBoxParams.placeholderID).remove()
			}
      	},
		width: 500, height: 500,
		dialogDivID: checkboxDialogSelector,
		panels: [newOrExistingFieldPanel,newFieldPanel],
		progressDivID: dialogProgressSelector,
	})
		
} // newLayoutContainer

function initNewCheckBoxDialog() {
	// Initialize the newTextBox dialog with the minimum parameters. This is necessary
	// to hide the dialog from view when the document is initially loaded. The
	// dialog is fully re-initialized just prior to it being opened.
	initWizardDialog(checkboxDialogSelector)
}