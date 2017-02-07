function loadRecordIntoCheckBox($checkboxContainer, recordRef) {
	
	console.log("loadRecordIntoCheckBox: loading record into text box: " + JSON.stringify(recordRef))
	
	var checkBoxObjectRef = getContainerObjectRef($checkboxContainer)
	var $checkBoxControl = getCheckboxControlFromCheckboxContainer($checkboxContainer);
	
	var componentLink = checkBoxObjectRef.properties.componentLink
	
	if(componentLink.linkedValType == linkedComponentValTypeField) {
		var checkBoxFieldID = componentLink.fieldID
	
		console.log("loadRecordIntoCheckBox: Field ID to load data:" + checkBoxFieldID)
		
		// Populate the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(checkBoxFieldID)) {

			var fieldVal = recordRef.fieldValues[checkBoxFieldID]

			if(fieldVal == true)
			{
				$checkBoxControl.prop("indeterminate", false)
				$checkBoxControl.prop("checked",true)
			}
			else {
				$checkBoxControl.prop("indeterminate", false)
				$checkBoxControl.prop("checked",false)
			}

		} // If record has a value for the current container's associated field ID.
		else
		{
			$checkBoxControl.prop("indeterminate", true)
		}
	} else {
		var checkBoxGlobalID = componentLink.globalID
		console.log("loadRecordIntoCheckBox: Global ID to load data:" + checkBoxGlobalID)
		
		if(checkBoxGlobalID in currGlobalVals) {
			var globalVal = currGlobalVals[checkBoxGlobalID]
			if(globalVal == true)
			{
				$checkBoxControl.prop("checked",true)
			}
			else {
				$checkBoxControl.prop("checked",false)
			}
		}
		else
		{
			$checkBoxControl.prop("indeterminate", true)
		}
		
	}
	
	
}

function initCheckBoxFieldEditBehavior($checkBox,componentContext,changeSetID,getRecordFunc, updateRecordFunc, checkBoxObjectRef) {
	
	var $checkboxControl = getCheckboxControlFromCheckboxContainer($checkBox)
	
	var componentLink = checkBoxObjectRef.properties.componentLink
	
	var fieldRef = getFieldRef(componentLink.fieldID)
	if(fieldRef.isCalcField) {
		$checkboxControl.checkbox()
		$checkboxControl.checkbox('set disabled')
		return;  // stop initialization, the check box is read only.
	}
	
	$checkboxControl.unbind("click")
  	$checkboxControl.click( function () {
		
			// Get the most recent copy of the object reference. It could have changed between
			// initialization time and the time the checkbox was changed.
			var objectRef = getContainerObjectRef($checkBox)
			var componentLink = objectRef.properties.componentLink
			
			var isChecked = $(this).prop("checked")
		
			var currRecordRef = getRecordFunc()
			var checkboxValueFormat = {
				context: "checkbox",
				format: "check"
			}
			var setRecordValParams = {
				parentDatabaseID:currRecordRef.parentDatabaseID,
				recordID:currRecordRef.recordID,
				changeSetID: changeSetID,
				fieldID:componentLink.fieldID, 
				value:isChecked,
				 valueFormat: checkboxValueFormat }
			console.log("Setting boolean value (record): " + JSON.stringify(setRecordValParams))
			jsonAPIRequest("recordUpdate/setBoolFieldValue",setRecordValParams,function(updatedRecordRef) {
				
				// After updating the record, the local cache of records in currentRecordSet will
				// be out of date. So after updating the record on the server, the locally cached
				// version of the record also needs to be updated.
				updateRecordFunc(updatedRecordRef)
			}) // set record's text field value
	})
	
}

function initCheckBoxGlobalEditBehavior($checkBox,componentContext,checkBoxObjectRef) {

	var $checkboxControl = getCheckboxControlFromCheckboxContainer($checkBox)
		
  	$checkboxControl.click( function () {
		
		// Get the most recent copy of the object reference. It could have changed between
		// initialization time and the time the checkbox was changed.
		var objectRef = getContainerObjectRef($checkBox)
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

function initCheckBoxRecordEditBehavior($checkBox,componentContext,changeSetID,
					getRecordFunc, updateRecordFunc, checkBoxObjectRef) {
		
	$checkBox.data("viewFormConfig", {
		loadRecord: loadRecordIntoCheckBox
	})
	
	var componentLink = checkBoxObjectRef.properties.componentLink

	if(componentLink.linkedValType == linkedComponentValTypeField) {
		initCheckBoxFieldEditBehavior($checkBox,componentContext,changeSetID,
				getRecordFunc, updateRecordFunc, checkBoxObjectRef)
	} else {
		initCheckBoxGlobalEditBehavior($checkBox,componentContext,checkBoxObjectRef)
	}


	
	
}