
function pruneComponentLayoutEmptyColsAndRows($parentLayout) {
	
	// Prune empty columns
	$parentLayout.find('.componentCol').each(function() {
		var colComponentCount = $(this).children().length
		if (colComponentCount <= 0) {
			$(this).remove()
		}
	})
	
	// Prune empty rows
	$parentLayout.find('.componentRow').each(function() {
		var rowColCount = $(this).children().length
		if (rowColCount <= 0) {
			$(this).remove()
		}
	})
	
}


// Reads the component layout from the DOM elements, using parentComponentLayoutSelector
// as the parent div of the layout elements.
function getComponentLayout($parentLayout) {
	var componentRows = []
	$parentLayout.children('.componentRow').each(function() { 
		
		var currRowCols = []
		
		$(this).children('.componentCol').each(function() {
			var colComponents = []
			$(this).children('.layoutContainer').each(function() {
				var $layoutContainer = $(this)
				var componentID = getContainerComponentID($layoutContainer)		
				colComponents.push(componentID)
			})
			if (colComponents.length > 0) {
				// Skip over empty/placeholder rows
				currRowCols.push({componentIDs: colComponents } )	
			}
			
		})
		
		if(currRowCols.length > 0) {
			componentRows.push({columns:currRowCols})
		}	
	});
	
	return componentRows

}

function createComponentRowNoInsert($parentLayout) {
	var rowHTML = '<div class="componentRow"></div>'
	var $componentRow = $(rowHTML)
	
	return $componentRow
}

// Create a new row in the layout.
function createComponentRow($parentLayout) {
	var $componentRow = createComponentRowNoInsert($parentLayout)	
	
	$parentLayout.append($componentRow)
	
	return $componentRow	
}

function createComponentColNoInsert() {
	var colHTML = '<div class="componentCol"></div>'
	var $componentCol = $(colHTML)
	
	return $componentCol
}

function createComponentCol($parentLayout,$parentRow) {
	
	var $componentCol = createComponentColNoInsert()
			
	$parentRow.append($componentCol)
	
	return $componentCol
}



function populateComponentLayout(componentLayout, parentLayoutSelector, compenentIDComponentMap) {

	var completedLayoutComponentIDs = {}
	
	var $parentLayout = $(parentLayoutSelector)
	
	for(var rowIndex = 0; rowIndex < componentLayout.length; rowIndex++) {
		
		var currRow = componentLayout[rowIndex]
		
		var $componentRow = createComponentRow($parentLayout)
		
		if(currRow.columns !== null) {
			for (var colIndex = 0; colIndex < currRow.columns.length; colIndex++) {
			
				var currCol = currRow.columns[colIndex]
				var currColComponents = currCol.componentIDs
			
				var $componentCol = createComponentCol($parentLayout,$componentRow)
				
				for(var componentIndex = 0; componentIndex<currColComponents.length; componentIndex++) {
					var componentID = currColComponents[componentIndex]
					console.log("Component layout: row=" + rowIndex + " col=" + colIndex +
						" component ID=" + componentID)
					// If the component has been deleted, then it won't be in the componentIDComponentMap.
					// In this case, skip initialiation for the deleted component.
					if(componentID in compenentIDComponentMap) {
						var initInfo = compenentIDComponentMap[componentID]
						console.log("Component layout: component info=" + JSON.stringify(initInfo.componentInfo))
						initInfo.initFunc($componentCol,initInfo.componentInfo)
						completedLayoutComponentIDs[componentID] = true			
					}
				} // for each component
				
							
			} // for each column			
			
		} // columns !== null


	} // for each row

	// Layout any "orphans" which may are not, for whatever reason in the
	// list of rows/cols and component IDs
	if(Object.keys(completedLayoutComponentIDs).length < Object.keys(compenentIDComponentMap).length) {
		console.log("populateComponentLayout: Layout orphan components")
		
		var $orphanLayoutRow = createComponentRow($parentLayout)
		var $orphanLayoutCol = createComponentCol($parentLayout,$orphanLayoutRow)
		
		for(var componentID in compenentIDComponentMap) {
			if(completedLayoutComponentIDs[componentID] != true) {
				var initInfo = compenentIDComponentMap[componentID]
				console.log("populateComponentLayout: Layout orphan component: " + componentID)
				initInfo.initFunc($orphanLayoutCol,initInfo.componentInfo)	
			}
		}	
	}


}
