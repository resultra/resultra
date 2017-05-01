

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
					// Value has been cleared
					$toggleLabel.removeClass("toggleStrikethroughCompleted")
					$toggleControl.prop("indeterminate", true)
			} else {
				$toggleLabel.removeClass("toggleStrikethroughCompleted")

				if(fieldVal == true)
				{
					$toggleControl.prop("indeterminate", false)
					$toggleControl.prop("checked",true)
					if(toggleObjectRef.properties.strikethroughCompleted) {
						$toggleLabel.addClass("toggleStrikethroughCompleted")
					} else {
						$toggleLabel.removeClass("toggleStrikethroughCompleted")
					}
				}
				else {
					$toggleLabel.removeClass("toggleStrikethroughCompleted")
					$toggleControl.prop("indeterminate", false)
					$toggleControl.prop("checked",false)
				}
			
			}


		} // If record has a value for the current container's associated field ID.
		else
		{
			// No value exits
			$toggleControl.prop("indeterminate", true)
			$toggleLabel.removeClass("toggleStrikethroughCompleted")
		}	
	
	}



	function initToggleFieldEditBehavior($toggle,componentContext,recordProxy, toggleObjectRef) {
	
		var $toggleControl = getToggleControlFromToggleContainer($toggle)
		
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
			setBoolValue(null)
		})
	
		
		$toggleControl.unbind("click")
	  	$toggleControl.click( function () {
			// Get the most recent copy of the object reference. It could have changed between
			// initialization time and the time the toggle was changed.
			var objectRef = getContainerObjectRef($toggle)
		
			var isChecked = $(this).prop("checked")
	
			setBoolValue(isChecked)	
		})
	
	}
		
	
			
	$toggle.data("viewFormConfig", {
		loadRecord: loadRecordIntoToggle,
		validateValue: validateToggleInput
	})
	initToggleFieldEditBehavior($toggle,componentContext,recordProxy, toggleObjectRef)
	
}