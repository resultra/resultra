
function initCaptionDesignInlineEditBehavior($captionContainer,captionObjectRef,enableDesignBehaviorCallback) {
	
	CKEDITOR.disableAutoInline = true;
	var $captionEditorControl = captionFromCaptionContainer($captionContainer)
	
	populateInlineDisplayContainerHTML($captionEditorControl,captionObjectRef.properties.caption)
	
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
					
					// Edit links in the caption to open a new window.
					populateInlineDisplayContainerHTML($captionEditorControl,editorInput)
				
					// Re-enable design behavior for this component
					enableDesignBehaviorCallback()
				
				})
			
				editor.focus()
			}
				
		}) 
		
	}
	
	initializeCaptionInlineEditing()
	
	

}

function initCaptionDesignControlBehavior($caption, captionObjectRef) {
  var formID = captionObjectRef.parentFormID
  var designFormLayoutConfig =  createFormLayoutDesignConfig(formID)
  var componentIDs = { formID: formID, componentID: captionObjectRef.captionID }
  
  // Call-back to enable design behavior. This is called initially when the caption is
  // created. It is also called when the caption exits edit mode in the form designer.
  function enableDesignBehavior() {
	  initFormComponentDesignBehavior($caption,componentIDs,
			captionObjectRef,formCaptionDesignFormConfig,designFormLayoutConfig)
	  }
  initCaptionDesignInlineEditBehavior($caption,captionObjectRef,enableDesignBehavior)
	  
  enableDesignBehavior() // initially enable design behavior
}


function initDesignFormCaption() {
	console.log("Init caption design form behavior")
//	initNewCheckBoxDialog()
}

function selectFormCaption($container,captionObjRef) {
	console.log("Selected caption: " + JSON.stringify(captionObjRef))
	loadFormCaptionProperties($container,captionObjRef)
}

function resizeFormCaption($container,geometry) {
	
	var captionRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		captionID: captionRef.captionID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/caption/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.captionID)
	})	
}

var formCaptionDesignFormConfig = {
	draggableHTMLFunc:	formCaptionContainerHTML,
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewFormCaptionDialog,
	resizeConstraints: elemResizeConstraints(125,1280,125,1280),
	resizeHandles: 'e,s,se',
	resizeFunc: resizeFormCaption,
	initFunc: initDesignFormCaption,
	selectionFunc: selectFormCaption
}
