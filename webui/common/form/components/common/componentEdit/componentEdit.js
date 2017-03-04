

function initFormComponentDesignBehavior($componentContainer, componentIDs, objectRef, designFormConfig,layoutDesignConfig) {
	
	console.log("initFormComponentDesignBehavior: params = " + JSON.stringify(componentIDs))
	
	initObjectGridEditBehavior($componentContainer,designFormConfig,layoutDesignConfig)
	
	var $designFormParentCanvas = $(formDesignCanvasSelector)
	
	initObjectSelectionBehavior($componentContainer, 
			$designFormParentCanvas,function(selectedCompenentID) {
		console.log("form design object selected: " + selectedCompenentID)
		var selectedObjRef	= getContainerObjectRef($componentContainer)
		designFormConfig.selectionFunc($componentContainer,selectedObjRef)
	})
		
	
}	  
