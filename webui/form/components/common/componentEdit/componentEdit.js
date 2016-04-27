function initFormComponentDesignBehavior(objectRef, designFormConfig) {
	
	
	// Store the object reference in the DOM element. This is needed for follow-on
	// property setting, resizing, etc.
	setElemObjectRef(objectRef.uniqueID.objectID,objectRef)
	
	initObjectEditBehavior(objectRef.uniqueID.parentID,
		objectRef.uniqueID.objectID,designFormConfig)
		
	initObjectSelectionBehavior($("#"+objectRef.uniqueID.objectID), 
			formDesignCanvasSelector,function(objectID) {
		console.log("form design object selected: " + objectID)
		var selectedObjRef	= getElemObjectRef(objectID)
		designFormConfig.selectionFunc(selectedObjRef)
	})	
	
}	  
