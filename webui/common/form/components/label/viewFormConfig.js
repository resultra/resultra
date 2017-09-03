

function initLabelRecordEditBehavior($labelContainer, componentContext,
		recordProxy, labelObjectRef, validateInputFunc) {

	var labelFieldID = labelObjectRef.properties.fieldID
		
	var $labelControl = labelControlFromLabelComponentContainer($labelContainer)

	var validateInput = function(validationCompleteCallback) {
		if($labelControl.prop('disabled')) {
			validationCompleteCallback(true)
			return
		}
		var currVal = $labelControl.val()
		validateInputFunc(currVal,function(validationResult) {
			setupFormComponentValidationPrompt($labelContainer,validationResult,validationCompleteCallback)			
		})	
	}

	function loadRecordIntoLabel(labelElem, recordRef) {

		var labelObjectRef = labelElem.data("objectRef")
		var $labelControl = labelControlFromLabelComponentContainer(labelElem)

		if(formComponentIsReadOnly(labelObjectRef.properties.permissions)) {
			$labelControl.prop('disabled',true);
		} else {
			$labelControl.prop('disabled',false);
		
		}

		var labelFieldID = labelObjectRef.properties.fieldID

		console.log("loadRecordIntoLabel: Field ID to load data:" + labelFieldID)

		// In other words, we are populating the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(labelFieldID)) {

			var fieldVal = recordRef.fieldValues[labelFieldID]
			if (fieldVal === null) {
				$labelControl.val([])
			} else {
				$labelControl.val(fieldVal)
			}

		} // If record has a value for the current container's associated field ID.
		else
		{
			$labelControl.val([])
		}
		
	}


	function initLabelEditBehavior() {
		function setLabelValue(selectedLabels) {
		
			validateInput(function(inputIsValid) {
				if (inputIsValid) {
					currRecordRef = recordProxy.getRecordFunc()

					var userFieldID = labelFieldID

					var userValueFormat = {
						context: "selectUser",
						format: "general"
					}
					var setRecordValParams = { 
						parentDatabaseID:currRecordRef.parentDatabaseID,
						recordID:currRecordRef.recordID,
						changeSetID: recordProxy.changeSetID, 
						fieldID:userFieldID, 
						labels:selectedLabels,
						valueFormat:userValueFormat}
					jsonAPIRequest("recordUpdate/setLabelFieldValue",setRecordValParams,function(updatedFieldVal) {
						// After updating the record, the local cache of records in currentRecordSet will
						// be out of date. So after updating the record on the server, the locally cached
						// version of the record also needs to be updated.
						recordProxy.updateRecordFunc(updatedFieldVal)
	
					}) // set record's number field value
				}
			})
		}
				
		$labelControl.on('change', function() {
			var selectedLabels = $(this).val()
			console.log('Label selection changed: ' + JSON.stringify(selectedLabels));
			setLabelValue(selectedLabels)
		});
		
	}
	initLabelEditBehavior()

	
	// When the user clicks on the control, prevent the click from propagating higher.
	// This allows the user to change the rating without selecting the form component itself.
	// The user can still select the component by clicking on the label or anywwhere outside
	// the control.
	$labelContainer.find(".formLabelControl").click(function (event){
		event.stopPropagation();
   	 	//   ... your code here
		return false;
	});
		
	$labelContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoLabel,
		validateValue: validateInput
	})
	

}


function initLabelFormRecordEditBehavior($container,componentContext,recordProxy, labelObjectRef) {
	
	function validateInput(inputVal,validationResultCallback) {
		var validationParams = {
			parentFormID: labelObjectRef.parentFormID,
			labelID: labelObjectRef.labelID,
			inputVal: inputVal
		}
		jsonAPIRequest("frm/label/validateInput", validationParams, function(validationResult) {
			validationResultCallback(validationResult)
		})
	}
	
	initLabelRecordEditBehavior($container,componentContext,recordProxy, 
			labelObjectRef,validateInput)
}



function initLabelTableRecordEditBehavior($container,componentContext,recordProxy, labelObjectRef) {
	
	function validateInput(inputVal,validationResultCallback) {
		var validationParams = {
			parentTableID: labelObjectRef.parentTableID,
			tagID: labelObjectRef.tagID,
			inputVal: inputVal
		}
		jsonAPIRequest("tableView/tag/validateInput", validationParams, function(validationResult) {
			validationResultCallback(validationResult)
		})
	}
		
	var labelWidth = 200
	initLabelSelectionControl($container, labelObjectRef,labelWidth)
			
	initLabelRecordEditBehavior($container,componentContext,recordProxy, 
			labelObjectRef,validateInput)
}


