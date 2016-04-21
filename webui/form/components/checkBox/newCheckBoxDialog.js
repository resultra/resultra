
var checkboxDialogSelector = "#newCheckbox"

function openNewCheckboxDialog(containerParams)
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
	
	// Enable Semantic UI checkboxes and popups
	$('.ui.checkbox').checkbox();
	$('.ui.radio.checkbox').checkbox();	
	
	function saveNewCheckbox(newCheckboxDialog) {
		console.log("New checkbox: done in dialog")
		
		var newOrExistingFormInfo = getFormFormInfoByPanelID(newCheckboxDialog,createNewOrExistingFieldDialogPanelID)
		if($(newOrExistingFormInfo.panelSelector).form('get field',newOrExistingFormInfo.newFieldRadio).prop('checked')) {
			console.log("saveNewCheckbox: New field selected - not implemented yet")
			$(newCheckboxDialog).dialog('close');
		} else {
			console.log("saveNewCheckbox: Existing field selected")
			console.log("saveNewCheckbox: getting field id from field = " + newOrExistingFormInfo.existingFieldSelection)
			var fieldID = $(newOrExistingFormInfo.panelSelector).form('get value',newOrExistingFormInfo.existingFieldSelection)
			console.log("saveNewCheckbox: Selected field ID: " + fieldID)
			
			var newCheckBoxAPIParams = {
				parentID: newCheckBoxParams.containerParams.parentLayoutID,
				geometry: newCheckBoxParams.containerParams.geometry,
				fieldID: fieldID
			}
			console.log("saveNewCheckbox: API params: " + JSON.stringify(newCheckBoxAPIParams))
			
			jsonAPIRequest("frm/checkBox/new",newCheckBoxAPIParams,function(newCheckBoxObjectRef) {
		          console.log("saveNewCheckbox: Done getting new ID:response=" + JSON.stringify(newCheckBoxObjectRef));
			  
				  $('#'+newCheckBoxParams.placeholderID).find('label').text(newCheckBoxObjectRef.fieldRef.fieldInfo.name)
				  $('#'+newCheckBoxParams.placeholderID).attr("id",newCheckBoxObjectRef.uniqueID.objectID)
			  
				  // Store the newly created object reference in the DOM element. This is needed for follow-on
				  // property setting, resizing, etc.
				  setElemObjectRef(newCheckBoxObjectRef.uniqueID.objectID,newCheckBoxObjectRef)
			
				  newCheckBoxParams.containerCreated = true				  
					  
				  newCheckBoxParams.dialogBox.dialog("close")

		       }) // newLayoutContainer API request
			
			
		} // Create check box with existing field
		
	} // saveNewCheckbox()
	
	
	var newOrExistingFieldPanel = createNewOrExistingFieldPanelConfig({
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