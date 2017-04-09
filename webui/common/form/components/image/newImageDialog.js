

function openNewImageDialog(databaseID,formID,containerParams)
{
			
	function createNewImageComponent($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/image/new",newComponentParams,function(newImageObjectRef) {
	          console.log("saveNewImage: Done getting new ID:response=" + JSON.stringify(newImageObjectRef));
			  
			  setAttachmentComponentLabel(containerParams.containerObj,newImageObjectRef)
			  			   
	  		  var newComponentSetupParams = {
				  parentFormID: formID,
	  		  	  $container: containerParams.containerObj,
				  componentID: newImageObjectRef.imageID,
				  componentObjRef: newImageObjectRef,
				  designFormConfig: imageDesignFormConfig
	  		  }
			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
			  			  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "image_",
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