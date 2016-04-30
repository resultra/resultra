

function initFormComponentDesignBehavior(componentIDs, objectRef, designFormConfig) {
	
	initObjectEditBehavior(componentIDs.formID,
		componentIDs.componentID,designFormConfig)
		
	initObjectSelectionBehavior($("#"+componentIDs.componentID), 
			formDesignCanvasSelector,function(selectedCompenentID) {
		console.log("form design object selected: " + selectedCompenentID)
		var selectedObjRef	= getElemObjectRef(selectedCompenentID)
		designFormConfig.selectionFunc(selectedObjRef)
	})	
	
}	  
