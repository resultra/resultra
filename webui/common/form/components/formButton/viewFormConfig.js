var FormButtonPopupBehaviorModal = "modal"

var viewFormInViewportEventName = "view-form-in-viewport"

function loadRecordIntoButton(buttonElem, recordRef) {	
	// no-op
}

function initFormButtonRecordEditBehavior($buttonContainer,componentContext,
			parentRecordProxy,buttonObjectRef,defaultValSrc,
		loadFormViewComponentFunc,loadRecordIntoFormLayoutFunc) {
	
	
	$buttonContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoButton
	})
	
	var $popupFormDialog = $('#formButtonPopupFormDialog')
			
	var popupMode = buttonObjectRef.properties.popupBehavior.popupMode
	
	var $formButton = buttonFromFormButtonContainer($buttonContainer)
	
	
	function showRecordInPopupView() {
		var currRecord = parentRecordProxy.getRecordFunc()

		var $popupFormViewCanvas = $('#formButtonPopupFormCanvas')
		$popupFormViewCanvas.empty()

		
		function getPopupFormRecordFunc() { return currRecord }
		function updatePopupFormRecordFunc(updatedRecordRef) {
			
			loadRecordIntoFormLayoutFunc($popupFormViewCanvas,updatedRecordRef)
			
			// Propagate the record update the parent form, allowing an update
			// to the parent layout. The changes are only propagated when the popup
			// form is not in a modal state and the changes are committed immediately.
			// If the popup is in a modal state, the changes are linked to a temporary
			// changeSetID and are not propagated until the "Save Changes" button is
			// pressed; however, if the "Cancel" button is pressed all these changes are
			// rolled back.
			if (popupMode !== FormButtonPopupBehaviorModal) {
				parentRecordProxy.updateRecordFunc(updatedRecordRef)			
			}
		}
		
		function showDialogAfterFormComponentLoaded() {
			loadRecordIntoFormLayoutFunc($popupFormViewCanvas ,currRecord)
			$popupFormDialog.modal('show')
		}
			
		var viewFormContext = {
			databaseID: componentContext.databaseID,
			formID: buttonObjectRef.properties.linkedFormID
		}
				
		if (popupMode === FormButtonPopupBehaviorModal) {
			
			$(".formButtonPopupModalPopupDialogButton").show()
			$(".formButtonPopupModelessPopupDialogButton").hide()
			
			
			jsonAPIRequest("record/allocateChangeSetID",{},function(changeSetIDResp) {
				
				var recordProxy = {
					changeSetID: changeSetIDResp.changeSetID,
					getRecordFunc: getPopupFormRecordFunc,
					updateRecordFunc: updatePopupFormRecordFunc
				}
				
				var defaultVals = buttonObjectRef.properties.defaultValues
				
				if (defaultVals !== undefined && defaultVals.length > 0) {
					// Apply the default values before loading the form.
					var defaultValRecord = getPopupFormRecordFunc()
					var defaultValParams = {
						parentDatabaseID: componentContext.databaseID,
						recordID: defaultValRecord.recordID,
						changeSetID: changeSetIDResp.changeSetID,
						defaultVals: defaultVals }
					jsonAPIRequest("recordUpdate/setDefaultValues",defaultValParams,function(updatedRecordRef) {
						
						currRecord = updatedRecordRef
						
						// loadFormViewComponentFunc is passed in as a parameter, since the loadFormViewComponents 
						// function is in 
						// a package which has a dependency on this package.
									 
						loadFormViewComponentFunc($popupFormViewCanvas, viewFormContext, recordProxy,
							 showDialogAfterFormComponentLoaded)
					})
					
					
				} else {
					loadFormViewComponentFunc($popupFormViewCanvas, viewFormContext, recordProxy,
						 showDialogAfterFormComponentLoaded)
				}
				
					
				initButtonClickHandler('#formButtonPopupFormDialogSaveChangesButton', function() {
					console.log("Modal Save changes button clicked: " + JSON.stringify(buttonObjectRef))
					// TODO - Remove the temporary changes set ID for any changes made while editing the record.
					
					var commitChangeParams = {
						recordID: getPopupFormRecordFunc().recordID,
						changeSetID: changeSetIDResp.changeSetID }
					jsonAPIRequest("recordUpdate/commitChangeSet",commitChangeParams,function(updatedRecordRef) {
						// If the popup form is modal, the parent form's record is not updated until the "Save Changes" button
						// is pressed.
						parentRecordProxy.updateRecordFunc(updatedRecordRef)
						$popupFormDialog.modal('hide')
					})
				})
				initButtonClickHandler('#formButtonPopupFormDialogCancelChangesButton', function() {
					console.log("Cancel button clicked: " + JSON.stringify(buttonObjectRef))
					$popupFormDialog.modal('hide')
				})
				
			
			})
			
		} else { // Popup shown in modeless mode
			
			// TBD - Should the default value be set when the popup form is modeless? It wouldn't
			// seem to make sense to set the default values in this case.
			
			$(".formButtonPopupModalPopupDialogButton").hide()
			$(".formButtonPopupModelessPopupDialogButton").show()
			initButtonClickHandler('#formButtonPopupFormDialogDoneButton', function() {
				console.log("Form button clicked: " + JSON.stringify(buttonObjectRef))
				$popupFormDialog.modal('hide')
			})
			
			var immediatelyCommitChangesChangeSetID = ""
			var recordProxy = {
				changeSetID: immediatelyCommitChangesChangeSetID,
				getRecordFunc: getPopupFormRecordFunc,
				updateRecordFunc: updatePopupFormRecordFunc
			}
			
			loadFormViewComponentFunc($popupFormViewCanvas, viewFormContext, recordProxy,
				 showDialogAfterFormComponentLoaded)
		}
		
	}
	
	function viewFormURL() {
		var currRecord = parentRecordProxy.getRecordFunc()
		return '/viewItem/' + buttonObjectRef.properties.linkedFormID + 
				'/' + currRecord.recordID + '?' + defaultValSrc
	}
	
	function showRecordInNewPage() {
		var win = window.open(viewFormURL(),"_blank")
		win.focus()
	}
	
	function showRecordInPage() {
		
		var getFormInfoParams = { formID: buttonObjectRef.properties.linkedFormID }
		
		jsonAPIRequest("frm/getFormInfo", getFormInfoParams, function(formInfo) {
			var currRecord = parentRecordProxy.getRecordFunc()
			var viewFormParams = {
				formID: buttonObjectRef.properties.linkedFormID,
				databaseID: componentContext.databaseID,
				title: formInfo.form.name,
				recordID: currRecord.recordID,
				defaultVals: buttonObjectRef.properties.defaultValues,
				saveMode: buttonObjectRef.properties.popupBehavior.popupMode	
			}
			console.log("Form button pressed: navigating to view form: " + JSON.stringify(viewFormParams))
			$buttonContainer.trigger(viewFormInViewportEventName,viewFormParams)
		})
		
		
	//	navigateToURL(viewFormURL())
	}
			
	initButtonControlClickHandler($formButton, function() {
		var showFormDest = buttonObjectRef.properties.popupBehavior.whereShowForm
		if(showFormDest === 'popup') {
			showRecordInPopupView()	
		} else if (showFormDest === 'page'){
			showRecordInPage()
		} else if (showFormDest === 'newPage') {
			showRecordInNewPage()
		}
	})
	
	
	
}
