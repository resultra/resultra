

function openSubmitFormDialog(viewFormContext) {
	console.log("Opening form submit dialog box")
	
	var $dialog = $('#submitFormDialog')
	
	var viewFormCanvasSelector = '#submitFormDialogLayoutCanvas'
	var $viewFormCanvas = $(viewFormCanvasSelector)
	$viewFormCanvas.empty()
	
	
	var newRecordsParams = {parentDatabaseID:viewFormContext.databaseID}
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
		

		
	}) // getRecord
	
	
	
	
	
}