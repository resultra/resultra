

function initRatingRecordEditBehavior($ratingContainer,componentContext,recordProxy, ratingObjectRef,remoteValidateInput) {

	var $ratingControl = getRatingControlFromRatingContainer($ratingContainer)
	var $clearValueButton = $ratingContainer.find(".ratingComponentClearValueButton")
	
	var validateInput = function(validationCompleteCallback) {
		
		if($ratingControl.prop('disabled')) {
			validationCompleteCallback(true)
			return
		}
		
		var currVal = getRatingValFromContainer($ratingContainer)
		remoteValidateInput(currVal,function(validationResult) {
			if (validationResult.validationSucceeded) {
				$ratingContainer.popover('destroy')
				validationCompleteCallback(true)
			} else {
				$ratingContainer.popover({
					html: 'true',
					content: function() { return escapeHTML(validationResult.errorMsg) },
					trigger: 'manual',
					placement: 'auto left'
				})
				$ratingContainer.popover('show')
				validationCompleteCallback(false)
			}
			
		})	
		
	}
	
	
	function loadRecordIntoRating(ratingElem, recordRef) {
	
		var ratingObjectRef = getContainerObjectRef(ratingElem)
		var $ratingControl = getRatingControlFromRatingContainer(ratingElem)
	
		var ratingFieldID = ratingObjectRef.properties.fieldID

		console.log("loadRecordIntoRating: Field ID to load data:" + ratingFieldID)

		// In other words, we are populating the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(ratingFieldID)) {

			var fieldVal = recordRef.fieldValues[ratingFieldID]

			console.log("loadRecordIntoRating: Load value into rating control: recordID: " + recordRef.recordID + " field ID:" + 
						ratingFieldID + "  value:" + fieldVal)
		
			if (fieldVal == null) {
				// A null field value corresponds to a value which has been cleared by the user.
				$ratingControl.rating('rate','')
			} else {
				var maxRating = 5
				if((fieldVal >= 0) && (fieldVal <= maxRating)) {
					$ratingControl.rating('rate',fieldVal)	
				} else {
					$ratingControl.rating('rate','')		
				}		
			}
		
		
		} // If record has a value for the current container's associated field ID.
		else
		{
			$ratingControl.rating('rate','')
		}
		
	
	}
	
	function initRatingEditBehavior() {
		function setRatingValue(ratingVal) {
			
			validateInput(function(inputIsValid) {
				if (inputIsValid) {
					currRecordRef = recordProxy.getRecordFunc()
					var ratingFieldID = ratingObjectRef.properties.fieldID

					var ratingValueFormat = { context: "rating", format: "star" }
					var setRecordValParams = { 
						parentDatabaseID:currRecordRef.parentDatabaseID,
						recordID:currRecordRef.recordID,
						changeSetID: recordProxy.changeSetID,
						fieldID:ratingFieldID, 
						value:ratingVal,
						valueFormat: ratingValueFormat}
					jsonAPIRequest("recordUpdate/setNumberFieldValue",setRecordValParams,function(replyData) {
						// After updating the record, the local cache of records in currentRecordSet will
						// be out of date. So after updating the record on the server, the locally cached
						// version of the record also needs to be updated.
						recordProxy.updateRecordFunc(replyData)
	
					}) // set record's number field value
					
				}
			})
		
		
		}
	
		if(formComponentIsReadOnly(ratingObjectRef.properties.permissions)) {
			$ratingControl.prop('disabled',true);
			$clearValueButton.hide()
		
		} else {
			$ratingControl.prop('disabled',false);
			$clearValueButton.show()
			// The rating control is initialized the same way for design and view mode, but in view mode
			// the event handlers need to be setup for when the user changes a rating value.
			$ratingControl.on('change', function() {
				var ratingVal = getRatingValFromContainer($ratingContainer)
				console.log('Rating changed: ' + ratingVal);
				setRatingValue(ratingVal)
			});
		
			initButtonControlClickHandler($clearValueButton,function() {
					console.log("Clear value clicked for rating")
					$ratingControl.rating('rate','')
					setRatingValue(null)
			})
		
		}
		
	}
	initRatingEditBehavior()

	
	// When the user clicks on the control, prevent the click from propagating higher.
	// This allows the user to change the rating without selecting the form component itself.
	// The user can still select the component by clicking on the label or anywwhere outside
	// the control.
	$ratingContainer.find(".formRatingControl").click(function (event){
		event.stopPropagation();
   	 	//   ... your code here
		return false;
	});
		
	$ratingContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoRating,
		validateValue: validateInput
	})
	

}

function initRatingFormRecordEditBehavior($ratingContainer,componentContext,recordProxy, ratingObjectRef) {
	
	function validateInput(inputVal,validationResultCallback) {
		var validationParams = {
			parentFormID: ratingObjectRef.parentFormID,
			ratingID: ratingObjectRef.ratingID,
			inputVal: inputVal
		}
		jsonAPIRequest("frm/rating/validateInput", validationParams, function(validationResult) {
			validationResultCallback(validationResult)
		})
	}
	
	initRatingRecordEditBehavior($ratingContainer,componentContext,recordProxy, ratingObjectRef,validateInput)
}

function initRatingTableCellRecordEditBehavior($ratingContainer,componentContext,recordProxy, ratingObjectRef) {
	
	function validateInput(inputVal,validationResultCallback) {
		var validationParams = {
			parentTableID: ratingObjectRef.parentTableID,
			ratingID: ratingObjectRef.ratingID,
			inputVal: inputVal
		}
		jsonAPIRequest("tableView/rating/validateInput", validationParams, function(validationResult) {
			validationResultCallback(validationResult)
		})
	}
	
	
	initRatingRecordEditBehavior($ratingContainer,componentContext,recordProxy, ratingObjectRef,validateInput)
}