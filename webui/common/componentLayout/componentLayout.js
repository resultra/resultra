
// Reads the component layout from the DOM elements, using parentComponentLayoutSelector
// as the parent div of the layout elements.
function getComponentLayout(parentComponentLayoutSelector) {
	var componentRows = []
	$(parentComponentLayoutSelector).children('.componentRow').each(function() { 
		var rowComponents = []
		$(this).children('.layoutContainer').each(function() {
			var componentID = $(this).attr("id")
			rowComponents.push(componentID)
		})
		if (rowComponents.length > 0) {
			// Skip over empty/placeholder rows
			componentRows.push({componentIDs: rowComponents } )	
		}
	});
	
	return componentRows

}

// Create a new row in the layout and setup the drag-and-drop functionality to insert new
// components, or re-order existing components. Whenever the row changes, the saveLayoutFunc
// is called with the updated overall layout.
function createComponentRow(parentComponentLayoutSelector, saveLayoutFunc) {
	var rowHTML = '<div class="componentRow">' +
	  '</div>'
	
	
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
					var updatedLayout = getComponentLayout(parentComponentLayoutSelector)
					saveLayoutFunc(updatedLayout)
				}
		};
		
		paletteConfig.dropComplete(droppedObjInfo)
		
	}
	
	var $componentRow = $(rowHTML)
	$componentRow.sortable({
		placeholder: "ui-sortable-placeholder",
		forcePlaceholderSize: true,
		connectWith:".layoutContainer",
		start: function(event, ui) { 
			// The next line is a work-around for horizontal sorting.
			ui.placeholder.html('&nbsp;');
		},
		stop: function(event,ui) {
				
			var $droppedObj = ui.item
			
			if($droppedObj.hasClass("newComponent")) {
				console.log("Adding new component to row")
				receiveNewComponent($droppedObj)				
			} else {
				console.log("Re-order existing component in row")
				var updatedLayout = getComponentLayout(parentComponentLayoutSelector)
				saveLayoutFunc(updatedLayout)
			}
			
		}
	})
	
	return $componentRow
}



function populateComponentLayout(componentLayout, parentLayoutSelector, compenentIDComponentMap,saveLayoutFunc) {

	var completedLayoutComponentIDs = {}
	for(var rowIndex = 0; rowIndex < componentLayout.length; rowIndex++) {

		var currRowComponents = componentLayout[rowIndex].componentIDs
		var $componentRow = createComponentRow(parentLayoutSelector,saveLayoutFunc)
		$(parentLayoutSelector).append($componentRow)

		for(var componentIndex = 0; componentIndex<currRowComponents.length; componentIndex++) {
			var componentID = currRowComponents[componentIndex]
			console.log("Component layout: row=" + rowIndex + " component ID=" + componentID)
			// If the component has been deleted, then it won't be in the componentIDComponentMap.
			// In this case, skip initialiation for the deleted component.
			if(componentID in compenentIDComponentMap) {
				var initInfo = compenentIDComponentMap[componentID]
				console.log("Component layout: component info=" + JSON.stringify(initInfo.componentInfo))
				initInfo.initFunc($componentRow,initInfo.componentInfo)
				completedLayoutComponentIDs[componentID] = true			
			}
		}

	}

	// Layout any "orphans" which may are not, for whatever reason in the
	// list of rows and component IDs
	if(Object.keys(completedLayoutComponentIDs).length < Object.keys(compenentIDComponentMap).length) {
		console.log("populateComponentLayout: Layout orphan components")
		var $orphanLayoutRow = createComponentRow(parentLayoutSelector,saveLayoutFunc)
		$(parentLayoutSelector).append($orphanLayoutRow)
		for(var componentID in compenentIDComponentMap) {
			if(completedLayoutComponentIDs[componentID] != true) {
				var initInfo = compenentIDComponentMap[componentID]
				console.log("populateComponentLayout: Layout orphan component: " + componentID)
				initInfo.initFunc($orphanLayoutRow,initInfo.componentInfo)	
			}
		}	
	}

	var $placeholderRowForDrop = createComponentRow(parentLayoutSelector,saveLayoutFunc)
	$(parentLayoutSelector).append($placeholderRowForDrop)

}
