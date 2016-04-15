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
			var reposParams = {
				parentID: parentID,
				objectID: event.target.id,
				position: {
				  positionTop: ui.position.top,			
				  positionLeft: ui.position.left
				}
			}
			console.log("Object reposition: params = " + JSON.stringify(reposParams))
			editConfig.reposFunc(reposParams)
		} // stop function
	})
	
	$(objSelector).resizable({
		aspectRatio: false,
		handles: 'e, w', // Only allow resizing horizontally
		maxWidth: editConfig.resizeConstraints.maxWidth,
		minWidth: editConfig.resizeConstraints.minWidth,
		grid: 20, // snap to grid during resize
		stop: function(event, ui) {  
			var resizeParams = {
				parentID: parentID,
				objectID: event.target.id,
				geometry: { 
					positionTop: ui.position.top,
					positionLeft: ui.position.left,
					sizeWidth: ui.size.width, 
					sizeHeight: ui.size.height }
			} 
			console.log("Object resize: params = " + JSON.stringify(resizeParams))
			editConfig.resizeFunc(resizeParams)
		} // stop function
	})
	
}