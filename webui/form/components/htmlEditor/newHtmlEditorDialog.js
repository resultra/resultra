

function openNewHtmlEditorDialog(databaseID,formID,parentTableID,containerParams)
{
				
	function createNewHtmlEditor($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/htmlEditor/new",newComponentParams,function(newHtmlEditorObjectRef) {
	          console.log("saveNewHtmlEditor: Done getting new ID:response=" + JSON.stringify(newHtmlEditorObjectRef));
		  
			  var fieldName = getFieldRef(newHtmlEditorObjectRef.properties.fieldID).name
			  
			  var placeholderSelector = '#'+containerParams.containerID
			  
			  $(placeholderSelector).find('label').text(fieldName)
			  $(placeholderSelector).attr("id",newHtmlEditorObjectRef.htmlEditorID)
		  
			  // Set up the newly created editor for resize, selection, etc.
			  var componentIDs = { formID: formID, componentID:newHtmlEditorObjectRef.htmlEditorID }
			  initFormComponentDesignBehavior(componentIDs,newHtmlEditorObjectRef,htmlEditorDesignFormConfig)
			  
			  // Put a reference to the check box's reference object in the check box's DOM element.
			  // This reference can be retrieved later for property setting, etc.
			  setElemObjectRef(newHtmlEditorObjectRef.htmlEditorID,newHtmlEditorObjectRef)
			  	
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
		
	var newFormComponentDialogParams = {
		elemPrefix: "htmlEditor_",
		parentTableID: parentTableID,
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeLongText],
		containerParams: containerParams,
		createNewFormComponent: createNewHtmlEditor
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)

		
} // newLayoutContainer

function initNewHtmlEditorDialog() {
}