function openNewFormButtonDialog(databaseID,formID,containerParams) {
	console.log("New form button dialog")
	
	var newButtonParams = {
		parentFormID: formID,
		geometry: containerParams.geometry}
	
	/*
	jsonAPIRequest("frm/formButton/new",newButtonParams,function(newButtonObjectRef) {
		
          console.log("create new form button: Done getting new ID:response=" + JSON.stringify(newButtonObjectRef));
    			  
		  var placeholderSelector = '#'+containerParams.containerID
		  
		  $(placeholderSelector).find('.formButton').text(newButtonObjectRef.properties.label)
		  $(placeholderSelector).attr("id",newHeaderObjectRef.headerID)
  
		  // Set up the newly created checkbox for resize, selection, etc.
		  var componentIDs = { formID: formID, componentID:newButtonObjectRef.buttonID }
		  
		  initFormComponentDesignBehavior(componentIDs,newButtonObjectRef,formButtonDesignFormConfig)
	  
		  // Put a reference to the check box's reference object in the check box's DOM element.
		  // This reference can be retrieved later for property setting, etc.
		  setElemObjectRef(newButtonObjectRef.buttonID,newButtonObjectRef)
	  			  
       }) // newLayoutContainer API request
	*/
}