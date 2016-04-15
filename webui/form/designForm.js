

function initLayoutEditFieldInfo(fieldRef)
{
	// Add the ability to select the field from the new layout container
	// dialog. This callback function is called for both number and 
	// text fields; so, text boxes can be used for both numbers and 
	// regular text.
 	var selectFieldOptionHTML = '<option value="' +
 		fieldRef.fieldID + '">' +
 		fieldRef.fieldInfo.name + '</option>'
 	$("#textBoxFieldSelection").append(selectFieldOptionHTML)	
}

function initCanvasComplete() {	
	
} // noop



function initContainerEditBehavior(container)
{

	// While in edit mode, disable input on the container
	container.find('input').prop('disabled',true);


	// TODO - This could be put into a common function, since these
	// properties should be the same for objects loaded with the page
	// and newly added objects.
	container.draggable ({
		grid: [20, 20], // snap to a grid
		cursor: "move",
		containment: "parent",
		clone: "original",				
		stop: function(event, ui) {
				  var layoutPos = {
					  positionLeft: ui.position.left,
					  positionTop: ui.position.top,
				   };
		  
				  console.log("drag stop: id: " + event.target.id);
				  console.log("drag stop: new position: " + JSON.stringify(layoutPos));
				  
				  // TODO: send ajax request to reposition the container
  
		      } // stop function
	})
	container.resizable({
		aspectRatio: false,
		handles: 'e, w', // Only allow resizing horizontally
		maxWidth: 600,
		minWidth: 100,
		grid: 20, // snap to grid during resize

		stop: function(event, ui) {
			  
			var resizeParams = {
				formID: layoutID,
				formItemID: event.target.id,
				geometry: { 
					positionTop: ui.position.top,
					positionLeft: ui.position.left,
					sizeWidth: ui.size.width, 
					sizeHeight: ui.size.height }
			} 
			console.log("Form Item resize: params = " + JSON.stringify(resizeParams))
			
				// TODO call resize function for item using the configuration above.
				  
	  			var layoutContainerParams = {
	  				parentLayoutID: layoutID,containerID: event.target.id,
	  				geometry: { positionTop: ui.position.top,positionLeft: ui.position.left,
						sizeWidth: ui.size.width, sizeHeight: ui.size.height }
	  			};
		
			 	jsonAPIRequest("resizeLayoutContainer",layoutContainerParams,function(replyData) {
			 		console.log("Done resizing text box")
			 	})	
				  
			  } // stop function
	});
	
}


var paletteItemsEditConfig = {
	paletteItemTextBox: textBoxEditConfig,
	paletteItemCheckBox: checkBoxEditConfig
}

function loadFormEditInfo()
{
	jsonAPIRequest("getLayoutEditInfo", { layoutID: layoutID }, function(layoutEditInfo) {
		
		var fieldRefsByID = {}		
		var textFields = layoutEditInfo.fieldsByType.textFields
		for (textFieldIter in textFields) {
			console.log("Text field: " + textFields[textFieldIter].fieldInfo.name)
			var textField = textFields[textFieldIter]
			fieldRefsByID[textField.fieldID] = textFields
		} // for each text field
		
		var numberFields = layoutEditInfo.fieldsByType.numberFields
		for (numberFieldIter in numberFields) {
			console.log("Number field: " + numberFields[numberFieldIter].fieldInfo.name)
			var numberField = numberFields[numberFieldIter]
			fieldRefsByID[numberField.fieldID] = numberField

		} // for each number field
		
		function initObjectFieldInfo(elemObj,fieldRef) {	
			console.log("Initialize field info for form object: " + JSON.stringify(fieldRef))
			elemObj.data("fieldID",fieldRef.fieldID)
			elemObj.data("fieldType",fieldRef.fieldInfo.type)
			elemObj.data("isCalcField",fieldRef.fieldInfo.isCalcField)			
		}		
		
		for (textBoxIter in layoutEditInfo.layoutContainers) {
			
			// Create an HTML block for the container
			textBox = layoutEditInfo.layoutContainers[textBoxIter]
			console.log("initializing text box: id=" + JSON.stringify(textBox))
			// TODO - textBoxContainerHTMl is specific to text boxes only. Need to use a callback
			// to create the right HTML for the containers.
			var containerHTML = textBoxContainerHTML(textBox.containerID);
			var containerObj = $(containerHTML)

			// Update the label to match the name of the field referenced by
			// the container.
			//
			// NOTE: There is a dependencey upon 'fieldsByID' to initialize the label
			//
			
			var fieldRef = fieldRefsByID[textBox.fieldID]
			
			// TODO - 'fieldsByID' should be refactored into a more robust object
			// which supports queries for fieldID and does error checking.
			containerObj.find('label').text(fieldRef.FieldInfo.name)
			
			// Store the field IDs and types for the container in the associated jQuery
			// object itself. This is needed for validation and to make the right API 			
			initObjectFieldInfo(containerObj,fieldRef)

			// Position the object withing the #layoutCanvas div
			$("#layoutCanvas").append(containerObj)
			setElemGeometry(containerObj,textBox.geometry)

			// Done with base initialization. Invoke the callback to finish any
			// specialized initialization for the client of this function.
//			containerInitCallback(containerObj)

		} // for each container
	})
}

$(document).ready(function() {
					
	var paletteConfig = {
		draggableItemHTML: function(placeholderID,paletteItemID) {
			return paletteItemsEditConfig[paletteItemID].draggableHTMLFunc(placeholderID)
		},
		
		dropComplete: function(droppedItemInfo) {
			console.log("designForm: Dashboard design pallete: drop item: " + JSON.stringify(droppedItemInfo))			
			
			// After the drag operation is complete, the draggable and resizable
			// properties need to be initialized.
			//
			// At this point, the placholder div for the bar chart will have just been inserted. However, the DOM may 
			// not be completely updated at this point. To ensure this, a small delay is needed before
			// drawing the dummy bar charts. See http://goo.gl/IloNM for more.
			var objEditConfig = paletteItemsEditConfig[droppedItemInfo.paletteItemID]
			
			setTimeout(function() {
				initObjectEditBehavior(layoutID,droppedItemInfo.placeholderID,objEditConfig) 
			}, 50);
			
			
			// "repackage" the dropped item paramaters for creating a new layout element. Also add the layoutID
			// to the parameters.
			var layoutContainerParams = {
				parentLayoutID: layoutID,
				geometry: droppedItemInfo.geometry,
				containerID: droppedItemInfo.placeholderID,
				};
				
			objEditConfig.createNewItemAfterDropFunc(layoutContainerParams,fieldsByID)
		},
		
		dropDestSelector: "#layoutCanvas",
		paletteSelector: "#gallerySidebar",
	}
	initDesignPalette(paletteConfig)			
	
	
	initNewTextBoxDialog()
	initNewCheckBoxDialog()
		
	// Initialize the page layout
	$('#layoutPage').layout({
		north: fixedUILayoutPaneParams(40),
		east: fixedUILayoutPaneParams(300),
		west: fixedUILayoutPaneParams(200),
		west__showOverflowOnHover:	true
	})	  
	  
	
	initCanvas(initContainerEditBehavior,initLayoutEditFieldInfo,initCanvasComplete)
	
	// loadFormEditInfo is WIP - need to finish refactoring the text box code on the server (specifically to enclose a
	// full-blown field reference in each text box rather than require a cross-reference). This will simplify initialization
	// and editing of text boxes and serve as a model for other form elements.
	// loadFormEditInfo(); 
});
