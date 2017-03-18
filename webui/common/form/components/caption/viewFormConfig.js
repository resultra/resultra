

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
