function initObjectEditBehavior(parentID, objID, editConfig) {
	
	console.log("Initialize object edit behavior: parent ID = " + parentID + " object ID = " + objID)
	var objSelector = "#"+objID
	
	// While in edit mode, disable input on the container
	$(objSelector).find('input').prop('disabled',true);

	$(objSelector).draggable ({
		grid: [20, 20], // snap to a grid
		cursor: "move",
		containment: "parent",
		clone: "original",				
		stop: function(event, ui) {
			var objectID = event.target.id
			var position = { top: ui.position.top,			
				left: ui.position.left }
			console.log("Object reposition: params: component id = " + objectID + " position=" + JSON.stringify(position))
			editConfig.repositionFunc(objectID,position)
		} // stop function
	})
	
	$(objSelector).resizable({
		aspectRatio: false,
		handles: 'e, w', // Only allow resizing horizontally
		maxWidth: editConfig.resizeConstraints.maxWidth,
		minWidth: editConfig.resizeConstraints.minWidth,
		grid: 20, // snap to grid during resize
		stop: function(event, ui) {
			var objectID = event.target.id  
			var resizeGeometry = {
				positionTop: ui.position.top,
				positionLeft: ui.position.left,
				sizeWidth: ui.size.width, 
				sizeHeight: ui.size.height }
			console.log("Object resize: component id = " + objectID + " geometry=" + JSON.stringify(resizeGeometry))
			editConfig.resizeFunc(objectID,resizeGeometry)
		} // stop function
	})
	
}