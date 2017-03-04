function openNewFormHeaderDialog(databaseID,formID,containerParams) {
	console.log("New form header dialog")
	
	var newHeaderParams = {
		parentFormID: formID,
		geometry: containerParams.geometry,
		label: "New Header"}
	
	jsonAPIRequest("frm/header/new",newHeaderParams,function(newHeaderObjectRef) {
          console.log("create new form header: Done getting new ID:response=" + JSON.stringify(newHeaderObjectRef));
		  
		  containerParams.containerObj.find('.formHeader').text(newHeaderObjectRef.properties.label)
  
  		  var newComponentSetupParams = {
			  parentFormID: formID,
  		  	  $container: containerParams.containerObj,
			  componentID: newHeaderObjectRef.headerID,
			  componentObjRef: newHeaderObjectRef,
			  designFormConfig: formHeaderDesignFormConfig
  		  }
		  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
  
	  			  
       }) // newLayoutContainer API request
	
}