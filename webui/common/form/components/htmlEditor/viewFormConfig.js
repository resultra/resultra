

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
		remoteValidationFunc(currentInputValText,function(validationResult) {
			if (validationResult.validationSucceeded) {
				$htmlEditor.popover('destroy')
				validationCompleteCallback(true)
			} else {
				$htmlEditor.popover({
					html: 'true',
					content: function() { return escapeHTML(validationResult.errorMsg) },
					trigger: 'manual',
					placement: 'auto left'
				})
				$htmlEditor.popover('show')
				validationCompleteCallback(false)
			}			
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
		
					var textBoxTextValueFormat = {
						context:"htmlEditor",
						format:"general"
					}

					var setRecordValParams = { 
						parentDatabaseID:currRecordRef.parentDatabaseID,
						recordID:currRecordRef.recordID,
						changeSetID: recordProxy.changeSetID,
						fieldID:htmlEditorFieldID, 
						value:editorVal,
						valueFormat:textBoxTextValueFormat 
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
	
		var $clearValueButton = $htmlEditor.find(".editorComponentClearValueButton")
		initButtonControlClickHandler($clearValueButton,function() {
			console.log("Clear value clicked for editor")
			$htmlEditorInput.html("")
			setEditorValue(null)
		})
	
	
		var $editButton = $htmlEditor.find(".startEditButton")
		initButtonControlClickHandler($editButton,function() {
			console.log("Starting inline editor")
		
			if (!inlineCKEditorEnabled($htmlEditorInput)) {
			    CKEDITOR.disableAutoInline = true;
	
				var editor = enableInlineCKEditor($htmlEditorInput)
		
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
		jsonAPIRequest("frm/numberInput/validateInput", validationParams, function(validationResult) {
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
		
	initHtmlEditorRecordEditBehavior($container,componentContext,recordProxy, 
			noteEditorObjectRef,validateInput)
}


function initNoteEditorTableCellEditBehavior($container,componentContext,recordProxy, noteEditorObjectRef) {

	// TBD - Needs a popup to display the editor.
	var validateInput = function(validationCompleteCallback) {
			validationCompleteCallback(true)
	}
	
	function loadRecordIntoHtmlEditor($htmlEditor, recordRef) {
		// no-op
	}
	
	console.log("Note editor table cell container: " + $container.html())
	
	$container.data("viewFormConfig", {
		loadRecord: loadRecordIntoHtmlEditor,
		validateValue: validateInput
	})
	
}