

function openSubmitFormDialog(viewFormContext) {
	console.log("Opening form submit dialog box")
	
	var $dialog = $('#submitFormDialog')
	
	var $submitFormViewCanvas = $('#submitFormDialogLayoutCanvas')
	$submitFormViewCanvas.empty()
	
	
	var newRecordsParams = {
		parentDatabaseID:viewFormContext.databaseID,
		isDraftRecord:true
	}
	jsonAPIRequest("recordUpdate/newRecord",newRecordsParams,function(newRecordRef) {

		var newRecord = newRecordRef
		
		function getCurrentRecord() {
			return newRecord
		}
	
		function updateCurrentRecord(updatedRecordRef) {
			newRecord = updatedRecordRef
			loadRecordIntoFormLayout($submitFormViewCanvas,newRecord)
		}
		
		function showDialogAfterFormComponentLoaded() {
			$dialog.modal('show')
		}
	
		var recordProxy = {
			changeSetID: MainLineFullyCommittedChangeSetID,
			getRecordFunc: getCurrentRecord,
			updateRecordFunc: updateCurrentRecord
		}
		
		
		loadFormViewComponents($submitFormViewCanvas,viewFormContext,recordProxy,showDialogAfterFormComponentLoaded)
		
		initButtonClickHandler('#submitFormSaveButton', function() {
			console.log("Saving form results")
			
			var recordDraftParams = {
				recordID: newRecord.recordID,
				isDraftRecord: false
			}
			
			jsonAPIRequest("record/setDraftStatus",recordDraftParams,function(response) {
				$dialog.modal('hide')
			})

		})
		
		var defaultVals = viewFormContext.formLink.properties.defaultValues
		if (defaultVals.length > 0) {
			// Apply the default values before loading the form.
			var defaultValParams = {
				parentDatabaseID: viewFormContext.databaseID,
				recordID: newRecordRef.recordID,
				changeSetID: MainLineFullyCommittedChangeSetID,
				defaultVals: defaultVals }
			jsonAPIRequest("recordUpdate/setDefaultValues",defaultValParams,function(updatedRecordRef) {				
				updateCurrentRecord(updatedRecordRef)
			})
		}
		
		

		
	}) // getRecord
	
	
	
	
	
}