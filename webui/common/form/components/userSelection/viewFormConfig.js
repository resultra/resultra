function loadRecordIntoUserSelection(userSelectionElem, recordRef) {
	
	
	var userSelectionObjectRef = userSelectionElem.data("objectRef")
	
	var $userSelectionControl = userSelectionControlFromUserSelectionComponentContainer(userSelectionElem)
	
	var userSelectionFieldID = userSelectionObjectRef.properties.fieldID

	console.log("loadRecordIntoUserSelection: Field ID to load data:" + userSelectionFieldID)

	// In other words, we are populating the "intersection" of field values in the record
	// with the fields shown by the layout's containers.
	if(recordRef.fieldValues.hasOwnProperty(userSelectionFieldID)) {

		var fieldVal = recordRef.fieldValues[userSelectionFieldID]

		setUserSelectionControlVal($userSelectionControl,fieldVal)

	} // If record has a value for the current container's associated field ID.
	else
	{
		clearUserSelectionControlVal($userSelectionControl)
	}
		
}


function initUserSelectionRecordEditBehavior($userSelectionContainer, componentContext,recordProxy, userSelectionObjectRef) {

	var selectionFieldID = userSelectionObjectRef.properties.fieldID
		
	var $userSelectionControl = userSelectionControlFromUserSelectionComponentContainer($userSelectionContainer)

	function setUserSelectionValue(selectedUserID) {
		
		currRecordRef = recordProxy.getRecordFunc()

		var userFieldID = selectionFieldID

		var userValueFormat = {
			context: "selectUser",
			format: "general"
		}
		var setRecordValParams = { 
			parentDatabaseID:currRecordRef.parentDatabaseID,
			recordID:currRecordRef.recordID,
			changeSetID: recordProxy.changeSetID, 
			fieldID:userFieldID, 
			userID:selectedUserID,
			valueFormat:userValueFormat}
		jsonAPIRequest("recordUpdate/setUserFieldValue",setRecordValParams,function(updatedFieldVal) {
			// After updating the record, the local cache of records in currentRecordSet will
			// be out of date. So after updating the record on the server, the locally cached
			// version of the record also needs to be updated.
			recordProxy.updateRecordFunc(updatedFieldVal)
	
		}) // set record's number field value

		
	}
	
	var selectionWidth = (userSelectionObjectRef.properties.geometry.sizeWidth - 15).toString() + "px"
	var userSelectionParams = {
		selectionInput: $userSelectionControl,
		dropdownParent: $userSelectionContainer,
		width: selectionWidth
	}
	
	initUserSelection(userSelectionParams)

	$userSelectionControl.on('change', function() {
		var selectedUserID = $(this).val()
		console.log('User selection changed: ' + selectedUserID);
		setUserSelectionValue(selectedUserID)
	});
	
	// When the user clicks on the control, prevent the click from propagating higher.
	// This allows the user to change the rating without selecting the form component itself.
	// The user can still select the component by clicking on the label or anywwhere outside
	// the control.
	$userSelectionContainer.find(".formUserSelectionControl").click(function (event){
		event.stopPropagation();
   	 	//   ... your code here
		return false;
	});
		
	$userSelectionContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoUserSelection
	})
	

}