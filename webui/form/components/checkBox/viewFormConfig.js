function loadRecordIntoCheckBox(checkBoxElem, recordRef) {
	
	console.log("loadRecordIntoCheckBox: loading record into text box: " + JSON.stringify(recordRef))
	
	var checkBoxObjectRef = checkBoxElem.data("objectRef")
	var checkBoxFieldID = checkBoxObjectRef.properties.fieldID
	
	console.log("loadRecordIntoCheckBox: Field ID to load data:" + checkBoxFieldID)


	var checkBoxContainerID = checkBoxObjectRef.checkBoxID
	var checkBoxID = checkBoxElemIDFromContainerElemID(checkBoxContainerID)
	var checkBoxSelector = '#' + checkBoxID;
	
	// Populate the "intersection" of field values in the record
	// with the fields shown by the layout's containers.
	if(recordRef.fieldValues.hasOwnProperty(checkBoxFieldID)) {

		var fieldVal = recordRef.fieldValues[checkBoxFieldID]

		console.log("loadRecordIntoCheckBox: Load value into container: " + $(checkBoxElem).attr("id") + " field ID:" + 
					checkBoxFieldID + "  value:" + fieldVal)

		if(fieldVal == true)
		{
			$(checkBoxSelector).prop("checked",true)
		}
		else {
			$(checkBoxSelector).prop("checked",false)
		}

	} // If record has a value for the current container's associated field ID.
	else
	{
		$(checkBoxSelector).prop("indeterminate", true)
	}
	
}


function initCheckBoxRecordEditBehavior(componentContext,checkBoxObjectRef) {
	
	var checkBoxContainerID = checkBoxObjectRef.checkBoxID
	var checkBoxID = checkBoxElemIDFromContainerElemID(checkBoxContainerID)
	var checkboxSelector = '#'+checkBoxID
	console.log("initCheckBoxRecordEditBehavior: container ID =  " +checkBoxContainerID + ' checkbox ID = '+ checkBoxID)
	
	var checkBoxContainer = $('#'+checkBoxContainerID)
	


	checkBoxContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoCheckBox
	})


	var fieldRef = getFieldRef(checkBoxObjectRef.properties.fieldID)
	if(fieldRef.isCalcField) {
		$(checkboxSelector).checkbox()
		$(checkboxSelector).checkbox('set disabled')
		return;  // stop initialization, the check box is read only.
	}
	

  	$(checkboxSelector).click( function () {
		
			// Get the most recent copy of the object reference. It could have changed between
			// initialization time and the time the checkbox was changed.
			var containerID = checkBoxObjectRef.checkBoxID
			var objectRef = getElemObjectRef(containerID)
			var checkBoxSelect = '#'+checkBoxElemIDFromContainerElemID(containerID)
			
			var isChecked = $(this).prop("checked")
	
			currRecordRef = currRecordSet.currRecordRef()
			var setRecordValParams = {
				parentTableID:viewFormContext.tableID,
				recordID:currRecordRef.recordID, 
				fieldID:objectRef.properties.fieldID, 
				value:isChecked }
			console.log("Setting boolean value: " + JSON.stringify(setRecordValParams))
			jsonAPIRequest("recordUpdate/setBoolFieldValue",setRecordValParams,function(updatedRecordRef) {
				
				// After updating the record, the local cache of records in currentRecordSet will
				// be out of date. So after updating the record on the server, the locally cached
				// version of the record also needs to be updated.
				currRecordSet.updateRecordRef(updatedRecordRef)
				// After changing the value, some of the calculated fields may have changed. For this
				// reason, it is necessary to reload the record into the layout/form, so the most
				// up to date values will be displayed.
				loadCurrRecordIntoLayout()
			}) // set record's text field value
	})

	
	
}