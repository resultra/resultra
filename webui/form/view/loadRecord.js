function loadRecordIntoFormLayout(parentFormLayoutSelector, recordRef) {

	console.log("Loading record into layout: record field values: " + JSON.stringify(recordRef.fieldValues))


	// Iterate through all the containers in the current layout (which may be a subset of the record's fields),
	// and populate the container's value with the field's value from the record.
	
	var formLayoutContainerSelector = parentFormLayoutSelector + " .layoutContainer"
	
	$(formLayoutContainerSelector).each(function() {
		
		// Each type of form object needs to set a "viewFormConfig" object on it's DOM element. The loadRecord()
		// function is called on each of these objects to perform per form object record initialization.
		var viewFormConfig = $(this).data("viewFormConfig")
		viewFormConfig.loadRecord($(this),recordRef)
		

	}) // for each container in the layout
	
}