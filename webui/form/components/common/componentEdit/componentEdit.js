

function initFormComponentDesignBehavior(componentIDs, objectRef, designFormConfig) {
	
	console.log("initFormComponentDesignBehavior: params = " + JSON.stringify(componentIDs))
	
	initObjectEditBehavior(componentIDs.formID,
		componentIDs.componentID,designFormConfig)
		
	initObjectSelectionBehavior($("#"+componentIDs.componentID), 
			formDesignCanvasSelector,function(selectedCompenentID) {
		console.log("form design object selected: " + selectedCompenentID)
		var selectedObjRef	= getElemObjectRef(selectedCompenentID)
		designFormConfig.selectionFunc(selectedObjRef)
	})	
	
}	  
