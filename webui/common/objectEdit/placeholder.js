
function hideAllComponentLayoutPlaceholders() {
		$(".newComponentRowPlaceholder").remove()	
		$(".newComponentColumnPlaceholder").remove()
		$(".componentColumnPlacementPlaceholder").remove()	
}

function highlightDroppablePlaceholder(currMouseOffset) {
	
	var placeholderFound = false
	
	if(hitExistingPlaceholder(".newComponentRowPlaceholder",currMouseOffset)) {
		return
	}
	if(hitExistingPlaceholder(".newComponentColumnPlaceholder",currMouseOffset)) {
		return
	}
	if(hitExistingPlaceholder(".componentColumnPlacementPlaceholder",currMouseOffset)) {
		return
	}
	
	$(".componentRow").each(function() {
		
		if (placeholderFound) { return } // short-circuit loop if placeholder already found
		
		var $currRow = $(this)
		
		var rowRect = createElemRect($currRow)
		var newRowAboveDropZoneRect = outerTopInsetRect(rowRect,5,20)
		if(hitTestLayoutRect(currMouseOffset,newRowAboveDropZoneRect)) {
			console.log("highlighting new placeholder row: drop zone: " + JSON.stringify(newRowAboveDropZoneRect))
			if(!$currRow.prev().hasClass("newComponentRowPlaceholder")) {
				hideAllComponentLayoutPlaceholders()
				$('<div class="newComponentRowPlaceholder"></div>').insertBefore($currRow)
			}
			placeholderFound = true					
		}
		
		if (placeholderFound) { return }
		
		var newRowBelowDropZoneRect = outerBottomInsetRect(rowRect,5,20)
		if(hitTestLayoutRect(currMouseOffset,newRowBelowDropZoneRect)) {
			console.log("highlighting new placeholder row: drop zone: " + JSON.stringify(newRowBelowDropZoneRect))
			if(!$currRow.next().hasClass("newComponentRowPlaceholder")) {
				hideAllComponentLayoutPlaceholders()
				$('<div class="newComponentRowPlaceholder"></div>').insertAfter($currRow)
			}
			placeholderFound = true					
		}

		if (placeholderFound) { return }
		
		$currRow.children(".componentCol").each( function() {
			if (placeholderFound) { return } // short-circuit loop if placeholder already found
			
			var $currCol = $(this)
			var colRect = createElemRect($currCol)
			var newColRightDropZoneRect = outerRightInsetRect(colRect,5,20)
			if(hitTestLayoutRect(currMouseOffset,newColRightDropZoneRect)) {
				console.log("highlighting new placeholder col: drop zone: " + JSON.stringify(newColRightDropZoneRect))
				if(!$currCol.next().hasClass("newComponentColumnPlaceholder")) {
					hideAllComponentLayoutPlaceholders()
					var $colPlaceholder = $('<div class="newComponentColumnPlaceholder"></div>')
					$colPlaceholder.css("min-height",$currRow.outerHeight() + "px")					
					$colPlaceholder.insertAfter($currCol)
				}
				placeholderFound = true					
			}
			
			if (placeholderFound) { return }
			
			var newColLeftDropZoneRect = outerLeftInsetRect(colRect,5,20)
			if(hitTestLayoutRect(currMouseOffset,newColLeftDropZoneRect)) {
				console.log("highlighting new placeholder col: drop zone: " + JSON.stringify(newColLeftDropZoneRect))
				if(!$currCol.prev().hasClass("newComponentColumnPlaceholder")) {
					hideAllComponentLayoutPlaceholders()
					var $colPlaceholder = $('<div class="newComponentColumnPlaceholder"></div>')
					$colPlaceholder.css("min-height",$currRow.outerHeight() + "px")					
					$colPlaceholder.insertBefore($currCol)
				}
				placeholderFound = true					
			}

			if (placeholderFound) { return }
			
			$currCol.children(".layoutContainer").each(function() {
				if (placeholderFound) { return } // short-circuit loop if placeholder already found
				
				var $currComponent = $(this)
				var componentRect = createElemRect($currComponent)
				// Span the entire column for purposes of drag and drop
				componentRect.left = colRect.left
				componentRect.width = colRect.width
				var repositionAboveComponentDropZoneRect = outerTopInsetRect(componentRect,5,5)
				if(hitTestLayoutRect(currMouseOffset,repositionAboveComponentDropZoneRect)) {
					
					console.log("highlighting new component position within column: drop zone: " + JSON.stringify(repositionAboveComponentDropZoneRect))
					if(!$currComponent.prev().hasClass("componentColumnPlacementPlaceholder")) {
						hideAllComponentLayoutPlaceholders()
						$('<div class="componentColumnPlacementPlaceholder"></div>').insertBefore($currComponent)
					}
					placeholderFound = true					
				}
				
				if (placeholderFound) { return }
				
				var repositionBelowComponentDropZoneRect = outerBottomInsetRect(componentRect,5,5)
				if(hitTestLayoutRect(currMouseOffset,repositionBelowComponentDropZoneRect)) {
					console.log("highlighting new component position within column: drop zone: " + JSON.stringify(repositionBelowComponentDropZoneRect))
					if(!$currComponent.next().hasClass("componentColumnPlacementPlaceholder")) {
						hideAllComponentLayoutPlaceholders()
						$('<div class="componentColumnPlacementPlaceholder"></div>').insertAfter($currComponent)
					}
					placeholderFound = true					
				}
				
				
			}) // each component (layout container) within the column.

			
		}) // each column	
		
	}) // Each row
	if (!placeholderFound) {
		hideAllComponentLayoutPlaceholders()
	}
	
}
