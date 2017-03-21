function initDashboardComponentDesignDashboardEditBehavior($component,componentID, designDashboardConfig,layoutDesignConfig) {
	console.log("initDashboardComponentDesignDashboardEditBehavior: component ID = " + componentID)

	initObjectGridEditBehavior($component,designDashboardConfig,layoutDesignConfig)		

	// When in design mode, dashboard components need to have dotted lines as their border.
	$component.addClass("layoutDesignContainer")
	
	var $parentDashboardCanvas = $(dashboardDesignCanvasSelector)
	initObjectSelectionBehavior($component, 
			$parentDashboardCanvas,function(selectedComponentID) {
		console.log("dashboard design object selected: " + selectedComponentID)
		var selectedObjRef	= getContainerObjectRef($component)
		designDashboardConfig.selectionFunc($component,selectedObjRef)
	})
	
	
}

function saveUpdatedDashboardComponentLayout(updatedLayout) {
	console.log("saveUpdatedDashboardComponentLayout: component layout = " + JSON.stringify(updatedLayout))		
	var setLayoutParams = {
		dashboardID: designDashboardContext.dashboardID,
		layout: updatedLayout
	}
	jsonAPIRequest("dashboard/setLayout", setLayoutParams, function(dashboardInfo) {})
	
}

function saveUpdatedDesignDashboardLayout() {
	
	var $parentLayoutContainer = $(dashboardDesignCanvasSelector)
	
	// There's a delay between the time the DOM is updated in this thread and when those 
	// changes are fully reflected in the DOM (see http://stackoverflow.com/questions/16876394/dom-refresh-on-long-running-function)
	// To accommodate this, the layout is pruned and saved after a small delay.
	setTimeout(function() {
		pruneComponentLayoutEmptyColsAndRows($parentLayoutContainer)
		var updatedLayout = getComponentLayout($parentLayoutContainer)
		
		saveUpdatedDashboardComponentLayout(updatedLayout)		
	 },20);	
	
}



function createDashboardLayoutDesignConfig() {
	// TODO - Move to function for dragging items
	
	var layoutDesignConfig = {
		parentLayoutSelector: dashboardDesignCanvasSelector,
		saveLayoutFunc: saveUpdatedDashboardComponentLayout
	}
	return layoutDesignConfig
}


function setupNewlyCreatedDashboardComponentInfo(params) {
	
	var layoutDesignConfig = createDashboardLayoutDesignConfig()
	
	initDashboardComponentDesignDashboardEditBehavior(params.$container,params.componentID,params.designFormConfig,layoutDesignConfig)
	
	// Save a reference to the object in the component's container
	setContainerComponentInfo(params.$container,params.componentObjRef,params.componentID)
	
	// Save the layout to reflect the newly created item
	saveUpdatedDesignDashboardLayout()
}
