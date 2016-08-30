

function openNewImageDialog(databaseID,formID,parentTableID,containerParams)
{
			
	function createNewImageComponent($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/image/new",newComponentParams,function(newImageObjectRef) {
	          console.log("saveNewImage: Done getting new ID:response=" + JSON.stringify(newImageObjectRef));
		  
			  var fieldName = getFieldRef(newImageObjectRef.properties.fieldID).name

			  var placeholderSelector = '#'+containerParams.containerID

			  $(placeholderSelector).find('label').text(fieldName)
			  $(placeholderSelector).attr("id",newImageObjectRef.imageID)
		  
			  // Set up the newly created editor for resize, selection, etc.
			  var componentIDs = { formID: formID, componentID:newImageObjectRef.imageID }
			  initFormComponentDesignBehavior(componentIDs,newImageObjectRef,imageDesignFormConfig)
			  
			  // Put a reference to the check box's reference object in the check box's DOM element.
			  // This reference can be retrieved later for property setting, etc.
			  setElemObjectRef(newImageObjectRef.imageID,newImageObjectRef)
			  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "image_",
		parentTableID: parentTableID,
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeFile],
		containerParams: containerParams,
		createNewFormComponent: createNewImageComponent
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
		
} // newLayoutContainer

function initNewImageDialog() {
}