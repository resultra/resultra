function initDeleteDashboardComponentPropertyPanel(params) {
	
	var $deleteButton = $('#'+ params.elemPrefix + 'DeleteDashboardComponentButton')
	
	initButtonControlClickHandler($deleteButton,function() {
		console.log("Delete component button clicked")
		openConfirmDeleteDialog(params.componentLabel,function() {
			
			var deleteParams = {
				parentDashboardID: params.parentDashboardID,
				componentID: params.componentID
			}
			jsonAPIRequest("dashboard/deleteComponent",deleteParams,function(replyStatus) {
				params.$componentContainer.remove()
				console.log("Delete confirmed")
				saveUpdatedDesignDashboardLayout()
			})
			
			
		})
	})
}