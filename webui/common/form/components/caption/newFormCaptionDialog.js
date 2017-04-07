function openNewFormCaptionDialog(databaseID,formID,containerParams) {
	console.log("New form caption dialog")
	
	var newCaptionParams = {
		parentFormID: formID,
		geometry: containerParams.geometry,
		label: "New Caption"}
	
	jsonAPIRequest("frm/caption/new",newCaptionParams,function(newCaptionObjectRef) {
          console.log("create new form header: Done getting new ID:response=" + JSON.stringify(newCaptionObjectRef));
		  
		  containerParams.containerObj.find('.formCaption').text(newCaptionObjectRef.properties.label)
 
  		  var newComponentSetupParams = {
			  parentFormID: formID,
  		  	  $container: containerParams.containerObj,
			  componentID: newCaptionObjectRef.captionID,
			  componentObjRef: newCaptionObjectRef,
			  designFormConfig: formCaptionDesignFormConfig
  		  }
		  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
		  
		  initCaptionDesignControlBehavior(containerParams.containerObj,newCaptionObjectRef)
		  
 
	  			  
       }) // newLayoutContainer API request
	
}