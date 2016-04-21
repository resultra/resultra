function loadRecordIntoCheckBox(checkBoxElem, recordRef) {
	
	console.log("loadRecordIntoCheckBox: loading record into text box: " + JSON.stringify(recordRef))
	
	var checkBoxObjectRef = checkBoxElem.data("objectRef")
	var checkBoxFieldID = checkBoxObjectRef.fieldRef.fieldID
	
	console.log("loadRecordIntoCheckBox: Field ID to load data:" + checkBoxFieldID)


	var checkBoxContainerID = checkBoxObjectRef.uniqueID.objectID
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
			$(checkBoxSelector).checkbox('set checked')	
		}
		else {
			$(checkBoxSelector).checkbox('set unchecked')		
		}

	} // If record has a value for the current container's associated field ID.
	else
	{
		$(checkBoxSelector).checkbox('set indeterminate')
	}
	
}


function initCheckBoxRecordEditBehavior(checkBoxObjectRef) {
	
	var checkBoxContainerID = checkBoxObjectRef.uniqueID.objectID
	var checkBoxID = checkBoxElemIDFromContainerElemID(checkBoxContainerID)
	console.log("initCheckBoxRecordEditBehavior: container ID =  " +checkBoxContainerID + ' checkbox ID = '+ checkBoxID)
	
	var checkBoxContainer = $('#'+checkBoxContainerID)
	var checkboxSelector = '#'+checkBoxID

	checkBoxContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoCheckBox
	})


	if(checkBoxObjectRef.fieldRef.fieldInfo.isCalcField) {
		$(checkboxSelector).checkbox()
		$(checkboxSelector).checkbox('set disabled')
		return;  // stop initialization, the check box is read only.
	}


	$(checkboxSelector).checkbox({
		onChange: function () {
			// Get the most recent copy of the object reference. It could have changed between
			// initialization time and the time the checkbox was changed.
			var containerID = checkBoxObjectRef.uniqueID.objectID
			var objectRef = getElemObjectRef(containerID)
			var checkBoxSelect = '#'+checkBoxElemIDFromContainerElemID(containerID)
			
			var isChecked = $(checkBoxSelect).checkbox('is checked')

			
			currRecordRef = currRecordSet.currRecordRef()
			var setRecordValParams = { recordID:currRecordRef.recordID, fieldID:objectRef.fieldRef.fieldID, value:isChecked }
			console.log("Setting boolean value: " + JSON.stringify(setRecordValParams))
			jsonAPIRequest("setBoolFieldValue",setRecordValParams,function(updatedRecordRef) {
				
				// After updating the record, the local cache of records in currentRecordSet will
				// be out of date. So after updating the record on the server, the locally cached
				// version of the record also needs to be updated.
				currRecordSet.updateRecordRef(updatedRecordRef)
				// After changing the value, some of the calculated fields may have changed. For this
				// reason, it is necessary to reload the record into the layout/form, so the most
				// up to date values will be displayed.
				loadCurrRecordIntoLayout()
			}) // set record's text field value
			
		}
	})

	
	
}