

function openNewHtmlEditorDialog(databaseID,formID,containerParams)
{
				
	function createNewHtmlEditor($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/htmlEditor/new",newComponentParams,function(newHtmlEditorObjectRef) {
	          console.log("saveNewHtmlEditor: Done getting new ID:response=" + JSON.stringify(newHtmlEditorObjectRef));
		  

	  		  var componentLink = newHtmlEditorObjectRef.properties.componentLink

			  var componentLabel
			  if(componentLink.linkedValType == linkedComponentValTypeField) {
				  componentLabel = getFieldRef(componentLink.fieldID).name;
			  } else {
			  	componentLabel = "Global Value"
			  }
			  
			  var placeholderSelector = '#'+containerParams.containerID
			  
			  $(placeholderSelector).find('label').text(componentLabel)
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
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeLongText],
		globalTypes: [],
		containerParams: containerParams,
		createNewFormComponent: createNewHtmlEditor
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)

		
} // newLayoutContainer

function initNewHtmlEditorDialog() {
}