

// Initialized below in newLayoutContainer
var newTextBoxParams;

var textBoxProgressDivID = '#newTextBoxProgress'

var textBoxDialogSelector = "#newTextBox"


function newFieldIsCalcField() {
	return $(newFieldPanelConfig.divID).form('get field','isCalcField').prop('checked')
}

function doCreateNewFieldWithTextBox() {
	
	return $(newOrExistingFieldPanelConfig.divID).form('get field','createNewFieldRadio').prop('checked')
}



function saveNewTextBox()
{
	// Complete the parameters to create the new textBox by including the 
	// fieldID selected in the dialog box.
	
	if(doCreateNewFieldWithTextBox()) {
		console.log("saveNewTextBox() with create new field - not implemented yet, closing dialog")
		newTextBoxParams.dialogBox.dialog("close")
	} else {
		var fieldID = $( "#newTextBoxDlgSelectOrNewFieldPanel" ).form('get value','textBoxFieldSelection')
		console.log("saveNewTextBox: Selected field ID: " + fieldID)
					
		var newTextBoxAPIParams = {
			parentID: newTextBoxParams.containerParams.parentLayoutID,
			geometry: newTextBoxParams.containerParams.geometry,
			fieldID: fieldID
		}

		// TODO - saveNewTextBox() depends on containerParams - need to pass in somehow
		jsonAPIRequest("frm/textBox/new",newTextBoxAPIParams,function(newTextBoxObjectRef) {
	          console.log("Done getting new ID:response=" + JSON.stringify(newTextBoxObjectRef));
			  
			  $('#'+newTextBoxParams.placeholderID).find('label').text(newTextBoxObjectRef.fieldRef.fieldInfo.name)
			  $('#'+newTextBoxParams.placeholderID).attr("id",newTextBoxObjectRef.uniqueID.objectID)
			  
			  // Set up the newly created checkbox for resize, selection, etc.
			  initFormComponentDesignBehavior(newTextBoxObjectRef,textBoxDesignFormConfig)
			
			  newTextBoxParams.containerCreated = true				  
					  
			  newTextBoxParams.dialogBox.dialog("close")

	       }) // newLayoutContainer API request
		
	} // An existing field was selected
} // save new text box


function openNewTextBoxDialog(containerParams)
{
	
	// Must be the same as designForm.go - this is the common prefix on all DOM element IDs to distinguish
	// checkbox related elements from other form elements.
	var textBoxElemPrefix = "textBox_"
	var dialogProgressSelector = '#' + textBoxElemPrefix + 'NewFormElemDialogProgress'
	
	newTextBoxParams = {
		containerParams: containerParams,
		containerCreated: false,
		placeholderID: containerParams.containerID,
		dialogBox: $( textBoxDialogSelector )
	}
		
	function saveNewTextBox(newTextBoxDialog) {
		console.log("New textBox: done in dialog")
		
		var newOrExistingFormInfo = getFormFormInfoByPanelID(newTextBoxDialog,createNewOrExistingFieldDialogPanelID)
		if($(newOrExistingFormInfo.panelSelector).form('get field',newOrExistingFormInfo.newFieldRadio).prop('checked')) {
			console.log("saveNewTextBox: New field selected - not implemented yet")
			$(newCheckboxDialog).dialog('close');
		} else {
			console.log("saveNewTextBox: Existing field selected")
			console.log("saveNewTextBox: getting field id from field = " + newOrExistingFormInfo.existingFieldSelection)
			var fieldID = $(newOrExistingFormInfo.panelSelector).form('get value',newOrExistingFormInfo.existingFieldSelection)
			console.log("saveNewTextBox: Selected field ID: " + fieldID)
			
			var newTextBoxAPIParams = {
				parentID: newTextBoxParams.containerParams.parentLayoutID,
				geometry: newTextBoxParams.containerParams.geometry,
				fieldID: fieldID
			}
			console.log("saveNewTextBox: API params: " + JSON.stringify(newTextBoxAPIParams))
			
			jsonAPIRequest("frm/textBox/new",newTextBoxAPIParams,function(newTextBoxObjectRef) {
		          console.log("saveNewTextBox: Done getting new ID:response=" + JSON.stringify(newTextBoxObjectRef));
			  
				  $('#'+newTextBoxParams.placeholderID).find('label').text(newTextBoxObjectRef.fieldRef.fieldInfo.name)
				  $('#'+newTextBoxParams.placeholderID).attr("id",newTextBoxObjectRef.uniqueID.objectID)
			  
				  // Set up the newly created checkbox for resize, selection, etc.
				  initFormComponentDesignBehavior(newTextBoxObjectRef,textBoxDesignFormConfig)
			
				  newTextBoxParams.containerCreated = true				  
					  
				  newTextBoxParams.dialogBox.dialog("close")

		       }) // newLayoutContainer API request
			
			
		} // Create check box with existing field
		
	} // saveNewCheckbox()
	
	
	var newOrExistingFieldPanel = createNewOrExistingFieldPanelConfig({
		elemPrefix:textBoxElemPrefix,
		fieldTypes: [fieldTypeText,fieldTypeNumber],
		doneIfSelectExistingField:true,
		doneFunc:saveNewTextBox})
	var newFieldPanel = createNewFieldDialogPanelConfig(textBoxElemPrefix)
	
	openWizardDialog({
		closeFunc: function() {
			console.log("Close dialog")
			if(!newTextBoxParams.containerCreated)
			{
			  // If the the text box creation is not complete, remove the placeholder
			  // from the canvas.
				$('#'+newTextBoxParams.placeholderID).remove()
			}
      	},
		width: 500, height: 500,
		dialogDivID: textBoxDialogSelector,
		panels: [newOrExistingFieldPanel,newFieldPanel],
		progressDivID: dialogProgressSelector,
	})
		
} // newLayoutContainer

function initNewTextBoxDialog() {
	// Initialize the newTextBox dialog with the minimum parameters. This is necessary
	// to hide the dialog from view when the document is initially loaded. The
	// dialog is fully re-initialized just prior to it being opened.
	initWizardDialog(textBoxDialogSelector)
}


