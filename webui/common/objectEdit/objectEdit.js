

function initObjectGridEditBehavior(objID, editConfig) {
	
	console.log("Initialize object edit behavior: object ID = " + objID)
	var objSelector = "#"+objID
	
	// While in edit mode, disable input on the container
	$(objSelector).find('input').prop('disabled',true);
	
	$(objSelector).draggable ({
		cursor: "move",
		//TODO - Components can be dragged between different parents, but need to be contained 
		// within the same overall parent layout. 
//		clone: "original",
		helper:'clone',
		opacity: 0.5,
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
			handleDropOnComponentLayoutPlaceholder(mouseOffset)
		 }
	})

		
	$(objSelector).resizable({
		aspectRatio: false,
		handles: 'e', // Only allow resizing horizontally
		maxWidth: editConfig.resizeConstraints.maxWidth,
		minWidth: editConfig.resizeConstraints.minWidth,
		grid: 20, // snap to grid during resize
		stop: function(event, ui) {
			var objectID = event.target.id  
			// TODO - remot top and left from layout geometry since layout is now relative
			var resizeGeometry = {
				positionTop: 0,
				positionLeft: 0,
				sizeWidth: ui.size.width, 
				sizeHeight: ui.size.height }
			console.log("Object resize: component id = " + objectID + " geometry=" + JSON.stringify(resizeGeometry))
			editConfig.resizeFunc(objectID,resizeGeometry)
		} // stop function
	})
	
}