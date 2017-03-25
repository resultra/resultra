function initDeleteFormComponentPropertyPanel(params) {
	
	var $deleteButton = $('#'+ params.elemPrefix + 'DeleteFormComponentButton')
	
	initButtonControlClickHandler($deleteButton,function() {
		console.log("Delete component button clicked")
		openFormComponentConfirmDeleteDialog(params.componentLabel,function() {
			
			var deleteParams = {
				parentFormID: params.parentFormID,
				componentID: params.componentID
			}
			jsonAPIRequest("frm/deleteComponent",deleteParams,function(replyStatus) {
				params.$componentContainer.remove()
				console.log("Delete confirmed")
				saveUpdatedDesignFormLayout(params.parentFormID)
			})
			
			
		})
	})
}