

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
				$labelControl.empty()
				$labelControl.val(null)
			} else {
				
				$labelControl.empty()
				var tagAdded = {}
				for(var tagIndex in fieldVal) {
					var currTag = fieldVal[tagIndex]
					if(tagAdded[currTag] === undefined) { // don't add duplicates
						var newOption = new Option(currTag,currTag);
						tagAdded[currTag] = true
						$labelControl.append(newOption)					
					}
				}
				
				$labelControl.val(fieldVal)
			}

		} // If record has a value for the current container's associated field ID.
		else
		{
				$labelControl.empty()
			$labelControl.val(null)
		}
		
	}


	function initLabelEditBehavior() {
		function setLabelValue(selectedLabels) {
		
			validateInput(function(inputIsValid) {
				if (inputIsValid) {
					currRecordRef = recordProxy.getRecordFunc()

					var userFieldID = labelFieldID

					var setRecordValParams = { 
						parentDatabaseID:currRecordRef.parentDatabaseID,
						recordID:currRecordRef.recordID,
						changeSetID: recordProxy.changeSetID, 
						fieldID:userFieldID, 
						labels:selectedLabels}
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



function initTagTablePopupRecordEditBehavior($container,componentContext,recordProxy, labelObjectRef) {
	
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
		
	var labelWidth = 250
	initLabelSelectionControl($container, labelObjectRef,labelWidth)
	initTagTablePopupDimensions($container)
			
	initLabelRecordEditBehavior($container,componentContext,recordProxy, 
			labelObjectRef,validateInput)
}



function initLabelTableRecordEditBehavior($container,componentContext,recordProxy, tagObjectRef) {

	var $tagPopupLink = $container.find(".tagEditPopop")

	// TBD - Needs a popup to display the editor.
	var validateInput = function(validationCompleteCallback) {
			validationCompleteCallback(true)
	}
	
	function formatTagPopupLinkText(recordRef) {
		var fieldID = tagObjectRef.properties.fieldID
		var tagsExist = recordRef.fieldValues.hasOwnProperty(fieldID)
		
		if(formComponentIsReadOnly(tagObjectRef.properties.permissions)) {
			if (tagsExist) {
				$tagPopupLink.css("display","")
				$tagPopupLink.text("View tags")
			} else {
				$tagPopupLink.css("display","none")
				$tagPopupLink.text("")
			}
		} else {
			$tagPopupLink.css("display","")
			if (tagsExist) {
				$tagPopupLink.text("Edit tags")
			} else {
				$tagPopupLink.text("Add tag")
			}
		}
	}
	
	var currRecordRef = null
	function loadRecordIntoHtmlEditor($htmlEditor, recordRef) {
		currRecordRef = recordRef
		formatTagPopupLinkText(recordRef)
	}
	
	
	$tagPopupLink.popover({
		html: 'true',
		content: function() { return labelTablePopupViewContainerHTML() },
		trigger: 'manual',
		placement: 'auto left'
	})
	
	$tagPopupLink.click(function(e) {
		$(this).popover('toggle')
		e.stopPropagation()
	})
	
	
	$tagPopupLink.on('shown.bs.popover', function()
	{
	    //get the actual shown popover
	    var $popover = $(this).data('bs.popover').tip();
		
		// By default the popover takes on the maximum size of it's containing
		// element. Overridding this size allows the size to grow as needed.
		$popover.css("max-width","300px")
		$popover.css("max-height","200px")
		
		var $tagEditorContainer = $popover.find(".labelTableCellContainer")
				
		var $closePopupButton = $tagEditorContainer.find(".closeTagEditorPopup")
		initButtonControlClickHandler($closePopupButton,function() {
			$tagPopupLink.popover('hide')
		})
			
		initTagTablePopupRecordEditBehavior($tagEditorContainer,componentContext,recordProxy, tagObjectRef)

		if(currRecordRef != null) {
			var viewConfig = $tagEditorContainer.data("viewFormConfig")
			viewConfig.loadRecord($tagEditorContainer,currRecordRef)
		}

	});
	
	$container.data("viewFormConfig", {
		loadRecord: loadRecordIntoHtmlEditor,
		validateValue: validateInput
	})
	
}

