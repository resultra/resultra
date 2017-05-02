function createNewDashboardHeader(headerParams) {
	
	
	var saveNewHeaderParams = {
		parentDashboardID: headerParams.dashboardContext.dashboardID,
		geometry: headerParams.geometry
	}
	console.log("Creating new header: " + JSON.stringify(saveNewHeaderParams))
	
	jsonAPIRequest("dashboard/header/new",saveNewHeaderParams,function(newHeaderObjectRef) {  
	  var newComponentSetupParams = {
		  parentDashboardID: headerParams.dashboardContext.dashboardID,
	  	  $container: headerParams.$componentContainer,
		  componentID: newHeaderObjectRef.headerID,
		  componentObjRef: newHeaderObjectRef,
		  designFormConfig: headerDashboardDesignConfig
	  }
	  setupNewlyCreatedDashboardComponentInfo(newComponentSetupParams)
	  			  
    }) 
	
	
}