

function initObjectCanvasSelectionBehavior(canvasElemSelector,selectionCallbackFunc)
{
	// jQuery UI selectable and draggable conflict with one another for click handling, so there is specialized
	// click handling for the selection and deselection of individual dashboard elements. When a click is made
	// on the canvas, all the other items/objects are deselected.
	$(canvasElemSelector).click(function(e) {
		console.log("click on dashboard canvas")
		
		// Unselect all the other divs within the canvas. This is done by 
		// removing the objectSelected CSS class.
	    $( canvasElemSelector + " > div" ).removeClass("objectSelected");
	
		// Toggle to the overall dashboard properties, hiding the other property panels
		selectionCallbackFunc()
	})
	
}


function selectObject(canvasElemSelector, objectElem)
{
	$( canvasElemSelector + " > div" ).removeClass("objectSelected");
	$(objectElem).addClass("objectSelected");
	
}


