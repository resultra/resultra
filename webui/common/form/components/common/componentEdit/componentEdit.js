

function initFormComponentDesignBehavior($componentContainer, componentIDs, objectRef, designFormConfig,layoutDesignConfig) {
	
	console.log("initFormComponentDesignBehavior: params = " + JSON.stringify(componentIDs))
	
	initObjectGridEditBehavior($componentContainer,designFormConfig,layoutDesignConfig)
	
	
	initObjectSelectionBehavior($componentContainer, 
			formDesignCanvasSelector,function(selectedCompenentID) {
		console.log("form design object selected: " + selectedCompenentID)
		var selectedObjRef	= getContainerObjectRef($componentContainer)
		designFormConfig.selectionFunc(selectedObjRef)
	})
		
	
}	  
