

function initUserSelectionRecordEditBehavior($userSelectionContainer, componentContext,
		recordProxy, userSelectionObjectRef, validateInputFunc) {

	var selectionFieldID = userSelectionObjectRef.properties.fieldID
		
	var $userSelectionControl = userSelectionControlFromUserSelectionComponentContainer($userSelectionContainer)

	var validateInput = function(validationCompleteCallback) {
		if($userSelectionControl.prop('disabled')) {
			validationCompleteCallback(true)
			return
		}
		var currVal = $userSelectionControl.val()
		validateInputFunc(currVal,function(validationResult) {
			setupFormComponentValidationPrompt($userSelectionContainer,validationResult,validationCompleteCallback)			
		})	
	}

	function loadRecordIntoUserSelection(userSelectionElem, recordRef) {

		var $userSelectionControl = userSelectionControlFromUserSelectionComponentContainer(userSelectionElem)

		if(formComponentIsReadOnly(userSelectionObjectRef.properties.permissions)) {
			$userSelectionControl.prop('disabled',true);
		} else {
			$userSelectionControl.prop('disabled',false);
		
		}

		var userSelectionFieldID = userSelectionObjectRef.properties.fieldID

		console.log("loadRecordIntoUserSelection: Field ID to load data:" + userSelectionFieldID)

		// In other words, we are populating the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(userSelectionFieldID)) {

			var fieldVal = recordRef.fieldValues[userSelectionFieldID]
			if (fieldVal === null) {
				clearUserSelectionControlVal($userSelectionControl)
			} else {
				setSingleUserSelectionControlVal($userSelectionControl,fieldVal)		
			}

		} // If record has a value for the current container's associated field ID.
		else
		{
			clearUserSelectionControlVal($userSelectionControl)
		}
		
	}


	function initUserSelectionEditBehavior() {
		
		function setUserSelectionValue(selectedUserID) {
		
			validateInput(function(inputIsValid) {
				if (inputIsValid) {
					currRecordRef = recordProxy.getRecordFunc()

					var userFieldID = selectionFieldID

					var setRecordValParams = { 
						parentDatabaseID:currRecordRef.parentDatabaseID,
						recordID:currRecordRef.recordID,
						changeSetID: recordProxy.changeSetID, 
						fieldID:userFieldID, 
						userID:selectedUserID}
					jsonAPIRequest("recordUpdate/setUserFieldValue",setRecordValParams,function(updatedFieldVal) {
						// After updating the record, the local cache of records in currentRecordSet will
						// be out of date. So after updating the record on the server, the locally cached
						// version of the record also needs to be updated.
						recordProxy.updateRecordFunc(updatedFieldVal)
	
					}) // set record's number field value
				}
			})
		}
	
		var userSelectionParams = {
			$selectionInput: $userSelectionControl,
			databaseID: componentContext.databaseID,
		}
		initCollaboratorUserSelection(userSelectionParams)
		
		function setSelectedUserFromDropdownMenu(userInfo) {
			var userIDSelection = userInfo.userID
			// For the selected user to be displayed in the selection,
			// it needs to be added as an option.
			addUserInfoSelectionOptionIfNotExists($userSelectionControl,userInfo)
			$userSelectionControl.val(userIDSelection).trigger("change")
			setUserSelectionValue(userIDSelection)
		}
		configureUserSelectionDropdown(componentContext,$userSelectionContainer,
					userSelectionObjectRef,setSelectedUserFromDropdownMenu)
		
	
		var $clearValueButton = $userSelectionContainer.find(".userSelectionComponentClearValueButton")
		initButtonControlClickHandler($clearValueButton,function() {
			console.log("Clear value clicked for user selection")
			clearUserSelectionControlVal($userSelectionControl)
			setUserSelectionValue(null)
		})

		$userSelectionControl.on('select2:select', function(e) {
			var selectedUserID = e.params.data.id
			console.log('User selection changed: ' + selectedUserID);
			setUserSelectionValue(selectedUserID)
		});
		
	}
	initUserSelectionEditBehavior()

	
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
		loadRecord: loadRecordIntoUserSelection,
		validateValue: validateInput
	})
	

}


function initUserSelectionFormRecordEditBehavior($container,componentContext,recordProxy, userSelectionObjectRef) {
	
	function validateInput(inputVal,validationResultCallback) {
		var validationParams = {
			parentFormID: userSelectionObjectRef.parentFormID,
			userSelectionID: userSelectionObjectRef.userSelectionID,
			inputVal: inputVal
		}
		jsonAPIRequest("frm/userSelection/validateInput", validationParams, function(validationResult) {
			validationResultCallback(validationResult)
		})
	}
	
	var selectionWidth = userSelectionObjectRef.properties.geometry.sizeWidth - 15
	
	initUserSelectionRecordEditBehavior($container,componentContext,recordProxy, 
			userSelectionObjectRef, validateInput)
}



function initUserSelectionTableRecordEditBehavior($container,componentContext,recordProxy, userSelectionObjectRef) {
	
	function validateInput(inputVal,validationResultCallback) {
		var validationParams = {
			parentTableID: userSelectionObjectRef.parentTableID,
			userSelectionID: userSelectionObjectRef.userSelectionID,
			inputVal: inputVal
		}
		jsonAPIRequest("tableView/userSelection/validateInput", validationParams, function(validationResult) {
			validationResultCallback(validationResult)
		})
	}
	
	// The width of the select2 control is setup during the controls initialization. However, the height
	// is configured via CSS. See the CSS file for corresponding CSS to configure the height.	
	initUserSelectionClearValueButton($container,userSelectionObjectRef)
	
	initUserSelectionRecordEditBehavior($container,componentContext,recordProxy, 
			userSelectionObjectRef, validateInput)
}

