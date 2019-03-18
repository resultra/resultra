// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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