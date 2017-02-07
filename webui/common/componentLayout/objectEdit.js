

function initObjectGridEditBehavior($container, editConfig,layoutDesignConfig) {
		
	// While in edit mode, disable input on the container
	$container.find('input').prop('disabled',true);
	
	$container.draggable ({
		cursor: "move",
		helper:'clone',
		opacity: 0.5,
		drag: function(e,ui) {			
			var mouseOffset = { 
				top: e.pageY, 
				left: e.pageX }
			console.log("Drag component: mouse offset: " + JSON.stringify(mouseOffset))
			highlightDroppablePlaceholder(mouseOffset,layoutDesignConfig.parentLayoutSelector)
		}, // drag
		 stop: function( e, ui ) {
 			var mouseOffset = { 
 				top: e.pageY, 
 				left: e.pageX }
			var $draggedComponent = $(e.target)
			var dropCompleteFunc = function($component) {} // no-op
			handleDropOnComponentLayoutPlaceholder(mouseOffset,layoutDesignConfig,$draggedComponent,dropCompleteFunc)
		 }
	})

		
	$container.resizable({
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