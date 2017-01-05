function openNewCommentComponentDialog(databaseID,formID,containerParams)
{
		
	function createNewCommentComponent($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/comment/new",newComponentParams,function(newCommentObjectRef) {
	          console.log("createNewComment: Done getting new ID:response=" + JSON.stringify(newCommentObjectRef));
		  			  
			  var fieldID = newCommentObjectRef.properties.fieldID
			  var componentLabel = getFieldRef(fieldID).name
			  
			  var placeholderSelector = '#'+containerParams.containerID
	
			  $(placeholderSelector).find('label').text(componentLabel)
			  $(placeholderSelector).attr("id",newCommentObjectRef.commentID)
		  
			  // Set up the newly created checkbox for resize, selection, etc.
			  var componentIDs = { formID: formID, componentID:newCommentObjectRef.commentID }
			  initFormComponentDesignBehavior(componentIDs,newCommentObjectRef,commentDesignFormConfig)
			  
			  // Put a reference to the check box's reference object in the check box's DOM element.
			  // This reference can be retrieved later for property setting, etc.
			  setElemObjectRef(newCommentObjectRef.commentID,newCommentObjectRef)
			  				  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "comment_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeComment],
		globalTypes: [],
		containerParams: containerParams,
		createNewFormComponent: createNewCommentComponent
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
			
} // newLayoutContainer

function initNewCommentComponentDialog() {
}