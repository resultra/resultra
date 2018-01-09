

function initHtmlEditorRecordEditBehavior($htmlEditor,componentContext,recordProxy,htmlEditorObjectRef,remoteValidationFunc) {
	
	var $htmlEditorInput = htmlInputFromHTMLEditorContainer($htmlEditor)
	
	var validateInput = function(validationCompleteCallback) {
		
		if(formComponentIsReadOnly(htmlEditorObjectRef.properties.permissions)) {
			validationCompleteCallback(true)
			return
		}
		
		
		// Although editing takes place with HTML, the contents of the editor are validated using the text
		// of what's in the editor. This will strip out any HTML elements, leaving only actual non-whitespace
		// characters.
		var currInputValText = $htmlEditorInput.text()
		remoteValidationFunc(currInputValText,function(validationResult) {
			setupFormComponentValidationPrompt($htmlEditor,validationResult,validationCompleteCallback)			
		})
		
	}
	
	function loadRecordIntoHtmlEditor($htmlEditor, recordRef) {
	
		console.log("loadRecordIntoHtmlEditor: loading record into html editor: " + JSON.stringify(recordRef))
	
		var htmlEditorObjectRef = $htmlEditor.data("objectRef")
	
	
		var $editButton = $htmlEditor.find(".startEditButton")
		if(formComponentIsReadOnly(htmlEditorObjectRef.properties.permissions)) {
			$editButton.prop('disabled',true);
			$editButton.hide()
		} else {
			$editButton.prop('disabled',false);
			$editButton.show()
		}
	
		var htmlEditorFieldID = htmlEditorObjectRef.properties.fieldID
		// Populate the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(htmlEditorFieldID)) {
		
			var fieldVal = recordRef.fieldValues[htmlEditorFieldID]
		
			if (fieldVal === null) {
				populateInlineDisplayContainerHTML($htmlEditorInput,"")
			} else {
				// If record has a value for the current container's associated field ID.
				populateInlineDisplayContainerHTML($htmlEditorInput,fieldVal)					
			}

	
		} else {
			// There's no value in the current record for this field, so clear the value in the container
			$htmlEditorInput.html("")
		}

	}
	
	
	
	function initEditorBehavior() {
		function setEditorValue(editorVal) {
			validateInput(function(inputIsValid) {
				if(inputIsValid) {
					var currRecordRef = recordProxy.getRecordFunc()
					var htmlEditorFieldID = htmlEditorObjectRef.properties.fieldID
		
	
					var setRecordValParams = { 
						parentDatabaseID:currRecordRef.parentDatabaseID,
						recordID:currRecordRef.recordID,
						changeSetID: recordProxy.changeSetID,
						fieldID:htmlEditorFieldID, 
						value:editorVal
					}
		
					jsonAPIRequest("recordUpdate/setLongTextFieldValue",setRecordValParams,function(updatedRecordRef) {

						// After updating the record, the local cache of records in currentRecordSet will
						// be out of date. So after updating the record on the server, the locally cached
						// version of the record also needs to be updated.
						recordProxy.updateRecordFunc(updatedRecordRef)
					}) // set record's text field value
					
				}
			})
		
		}
		
		var $editButton = $htmlEditor.find(".startEditButton")
		initButtonControlClickHandler($editButton,function() {
			console.log("Starting inline editor")
		
			if (!inlineCKEditorEnabled($htmlEditorInput)) {
			    CKEDITOR.disableAutoInline = true;
	
				var editor = enableInlineCKEditor($htmlEditorInput)
				editor.focus()
		
				editor.on('blur', function(event) {
					var inputVal = editor.getData();
					setEditorValue(inputVal)				
					disableInlineCKEditor($htmlEditorInput,editor)	
				})
		
				$htmlEditorInput.focus()
			}
		
		})
		
	}
	initEditorBehavior()
	
		
	$htmlEditor.data("viewFormConfig", {
		loadRecord: loadRecordIntoHtmlEditor,
		validateValue: validateInput
	})
	
	$htmlEditor.find(".htmlEditorContent").click(function(e) {
		
		// This is important - if a click hits an object, then stop the propagation of the click
		// to the parent div(s), including the canvas itself. If the parent canvas
		// gets a click, it will deselect all the items (see initObjectCanvasSelectionBehavior)
		e.stopPropagation();
	})
	

	
	
}

function initNoteEditorFormRecordEditBehavior($container,componentContext,recordProxy, noteEditorObjectRef) {
	
	function validateInput(inputVal,validationResultCallback) {
		var validationParams = {
			parentFormID: noteEditorObjectRef.parentFormID,
			htmlEditorID: noteEditorObjectRef.htmlEditorID,
			inputVal: inputVal
		}
		jsonAPIRequest("frm/htmlEditor/validateInput", validationParams, function(validationResult) {
			validationResultCallback(validationResult)
		})
	}
	
	initHtmlEditorRecordEditBehavior($container,componentContext,recordProxy, 
			noteEditorObjectRef,validateInput)
}

function initNoteEditorTableRecordEditBehavior($container,componentContext,recordProxy, noteEditorObjectRef) {
	
	function validateInput(inputVal,validationResultCallback) {
		var validationParams = {
			parentTableID: noteEditorObjectRef.parentTableID,
			noteID: noteEditorObjectRef.noteID,
			inputVal: inputVal
		}
		jsonAPIRequest("tableView/note/validateInput", validationParams, function(validationResult) {
			validationResultCallback(validationResult)
		})
	}
	
	setContainerComponentInfo($container,noteEditorObjectRef,noteEditorObjectRef.noteID)
		
	initHtmlEditorRecordEditBehavior($container,componentContext,recordProxy, 
			noteEditorObjectRef,validateInput)
}


function initNoteEditorTableCellEditBehavior($container,componentContext,recordProxy, noteEditorObjectRef) {

	var $notePopupLink = $container.find(".noteEditPopop")

	// TBD - Needs a popup to display the editor.
	var validateInput = function(validationCompleteCallback) {
			validationCompleteCallback(true)
	}
	
	function formatNotePopupLinkText(recordRef) {
		var fieldID = noteEditorObjectRef.properties.fieldID
		var noteExists = recordRef.fieldValues.hasOwnProperty(fieldID)
		
		if(formComponentIsReadOnly(noteEditorObjectRef.properties.permissions)) {
			if (noteExists) {
				$notePopupLink.css("display","")
				$notePopupLink.text("View note")
			} else {
				$notePopupLink.css("display","none")
				$notePopupLink.text("")
			}
		} else {
			$notePopupLink.css("display","")
			if (noteExists) {
				$notePopupLink.text("Edit note")
			} else {
				$notePopupLink.text("Add note")
			}
		}
	}
	
	var currRecordRef = null
	function loadRecordIntoHtmlEditor($htmlEditor, recordRef) {
		currRecordRef = recordRef
		formatNotePopupLinkText(recordRef)
	}
	
	console.log("Note editor table cell container: " + $container.html())
	
	
	$notePopupLink.popover({
		html: 'true',
		content: function() { return noteEditorTableViewContainerHTML() },
		trigger: 'manual',
		placement: 'auto left',
		container:'body'
	})
	
	$notePopupLink.click(function(e) {
		$(this).popover('toggle')
		e.stopPropagation()
	})
	
	
	$notePopupLink.on('shown.bs.popover', function()
	{
	    //get the actual shown popover
	    var $popover = $(this).data('bs.popover').tip();
		
		// By default the popover takes on the maximum size of it's containing
		// element. Overridding this size allows the size to grow as needed.
		$popover.css("max-width","300px")
		$popover.css("max-height","300px")
		console.log("Popover html: " + $popover.html())
		
		var $noteEditorContainer = $popover.find(".noteEditorPopupContainer")
		
		initHTMLEditorTextCellComponentViewModeGeometry($noteEditorContainer)
		
		var $closePopupButton = $noteEditorContainer.find(".closeEditorPopup")
		initButtonControlClickHandler($closePopupButton,function() {
			$notePopupLink.popover('hide')
		})
		
		
		console.log("Popover html: " + $noteEditorContainer.html())
		initNoteEditorTableRecordEditBehavior($noteEditorContainer,componentContext,recordProxy, noteEditorObjectRef)
		if(currRecordRef != null) {
			var viewConfig = $noteEditorContainer.data("viewFormConfig")
			viewConfig.loadRecord($noteEditorContainer,currRecordRef)
		}

	});
	
	$container.data("viewFormConfig", {
		loadRecord: loadRecordIntoHtmlEditor,
		validateValue: validateInput
	})
	
}