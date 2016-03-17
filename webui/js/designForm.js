

function initLayoutEditFieldInfo(fieldRef)
{
	// Add the ability to select the field from the new layout container
	// dialog. This callback function is called for both number and 
	// text fields; so, text boxes can be used for both numbers and 
	// regular text.
 	var selectFieldOptionHTML = '<option value="' +
 		fieldRef.fieldID + '">' +
 		fieldRef.fieldInfo.name + '</option>'
 	$("#textBoxFieldSelection").append(selectFieldOptionHTML)	
}

function initCanvasComplete() {	
	
} // noop



function initContainerEditBehavior(container)
{

	// While in edit mode, disable input on the container
	container.find('input').prop('disabled',true);


	// TODO - This could be put into a common function, since these
	// properties should be the same for objects loaded with the page
	// and newly added objects.
	container.draggable ({
		grid: [20, 20], // snap to a grid
		cursor: "move",
		containment: "parent",
		clone: "original",				
		stop: function(event, ui) {
				  var layoutPos = {
					  positionLeft: ui.position.left,
					  positionTop: ui.position.top,
				   };
		  
				  console.log("drag stop: id: " + event.target.id);
				  console.log("drag stop: new position: " + JSON.stringify(layoutPos));
				  
				  // TODO: send ajax request to reposition the container
  
		      } // stop function
	})
	container.resizable({
		aspectRatio: false,
		handles: 'e, w', // Only allow resizing horizontally
		maxWidth: 600,
		minWidth: 100,
		grid: 20, // snap to grid during resize

		stop: function(event, ui) {  
				  
	  			var layoutContainerParams = {
	  				parentLayoutID: layoutID,containerID: event.target.id,
	  				geometry: { positionTop: ui.position.top,positionLeft: ui.position.left,
						sizeWidth: ui.size.width, sizeHeight: ui.size.height }
	  			};
		
			 	jsonAPIRequest("resizeLayoutContainer",layoutContainerParams,function(replyData) {})	
				  
			  } // stop function
	});
	
}


function droppedObjGeometry(dropDest,droppedObj,ui)
{
	// TODO - Support snapping of the top and left
	var relTop = ui.offset.top-$(dropDest).offset().top
	var relLeft = ui.offset.left-$(dropDest).offset().left
	var objWidth = droppedObj.width()
	var objHeight = droppedObj.height()
		
	return { top: relTop, left: relLeft, width: objWidth, height: objHeight }
}


$(document).ready(function() {
				
	
	$(".paletteItem").draggable({
		
		revert: 'invalid',
		cursor: "move",
		helper: function() { 
			 // Instead of dragging the icon in the gallery, drag a "placeholder image" of the
			 // field itself. This allows the user to see the proporitions of the field relative
			 // to the other elements already on the canvas.
			var newFieldDraggablePlaceholderHTML = fieldContainerHTML(allocNextPlaceholderID());
			var placeholder = $(newFieldDraggablePlaceholderHTML);
						
			 // While in layout mode, disable entry into the fields
			 // Interestingly, there's not CSS for this.
			placeholder.find('input').prop('disabled',true);
			 			 
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
		appendTo: "#gallerySidebar",
		zIndex: 999
		
	});
	
    $("#layoutCanvas").droppable({
        accept: ".paletteItem",
        activeClass: "ui-state-highlight",
        drop: function( event, ui ) {
 			var theClone = $(ui.helper).clone()
			
			// After the drag operation is complete, the draggable and resizable
			// properties need to be initialized.
			initContainerEditBehavior(theClone)
					
			$(this).append(theClone);

			// IMPORTANT - Don't call droppedObjGeometry until after theClone has been appended
			// Otherwise the width and height will be returned as 0.
			var placeholderID = $(theClone).attr("id");
			var droppedObjGeom = droppedObjGeometry(this,theClone,ui)
			
			theClone.css({top: droppedObjGeom.top, left: droppedObjGeom.left, position:"absolute"});
			
		    console.log("End Drag and drop: placeholder ID=" + placeholderID + 
				" drop loc =" + JSON.stringify(droppedObjGeom));
			
			var layoutContainerParams = {
				parentLayoutID: layoutID, containerID: placeholderID,
				geometry: {positionTop: droppedObjGeom.top, positionLeft: droppedObjGeom.left,
				sizeWidth: droppedObjGeom.width,sizeHeight: droppedObjGeom.height}
				};
			
			newLayoutContainer(layoutContainerParams,fieldsByID)
						
        }
    }); // #layoutCanvas droppable
	
	
	initNewTextBoxDialog() 
		
	// Initialize the page layout
	$('#layoutPage').layout({
		north: fixedUILayoutPaneParams(40),
		east: fixedUILayoutPaneParams(300),
		west: fixedUILayoutPaneParams(200),
		west__showOverflowOnHover:	true
	})	  
	  
	
	initCanvas(initContainerEditBehavior,initLayoutEditFieldInfo,initCanvasComplete)
	  
});
