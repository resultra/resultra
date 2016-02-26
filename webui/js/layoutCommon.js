

function fieldContainerHTML(id)
{
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer layoutField draggable resizable" id="'+id+'">' +
			'<div class="field">'+
				'<label>New Field</label>'+
				'<input type="text" name="symbol" class="layoutInput" placeholder="Enter">'+
			'</div>'+
		'</div>';
	return containerHTML
}

// Map of field ID's to the fieldRef object: see initialization below
var fieldsByID = {}

function initCanvas(containerInitCallback, fieldInitCallback, initCompleteCallback) {
	var jsonReqData = jsonAPIRequest("getLayoutEditInfo", {
			layoutID: layoutID
		},
		function(replyData) {
			// Populate the selection boxes used in the dialogs to create new
			// text boxes.
			var textFields = replyData.fieldsByType.textFields
			for (textFieldIter in textFields) {
				console.log("Text field: " + textFields[textFieldIter].fieldInfo.name)

				// Populate a map/dictionary of field IDs to the field information.
				// This is needed when creating new layout elements (text boxes, etc.),
				// so the fields information can be used after creation of the layout
				// element.
				fieldsByID[textFields[textFieldIter].fieldID] = textFields[textFieldIter].fieldInfo

				// Done with base initialization. Invoke the callback to finish any
				// specialized initialization for the client of this function.
				fieldInitCallback(textFields[textFieldIter])

			} // for each text field
			
			var numberFields = replyData.fieldsByType.numberFields
			for (numberFieldIter in numberFields) {
				console.log("Number field: " + numberFields[numberFieldIter].fieldInfo.name)

				var numberField = numberFields[numberFieldIter]
				// Populate a map/dictionary of field IDs to the field information.
				// This is needed when creating new layout elements (text boxes, etc.),
				// so the fields information can be used after creation of the layout
				// element.
				fieldsByID[numberField.fieldID] = numberField.fieldInfo

				// Done with base initialization. Invoke the callback to finish any
				// specialized initialization for the client of this function.
				fieldInitCallback(numberField)

			} // for each number field

			for (containerIter in replyData.layoutContainers) {
				
				// Create an HTML block for the container
				container = replyData.layoutContainers[containerIter]
				console.log("initializing container: id=" + JSON.stringify(container))
				var containerHTML = fieldContainerHTML(container.containerID);
				var containerObj = $(containerHTML)

				// Update the label to match the name of the field referenced by
				// the container.
				//
				// NOTE: There is a dependencey upon 'fieldsByID' to initialize the label
				//
				// TODO - 'fieldsByID' should be refactored into a more robust object
				// which supports queries for fieldID and does error checking.
				containerObj.find('label').text(fieldsByID[container.fieldID].name)
				
				// Store the field IDs and types for the container in the associated jQuery
				// object itself. This is needed for validation and to make the right API 
				// call when setting values.
				containerObj.data("fieldID",container.fieldID)
				containerObj.data("fieldType",fieldsByID[container.fieldID].type)
				containerObj.data("isCalcField",fieldsByID[container.fieldID].isCalcField)

				// Position the object withing the #layoutCanvas div
				$("#layoutCanvas").append(containerObj)
				var geometry = container.geometry
				containerObj.css({
					top: geometry.positionTop,
					left: geometry.positionLeft,
					width: geometry.sizeWidth,
					height: geometry.sizeHeight,
					position: "absolute"
				});

				// Done with base initialization. Invoke the callback to finish any
				// specialized initialization for the client of this function.
				containerInitCallback(containerObj)

			} // for each container
			
			initCompleteCallback();

		})
}
