

function initSocialButtonRecordEditBehavior($socialButtonContainer,componentContext,recordProxy,
		 	socialButtonObjectRef,remoteValidateInput) {

	var $socialButtonControl = getSocialButtonControlFromContainer($socialButtonContainer)
	
	var validateInput = function(validationCompleteCallback) {
		
		if($socialButtonControl.prop('disabled')) {
			validationCompleteCallback(true)
			return
		}
		
		var currVal = getRatingValFromContainer($socialButtonContainer)
		remoteValidateInput(currVal,function(validationResult) {
			setupFormComponentValidationPrompt($socialButtonContainer,validationResult,validationCompleteCallback)			
		})	
		
	}
	
	
	function loadRecordIntoRating(socialButtonElem, recordRef) {
	
		var socialButtonObjectRef = getContainerObjectRef(socialButtonElem)
		var $socialButtonControl = getSocialButtonControlFromContainer(socialButtonElem)
	
		var socialButtonFieldID = socialButtonObjectRef.properties.fieldID

		console.log("loadRecordIntoRating: Field ID to load data:" + socialButtonFieldID)

		// In other words, we are populating the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(socialButtonFieldID)) {

			var fieldVal = recordRef.fieldValues[socialButtonFieldID]

			console.log("loadRecordIntoRating: Load value into social button control: recordID: " + recordRef.recordID + " field ID:" + 
						socialButtonFieldID + "  value:" + fieldVal)
		
			if (fieldVal == null) {
				// A null field value corresponds to a value which has been cleared by the user.
				$socialButtonControl.rating('rate','')
			} else {
				var maxRating = 1
				if((fieldVal >= 0) && (fieldVal <= maxRating)) {
					$socialButtonControl.rating('rate',fieldVal)	
				} else {
					$socialButtonControl.rating('rate','')		
				}		
			}
		
		
		} // If record has a value for the current container's associated field ID.
		else
		{
			$socialButtonControl.rating('rate','')
		}
		
	
	}
	
	function initRatingEditBehavior() {
		function setRatingValue(ratingVal) {
			
			validateInput(function(inputIsValid) {
				if (inputIsValid) {
					currRecordRef = recordProxy.getRecordFunc()
					var socialButtonFieldID = socialButtonObjectRef.properties.fieldID

					var valueFormat = { context: "social", format: "button" }
					var setRecordValParams = { 
						parentDatabaseID:currRecordRef.parentDatabaseID,
						recordID:currRecordRef.recordID,
						changeSetID: recordProxy.changeSetID,
						fieldID:socialButtonFieldID, 
						value:ratingVal,
						valueFormat: valueFormat}
					jsonAPIRequest("recordUpdate/setNumberFieldValue",setRecordValParams,function(replyData) {
						// After updating the record, the local cache of records in currentRecordSet will
						// be out of date. So after updating the record on the server, the locally cached
						// version of the record also needs to be updated.
						recordProxy.updateRecordFunc(replyData)
	
					}) // set record's number field value
					
				}
			})
		
		
		}
	
		if(formComponentIsReadOnly(socialButtonObjectRef.properties.permissions)) {
			$socialButtonControl.prop('disabled',true);
		} else {
			$socialButtonControl.prop('disabled',false);
			// The rating control is initialized the same way for design and view mode, but in view mode
			// the event handlers need to be setup for when the user changes a rating value.
			$socialButtonControl.on('change', function() {
				var ratingVal = getRatingValFromContainer($socialButtonContainer)
				console.log('Rating changed: ' + ratingVal);
				setRatingValue(ratingVal)
			});		
		}
		
	}
	initRatingEditBehavior()

	
	// When the user clicks on the control, prevent the click from propagating higher.
	// This allows the user to change the rating without selecting the form component itself.
	// The user can still select the component by clicking on the label or anywwhere outside
	// the control.
	$socialButtonContainer.find(".socialButtonControl").click(function (event){
		event.stopPropagation();
   	 	//   ... your code here
		return false;
	});
		
	$socialButtonContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoRating,
		validateValue: validateInput
	})
	

}

function initSocialButtonFormRecordEditBehavior($socialButtonContainer,componentContext,recordProxy, socialButtonObjectRef) {
	
	function validateInput(inputVal,validationResultCallback) {
		var validationParams = {
			parentFormID: socialButtonObjectRef.parentFormID,
			socialButtonID: socialButtonObjectRef.socialButtonID,
			inputVal: inputVal
		}
		jsonAPIRequest("frm/socialButton/validateInput", validationParams, function(validationResult) {
			validationResultCallback(validationResult)
		})
	}
	
	initRatingRecordEditBehavior($socialButtonContainer,componentContext,recordProxy, socialButtonObjectRef,validateInput)
}

function initSocialButtonTableCellRecordEditBehavior($socialButtonContainer,componentContext,recordProxy, socialButtonObjectRef) {
	
	function validateInput(inputVal,validationResultCallback) {
		var validationParams = {
			parentTableID: socialButtonObjectRef.parentTableID,
			socialButtonID: socialButtonObjectRef.socialButtonID,
			inputVal: inputVal
		}
		jsonAPIRequest("tableView/socialButton/validateInput", validationParams, function(validationResult) {
			validationResultCallback(validationResult)
		})
	}
	
	initSocialButtonFormComponentControl($socialButtonContainer,socialButtonObjectRef)
	
	initSocialButtonRecordEditBehavior($socialButtonContainer,componentContext,recordProxy, socialButtonObjectRef,validateInput)
}