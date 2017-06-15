function initSubmitFormUILayoutPanes()
{
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }
	
	// Initialize the page layout
	var mainLayout = $('#layoutPage').layout({
		inset: zeroPaddingInset,
		north: fixedUILayoutPaneParams(40),
	})
				
}


function initFormPageSubmitForm(params) {
	
	var $submitButton = $('#submitFormPageSubmitButton')
	var $saveConfirmation = $('#newItemSaveConfirmation')
	
	// The form may be re-initialized in the case of a user who submits one form,
	// then chooses to submit another immediately thereafter.
	params.$parentFormCanvas.empty()	
	$saveConfirmation.hide()
	params.$parentFormCanvas.show()
	$submitButton.prop('disabled', false);
	
	var newRecordsParams = {
		parentDatabaseID:params.databaseID,
		isDraftRecord:true
	}
	jsonAPIRequest("recordUpdate/newRecord",newRecordsParams,function(newRecordRef) {

		var newRecord = newRecordRef
		
		function getCurrentRecord() {
			return newRecord
		}
	
		function updateCurrentRecord(updatedRecordRef) {
			newRecord = updatedRecordRef
			loadRecordIntoFormLayout(params.$parentFormCanvas,newRecord)
		}
		
		function finalizePageAfterFormComponentsLoaded() {
			console.log("Submit form: form components loaded")
		}
	
		var recordProxy = {
			changeSetID: MainLineFullyCommittedChangeSetID,
			getRecordFunc: getCurrentRecord,
			updateRecordFunc: updateCurrentRecord
		}
		
		
		var formContext = {
			databaseID: params.databaseID,
			formID: params.formID
		}
		
		loadFormViewComponents(params.$parentFormCanvas,formContext,recordProxy,
						finalizePageAfterFormComponentsLoaded)
		
		
		initButtonControlClickHandler($submitButton, function() {
					
			validateFormValues(params.$parentFormCanvas,function(validationResult) {
				if(validationResult === true) {
					console.log("Saving form results")
					
					 $submitButton.prop('disabled', true);
			
					var recordDraftParams = {
						recordID: newRecord.recordID,
						isDraftRecord: false
					}
			
					jsonAPIRequest("record/setDraftStatus",recordDraftParams,function(response) {
						console.log("Submit form: form data submitted/saved")
						// Reset the form and re-load a different record.
						jsonAPIRequest("recordUpdate/newRecord",newRecordsParams,function(newRecordRef) {
							updateCurrentRecord(newRecordRef)
							
							params.$parentFormCanvas.hide()
							$saveConfirmation.show()
						})
					})
				} else {
					console.log("Form validation failed: not saving form results")
				}
			})	

		})
		
		var formLinkParams = { formLinkID: params.formLinkID }
		jsonAPIRequest("formLink/get",formLinkParams,function(formLink) {
			var defaultVals = formLink.properties.defaultValues
			if (defaultVals.length > 0) {
				// Apply the default values before loading the form.
				var defaultValParams = {
					parentDatabaseID: params.databaseID,
					recordID: newRecordRef.recordID,
					changeSetID: MainLineFullyCommittedChangeSetID,
					defaultVals: defaultVals }
				jsonAPIRequest("recordUpdate/setDefaultValues",defaultValParams,function(updatedRecordRef) {				
					updateCurrentRecord(updatedRecordRef)
				})
			}
		})
		
	}) // getRecord
	
}
