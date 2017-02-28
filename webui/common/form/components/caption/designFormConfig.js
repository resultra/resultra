
function initCaptionDesignControlBehavior($captionContainer,captionObjectRef,enableDesignBehaviorCallback) {
	
	CKEDITOR.disableAutoInline = true;
	var $captionEditorControl = captionFromCaptionContainer($captionContainer)
	$captionEditorControl.html(captionObjectRef.properties.caption)
	
	$captionEditorControl.dblclick(function(e) {
	
		e.stopPropagation()
		
		var currentlyEditable = $captionEditorControl.attr("contenteditable")
		if (currentlyEditable !== "true") {
			
			// Disable drag-and-drop design behavior for this component. This interferes with
			// selections, cut-and-paste, etc in the editor.
			disableObjectEditBehavior($captionContainer)
			
			$captionEditorControl.attr("contenteditable","true")
			var captionEditorInputDomElem = $captionEditorControl.get(0)
			var editor = CKEDITOR.inline( captionEditorInputDomElem )
			
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
				
				$captionEditorControl.attr("contenteditable","false")
				editor.destroy()
				
				// Re-enable design behavior for this component
				enableDesignBehaviorCallback()
				
			})
			
			
			$captionEditorControl.focus()
		}
		console.log("editor control click")
				
	}) 

	
	
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
