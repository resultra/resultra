function initDeleteFormComponentPropertyPanel(params) {
	
	var $deleteButton = $('#'+ params.elemPrefix + 'DeleteFormComponentButton')
	
	initButtonControlClickHandler($deleteButton,function() {
		console.log("Delete component button clicked")
		openFormComponentConfirmDeleteDialog(params.componentLabel,function() {
			console.log("Delete confirmed")
		})
	})
}