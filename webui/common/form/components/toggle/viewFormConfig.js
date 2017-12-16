


function initToggleRecordEditBehavior($toggle,componentContext,recordProxy, toggleObjectRef,remoteValidate) {
	
	var validateToggleInput = function(validationCompleteCallback) {
		
		if(toggleComponentIsDisabled($toggle)) {
			validationCompleteCallback(true)
			return
		}
		
		var currVal = getCurrentToggleComponentValue($toggle)
		remoteValidate(currVal,function(validationResult) {
			setupFormComponentValidationPrompt($toggle,validationResult,validationCompleteCallback)			
		})	
		
	}
	
	

	function loadRecordIntoToggle($toggleContainer, recordRef) {
	
		console.log("loadRecordIntoToggle: loading record into text box: " + JSON.stringify(recordRef))
	
		var toggleObjectRef = getContainerObjectRef($toggleContainer)
		var $toggleControl = getToggleControlFromToggleContainer($toggleContainer);
		var $toggleLabel = $toggleContainer.find("label")
	
		var toggleFieldID = toggleObjectRef.properties.fieldID

		console.log("loadRecordIntoToggle: Field ID to load data:" + toggleFieldID)
		
		function setToggleIndeterminate() {
			$toggleControl.bootstrapSwitch('indeterminate',true)
			
			// NOTE: There's an issue with bootstrap switch which prevents the change
			// event from firing if the value is already set to false. The following 
			// workaround ensures the event is fired.
			// See: https://github.com/Bttstrp/bootstrap-switch/issues/426
			$toggleControl.data('bootstrap-switch').options.state = null; 

		}
		
		// To initialize the control the readonly option needs to first be
		// disabled. Then, after setting the value, the readonly value can be restored (see below)
		var isReadonly = $toggleControl.bootstrapSwitch("readonly")
		$toggleControl.bootstrapSwitch("readonly",false)
		
	
		// Populate the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(toggleFieldID)) {

			var fieldVal = recordRef.fieldValues[toggleFieldID]
			
			// When initially loading the record, don't fire the change event.
			var skipSwitchChangeEventFiring = true
		
			if (fieldVal === null) {
				setToggleIndeterminate()
			} else {
				if(fieldVal == true)
				{
					$toggleControl.bootstrapSwitch('indeterminate',false)
					$toggleControl.bootstrapSwitch('state',true,skipSwitchChangeEventFiring)
				}
				else {
					$toggleLabel.removeClass("toggleStrikethroughCompleted")
					$toggleControl.bootstrapSwitch('indeterminate',false)
					$toggleControl.bootstrapSwitch('state',false,skipSwitchChangeEventFiring)
				}
			
			}
		} // If record has a value for the current container's associated field ID.
		else
		{
			// No value exits
			setToggleIndeterminate()
		}
		
		// Restore the read-only state of the control.
		$toggleControl.bootstrapSwitch("readonly",isReadonly)
			
	
	}



	function initToggleFieldEditBehavior($toggle,componentContext,recordProxy, toggleObjectRef) {
			
		var fieldID = toggleObjectRef.properties.fieldID
		var fieldRef = getFieldRef(fieldID)
		if(fieldRef.isCalcField || formComponentIsReadOnly(toggleObjectRef.properties.permissions)) {
			var $toggleControl = getToggleControlFromToggleContainer($toggle)
//			$toggleControl.bootstrapSwitch("disabled",true)
			$toggleControl.bootstrapSwitch("readonly",true)
			return;  // stop initialization, the toggle is read only.
		}
	
		function setBoolValue(boolVal) {
			
			validateToggleInput(function(inputIsValid) {
				
				if(inputIsValid) {
					var currRecordRef = recordProxy.getRecordFunc()
					var setRecordValParams = {
						parentDatabaseID:currRecordRef.parentDatabaseID,
						recordID:currRecordRef.recordID,
						changeSetID: recordProxy.changeSetID,
						fieldID:fieldID, 
						value:boolVal}
					console.log("Setting boolean value (record): " + JSON.stringify(setRecordValParams))
					jsonAPIRequest("recordUpdate/setBoolFieldValue",setRecordValParams,function(updatedRecordRef) {
			
						// After updating the record, the local cache of records in currentRecordSet will
						// be out of date. So after updating the record on the server, the locally cached
						// version of the record also needs to be updated.
						recordProxy.updateRecordFunc(updatedRecordRef)
					}) // set record's text field value
					
				}
				
			})
		
		}
	
	
		var $clearValueButton = $toggle.find(".toggleComponentClearValueButton")
		initButtonControlClickHandler($clearValueButton,function() {
			console.log("Clear value clicked for toggle")
			$toggleControl.bootstrapSwitch('indeterminate',true)
			setBoolValue(null)
		})
	
		var $toggleControl = getToggleControlFromToggleContainer($toggle)
		
	  	$toggleControl.on('switchChange.bootstrapSwitch', function(event, state) {
			var toggleVal = getCurrentToggleComponentValue($toggle)
			setBoolValue(toggleVal)	
		})
	
	}
			
	$toggle.data("viewFormConfig", {
		loadRecord: loadRecordIntoToggle,
		validateValue: validateToggleInput
	})
	initToggleFieldEditBehavior($toggle,componentContext,recordProxy, toggleObjectRef)
	
}


function initToggleFormRecordEditBehavior($toggle,componentContext,recordProxy, toggleObjectRef) {
	function validatInput(currVal, validateResultsCallback) {
		var validationParams = {
			parentFormID: toggleObjectRef.parentFormID,
			toggleID: toggleObjectRef.toggleID,
			inputVal: currVal
		}
		jsonAPIRequest("frm/toggle/validateInput", validationParams, function(validationResult) {
			validateResultsCallback(validationResult)
		})
	}
	initToggleRecordEditBehavior($toggle,componentContext,recordProxy, toggleObjectRef,validatInput)
}



function initToggleTableCellRecordEditBehavior($toggle,componentContext,recordProxy, toggleObjectRef) {
	
	function validatInput(currVal, validateResultsCallback) {
		var validationParams = {
			parentTableID: toggleObjectRef.parentTableID,
			toggleID: toggleObjectRef.toggleID,
			inputVal: currVal
		}
		jsonAPIRequest("tableView/toggle/validateInput", validationParams, function(validationResult) {
			validateResultsCallback(validationResult)
		})
	}
	
		initToggleComponentTableViewComponentContainer($toggle,toggleObjectRef)
		
	
	
	initToggleRecordEditBehavior($toggle,componentContext,recordProxy, toggleObjectRef,validatInput)
}

