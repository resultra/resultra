function loadRecordIntoCheckBox(checkBoxElem, recordRef) {
	
	console.log("loadRecordIntoCheckBox: loading record into text box: " + JSON.stringify(recordRef))
	
	var checkBoxObjectRef = checkBoxElem.data("objectRef")
	var checkBoxContainerID = checkBoxObjectRef.checkBoxID
	var checkBoxID = checkBoxElemIDFromContainerElemID(checkBoxContainerID)
	var checkBoxSelector = '#' + checkBoxID;
	
	var componentLink = checkBoxObjectRef.properties.componentLink
	
	if(componentLink.linkedValType == linkedComponentValTypeField) {
		var checkBoxFieldID = componentLink.fieldID
	
		console.log("loadRecordIntoCheckBox: Field ID to load data:" + checkBoxFieldID)


	
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
	} else {
		var checkBoxGlobalID = componentLink.globalID
		console.log("loadRecordIntoCheckBox: Global ID to load data:" + checkBoxGlobalID)
		
		if(checkBoxGlobalID in currGlobalVals) {
			var globalVal = currGlobalVals[checkBoxGlobalID]
			if(globalVal == true)
			{
				$(checkBoxSelector).prop("checked",true)
			}
			else {
				$(checkBoxSelector).prop("checked",false)
			}
		}
		else
		{
			$(checkBoxSelector).prop("indeterminate", true)
		}
		
	}
	
	
}

function initCheckBoxFieldEditBehavior(componentContext,checkBoxObjectRef) {
	
	var checkboxSelector = '#'+checkBoxElemIDFromContainerElemID(checkBoxObjectRef.checkBoxID)
	
	var componentLink = checkBoxObjectRef.properties.componentLink
	
	var fieldRef = getFieldRef(componentLink.fieldID)
	if(fieldRef.isCalcField) {
		$(checkboxSelector).checkbox()
		$(checkboxSelector).checkbox('set disabled')
		return;  // stop initialization, the check box is read only.
	}
	

  	$(checkboxSelector).click( function () {
		
			// Get the most recent copy of the object reference. It could have changed between
			// initialization time and the time the checkbox was changed.
			var objectRef = getElemObjectRef(checkBoxObjectRef.checkBoxID)
			var componentLink = objectRef.properties.componentLink
			
			var isChecked = $(this).prop("checked")
		
			var currRecordRef = currRecordSet.currRecordRef()
			var checkboxValueFormat = {
				context: "checkbox",
				format: "check"
			}
			var setRecordValParams = {
				parentDatabaseID:viewListContext.databaseID,
				recordID:currRecordRef.recordID, 
				fieldID:componentLink.fieldID, 
				value:isChecked,
				 valueFormat: checkboxValueFormat }
			console.log("Setting boolean value (record): " + JSON.stringify(setRecordValParams))
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

function initCheckBoxGlobalEditBehavior(componentContext,checkBoxObjectRef) {

	var checkboxSelector = '#'+checkBoxElemIDFromContainerElemID(checkBoxObjectRef.checkBoxID)
		
  	$(checkboxSelector).click( function () {
		
		// Get the most recent copy of the object reference. It could have changed between
		// initialization time and the time the checkbox was changed.
		var objectRef = getElemObjectRef(checkBoxObjectRef.checkBoxID)
		var componentLink = objectRef.properties.componentLink
		
		var isChecked = $(this).prop("checked")
		
		var setGlobalValParams = {
			parentDatabaseID: componentContext.databaseID,
			globalID: componentLink.globalID,
			value: isChecked }
		console.log("Setting boolean value (global): " + JSON.stringify(setGlobalValParams))
		jsonAPIRequest("global/setBoolValue",setGlobalValParams,function(updatedGlobalVal) {
			
			// TODO - Update global structure with updated value.
		}) // set record's text field value
		
	})
	
}

function initCheckBoxRecordEditBehavior(componentContext,checkBoxObjectRef) {

	var $checkBoxContainer = $('#'+checkBoxObjectRef.checkBoxID)
		
	$checkBoxContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoCheckBox
	})
	
	var componentLink = checkBoxObjectRef.properties.componentLink

	if(componentLink.linkedValType == linkedComponentValTypeField) {
		initCheckBoxFieldEditBehavior(componentContext,checkBoxObjectRef)
	} else {
		initCheckBoxGlobalEditBehavior(componentContext,checkBoxObjectRef)
	}


	
	
}