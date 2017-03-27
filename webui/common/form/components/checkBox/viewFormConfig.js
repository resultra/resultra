function loadRecordIntoCheckBox($checkboxContainer, recordRef) {
	
	console.log("loadRecordIntoCheckBox: loading record into text box: " + JSON.stringify(recordRef))
	
	var checkBoxObjectRef = getContainerObjectRef($checkboxContainer)
	var $checkBoxControl = getCheckboxControlFromCheckboxContainer($checkboxContainer);
	var $checkboxLabel = $checkboxContainer.find("label")
	
	var checkBoxFieldID = checkBoxObjectRef.properties.fieldID

	console.log("loadRecordIntoCheckBox: Field ID to load data:" + checkBoxFieldID)
	
	// Populate the "intersection" of field values in the record
	// with the fields shown by the layout's containers.
	if(recordRef.fieldValues.hasOwnProperty(checkBoxFieldID)) {

		var fieldVal = recordRef.fieldValues[checkBoxFieldID]
		
		if (fieldVal === null) {
				// Value has been cleared
				$checkboxLabel.removeClass("checkboxStrikethroughCompleted")
				$checkBoxControl.prop("indeterminate", true)
		} else {
			$checkboxLabel.removeClass("checkboxStrikethroughCompleted")

			if(fieldVal == true)
			{
				$checkBoxControl.prop("indeterminate", false)
				$checkBoxControl.prop("checked",true)
				if(checkBoxObjectRef.properties.strikethroughCompleted) {
					$checkboxLabel.addClass("checkboxStrikethroughCompleted")
				} else {
					$checkboxLabel.removeClass("checkboxStrikethroughCompleted")
				}
			}
			else {
				$checkboxLabel.removeClass("checkboxStrikethroughCompleted")
				$checkBoxControl.prop("indeterminate", false)
				$checkBoxControl.prop("checked",false)
			}
			
		}


	} // If record has a value for the current container's associated field ID.
	else
	{
		// No value exits
		$checkBoxControl.prop("indeterminate", true)
		$checkboxLabel.removeClass("checkboxStrikethroughCompleted")
	}	
	
}

function initCheckBoxFieldEditBehavior($checkBox,componentContext,recordProxy, checkBoxObjectRef) {
	
	var $checkboxControl = getCheckboxControlFromCheckboxContainer($checkBox)
		
	var fieldID = checkBoxObjectRef.properties.fieldID
	var fieldRef = getFieldRef(fieldID)
	if(fieldRef.isCalcField || formComponentIsReadOnly(checkBoxObjectRef.properties.permissions)) {
		$checkboxControl.prop('disabled',true)
		return;  // stop initialization, the check box is read only.
	}
	
	function setBoolValue(boolVal) {
		var currRecordRef = recordProxy.getRecordFunc()
		var checkboxValueFormat = {
			context: "checkbox",
			format: "check"
		}
		var setRecordValParams = {
			parentDatabaseID:currRecordRef.parentDatabaseID,
			recordID:currRecordRef.recordID,
			changeSetID: recordProxy.changeSetID,
			fieldID:fieldID, 
			value:boolVal,
			 valueFormat: checkboxValueFormat }
		console.log("Setting boolean value (record): " + JSON.stringify(setRecordValParams))
		jsonAPIRequest("recordUpdate/setBoolFieldValue",setRecordValParams,function(updatedRecordRef) {
			
			// After updating the record, the local cache of records in currentRecordSet will
			// be out of date. So after updating the record on the server, the locally cached
			// version of the record also needs to be updated.
			recordProxy.updateRecordFunc(updatedRecordRef)
		}) // set record's text field value
		
	}
	
	
	var $clearValueButton = $checkBox.find(".checkBoxComponentClearValueButton")
	initButtonControlClickHandler($clearValueButton,function() {
		console.log("Clear value clicked for check box")
		setBoolValue(null)
	})
	
		
	$checkboxControl.unbind("click")
  	$checkboxControl.click( function () {
		// Get the most recent copy of the object reference. It could have changed between
		// initialization time and the time the checkbox was changed.
		var objectRef = getContainerObjectRef($checkBox)
		
		var isChecked = $(this).prop("checked")
	
		setBoolValue(isChecked)	
	})
	
}

function initCheckBoxRecordEditBehavior($checkBox,componentContext,recordProxy, checkBoxObjectRef) {
		
	$checkBox.data("viewFormConfig", {
		loadRecord: loadRecordIntoCheckBox
	})
	initCheckBoxFieldEditBehavior($checkBox,componentContext,recordProxy, checkBoxObjectRef)
	
}