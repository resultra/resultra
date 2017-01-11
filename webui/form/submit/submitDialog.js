

function openSubmitFormDialog(viewFormContext) {
	console.log("Opening form submit dialog box")
	
	var $dialog = $('#submitFormDialog')
	
	var viewFormCanvasSelector = '#submitFormDialogLayoutCanvas'
	var $viewFormCanvas = $(viewFormCanvasSelector)
	$viewFormCanvas.empty()
	
	loadFormViewComponents(viewFormCanvasSelector,viewFormContext,function() {
		$dialog.modal('show')
	})
	
	
}