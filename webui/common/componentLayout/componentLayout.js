
// Reads the component layout from the DOM elements, using parentComponentLayoutSelector
// as the parent div of the layout elements.
function getComponentLayout($parentLayout) {
	var componentRows = []
	$parentLayout.children('.componentRow').each(function() { 
		
		var currRowCols = []
		
		$(this).children('.componentCol').each(function() {
			var colComponents = []
			$(this).children('.layoutContainer').each(function() {
				var componentID = $(this).attr("id")
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


// Create a new row in the layout and setup the drag-and-drop functionality to insert new
// components, or re-order existing components. Whenever the row changes, the saveLayoutFunc
// is called with the updated overall layout.
function createComponentRow($parentLayout) {
	var rowHTML = '<div class="componentRow"></div>'
	var $componentRow = $(rowHTML)	
	
	$parentLayout.append($componentRow)
	
	return $componentRow	
}

function createComponentCol($parentLayout,$parentRow, saveLayoutFunc) {
	
	var colHTML = '<div class="componentCol"></div>'
	var $componentCol = $(colHTML)
	
	function receiveNewComponent($droppedObj) {
		
		$droppedObj.removeClass("newComponent")
		
		var placeholderID = $droppedObj.attr('id')
		assert(placeholderID !== undefined, "receiveNewComponent: palette item missing element id")
		console.log("receiveNewComponent: drop: placeholder ID of palette item: " + placeholderID)
	
		var objWidth = $droppedObj.width()
		var objHeight = $droppedObj.height()
	
		var paletteItemID = $droppedObj.data("paletteItemID")
		console.log("receiveNewComponent: drop: palette item ID/type: " + paletteItemID)
	
		var paletteConfig = $droppedObj.data("paletteConfig")
		
		var componentParentLayoutSelector = paletteConfig.dropDestSelector
		
		var droppedObjInfo = {
			droppedElem: $droppedObj,
			paletteItemID: paletteItemID,
			placeholderID: placeholderID,
			geometry: {positionTop: 0, positionLeft: 0,
			sizeWidth: objWidth,sizeHeight: objHeight},
			finalizeLayoutIncludingNewComponentFunc: function() {
					console.log("receiveNewComponent: finalizing layout with new component")
					var updatedLayout = getComponentLayout($parentLayout)
					saveLayoutFunc(updatedLayout)
				}
		};
		
		paletteConfig.dropComplete(droppedObjInfo)
		
	}
	
	/*
	$componentCol.sortable({
		placeholder: "component-placeholder",
		connectWith:".layoutContainer",
		stop: function(event,ui) {
				
			var $droppedObj = ui.item
			
			if($droppedObj.hasClass("newComponent")) {
				console.log("Adding new component to existing column")
				receiveNewComponent($droppedObj)				
			} else {
				console.log("Re-order existing component in existing column")
				var updatedLayout = getComponentLayout($parentLayout)
				saveLayoutFunc(updatedLayout)
			}
			
		}
	})
*/
		
	$parentRow.append($componentCol)
	
	return $componentCol
}



function populateComponentLayout(componentLayout, parentLayoutSelector, compenentIDComponentMap,saveLayoutFunc) {

	var completedLayoutComponentIDs = {}
	
	$parentLayout = $(parentLayoutSelector)
	
	for(var rowIndex = 0; rowIndex < componentLayout.length; rowIndex++) {
		
		var currRow = componentLayout[rowIndex]
		
		var $componentRow = createComponentRow($parentLayout)
		
		if(currRow.columns !== null) {
			for (var colIndex = 0; colIndex < currRow.columns.length; colIndex++) {
			
				var currCol = currRow.columns[colIndex]
				var currColComponents = currCol.componentIDs
			
				var $componentCol = createComponentCol($parentLayout,$componentRow,saveLayoutFunc)
				
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
		var $orphanLayoutCol = createComponentCol($parentLayout,$orphanLayoutRow,saveLayoutFunc)
		
		for(var componentID in compenentIDComponentMap) {
			if(completedLayoutComponentIDs[componentID] != true) {
				var initInfo = compenentIDComponentMap[componentID]
				console.log("populateComponentLayout: Layout orphan component: " + componentID)
				initInfo.initFunc($orphanLayoutCol,initInfo.componentInfo)	
			}
		}	
	}


}
