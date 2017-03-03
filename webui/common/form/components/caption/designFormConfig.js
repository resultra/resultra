
function initCaptionDesignControlBehavior($captionContainer,captionObjectRef,enableDesignBehaviorCallback) {
	
	CKEDITOR.disableAutoInline = true;
	var $captionEditorControl = captionFromCaptionContainer($captionContainer)
	$captionEditorControl.html(captionObjectRef.properties.caption)
	
	function initializeCaptionInlineEditing() {
		
		$captionEditorControl.unbind("dblclick")
		$captionEditorControl.dblclick(function(e) {
			
			console.log("Caption editor control area double clicked")
			
			e.stopPropagation()
		
			if (!inlineCKEditorEnabled($captionEditorControl)) {
			
				// Disable drag-and-drop design behavior for this component. This interferes with
				// selections, cut-and-paste, etc in the editor.
				disableObjectEditBehavior($captionContainer)
			
				console.log("Starting inline editor for caption")
				var editor = enableInlineCKEditor($captionEditorControl)
				
				editor.setData($captionEditorControl.html())
			
				editor.on('blur', function(event) {
					var editorInput = editor.getData();
		
					var setCaptionParams = {
						parentFormID: captionObjectRef.parentFormID,
						captionID: captionObjectRef.captionID,
						caption: editorInput
					}
					console.log("Caption edit complete: " + JSON.stringify(setCaptionParams))
					jsonAPIRequest("frm/caption/setCaption",setCaptionParams,function(updatedCaption) {
					})
				
					disableInlineCKEditor($captionEditorControl,editor)
				
					// Re-enable design behavior for this component
					enableDesignBehaviorCallback()
				
				})
			
				editor.focus()
			}
				
		}) 
		
	}
	
	initializeCaptionInlineEditing()
	
	

}


function initDesignFormCaption() {
	console.log("Init caption design form behavior")
//	initNewCheckBoxDialog()
}

function selectFormCaption(captionObjRef) {
	console.log("Selected caption: " + JSON.stringify(captionObjRef))
	loadFormCaptionProperties(captionObjRef)
}

function resizeFormCaption(captionID,geometry) {
	var resizeParams = {
		parentFormID: designFormContext.formID,
		captionID: captionID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/caption/resize", resizeParams, function(updatedObjRef) {
		setElemObjectRef(captionID,updatedObjRef)
	})	
}

var formCaptionDesignFormConfig = {
	draggableHTMLFunc:	formCaptionContainerHTML,
	startPaletteDrag: function(placeholderID,$paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewFormCaptionDialog,
	resizeConstraints: elemResizeConstraints(320,640,50,50),
	resizeFunc: resizeFormCaption,
	initFunc: initDesignFormCaption,
	selectionFunc: selectFormCaption
}
