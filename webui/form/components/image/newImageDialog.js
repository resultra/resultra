
var imageDialogSelector = "#newImage"

function openNewImageDialog(formID,parentTableID,containerParams)
{
	
	// Must be the same as designForm.go - this is the common prefix on all DOM element IDs to distinguish
	// checkbox related elements from other form elements.
	var imageElemPrefix = "image_"
	var dialogProgressSelector = '#' + imageElemPrefix + 'NewFormElemDialogProgress'
	
	newImageParams = {
		containerParams: containerParams,
		containerCreated: false,
		placeholderID: containerParams.containerID,
		dialogBox: $( imageDialogSelector )
	}
		
	function saveNewImage(newImageDialog) {
		console.log("New html editor: done in dialog")
		
		var newOrExistingFormInfo = getFormFormInfoByPanelID(newImageDialog,createNewOrExistingFieldDialogPanelID)
		var doCreateNewField = $(newOrExistingFormInfo.newFieldRadioSelector).prop("checked")
		
		if(doCreateNewField) {
			console.log("saveNewImage: New field selected - not implemented yet")
			$(newImageDialog).dialog('close');
		} else {
			console.log("saveNewImage: Existing field selected")
			console.log("saveNewImage: getting field id from field = " + newOrExistingFormInfo.existingFieldSelection)
			var fieldID = $(newOrExistingFormInfo.existingFieldSelectionSelector).val()
			console.log("saveNewImage: Selected field ID: " + fieldID)
			
			var newImageAPIParams = {
				parentID: newImageParams.containerParams.parentFormID,
				geometry: newImageParams.containerParams.geometry,
				fieldID: fieldID
			}
			console.log("saveNewImage: API params: " + JSON.stringify(newImageAPIParams))
			
			jsonAPIRequest("frm/image/new",newImageAPIParams,function(newImageObjectRef) {
		          console.log("saveNewImage: Done getting new ID:response=" + JSON.stringify(newImageObjectRef));
			  
				  var fieldName = getFieldRef(newImageObjectRef.properties.fieldID).name
				  $('#'+newImageParams.placeholderID).find('label').text(fieldName)
				  $('#'+newImageParams.placeholderID).attr("id",newImageObjectRef.imageID)
			  
				  // Set up the newly created editor for resize, selection, etc.
				  var componentIDs = { formID: formID, componentID:newImageObjectRef.imageID }
				  initFormComponentDesignBehavior(componentIDs,newImageObjectRef,imageDesignFormConfig)
				  
				  // Put a reference to the check box's reference object in the check box's DOM element.
				  // This reference can be retrieved later for property setting, etc.
				  setElemObjectRef(newImageObjectRef.imageID,newImageObjectRef)
				  
			
				  newImageParams.containerCreated = true				  
					  
				  newImageParams.dialogBox.dialog("close")

		       }) // newLayoutContainer API request
			
			
		} // Create check box with existing field
		
	} // saveNewCheckbox()
		
	var newOrExistingFieldPanel = createNewOrExistingFieldPanelConfig({
		parentTableID: parentTableID,
		elemPrefix:imageElemPrefix,
		fieldTypes: [fieldTypeFile],
		doneIfSelectExistingField:true,
		doneFunc:saveNewImage})
	var newFieldPanel = createNewFieldDialogPanelConfig(imageElemPrefix)
	
	openWizardDialog({
		closeFunc: function() {
			console.log("Close dialog")
			if(!newImageParams.containerCreated)
			{
			  // If the the text box creation is not complete, remove the placeholder
			  // from the canvas.
				$('#'+newImageParams.placeholderID).remove()
			}
      	},
		width: 500, height: 500,
		dialogDivID: imageDialogSelector,
		panels: [newOrExistingFieldPanel,newFieldPanel],
		progressDivID: dialogProgressSelector,
	})
		
} // newLayoutContainer

function initNewImageDialog() {
	// Initialize the newTextBox dialog with the minimum parameters. This is necessary
	// to hide the dialog from view when the document is initially loaded. The
	// dialog is fully re-initialized just prior to it being opened.
	initWizardDialog(imageDialogSelector)
}