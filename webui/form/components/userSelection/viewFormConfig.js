function loadRecordIntoUserSelection(userSelectionElem, recordRef) {
	
	
	var userSelectionObjectRef = userSelectionElem.data("objectRef")
	var userSelectionContainerID = userSelectionObjectRef.userSelectionID
	var userSelectionControlID = userSelectionIDFromElemID(userSelectionContainerID)
	var userSelectionControlSelector = '#' + userSelectionControlID;
	
	var componentLink = userSelectionObjectRef.properties.componentLink
	
	if(componentLink.linkedValType == linkedComponentValTypeField) {
		var userSelectionFieldID = componentLink.fieldID
	
		console.log("loadRecordIntoUserSelection: Field ID to load data:" + userSelectionFieldID)
	
		// In other words, we are populating the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(userSelectionFieldID)) {

			var fieldVal = recordRef.fieldValues[userSelectionFieldID]

			console.log("loadRecordIntoTextBox: Load value into container: " + $(this).attr("id") + " field ID:" + 
						userSelectionFieldID + "  value:" + fieldVal)
			
			// TODO - set user settings
			$(userSelectionControlSelector).val("TBD")
			
			
		} // If record has a value for the current container's associated field ID.
		else
		{
			$(userSelectionControlSelector).val("TBD")
		}
	} else {
		var userSelectionGlobalID = componentLink.globalID
		console.log("loadRecordIntoUserSelection: Global ID to load data:" + userSelectionGlobalID)
		
		if(userSelectionGlobalID in currGlobalVals) {
			var globalVal = currGlobalVals[userSelectionGlobalID]
			
			$(userSelectionControlSelector).val("Global value - TBD")
			
		}
		else
		{
			$(userSelectionControlSelector).val("")
		}				
	}
	
	
}


function initUserSelectionRecordEditBehavior(componentContext,userSelectionObjectRef) {

	var $userSelectionContainer = $('#'+userSelectionObjectRef.userSelectionID)
	
	var componentLink = userSelectionObjectRef.properties.componentLink
	
	var userSelectionControlSelector = '#' + userSelectionIDFromElemID(userSelectionObjectRef.userSelectionID)
	

	function setUserSelectionValue(userVal) {
		
		currRecordRef = currRecordSet.currRecordRef()
	
		if(componentLink.linkedValType == linkedComponentValTypeField) {
			var userFieldID = componentLink.fieldID
	
			var setRecordValParams = { 
				parentTableID:viewFormContext.tableID,
				recordID:currRecordRef.recordID, 
				fieldID:userFieldID, 
				value:userVal }
			jsonAPIRequest("recordUpdate/setUserFieldValue",setRecordValParams,function(updatedFieldVal) {
				// After updating the record, the local cache of records in currentRecordSet will
				// be out of date. So after updating the record on the server, the locally cached
				// version of the record also needs to be updated.
				currRecordSet.updateRecordRef(updatedFieldVal)
		
				// After changing the value, some of the calculated fields may have changed. For this
				// reason, it is necessary to reload the record into the layout/form, so the most
				// up to date values will be displayed.
				loadCurrRecordIntoLayout()
			}) // set record's number field value

			// TBD - initialize control
		
		} else {
			var userGlobalID = componentLink.globalID
			console.log("loadRecordIntoUserSelection: Global ID to load data:" + userGlobalID)
			
			var setGlobalValParams = {
				parentDatabaseID: componentContext.databaseID,
				globalID: componentLink.globalID,
				value: userVal
			}
			console.log("Setting global value (user): " + JSON.stringify(setGlobalValParams))
			jsonAPIRequest("global/setUserValue",setGlobalValParams,function(replyData) {
			})
					
		}
		
	}

	var $userSelectionControl = $(userSelectionControlSelector)
	
	var userSelectionParams = {
		selectionInput: $userSelectionControl,
		dropdownParent: $userSelectionContainer,
		width: '150px'
	}
	
	initUserSelection(userSelectionParams)

	$userSelectionControl.on('change', function() {
		var userVal = $(this).val()
		console.log('User selection changed: ' + userVal);
		setUserSelectionValue(userVal)
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