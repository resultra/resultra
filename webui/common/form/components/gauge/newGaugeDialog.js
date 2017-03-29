function openNewGaugeDialog(databaseID,formID,containerParams) {
	
	function createNewGaugeComponent($parentDialog, newComponentParams) {
		
		jsonAPIRequest("frm/gauge/new",newComponentParams,function(newGaugeObjectRef) {
	          console.log("openNewGaugeDialog: Done getting new gauge component:response=" 
						+ JSON.stringify(newGaugeObjectRef));
						
				setGaugeComponentLabel(newGaugeObjectRef, newGaugeObjectRef)
	  	  						 
  	  		  var newComponentSetupParams = {
  				  parentFormID: formID,
  	  		  	  $container: containerParams.containerObj,
  				  componentID: newGaugeObjectRef.gaugeID,
  				  componentObjRef: newGaugeObjectRef,
  				  designFormConfig: gaugeDesignFormConfig
  	  		  }
  			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
	  		  			  
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
		
	}
	
	var newFormComponentDialogParams = {
		elemPrefix: "gauge_",
		databaseID: databaseID,
		formID: formID,
		fieldTypes: [fieldTypeNumber],
		containerParams: containerParams,
		createNewFormComponent: createNewGaugeComponent
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)
	
}