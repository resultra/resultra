// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function openAttachmentConfirmDeleteDialog(deleteAttachmentCallback) {
	
	var $dialog = $('#attachmentConfirmDeleteDialog')
	
	$dialog.modal("show")
	
	var $confirmDeleteButton = $dialog.find(".attachmentConfirmDeleteButton")
	
	initButtonControlClickHandler($confirmDeleteButton,function() {
		deleteAttachmentCallback()
		$dialog.modal("hide")
	})
	
	
}