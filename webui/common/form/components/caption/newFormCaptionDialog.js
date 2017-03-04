function openNewFormCaptionDialog(databaseID,formID,containerParams) {
	console.log("New form caption dialog")
	
	var newCaptionParams = {
		parentFormID: formID,
		geometry: containerParams.geometry,
		label: "New Caption"}
	
	jsonAPIRequest("frm/caption/new",newCaptionParams,function(newCaptionObjectRef) {
          console.log("create new form header: Done getting new ID:response=" + JSON.stringify(newHeaderObjectRef));
		  
		  containerParams.containerObj.find('.formCaption').text(newCaptionObjectRef.properties.label)
 
		  // Set up the newly created checkbox for resize, selection, etc.
		  var componentIDs = { formID: formID, componentID:newCaptionObjectRef.captionID }
		  initFormComponentDesignBehavior(containerParams.containerObj,componentIDs,newCaptionObjectRef,formCaptionDesignFormConfig)
	  
		  // Put a reference to the caption's reference object in the check box's DOM element.
		  // This reference can be retrieved later for property setting, etc.
		  setContainerComponentInfo(containerParams.containerObj,newCaptionObjectRef,newCaptionObjectRef.captionID)
	  			  
       }) // newLayoutContainer API request
	
}