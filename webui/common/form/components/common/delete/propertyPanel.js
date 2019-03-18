// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initDeleteFormComponentPropertyPanel(params) {
	
	var $deleteButton = $('#'+ params.elemPrefix + 'DeleteFormComponentButton')
	
	initButtonControlClickHandler($deleteButton,function() {
		console.log("Delete component button clicked")
		openConfirmDeleteDialog(params.componentLabel,function() {
			
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