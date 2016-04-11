

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


$(document).ready(function() {
				
	
	var paletteConfig = {
		draggableItemHTML: function(placeholderID,paletteItemID) { 
			return fieldContainerHTML(placeholderID)
		},
		
		dropComplete: function(droppedItemInfo) {
			console.log("designForm: Dashboard design pallete: drop item: " + JSON.stringify(droppedItemInfo))			
			
			// After the drag operation is complete, the draggable and resizable
			// properties need to be initialized.
			initContainerEditBehavior(droppedItemInfo.droppedElem)
			
			// "repackage" the dropped item paramaters for creating a new layout element. Also add the layoutID
			// to the parameters.
			var layoutContainerParams = {
				parentLayoutID: layoutID,
				geometry: droppedItemInfo.geometry,
				containerID: droppedItemInfo.placeholderID,
				};
			newLayoutContainer(layoutContainerParams,fieldsByID)
			//	openNewCheckboxDialog(layoutContainerParams,fieldsByID)
		},
		
		dropDestSelector: "#layoutCanvas",
		paletteSelector: "#gallerySidebar",
	}
	initDesignPalette(paletteConfig)			
	
	
	initNewTextBoxDialog()
	initNewCheckBoxDialog()
		
	// Initialize the page layout
	$('#layoutPage').layout({
		north: fixedUILayoutPaneParams(40),
		east: fixedUILayoutPaneParams(300),
		west: fixedUILayoutPaneParams(200),
		west__showOverflowOnHover:	true
	})	  
	  
	
	initCanvas(initContainerEditBehavior,initLayoutEditFieldInfo,initCanvasComplete)
	  
});
