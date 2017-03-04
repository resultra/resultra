

function openNewHtmlEditorDialog(databaseID,formID,containerParams)
{
				
	function createNewHtmlEditor($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/htmlEditor/new",newComponentParams,function(newHtmlEditorObjectRef) {
	          console.log("saveNewHtmlEditor: Done getting new ID:response=" + JSON.stringify(newHtmlEditorObjectRef));
		  	  
			  
	  		  var componentLabel = getFieldRef(newHtmlEditorObjectRef.properties.fieldID).name
			  containerParams.containerObj.find('label').text(componentLabel)
		  
			  // Set up the newly created editor for resize, selection, etc.
			  var componentIDs = { formID: formID, componentID:newHtmlEditorObjectRef.htmlEditorID }
			  initFormComponentDesignBehavior(containerParams.containerObj,componentIDs,newHtmlEditorObjectRef,htmlEditorDesignFormConfig)
			  
			  // Put a reference to the check box's reference object in the check box's DOM element.
			  // This reference can be retrieved later for property setting, etc.
			  setContainerComponentInfo(containerParams.containerObj,newHtmlEditorObjectRef,newHtmlEditorObjectRef.htmlEditorID)
			  	
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
		
	var newFormComponentDialogParams = {
		elemPrefix: "htmlEditor_",
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