

function initUserSelectionRecordEditBehavior($userSelectionContainer, componentContext,
		recordProxy, userSelectionObjectRef, controlWidth, validateInputFunc) {

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
				setUserSelectionControlVal($userSelectionControl,fieldVal)		
			}

		} // If record has a value for the current container's associated field ID.
		else
		{
			clearUserSelectionControlVal($userSelectionControl)
		}
		
	}


	function initUserSelectionEditBehavior() {
		function setUserSelectionValue(selectedUserIDs) {
		
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
					jsonAPIRequest("recordUpdate/setUserFieldValue",setRecordValParams,function(updatedFieldVal) {
						// After updating the record, the local cache of records in currentRecordSet will
						// be out of date. So after updating the record on the server, the locally cached
						// version of the record also needs to be updated.
						recordProxy.updateRecordFunc(updatedFieldVal)
	
					}) // set record's number field value
				}
			})
		}
	
		var selectionWidth = controlWidth.toString() + "px"
		var userSelectionParams = {
			$selectionInput: $userSelectionControl,
			databaseID: componentContext.databaseID,
			width: selectionWidth
		}
		initCollaboratorUserSelection(userSelectionParams)
		
		function setSelectedUserFromDropdownMenu(userInfo) {
			var userIDSelection = [userInfo.userID]
			// For the selected user to be displayed in the selection,
			// it needs to be added as an option.
			var newOption = new Option('@'+userInfo.userID, userInfo.userID, true, true);
			$userSelectionControl.append(newOption)
			$userSelectionControl.val(userIDSelection)
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

		$userSelectionControl.on('change', function() {
			var selectedUserID = $(this).val()
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
			userSelectionObjectRef,selectionWidth, validateInput)
}



function initUserSelectionTableRecordPopupEditBehavior($container,componentContext,recordProxy, userSelectionObjectRef) {
	
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
	var selectionWidth = 250
	
	initUserSelectionClearValueButton($container,userSelectionObjectRef)
	
	initUserSelectionRecordEditBehavior($container,componentContext,recordProxy, 
			userSelectionObjectRef,selectionWidth, validateInput)
}


function initUserSelectionTableRecordEditBehavior($container,componentContext,recordProxy, userSelectionObjectRef) {

	var $userPopupLink = $container.find(".userSelectionEditPopop")

	// TBD - Needs a popup to display the editor.
	var validateInput = function(validationCompleteCallback) {
			validationCompleteCallback(true)
	}
	
	function formatPopupLinkText(recordRef) {
		
		var fieldID = userSelectionObjectRef.properties.fieldID
		var usersExist = recordRef.fieldValues.hasOwnProperty(fieldID)
		
		if(formComponentIsReadOnly(userSelectionObjectRef.properties.permissions)) {
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
		content: function() { return userSelectionTablePopupEditorContainerHTML() },
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
		
		var $userEditorContainer = $popover.find(".userSelectionTableCellContainer")
		
//		initHTMLEditorTextCellComponentViewModeGeometry($noteEditorContainer)
		
		var $closePopupButton = $userEditorContainer.find(".closeEditorPopup")
		initButtonControlClickHandler($closePopupButton,function() {
			$userPopupLink.popover('hide')
		})
			
		initUserSelectionTableRecordPopupEditBehavior($userEditorContainer,componentContext,
					recordProxy, userSelectionObjectRef)

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


