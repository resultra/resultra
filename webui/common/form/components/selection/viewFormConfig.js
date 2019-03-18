// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function loadRecordIntoSelection(selectionElem, recordRef) {

	var selectionObjectRef = selectionElem.data("objectRef")

	var $selectionControl = selectionFormControlFromSelectionFormComponent(selectionElem)
	
	if(formComponentIsReadOnly(selectionObjectRef.properties.permissions)) {
		$selectionControl.prop('disabled',true);
	} else {
		$selectionControl.prop('disabled',false);
		
	}	
	
	var selectionFieldID = selectionObjectRef.properties.fieldID
	console.log("loadRecordIntoSelection: Field ID to load data:" + selectionFieldID)	

	// In other words, we are populating the "intersection" of field values in the record
	// with the fields shown by the layout's containers.
	if(recordRef.fieldValues.hasOwnProperty(selectionFieldID)) {

		var fieldVal = recordRef.fieldValues[selectionFieldID]
		
		if (fieldVal === null) {
			$selectionControl.val("")
		} else {
			console.log("loadRecordIntoSelection: Load value into container: " + " field ID:" + 
						selectionFieldID + "  value:" + fieldVal)

			$selectionControl.val(fieldVal.toString())
		}

	} // If record has a value for the current container's associated field ID.
	else
	{
		$selectionControl.val("")
	}	
	
}

function initSelectionRecordEditBehavior($selectionContainer,
			componentContext,recordProxy, selectionObjectRef,initDoneCallback) {
	
	$selectionContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoSelection
	})
	
	var $selectionControl = selectionFormControlFromSelectionFormComponent($selectionContainer)
	
		
	
	// When the user clicks on the control, prevent the click from propagating higher.
	// This allows the user to change the rating without selecting the form component itself.
	// The user can still select the component by clicking on the label or anywwhere outside
	// the control.
	$selectionContainer.find(".selectionFormControl").click(function (event){
		event.stopPropagation();
		return false;
	});
	
	var selectionFieldID = selectionObjectRef.properties.fieldID
	var selectionFieldType = getFieldRef(selectionFieldID).type
	
	
	function setTextVal(newValue) {
		var currRecordRef = recordProxy.getRecordFunc()	
		var setRecordValParams = { 
			parentDatabaseID:currRecordRef.parentDatabaseID,
			recordID:currRecordRef.recordID, 
			changeSetID: recordProxy.changeSetID,
			fieldID:selectionFieldID, 
			value:newValue}
		jsonAPIRequest("recordUpdate/setTextFieldValue",setRecordValParams,function(replyData) {
			// After updating the record, the local cache of records in currentRecordSet will
			// be out of date. So after updating the record on the server, the locally cached
			// version of the record also needs to be updated.
			recordProxy.updateRecordFunc(replyData)
		}) // set record's text field value
		
	}
	
	function setNumberVal(numberVal) {
		var currRecordRef = recordProxy.getRecordFunc()
		var setRecordValParams = { 
			parentDatabaseID:currRecordRef.parentDatabaseID,
			recordID:currRecordRef.recordID, 
			changeSetID: recordProxy.changeSetID,
			fieldID:selectionFieldID, 
			value:numberVal
		}
		jsonAPIRequest("recordUpdate/setNumberFieldValue",setRecordValParams,function(replyData) {
			// After updating the record, the local cache of records will
			// be out of date. So after updating the record on the server, the locally cached
			// version of the record also needs to be updated.
			recordProxy.updateRecordFunc(replyData)
		}) // set record's number field value
		
	}
	
	var $clearValueButton = $selectionContainer.find(".selectComponentClearValueButton")
	initButtonControlClickHandler($clearValueButton,function() {
		console.log("Clear value clicked for text box")
		if(selectionFieldType == fieldTypeNumber) {
			setNumberVal(null)
		} else {
			setTextVal(null)
		}		
	})
	
	initSelectControlChangeHandler($selectionControl,function(newValue) {
		if(selectionFieldType == "text") {
			setTextVal(newValue)
		} else if (selectionFieldType == "number") {
			var numberVal = Number(newValue)
			if(!isNaN(numberVal)) {
				setNumberVal(numberVal)
			}		
		}	
	})
	
	// Populate the selection with values from the value list. Initialization of the selection when 
	// in view mode happens asynchronously, since the value list is retrieved from the server. 
	// This value list must be in place before any records are loaded into the control. 
	$selectionControl.empty()
	$selectionControl.append(defaultSelectOptionPromptHTML("Select a Value"))
	var valueListID = selectionObjectRef.properties.valueListID
	if (valueListID !== undefined && valueListID !== null) {
		var getValListParams = { valueListID: valueListID }
		jsonAPIRequest("valueList/get",getValListParams,function(valListInfo) {
			var values = valListInfo.properties.values
			for(var valIndex = 0; valIndex < values.length; valIndex++) {
				var val = values[valIndex]
				$selectionControl.append(selectOptionHTML(val.textValue,val.textValue))
			}
			initDoneCallback()					
		})
	} else {
		initDoneCallback()
	}
	
	
}

function initTextSelectionRecordEditTableBehavior($container,componentContext,recordProxy, textSelectionObjectRef) {
	
	function initDoneCallback() {
		// TBD
	}
	
	initSelectionRecordEditBehavior($container,
				componentContext,recordProxy, textSelectionObjectRef,initDoneCallback)
	
}