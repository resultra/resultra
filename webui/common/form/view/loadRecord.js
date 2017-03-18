function loadRecordIntoFormLayout($parentFormLayout, recordRef) {

	console.log("Loading record into layout: record field values: " + JSON.stringify(recordRef.fieldValues))

	// TODO - Conditionally show or hide the form components.

	var hiddenComponents = new IDLookupTable(recordRef.hiddenFormComponents)


	// Iterate through all the containers in the current layout (which may be a subset of the record's fields),
	// and populate the container's value with the field's value from the record.
	
	$parentFormLayout.find(".layoutContainer").each(function() {
		
		
		var $container = $(this)
		
		var currRecordID = $container.attr("data-recordID")
		var recordBeingLoadedSameAsOneAlreadyInContainer = (currRecordID === recordRef.recordID)
		
		var componentID = getContainerComponentID($container)
		
		if (hiddenComponents.hasID(componentID)) {
			if (elemIsDisplayed($container)) {
				if (recordBeingLoadedSameAsOneAlreadyInContainer) {
					// If currRecordID is the same as the record ID of the record being loaded,
					// then fade out the component. In this case, loading the record is due to some type
					// of value change in the record which necessitates updating the visibility of different components.
					// Otherwise, immediately hide the component.
					$container.animate({opacity:0},500,function() {
						// fade out, then hide completely
					$container.hide()
					})								
				} else {
					$container.css("opacity",'0')
					$container.hide()
				}
			}
		
		} else {
			if (!elemIsDisplayed($container)) {
				if (recordBeingLoadedSameAsOneAlreadyInContainer) {
					$container.show() // show it but opacity will still be 0
					$container.animate({opacity:1},500) // fade in
				} else {
					$container.show() // show it but opacity will still be 0
					$container.css("opacity",'1')
				}
			}
		}
		
		// Each type of form object needs to set a "viewFormConfig" object on it's DOM element. The loadRecord()
		// function is called on each of these objects to perform per form object record initialization.
		var viewFormConfig = $container.data("viewFormConfig")
		viewFormConfig.loadRecord($container,recordRef)
		
		$container.attr("data-recordID",recordRef.recordID)
		
	})
		
}