

// A placeholderID is a temporary ID to assign to a div. After saving a 
// new object via JSON call, it is replaced with a unique ID created by the server.
var paletteItemPlaceholderNum = 1
function allocNextPaletteItemPlaceholderID()
{
	placeholderID = "paletteItemPlaceholderID" + paletteItemPlaceholderNum.toString()
	paletteItemPlaceholderNum = paletteItemPlaceholderNum + 1
	return placeholderID
}

function droppedPaletteItemGeometry(dropDest,droppedObj,ui)
{
	// TODO - Support snapping of the top and left
	var relTop = ui.offset.top-$(dropDest).offset().top
	var relLeft = ui.offset.left-$(dropDest).offset().left
	var objWidth = droppedObj.width()
	var objHeight = droppedObj.height()
		
	return { top: relTop, left: relLeft, width: objWidth, height: objHeight }
}


function initDesignPalette(paletteConfig) {
	$(".paletteItem").draggable({
	
		revert: 'invalid',
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
			var placeholder = $(newItemDraggablePlaceholderHTML);
					
			 // While in layout mode, disable entry into the fields
			 // Interestingly, there's not CSS for this.
			placeholder.find('input').prop('disabled',true);
			
			
			// TODO - Test & integrate the following code to make the palette items drag and drop
			// more seamlessly with the jQuery UI layout plug-in. This is the recommended way
			// to do drag and drop under jQuery UI Layout (see http://layout.jquery-dev.com/tips.cfm#Widget_Draggable)
			//placeholder.appendTo('body').css('zIndex',5).show()
		 			 
			 return placeholder
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
		appendTo: paletteConfig.paletteSelector,
		zIndex: 999
	
	});

	$(paletteConfig.dropDestSelector).droppable({
	    accept: ".paletteItem",
	    activeClass: "ui-state-highlight",
	    drop: function( event, ui ) {
			var theClone = $(ui.helper).clone()
			
			var paletteItemID = ui.draggable.attr('id')
			assert(paletteItemID !== undefined, "designPalette: palette item missing element id")
			console.log("designPallete: drop: ID of palette item: " + paletteItemID)
						
			$(this).append(theClone);

			// IMPORTANT - Don't call droppedPaletteItemGeometry until after theClone has been appended
			// Otherwise the width and height will be returned as 0.
			var placeholderID = $(theClone).attr("id");
			var droppedObjGeom = droppedPaletteItemGeometry(this,theClone,ui)
		
			theClone.css({top: droppedObjGeom.top, left: droppedObjGeom.left, position:"absolute"});
		
		    console.log("designPalette: End Palette item drag and drop: placeholder ID=" + placeholderID + 
				" drop geometry =" + JSON.stringify(droppedObjGeom));
		
			var droppedObjInfo = {
				droppedElem: theClone,
				paletteItemID: paletteItemID,
				placeholderID: placeholderID,
				geometry: {positionTop: droppedObjGeom.top, positionLeft: droppedObjGeom.left,
				sizeWidth: droppedObjGeom.width,sizeHeight: droppedObjGeom.height}
				};
		
			paletteConfig.dropComplete(droppedObjInfo)
					
	    }
	}); // droppable
	
} // initDesignPalette

