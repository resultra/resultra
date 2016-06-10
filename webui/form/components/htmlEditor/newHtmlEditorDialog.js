
var htmlEditorDialogSelector = "#newHtmlEditor"

function openNewHtmlEditorDialog(formID,parentTableID,containerParams)
{
	
	// Must be the same as designForm.go - this is the common prefix on all DOM element IDs to distinguish
	// checkbox related elements from other form elements.
	var htmlEditorElemPrefix = "htmlEditor_"
	var dialogProgressSelector = '#' + htmlEditorElemPrefix + 'NewFormElemDialogProgress'
	
	newHtmlEditorParams = {
		containerParams: containerParams,
		containerCreated: false,
		placeholderID: containerParams.containerID,
		dialogBox: $( htmlEditorDialogSelector )
	}
		
	function saveNewHtmlEditor(newHtmlEditorDialog) {
		console.log("New html editor: done in dialog")
		
		var newOrExistingFormInfo = getFormFormInfoByPanelID(newHtmlEditorDialog,createNewOrExistingFieldDialogPanelID)
		var doCreateNewField = $(newOrExistingFormInfo.newFieldRadioSelector).prop("checked")

		if(doCreateNewField) {
			console.log("saveNewHtmlEditor: New field selected - not implemented yet")
			$(newHtmlEditorDialog).dialog('close');
		} else {
			console.log("saveNewHtmlEditor: Existing field selected")
			console.log("saveNewHtmlEditor: getting field id from field = " + newOrExistingFormInfo.existingFieldSelection)
			var fieldID = $(newOrExistingFormInfo.existingFieldSelectionSelector).val()
			console.log("saveNewHtmlEditor: Selected field ID: " + fieldID)
			
			var newHtmlEditorAPIParams = {
				fieldParentTableID: designFormContext.tableID,
				parentFormID: newHtmlEditorParams.containerParams.parentFormID,
				geometry: newHtmlEditorParams.containerParams.geometry,
				fieldID: fieldID
			}
			console.log("saveNewHtmlEditor: API params: " + JSON.stringify(newHtmlEditorAPIParams))
			
			jsonAPIRequest("frm/htmlEditor/new",newHtmlEditorAPIParams,function(newHtmlEditorObjectRef) {
		          console.log("saveNewHtmlEditor: Done getting new ID:response=" + JSON.stringify(newHtmlEditorObjectRef));
			  
				  var fieldName = getFieldRef(newHtmlEditorObjectRef.properties.fieldID).name
				  $('#'+newHtmlEditorParams.placeholderID).find('label').text(fieldName)
				  $('#'+newHtmlEditorParams.placeholderID).attr("id",newHtmlEditorObjectRef.htmlEditorID)
			  
				  // Set up the newly created editor for resize, selection, etc.
				  var componentIDs = { formID: formID, componentID:newHtmlEditorObjectRef.htmlEditorID }
				  initFormComponentDesignBehavior(componentIDs,newHtmlEditorObjectRef,htmlEditorDesignFormConfig)
				  
				  // Put a reference to the check box's reference object in the check box's DOM element.
				  // This reference can be retrieved later for property setting, etc.
				  setElemObjectRef(newHtmlEditorObjectRef.htmlEditorID,newHtmlEditorObjectRef)
				  
			
				  newHtmlEditorParams.containerCreated = true				  
					  
				  newHtmlEditorParams.dialogBox.dialog("close")

		       }) // newLayoutContainer API request
			
			
		} // Create check box with existing field
		
	} // saveNewCheckbox()
		
	var newOrExistingFieldPanel = createNewOrExistingFieldPanelConfig({
		parentTableID: parentTableID,
		elemPrefix:htmlEditorElemPrefix,
		fieldTypes: [fieldTypeLongText],
		doneIfSelectExistingField:true,
		doneFunc:saveNewHtmlEditor})
	var newFieldPanel = createNewFieldDialogPanelConfig(htmlEditorElemPrefix)
	
	openWizardDialog({
		closeFunc: function() {
			console.log("Close dialog")
			if(!newHtmlEditorParams.containerCreated)
			{
			  // If the the text box creation is not complete, remove the placeholder
			  // from the canvas.
				$('#'+newHtmlEditorParams.placeholderID).remove()
			}
      	},
		width: 500, height: 500,
		dialogDivID: htmlEditorDialogSelector,
		panels: [newOrExistingFieldPanel,newFieldPanel],
		progressDivID: dialogProgressSelector,
	})
		
} // newLayoutContainer

function initNewHtmlEditorDialog() {
	// Initialize the newTextBox dialog with the minimum parameters. This is necessary
	// to hide the dialog from view when the document is initially loaded. The
	// dialog is fully re-initialized just prior to it being opened.
	initWizardDialog(htmlEditorDialogSelector)
}