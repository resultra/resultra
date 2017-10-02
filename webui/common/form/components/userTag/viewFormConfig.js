

function initUserTagRecordEditBehavior($userTagContainer, componentContext,
		recordProxy, userTagObjectRef, validateInputFunc) {

	var selectionFieldID = userTagObjectRef.properties.fieldID
		
	var $userTagControl = userTagControlFromUserTagComponentContainer($userTagContainer)

	var validateInput = function(validationCompleteCallback) {
		if($userTagControl.prop('disabled')) {
			validationCompleteCallback(true)
			return
		}
		var currVal = $userTagControl.val()
		validateInputFunc(currVal,function(validationResult) {
			setupFormComponentValidationPrompt($userTagContainer,validationResult,validationCompleteCallback)			
		})	
	}

	function loadRecordIntoUserTag(userTagElem, recordRef) {

		var $userTagControl = userTagControlFromUserTagComponentContainer(userTagElem)

		if(formComponentIsReadOnly(userTagObjectRef.properties.permissions)) {
			$userTagControl.prop('disabled',true);
		} else {
			$userTagControl.prop('disabled',false);
		
		}

		var userTagFieldID = userTagObjectRef.properties.fieldID

		console.log("loadRecordIntoUserTag: Field ID to load data:" + userTagFieldID)

		// In other words, we are populating the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(userTagFieldID)) {

			var fieldVal = recordRef.fieldValues[userTagFieldID]
			if (fieldVal === null) {
				clearUserTagControlVal($userTagControl)
			} else {
				setUserTagControlVal($userTagControl,fieldVal)		
			}

		} // If record has a value for the current container's associated field ID.
		else
		{
			clearUserTagControlVal($userTagControl)
		}
		
	}


	function initUserTagEditBehavior() {
		function setUserTagValue(selectedUserIDs) {
		
			validateInput(function(inputIsValid) {
				if (inputIsValid) {
					currRecordRef = recordProxy.getRecordFunc()

					var userFieldID = selectionFieldID

					var setRecordValParams = { 
						parentDatabaseID:currRecordRef.parentDatabaseID,
						recordID:currRecordRef.recordID,
						changeSetID: recordProxy.changeSetID, 
						fieldID:userFieldID, 
						userIDs:selectedUserIDs}
					jsonAPIRequest("recordUpdate/setUsersFieldValue",setRecordValParams,function(updatedFieldVal) {
						// After updating the record, the local cache of records in currentRecordSet will
						// be out of date. So after updating the record on the server, the locally cached
						// version of the record also needs to be updated.
						recordProxy.updateRecordFunc(updatedFieldVal)
	
					}) // set record's number field value
				}
			})
		}
	
		var userTagParams = {
			$selectionInput: $userTagControl,
			databaseID: componentContext.databaseID,
		}
		initCollaboratorUserTag(userTagParams)
		
		function setSelectedUserFromDropdownMenu(userInfo) {
			var userIDSelection = [userInfo.userID]
			// For the selected user to be displayed in the selection,
			// it needs to be added as an option.
			var newOption = new Option('@'+userInfo.userID, userInfo.userID, true, true);
			$userTagControl.append(newOption)
			$userTagControl.val(userIDSelection)
			setUserTagValue(userIDSelection)
		}
		configureUserTagDropdown(componentContext,$userTagContainer,
					userTagObjectRef,setSelectedUserFromDropdownMenu)
		
	
		var $clearValueButton = $userTagContainer.find(".userTagComponentClearValueButton")
		initButtonControlClickHandler($clearValueButton,function() {
			console.log("Clear value clicked for user selection")
			clearUserTagControlVal($userTagControl)
			setUserTagValue(null)
		})

		$userTagControl.on('change', function() {
			var selectedUserID = $(this).val()
			console.log('User selection changed: ' + selectedUserID);
			setUserTagValue(selectedUserID)
		});
		
	}
	initUserTagEditBehavior()

	
	// When the user clicks on the control, prevent the click from propagating higher.
	// This allows the user to change the rating without selecting the form component itself.
	// The user can still select the component by clicking on the label or anywwhere outside
	// the control.
	$userTagContainer.find(".formUserTagControl").click(function (event){
		event.stopPropagation();
   	 	//   ... your code here
		return false;
	});
		
	$userTagContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoUserTag,
		validateValue: validateInput
	})
	

}


function initUserTagFormRecordEditBehavior($container,componentContext,recordProxy, userTagObjectRef) {
	
	function validateInput(inputVal,validationResultCallback) {
		var validationParams = {
			parentFormID: userTagObjectRef.parentFormID,
			userTagID: userTagObjectRef.userTagID,
			inputVal: inputVal
		}
		jsonAPIRequest("frm/userTag/validateInput", validationParams, function(validationResult) {
			validationResultCallback(validationResult)
		})
	}
	
	var selectionWidth = userTagObjectRef.properties.geometry.sizeWidth - 15
	
	initUserTagRecordEditBehavior($container,componentContext,recordProxy, 
			userTagObjectRef, validateInput)
}



function initUserTagTableRecordPopupEditBehavior($container,componentContext,recordProxy, userTagObjectRef) {
	
	function validateInput(inputVal,validationResultCallback) {
		var validationParams = {
			parentTableID: userTagObjectRef.parentTableID,
			userTagID: userTagObjectRef.userTagID,
			inputVal: inputVal
		}
		jsonAPIRequest("tableView/userTag/validateInput", validationParams, function(validationResult) {
			validationResultCallback(validationResult)
		})
	}
	
	// The width of the select2 control is setup during the controls initialization. However, the height
	// is configured via CSS. See the CSS file for corresponding CSS to configure the height.	
	initUserTagClearValueButton($container,userTagObjectRef)
	
	initUserTagRecordEditBehavior($container,componentContext,recordProxy, 
			userTagObjectRef, validateInput)
}


function initUserTagTableRecordEditBehavior($container,componentContext,recordProxy, userTagObjectRef) {

	var $userPopupLink = $container.find(".userTagEditPopop")

	// TBD - Needs a popup to display the editor.
	var validateInput = function(validationCompleteCallback) {
			validationCompleteCallback(true)
	}
	
	function formatPopupLinkText(recordRef) {
		
		var fieldID = userTagObjectRef.properties.fieldID
		var usersExist = recordRef.fieldValues.hasOwnProperty(fieldID)
		
		if(formComponentIsReadOnly(userTagObjectRef.properties.permissions)) {
			if (usersExist) {
				$userPopupLink.css("display","")
				$userPopupLink.text("View tags")
			} else {
				$userPopupLink.css("display","none")
				$userPopupLink.text("")
			}
		} else {
			$userPopupLink.css("display","")
			if (usersExist) {
				$userPopupLink.text("Edit users")
			} else {
				$userPopupLink.text("Add user")
			}
		}
	}
	
	var currRecordRef = null
	function loadRecordIntoHtmlEditor($htmlEditor, recordRef) {
		currRecordRef = recordRef
		formatPopupLinkText(recordRef)
	}
	
	
	$userPopupLink.popover({
		html: 'true',
		content: function() { return userTagTablePopupEditorContainerHTML() },
		trigger: 'manual',
		placement: 'auto left'
	})
	
	$userPopupLink.click(function(e) {
		$(this).popover('toggle')
		e.stopPropagation()
	})
	
	
	$userPopupLink.on('shown.bs.popover', function()
	{
	    //get the actual shown popover
	    var $popover = $(this).data('bs.popover').tip();
		
		// By default the popover takes on the maximum size of it's containing
		// element. Overridding this size allows the size to grow as needed.
		$popover.css("max-width","300px")
		$popover.css("max-height","200px")
		
		var $userEditorContainer = $popover.find(".userTagTableCellContainer")
		
//		initHTMLEditorTextCellComponentViewModeGeometry($noteEditorContainer)
		
		var $closePopupButton = $userEditorContainer.find(".closeEditorPopup")
		initButtonControlClickHandler($closePopupButton,function() {
			$userPopupLink.popover('hide')
		})
			
		initUserTagTableRecordPopupEditBehavior($userEditorContainer,componentContext,
					recordProxy, userTagObjectRef)

		if(currRecordRef != null) {
			var viewConfig = $userEditorContainer.data("viewFormConfig")
			viewConfig.loadRecord($userEditorContainer,currRecordRef)
		}

	});
	
	$container.data("viewFormConfig", {
		loadRecord: loadRecordIntoHtmlEditor,
		validateValue: validateInput
	})
	
}


