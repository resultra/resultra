

function loadRecordIntoCaption(headerElem, recordRef) {
	// no-op
}

function initCaptionRecordEditBehavior($captionContainer,componentContext,captionObjectRef) {	
	$captionContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoCaption
	})
	
	console.log("Initializing caption: " + JSON.stringify(captionObjectRef))
	var $captionEditorControl = captionFromCaptionContainer($captionContainer)
	
	// When viewing a caption always open a new window for all links.
	var $caption = $(captionObjectRef.properties.caption)
	$caption.find('a').attr("target","_blank")
	
	$captionEditorControl.html($caption)
	
	
}
