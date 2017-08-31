

function initSocialButtonRecordEditBehavior($socialButtonContainer,componentContext,recordProxy,
		 	socialButtonObjectRef) {

	var $socialButtonControl = getSocialButtonControlFromContainer($socialButtonContainer)
	
	var validateInput = function(validationCompleteCallback) {
		validationCompleteCallback(true) // validation is a no-op		
	}
	
	
	function loadRecordIntoSocialButton(socialButtonElem, recordRef) {
	
		var socialButtonObjectRef = getContainerObjectRef(socialButtonElem)
		var $socialButtonControl = getSocialButtonControlFromContainer(socialButtonElem)
	
		var socialButtonFieldID = socialButtonObjectRef.properties.fieldID

		console.log("loadRecordIntoRating: Field ID to load data:" + socialButtonFieldID)

		// In other words, we are populating the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(socialButtonFieldID)) {

			var fieldVal = recordRef.fieldValues[socialButtonFieldID]
		
			if (fieldVal == null) {
				// A null field value corresponds to a value which has been cleared by the user.
				console.log("loadRecordIntoSocialButton: clearing social button value for null field value")
			} else {
				console.log("loadRecordIntoSocialButton: setting social button value for field value: " + JSON.stringify(fieldVal))
			}
		
		
		} // If record has a value for the current container's associated field ID.
		else
		{
			console.log("loadRecordIntoSocialButton: clearing social button value for undefined field value")
		}
		
	
	}
	
	function initSocialButtonEditBehavior() {
		
		function toggleCurrUserSocialButtonVal(selectedUsersCallback) {
			
			
			function getToggledSocialButtonVal(selectedUsersCallback) {
				
				var getUserInfoParams = {}
				jsonRequest("/auth/getCurrentUserInfo",getUserInfoParams,function(currUserInfo) {
					
					var currRecordRef = recordProxy.getRecordFunc()
					var socialButtonFieldID = socialButtonObjectRef.properties.fieldID
						
					if(currRecordRef.fieldValues.hasOwnProperty(socialButtonFieldID)) {

						var fieldVal = currRecordRef.fieldValues[socialButtonFieldID]
		
						if (fieldVal == null) {
							// A null field value corresponds to a value which has been cleared by the user.
							console.log("getToggledSocialButtonVal: null field value: adding user via button toggle")
							var selectedUserIDs = [currUserInfo.userID]
							selectedUsersCallback(selectedUserIDs)
							
						} else {
							console.log("getToggledSocialButtonVal: setting social button value for field value: " + JSON.stringify(fieldVal))
							var selectedUserIDLookup = new IDLookupTable(fieldVal)
							if (selectedUserIDLookup.hasID(currUserInfo.userID)) {
								selectedUserIDLookup.removeID(currUserInfo.userID)
								var selectedUserIDs = selectedUserIDLookup.getIDList()
								selectedUsersCallback(selectedUserIDs)
							} else {
								var selectedUserIDs = fieldVal
								selectedUserIDs.push(currUserInfo.userID)
								selectedUsersCallback(selectedUserIDs)
							}
						}
		
		
					} // If record has a value for the current container's associated field ID.
					else
					{
						console.log("getToggledSocialButtonVal: undefined field value: adding user via button toggle")
						var selectedUserIDs = [currUserInfo.userID]
						selectedUsersCallback(selectedUserIDs)
					}
				})
				
			} // getToggledSocialButtonVal
			
			getToggledSocialButtonVal(function(selectedUserIDs) {
				var currRecordRef = recordProxy.getRecordFunc()
				var socialButtonFieldID = socialButtonObjectRef.properties.fieldID
				var userValueFormat = { context: "social", format: "button" }
				var setRecordValParams = { 
					parentDatabaseID:currRecordRef.parentDatabaseID,
					recordID:currRecordRef.recordID,
					changeSetID: recordProxy.changeSetID, 
					fieldID:socialButtonFieldID, 
					userIDs:selectedUserIDs,
					valueFormat:userValueFormat}
				jsonAPIRequest("recordUpdate/setUserFieldValue",setRecordValParams,function(updatedFieldVal) {
					// After updating the record, the local cache of records in currentRecordSet will
					// be out of date. So after updating the record on the server, the locally cached
					// version of the record also needs to be updated.
					recordProxy.updateRecordFunc(updatedFieldVal)

				}) // set record's number field value
				
			})

		} // toggleCurrUserSocialButtonVal	
	
		if(formComponentIsReadOnly(socialButtonObjectRef.properties.permissions)) {
			$socialButtonControl.prop('disabled',true);
		} else {
			$socialButtonControl.prop('disabled',false);
			initButtonControlClickHandler($socialButtonControl,function() {
				console.log('Toggling social button values');
				toggleCurrUserSocialButtonVal()
			})
		}
		
	}
	initSocialButtonEditBehavior()
		
	$socialButtonContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoSocialButton,
		validateValue: validateInput
	})
	

}

function initSocialButtonFormRecordEditBehavior($socialButtonContainer,componentContext,recordProxy, socialButtonObjectRef) {
		
	initSocialButtonRecordEditBehavior($socialButtonContainer,componentContext,recordProxy, socialButtonObjectRef)
}

function initSocialButtonTableCellRecordEditBehavior($socialButtonContainer,componentContext,recordProxy, socialButtonObjectRef) {
		
	initSocialButtonFormComponentControl($socialButtonContainer,socialButtonObjectRef)
	
	initSocialButtonRecordEditBehavior($socialButtonContainer,componentContext,recordProxy, socialButtonObjectRef)
}