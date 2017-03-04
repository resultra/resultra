function openNewCommentComponentDialog(databaseID,formID,containerParams)
{
		
	function createNewCommentComponent($parentDialog, newComponentParams) {
		
		var newCommentBoxParams = {
			fieldID: newComponentParams.fieldID,
			geometry: newComponentParams.geometry,
			parentFormID: newComponentParams.parentFormID
		}
		
		jsonAPIRequest("frm/comment/new",newCommentBoxParams,function(newCommentObjectRef) {
	          console.log("createNewComment: Done getting new ID:response=" + JSON.stringify(newCommentObjectRef));
		  			  
			  var fieldID = newCommentObjectRef.properties.fieldID
			  var componentLabel = getFieldRef(fieldID).name
			  containerParams.containerObj.find('label').text(componentLabel)
			  
			  
	  		  var newComponentSetupParams = {
				  parentFormID: formID,
	  		  	  $container: containerParams.containerObj,
				  componentID: newCommentObjectRef.commentID,
				  componentObjRef: newCommentObjectRef,
				  designFormConfig: commentDesignFormConfig
	  		  }
			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
			  			  				  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "comment_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeComment],
		containerParams: containerParams,
		createNewFormComponent: createNewCommentComponent
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
			
} // newLayoutContainer

function initNewCommentComponentDialog() {
}