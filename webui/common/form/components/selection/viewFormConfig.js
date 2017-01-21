function loadRecordIntoSelection(selectionElem, recordRef) {

	var selectionObjectRef = selectionElem.data("objectRef")
	var selectionControlSelector = '#'+selectionFormControlID(selectionObjectRef.selectionID)
	var $selectionControl = $(selectionControlSelector)

	
	console.log("loadRecordIntoSelection: loading record into selection: " + JSON.stringify(recordRef))
	console.log("loadRecordIntoSelection: loading record into selection: " + 
		" selectionID = " + selectionObjectRef.selectionID + 
		" selector=" + selectionControlSelector)
	
	
	var componentLink = selectionObjectRef.properties.componentLink
	
	if(componentLink.linkedValType == linkedComponentValTypeField) {
		// text box is linked to a field value
		var selectionFieldID = componentLink.fieldID
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
		
	} else {
		// Text box is linked to a global value
		
		var selectionGlobalID = componentLink.globalID
		if(selectionGlobalID in currGlobalVals) {
			var globalVal = currGlobalVals[selectionGlobalID]
			$selectionControl.val(globalVal.toString())
		}
		else
		{
			$selectionControl.val("")
		}		
	}
	
	
	
}

function initSelectionRecordEditBehavior(componentContext,getRecordFunc, updateRecordFunc, selectionObjectRef) {
	
	var container = $('#'+selectionObjectRef.selectionID)

	container.data("viewFormConfig", {
		loadRecord: loadRecordIntoSelection
	})
	
	var selectionControlSelector = '#' + selectionFormControlID(selectionObjectRef.selectionID)
	var $selectionControl = $(selectionControlSelector)
	$selectionControl.append(defaultSelectOptionPromptHTML("Select a Value"))
	for(var selValIndex = 0; selValIndex < selectionObjectRef.properties.selectableVals.length; selValIndex++) {
		var selectableVal = selectionObjectRef.properties.selectableVals[selValIndex]	
		$selectionControl.append(selectOptionHTML(selectableVal.val,selectableVal.label))
	}
	
	// When the user clicks on the control, prevent the click from propagating higher.
	// This allows the user to change the rating without selecting the form component itself.
	// The user can still select the component by clicking on the label or anywwhere outside
	// the control.
	container.find(".selectionFormControl").click(function (event){
		event.stopPropagation();
		return false;
	});



	var componentLink = selectionObjectRef.properties.componentLink
	var currRecordRef = 
	
	initSelectionChangedHandler(selectionControlSelector,function(newValue) {
		if(componentLink.linkedValType == linkedComponentValTypeField) {
			var currRecordRef = getRecordFunc()	
			var fieldID = componentLink.fieldID
			var fieldRef = getFieldRef(fieldID)
			var fieldType = fieldRef.type
			if(fieldType == "text") {
				currRecordRef.fieldValues[fieldID] = newValue
				var setTextFieldValueFormat = {
					context: "select",
					format:"general" 
				}
				var setRecordValParams = { 
					parentDatabaseID:currRecordRef.parentDatabaseID,
					recordID:currRecordRef.recordID, 
					fieldID:fieldID, value:newValue,
					 valueFormat:setTextFieldValueFormat}
				jsonAPIRequest("recordUpdate/setTextFieldValue",setRecordValParams,function(replyData) {
					// After updating the record, the local cache of records in currentRecordSet will
					// be out of date. So after updating the record on the server, the locally cached
					// version of the record also needs to be updated.
					updateRecordFunc(replyData)
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
						fieldID:fieldID, 
						value:numberVal,
						 valueFormat:setNumberFieldValueFormat}
					jsonAPIRequest("recordUpdate/setNumberFieldValue",setRecordValParams,function(replyData) {
						// After updating the record, the local cache of records in currentRecordSet will
						// be out of date. So after updating the record on the server, the locally cached
						// version of the record also needs to be updated.
						updateRecordFunc(replyData)
					}) // set record's number field value
				}
			
			}
		
		} else { 
			assert(componentLink.linkedValType == linkedComponentValTypeGlobal)
			if(globalInfo.type == globalTypeNumber) {
				var numberVal = Number(inputVal)
				if(!isNaN(numberVal)) {
					var setGlobalValParams = {
						parentDatabaseID: componentContext.databaseID,
						globalID: componentLink.globalID,
						value: numberVal
					}
					console.log("Setting global value (number): " + JSON.stringify(setGlobalValParams))
					jsonAPIRequest("global/setNumberValue",setGlobalValParams,function(replyData) {
						// TODO - Update record after recalculation based upon global values
					})
				
				}
			} else {
				var setGlobalValParams = {
					parentDatabaseID: componentContext.databaseID,
					globalID: componentLink.globalID,
					value: inputVal
				}
				console.log("Setting global value (text): " + JSON.stringify(setGlobalValParams))
				jsonAPIRequest("global/setTextValue",setGlobalValParams,function(replyData) {
						// TODO - Update record after recalculation based upon global values
				})	
			}
		
		}
		
	})
	
	
}