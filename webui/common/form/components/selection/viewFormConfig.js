function loadRecordIntoSelection(selectionElem, recordRef) {

	var selectionObjectRef = selectionElem.data("objectRef")

	var $selectionControl = selectionFormControlFromSelectionFormComponent(selectionElem)	
	
	var selectionFieldID = selectionObjectRef.properties.fieldID
	console.log("loadRecordIntoSelection: Field ID to load data:" + selectionFieldID)	

	// In other words, we are populating the "intersection" of field values in the record
	// with the fields shown by the layout's containers.
	if(recordRef.fieldValues.hasOwnProperty(selectionFieldID)) {

		var fieldVal = recordRef.fieldValues[selectionFieldID]

		console.log("loadRecordIntoSelection: Load value into container: " + " field ID:" + 
					selectionFieldID + "  value:" + fieldVal)

		$selectionControl.val(fieldVal.toString())
	} // If record has a value for the current container's associated field ID.
	else
	{
		$selectionControl.val("")
	}	
	
}

function initSelectionRecordEditBehavior($selectionContainer,componentContext,recordProxy, selectionObjectRef) {
	
	$selectionContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoSelection
	})
	
	var $selectionControl = selectionFormControlFromSelectionFormComponent($selectionContainer)
	
	
	$selectionControl.append(defaultSelectOptionPromptHTML("Select a Value"))
	for(var selValIndex = 0; selValIndex < selectionObjectRef.properties.selectableVals.length; selValIndex++) {
		var selectableVal = selectionObjectRef.properties.selectableVals[selValIndex]	
		$selectionControl.append(selectOptionHTML(selectableVal.val,selectableVal.label))
	}
	
	// When the user clicks on the control, prevent the click from propagating higher.
	// This allows the user to change the rating without selecting the form component itself.
	// The user can still select the component by clicking on the label or anywwhere outside
	// the control.
	$selectionContainer.find(".selectionFormControl").click(function (event){
		event.stopPropagation();
		return false;
	});



	var selectionFieldID = selectionObjectRef.properties.fieldID
	
	initSelectControlChangeHandler($selectionControl,function(newValue) {
		
		var currRecordRef = recordProxy.getRecordFunc()	
		var fieldID = selectionFieldID
		var fieldRef = getFieldRef(selectionFieldID)
		var fieldType = fieldRef.type
		if(fieldType == "text") {
			currRecordRef.fieldValues[selectionFieldID] = newValue
			var setTextFieldValueFormat = {
				context: "select",
				format:"general" 
			}
			var setRecordValParams = { 
				parentDatabaseID:currRecordRef.parentDatabaseID,
				recordID:currRecordRef.recordID, 
				changeSetID: recordProxy.changeSetID,
				fieldID:selectionFieldID, value:newValue,
				 valueFormat:setTextFieldValueFormat}
			jsonAPIRequest("recordUpdate/setTextFieldValue",setRecordValParams,function(replyData) {
				// After updating the record, the local cache of records in currentRecordSet will
				// be out of date. So after updating the record on the server, the locally cached
				// version of the record also needs to be updated.
				recordProxy.updateRecordFunc(replyData)
			}) // set record's text field value
		
		} else if (fieldType == "number") {
			var numberVal = Number(newValue)
			if(!isNaN(numberVal)) {
				console.log("Change number val: "
					+ "fieldID: " + fieldID
				    + " ,number = " + numberVal)
				currRecordRef.fieldValues[fieldID] = numberVal
				var setNumberFieldValueFormat = {
					context: "select",
					format:"general" 
				}			
				var setRecordValParams = { 
					parentDatabaseID:currRecordRef.parentDatabaseID,
					recordID:currRecordRef.recordID, 
					changeSetID: recordProxy.changeSetID,
					fieldID:selectionFieldID, 
					value:numberVal,
					 valueFormat:setNumberFieldValueFormat}
				jsonAPIRequest("recordUpdate/setNumberFieldValue",setRecordValParams,function(replyData) {
					// After updating the record, the local cache of records will
					// be out of date. So after updating the record on the server, the locally cached
					// version of the record also needs to be updated.
					recordProxy.updateRecordFunc(replyData)
				}) // set record's number field value
			}
		
		}
				
	})
	
	
}