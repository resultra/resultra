

function initObjectCanvasSelectionBehavior(canvasElemSelector,selectionCallbackFunc)
{
	// jQuery UI selectable and draggable conflict with one another for click handling, so there is specialized
	// click handling for the selection and deselection of individual dashboard elements. When a click is made
	// on the canvas, all the other items/objects are deselected.
	$(canvasElemSelector).click(function(e) {
		console.log("Selection click on object canvas: " + canvasElemSelector)
		
		// Unselect all the other divs within the canvas. This is done by 
		// removing the objectSelected CSS class.
	    $(canvasElemSelector).find("div").removeClass("objectSelected");
	
		// Toggle to the overall dashboard properties, hiding the other property panels
		selectionCallbackFunc()
	})
	
}




// jQuery UI draggable and selectable functionality conflict with one another (using draggable masks
// the click behavior for selectable). So, the click handling to select an object/item needs to 
// be done directly.
function initObjectSelectionBehavior(objectElem, $parentCanvas,objectSelectedCallbackFunc) {

	function selectObject(objectElem)
	{
		$parentCanvas.find("div").removeClass("objectSelected");
		$(objectElem).addClass("objectSelected");
	
	}


	objectElem.click(function(e) {
		
		// This is important - if a click hits an object, then stop the propagation of the click
		// to the parent div(s), including the canvas itself. If the parent canvas
		// gets a click, it will deselect all the items (see initObjectCanvasSelectionBehavior)
		e.stopPropagation();
		
		var objectID = $(this).attr("id")
		console.log("Selection click on object: object ID = " + $(this).attr("id"))
		
		selectObject(this)
		
		// TODO - need to make object selection based upon the object itself.
		objectSelectedCallbackFunc(objectID)
	})
	
}

// Re-triggers the selection of any currently selected object. If the same layout is loaded with different data, this
// is used to re-initialize any selection-dependent behavior for the new data.
function reselectCurrentObjectSelection() {
	
	$(".objectSelected").each(function() {
		$(this).trigger("click")		
	})	
}
