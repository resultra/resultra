// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function loadRecordIntoCaption($captionContainer, recordRef) {
	// no-op	
}

function initCaptionRecordEditBehavior($captionContainer,componentContext,captionObjectRef) {	
	$captionContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoCaption
	})
	
	console.log("Initializing caption: " + JSON.stringify(captionObjectRef))
	var $captionEditorControl = captionFromCaptionContainer($captionContainer)
	
	populateInlineDisplayContainerHTML($captionEditorControl,captionObjectRef.properties.caption)
	
	
}
