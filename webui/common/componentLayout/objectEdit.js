// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


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

	var horizGridSize = 10
	var vertGridSize = 5
		
	$container.resizable({
		aspectRatio: false,
		handles: resizeHandles,
		maxWidth: editConfig.resizeConstraints.maxWidth,
		minWidth: editConfig.resizeConstraints.minWidth,
		maxHeight: editConfig.resizeConstraints.maxHeight,
		minHeight: editConfig.resizeConstraints.minHeight,
		grid: [ horizGridSize, vertGridSize ], // snap to grid during resize
		resize: function(event,ui) {
			
			// If there is a resizeInProgressFunc, the "unsnapped" ui size is used.
			if(editConfig.resizeInProgressFunc !== undefined) {
				// With the resize in 
				var resizeGeometry = {
					positionTop: 0,
					positionLeft: 0,
					sizeWidth: ui.size.width, 
					sizeHeight: ui.size.height 
				}
				editConfig.resizeInProgressFunc($container,resizeGeometry)	
			}
			
			
			
			// For some reason the resizable plugin doesn't snap to the grid
			// based upon the grid size. The logic below will snap the width
			// to the nearest width which is an exact multiple of the gridSize.
			//
			// TODO - There's a conflict in the resize logic when the height
			// is also snapped. In the case the container has 'auto' for height,
			// or more specifically the 's' or 'se' resize handles are disabled,
			// we don't want any resizing to take place.  
			var snapWidth = Math.round(ui.size.width/horizGridSize)*horizGridSize
//			var snapHeight = Math.round(ui.size.height/gridSize)*gridSize
			ui.size.width = snapWidth
//			ui.size.height = snapHeight
			
			
		},
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