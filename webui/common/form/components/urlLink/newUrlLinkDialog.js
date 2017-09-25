

function openNewUrlLinkDialog(databaseID,formID,containerParams)
{
		
	function createNewUrlLink($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/urlLink/new",newComponentParams,function(newUrlLinkObjectRef) {
	          console.log("saveNewUrlLink: Done getting new ID:response=" + JSON.stringify(newUrlLinkObjectRef));
		  		  
			  // The new text box has been saved on the server, but only a placeholder of the text box 
			  // is currently shown in the layout. The following code is needed to update and finalize the placeholder
			  // as a complete and fully-functional text box.
			  
			  var fieldName = getFieldRef(newUrlLinkObjectRef.properties.fieldID).name
			  containerParams.containerObj.find('label').text(fieldName)			  	
			  	  
  	  		  var newComponentSetupParams = {
  				  parentFormID: formID,
  	  		  	  $container: containerParams.containerObj,
  				  componentID: newUrlLinkObjectRef.urlLinkID,
  				  componentObjRef: newUrlLinkObjectRef,
  				  designFormConfig: urlLinkDesignFormConfig
  	  		  }
  			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
		
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "urlLink_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeURL],
		hideCreateCalcFieldCheckbox: true,
		containerParams: containerParams,
		createNewFormComponent: createNewUrlLink
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
			
} // newLayoutContainer

function initNewUrlLinkDialog() {
}


