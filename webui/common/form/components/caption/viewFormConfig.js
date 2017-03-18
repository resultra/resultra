

function loadRecordIntoCaption($captionContainer, recordRef) {
	// no-op
	
	var viewConfig = $captionContainer.data("viewFormConfig")
	var componentID = viewConfig.captionRef.captionID
	
	var hiddenComponents = new IDLookupTable(recordRef.hiddenFormComponents)
	if (hiddenComponents.hasID(componentID)) {
		if (elemIsDisplayed($captionContainer)) {
			$captionContainer.animate({opacity:0},500,function() {
				// fade out, then hide completely
				$captionContainer.hide()
			})			
		}
		
	} else {
		if (!elemIsDisplayed($captionContainer)) {
			$captionContainer.show() // show it but opacity will still be 0
			$captionContainer.animate({opacity:1},500) // fade in
		}
	}
	
}

function initCaptionRecordEditBehavior($captionContainer,componentContext,captionObjectRef) {	
	$captionContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoCaption,
		captionRef: captionObjectRef
	})
	
	console.log("Initializing caption: " + JSON.stringify(captionObjectRef))
	var $captionEditorControl = captionFromCaptionContainer($captionContainer)
	
	populateInlineDisplayContainerHTML($captionEditorControl,captionObjectRef.properties.caption)
	
	
}
