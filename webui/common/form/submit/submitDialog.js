

function openSubmitFormDialog(viewFormContext) {
	console.log("Opening form submit dialog box")
	
	var $dialog = $('#submitFormDialog')
	
	var viewFormCanvasSelector = '#submitFormDialogLayoutCanvas'
	var $viewFormCanvas = $(viewFormCanvasSelector)
	$viewFormCanvas.empty()
	
	
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
			loadRecordIntoFormLayout(viewFormCanvasSelector,newRecord)
		}
		
		function showDialogAfterFormComponentLoaded() {
			$dialog.modal('show')
		}
	
		loadFormViewComponents(viewFormCanvasSelector,viewFormContext,
			getCurrentRecord,updateCurrentRecord,showDialogAfterFormComponentLoaded)
		
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
		

		
	}) // getRecord
	
	
	
	
	
}