

function initObjectGridEditBehavior($container, editConfig,layoutDesignConfig) {
		
	// While in edit mode, disable input on the container
	$container.find('input').prop('disabled',true);
	
	var resizeHandles = 'e' // By default only resize horizontally
	if (editConfig.hasOwnProperty('resizeHandles')) {
		resizeHandles = editConfig.resizeHandles
	}
	
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
		handles: resizeHandles,
		maxWidth: editConfig.resizeConstraints.maxWidth,
		minWidth: editConfig.resizeConstraints.minWidth,
		grid: 5, // snap to grid during resize
		stop: function(event, ui) {
			var resizeGeometry = {
				positionTop: 0,
				positionLeft: 0,
				sizeWidth: ui.size.width, 
				sizeHeight: ui.size.height }
			editConfig.resizeFunc($container,resizeGeometry)
		} // stop function
	})
	
}

function disableObjectEditBehavior($container) {
	$container.resizable('destroy')
	$container.draggable('destroy')
}