

function initToggleRecordEditBehavior($toggle,componentContext,recordProxy, toggleObjectRef) {
	
	var validateToggleInput = function(validationCompleteCallback) {
		
		if(toggleComponentIsDisabled($toggle)) {
			validationCompleteCallback(true)
			return
		}
		
		var currVal = getCurrentToggleComponentValue($toggle)
		var validationParams = {
			parentFormID: toggleObjectRef.parentFormID,
			toggleID: toggleObjectRef.toggleID,
			inputVal: currVal
		}
		jsonAPIRequest("frm/toggle/validateInput", validationParams, function(validationResult) {
			if (validationResult.validationSucceeded) {
				$toggle.popover('destroy')
				validationCompleteCallback(true)
			} else {
				$toggle.popover({
					html: 'true',
					content: function() { return escapeHTML(validationResult.errorMsg) },
					trigger: 'manual',
					placement: 'auto left'
				})
				$toggle.popover('show')
				validationCompleteCallback(false)
			}
			
		})	
		
	}
	
	

	function loadRecordIntoToggle($toggleContainer, recordRef) {
	
		console.log("loadRecordIntoToggle: loading record into text box: " + JSON.stringify(recordRef))
	
		var toggleObjectRef = getContainerObjectRef($toggleContainer)
		var $toggleControl = getToggleControlFromToggleContainer($toggleContainer);
		var $toggleLabel = $toggleContainer.find("label")
	
		var toggleFieldID = toggleObjectRef.properties.fieldID

		console.log("loadRecordIntoToggle: Field ID to load data:" + toggleFieldID)
	
		// Populate the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(toggleFieldID)) {

			var fieldVal = recordRef.fieldValues[toggleFieldID]
		
			if (fieldVal === null) {
				$toggleControl.bootstrapSwitch('indeterminate',true)
			} else {
				if(fieldVal == true)
				{
					$toggleControl.bootstrapSwitch('indeterminate',false)
					$toggleControl.bootstrapSwitch('state',true)
				}
				else {
					$toggleLabel.removeClass("toggleStrikethroughCompleted")
					$toggleControl.bootstrapSwitch('indeterminate',false)
					$toggleControl.bootstrapSwitch('state',false)
				}
			
			}
		} // If record has a value for the current container's associated field ID.
		else
		{
			// No value exits
			$toggleControl.bootstrapSwitch('indeterminate',true)
		}	
	
	}



	function initToggleFieldEditBehavior($toggle,componentContext,recordProxy, toggleObjectRef) {
	
		var $toggleControl = getToggleControlFromToggleContainer($toggle)
		
		 $toggleControl.bootstrapSwitch({
			handleWidth:40,
			indeterminate:true,
			onText:'Yes',
			offText:'No',
			labelWidth:5 ,
			onColor:'success',
			offColor:'warning'
		});
		
		var fieldID = toggleObjectRef.properties.fieldID
		var fieldRef = getFieldRef(fieldID)
		if(fieldRef.isCalcField || formComponentIsReadOnly(toggleObjectRef.properties.permissions)) {
			$toggleControl.prop('disabled',true)
			return;  // stop initialization, the check box is read only.
		}
	
		function setBoolValue(boolVal) {
			
			validateToggleInput(function(inputIsValid) {
				
				if(inputIsValid) {
					var currRecordRef = recordProxy.getRecordFunc()
					var toggleValueFormat = {
						context: "toggle",
						format: "toggle"
					}
					var setRecordValParams = {
						parentDatabaseID:currRecordRef.parentDatabaseID,
						recordID:currRecordRef.recordID,
						changeSetID: recordProxy.changeSetID,
						fieldID:fieldID, 
						value:boolVal,
						 valueFormat: toggleValueFormat }
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