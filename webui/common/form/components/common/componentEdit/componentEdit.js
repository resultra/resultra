

function initFormComponentDesignBehavior(componentIDs, objectRef, designFormConfig,layoutDesignConfig) {
	
	console.log("initFormComponentDesignBehavior: params = " + JSON.stringify(componentIDs))
	
	initObjectGridEditBehavior(componentIDs.componentID,designFormConfig,layoutDesignConfig)
	
	
	initObjectSelectionBehavior($("#"+componentIDs.componentID), 
			formDesignCanvasSelector,function(selectedCompenentID) {
		console.log("form design object selected: " + selectedCompenentID)
		var selectedObjRef	= getElemObjectRef(selectedCompenentID)
		designFormConfig.selectionFunc(selectedObjRef)
	})
		
	
}	  
