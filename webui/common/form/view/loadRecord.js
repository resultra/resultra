function loadRecordIntoFormLayout($parentFormLayout, recordRef) {

	console.log("Loading record into layout: record field values: " + JSON.stringify(recordRef.fieldValues))

	// TODO - Conditionally show or hide the form components.

	var hiddenComponents = new IDLookupTable(recordRef.hiddenFormComponents)


	// Iterate through all the containers in the current layout (which may be a subset of the record's fields),
	// and populate the container's value with the field's value from the record.
	
	$parentFormLayout.find(".layoutContainer").each(function() {
		
		
		var $container = $(this)
		
		var componentID = getContainerComponentID($container)
		
		if (hiddenComponents.hasID(componentID)) {
			if (elemIsDisplayed($container)) {
				$container.animate({opacity:0},500,function() {
					// fade out, then hide completely
					$container.hide()
				})			
			}
		
		} else {
			if (!elemIsDisplayed($container)) {
				$container.show() // show it but opacity will still be 0
				$container.animate({opacity:1},500) // fade in
			}
		}
		
		// Each type of form object needs to set a "viewFormConfig" object on it's DOM element. The loadRecord()
		// function is called on each of these objects to perform per form object record initialization.
		var viewFormConfig = $container.data("viewFormConfig")
		viewFormConfig.loadRecord($container,recordRef)
		
	})
		
}