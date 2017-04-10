

function openNewNumberInputDialog(databaseID,formID,containerParams)
{
		
	function createNewNumberInput($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/numberInput/new",newComponentParams,function(newNumberInputObjectRef) {
	          console.log("saveNewNumberInput: Done getting new ID:response=" + JSON.stringify(newNumberInputObjectRef));
		  		  
			  // The new number input has been saved on the server, but only a placeholder of the text box 
			  // is currently shown in the layout. The following code is needed to update and finalize the placeholder
			  // as a complete and fully-functional text box.
			  			  
			  setNumberInputComponentLabel(containerParams.containerObj,newNumberInputObjectRef)
			  configureNumberInputButtonSpinner(containerParams.containerObj,newNumberInputObjectRef)			  	
			  	  
  	  		  var newComponentSetupParams = {
  				  parentFormID: formID,
  	  		  	  $container: containerParams.containerObj,
  				  componentID: newNumberInputObjectRef.numberInputID,
  				  componentObjRef: newNumberInputObjectRef,
  				  designFormConfig: numberInputDesignFormConfig
  	  		  }
  			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
		
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "numberInput_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeNumber],
		containerParams: containerParams,
		createNewFormComponent: createNewNumberInput
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
			
} // newLayoutContainer

function initNewNumberInputDialog() {
}


