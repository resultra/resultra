function openNewFormHeaderDialog(databaseID,formID,containerParams) {
	console.log("New form header dialog")
	
	var newHeaderParams = {
		parentFormID: formID,
		geometry: containerParams.geometry,
		label: "New Header"}
	
	jsonAPIRequest("frm/header/new",newHeaderParams,function(newHeaderObjectRef) {
          console.log("create new form header: Done getting new ID:response=" + JSON.stringify(newHeaderObjectRef));
    			  
		  var placeholderSelector = '#'+containerParams.containerID
		  
		  $(placeholderSelector).find('.formHeader').text(newHeaderObjectRef.properties.label)
		  $(placeholderSelector).attr("id",newHeaderObjectRef.headerID)
  
		  // Set up the newly created checkbox for resize, selection, etc.
		  var componentIDs = { formID: formID, componentID:newHeaderObjectRef.headerID }
		  initFormComponentDesignBehavior(componentIDs,newHeaderObjectRef,formHeaderDesignFormConfig)
	  
		  // Put a reference to the check box's reference object in the check box's DOM element.
		  // This reference can be retrieved later for property setting, etc.
		  setElemObjectRef(newHeaderObjectRef.headerID,newHeaderObjectRef)
	  			  
       }) // newLayoutContainer API request
	
}