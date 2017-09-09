

function openNewFileDialog(databaseID,formID,containerParams)
{
		
	function createNewFile($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/file/new",newComponentParams,function(newFileObjectRef) {
	          console.log("saveNewFile: Done getting new ID:response=" + JSON.stringify(newFileObjectRef));
		  		  
			  // The new text box has been saved on the server, but only a placeholder of the text box 
			  // is currently shown in the layout. The following code is needed to update and finalize the placeholder
			  // as a complete and fully-functional text box.
			  
			  var fieldName = getFieldRef(newFileObjectRef.properties.fieldID).name
			  containerParams.containerObj.find('label').text(fieldName)			  	
			  	  
  	  		  var newComponentSetupParams = {
  				  parentFormID: formID,
  	  		  	  $container: containerParams.containerObj,
  				  componentID: newFileObjectRef.fileID,
  				  componentObjRef: newFileObjectRef,
  				  designFormConfig: fileDesignFormConfig
  	  		  }
  			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
		
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "file_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeFile],
		containerParams: containerParams,
		createNewFormComponent: createNewFile
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
			
} // newLayoutContainer

function initNewFileDialog() {
}


