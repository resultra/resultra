

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

function createFormLayoutDesignConfig(parentFormID) {
	
	function saveUpdatedFormComponentLayout(updatedLayout) {
		console.log("createFormLayoutDesignConfig: component layout = " + JSON.stringify(updatedLayout))		
		var setLayoutParams = {
			formID: parentFormID,
			layout: updatedLayout
		}
		jsonAPIRequest("frm/setLayout", setLayoutParams, function(formInfo) {
		})
	}		
	
	
	var designFormLayoutConfig =  {
		parentLayoutSelector: formDesignCanvasSelector,
		saveLayoutFunc: saveUpdatedFormComponentLayout
	}
	
	return designFormLayoutConfig
}


function saveUpdatedDesignFormLayout(parentFormID) {
	
	var $parentLayoutContainer = $(formDesignCanvasSelector)
	
	// There's a delay between the time the DOM is updated in this thread and when those 
	// changes are fully reflected in the DOM (see http://stackoverflow.com/questions/16876394/dom-refresh-on-long-running-function)
	// To accommodate this, the layout is pruned and saved after a small delay.
	setTimeout(function() {
		pruneComponentLayoutEmptyColsAndRows($parentLayoutContainer)
		var updatedLayout = getComponentLayout($parentLayoutContainer)
		
		var setLayoutParams = {
			formID: designFormContext.formID,
			layout: updatedLayout
		}
	
		jsonAPIRequest("frm/setLayout", setLayoutParams, function(formInfo) {
		})
		
	 },20);	
	
}

function setupNewlyCreatedFormComponentInfo(setupParams) {
	
  // Set up the newly created checkbox for resize, selection, etc.
  var componentIDs = { 
	  formID: setupParams.parentFormID, 
	  componentID:setupParams.componentID }
	  
  var formLayoutConfig = createFormLayoutDesignConfig(setupParams.parentFormID)
	  
  initFormComponentDesignBehavior(setupParams.$container,componentIDs,setupParams.componentObjRef,setupParams.designFormConfig)

  // Put a reference to the check box's reference object in the check box's DOM element.
  // This reference can be retrieved later for property setting, etc.
  setContainerComponentInfo(setupParams.$container,setupParams.componentObjRef,setupParams.componentID)
  
  // Now that the form component has been fully created, save the 
  // updated form layout to include the component.
  saveUpdatedDesignFormLayout(setupParams.parentFormID)
	
}
