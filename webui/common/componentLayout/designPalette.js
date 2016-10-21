

// A placeholderID is a temporary ID to assign to a div. After saving a 
// new object via JSON call, it is replaced with a unique ID created by the server.
var paletteItemPlaceholderNum = 1
function allocNextPaletteItemPlaceholderID()
{
	placeholderID = "paletteItemPlaceholderID" + paletteItemPlaceholderNum.toString()
	paletteItemPlaceholderNum = paletteItemPlaceholderNum + 1
	return placeholderID
}

function initDesignPalette(paletteConfig,layoutDesignConfig) {
	$(".paletteItem").draggable({
	
		cursor: "move",
		// The cursorAt option aligns the draggable object at an offset from the cursor, rather than from
		// the top corner of the palette item.
		cursorAt: { top: 10, left: 10 },
				
		helper: function() { 
			
			var paletteItemID = $(this).attr('id')
			assert(paletteItemID !== undefined, "designPalette: palette item missing element id")
			console.log("designPallete: helper: ID of palette item: " + paletteItemID)
			
			 // Instead of dragging the icon in the palette, drag a "placeholder image" of the
			 // item itself. This allows the user to see the proporitions of the field relative
			 // to the other elements already on the canvas.
			var newItemDraggablePlaceholderHTML = paletteConfig.draggableItemHTML(allocNextPaletteItemPlaceholderID(),paletteItemID);
			var $placeholder = $(newItemDraggablePlaceholderHTML);
					
			 // While in layout mode, disable entry into the fields
			 // Interestingly, there's not CSS for this.
			$placeholder.find('input').prop('disabled',true);
			
			// Set some data within the DOM element being dragged strictly for passing parameters
			// for drag and drop. Adding the class "newComponent" allows the recipient of the drop
			// to distinguish between existing components being dragged to new position or new components.
			$placeholder.data("paletteConfig",paletteConfig)
			$placeholder.data("paletteItemID",paletteItemID)
			$placeholder.addClass("newComponent")
			
			
			// TODO - Test & integrate the following code to make the palette items drag and drop
			// more seamlessly with the jQuery UI layout plug-in. This is the recommended way
			// to do drag and drop under jQuery UI Layout (see http://layout.jquery-dev.com/tips.cfm#Widget_Draggable)
			//placeholder.appendTo('body').css('zIndex',5).show()
		 			 
			 return $placeholder
		 },
		 
		 start: function(event,ui) {
		 	console.log("drag start")
			 
			 var placeholderID = $(ui.helper).attr('id')
			 var paletteItemID = $(this).attr('id')
			 
			 // The startPaletteDrag function is called so any dynamic Javascript initialization
			 // which needs to take place after the helper has been added can take place.
			 paletteConfig.startPaletteDrag(placeholderID,paletteItemID)
			 
		 },
		 
 		drag: function(e,ui) {			
 			var mouseOffset = { 
 				top: e.pageY, 
 				left: e.pageX }
 			console.log("Drag component: mouse offset: " + JSON.stringify(mouseOffset))
 			highlightDroppablePlaceholder(mouseOffset)
 		}, // drag
		
		 stop: function( e, ui ) {
				var mouseOffset = { 
					top: e.pageY, 
					left: e.pageX }
					
			var paletteConfig = $(ui.helper).data("paletteConfig")
			var paletteItemID = $(ui.helper).data("paletteItemID")
			var componentHTML = paletteConfig.draggableItemHTML(allocNextPaletteItemPlaceholderID(),paletteItemID)
					
			var $draggedPlaceholderComponent = $(componentHTML);
			$draggedPlaceholderComponent.data("paletteConfig",paletteConfig)
			$draggedPlaceholderComponent.data("paletteItemID",paletteItemID)
				
			function newComponentDropComplete($droppedObj) {
		
				$droppedObj.removeClass("newComponent")
		
				var placeholderID = $droppedObj.attr('id')
				assert(placeholderID !== undefined, "receiveNewComponent: palette item missing element id")
				console.log("receiveNewComponent: drop: placeholder ID of palette item: " + placeholderID)
	
				var objWidth = $droppedObj.width()
				var objHeight = $droppedObj.height()
	
				var paletteItemID = $droppedObj.data("paletteItemID")
				console.log("receiveNewComponent: drop: palette item ID/type: " + paletteItemID)
	
				var paletteConfig = $droppedObj.data("paletteConfig")
		
				var componentParentLayoutSelector = paletteConfig.dropDestSelector
		
				var droppedObjInfo = {
					droppedElem: $droppedObj,
					paletteItemID: paletteItemID,
					placeholderID: placeholderID,
					geometry: {positionTop: 0, positionLeft: 0,
					sizeWidth: objWidth,sizeHeight: objHeight},
					finalizeLayoutIncludingNewComponentFunc: function() {
							console.log("receiveNewComponent: finalizing layout with new component")
							var updatedLayout = getComponentLayout($parentLayout)
							saveLayoutFunc(updatedLayout)
						}
				};
		
				paletteConfig.dropComplete(droppedObjInfo)
		
			}	
			handleDropOnComponentLayoutPlaceholder(mouseOffset,layoutDesignConfig,$draggedPlaceholderComponent,newComponentDropComplete)
		 },
		
		 
		// The following are needed to keep the draggable above the other elements. Getting Semantic UI,
		// jQuery UI Layout and jQuery UI drag-and-drop working together is really tricky. Some issues:
		//
		// 1. jQuery UI layout has difficulty with the z-index for items which are opened in one pane,
		//    but overlap another. The recommended solution is to attach items which need to overlap all
		//    the panes to <body>. If this isn't possible, the west__showOverflowOnHover option is provided
		//    for initializing the panes.
		// 2. The Semantic UI form elements inherit some of their styling from the <form> styling. So,
		//    if a form element is not embeded in a <form> tag, the styling isn't correct. To work around this,
		//    the entire page is embedded in a <form> tag. However, it wasn't possible to attach the item to
		//    <body> and still have it inherit Semantic UI's <form> styling.
		//
		// So, the following solution is the only solution which could be found to address both issues.
		// described above.
		stack: "div",
		appendTo: 'body',
		zIndex: 999
	
	});
	
} // initDesignPalette

