

function loadFormComponents(loadFormConfig) {
	
	function createComponentRow() {
		var rowHTML = '<div class="componentRow">' +
		  '</div>'
		
		function saveUpdatedFormComponentLayout(parentComponentLayoutSelector, formID) {
			console.log("saveUpdatedFormComponentLayout: saving updated layout " 
				+ parentComponentLayoutSelector + ", form id = " + formID)
			
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
			
			console.log("saveUpdatedFormComponentLayout: component layout = " + JSON.stringify(componentRows))
			
			var setLayoutParams = {
				formID: loadFormConfig.formID,
				layout: componentRows
			}
			jsonAPIRequest("frm/setLayout", setLayoutParams, function(formInfo) {
			})
		}		
		
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
						saveUpdatedFormComponentLayout(loadFormConfig.formParentElemID,loadFormConfig.formID)
					}
			};
			
			paletteConfig.dropComplete(droppedObjInfo)
			
		}
		
		var $componentRow = $(rowHTML)
		$componentRow.sortable({
			revert:true,
			placeholder: "ui-sortable-placeholder",
			forcePlaceholderSize: true,
			connectWith:".layoutContainer",
			axis: 'x',
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
					saveUpdatedFormComponentLayout(loadFormConfig.formParentElemID,loadFormConfig.formID)
				}
				
			}
		})
		
		return $componentRow
	}
	
	// TODO: formID is accessing a global. Pass it into this function instead.
	
	jsonAPIRequest("frm/getFormInfo", { formID: formID }, function(formInfo) {
													
		var compenentIDComponentMap = {}	

		function initTextBox($componentRow,textBox) {
			// Create an HTML block for the container
			console.log("loadFormComponents: initializing text box: " + JSON.stringify(textBox))
		
			var containerHTML = textBoxContainerHTML(textBox.textBoxID);
			var containerObj = $(containerHTML)
		
			// Set the label to the field name
			var fieldName = getFieldRef(textBox.properties.fieldID).name
			containerObj.find('label').text(fieldName)
		
			$componentRow.append(containerObj)
			
			setElemDimensions(containerObj,textBox.properties.geometry)
		
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(textBox.textBoxID,textBox)
		
			// Callback for any specific initialization for either the form design or view mode
			loadFormConfig.initTextBoxFunc(textBox)
			
		}

		function initCheckBox($componentRow,checkBox) {
			// Create an HTML block for the container
			
			var containerHTML = checkBoxContainerHTML(checkBox.checkBoxID);
			var containerObj = $(containerHTML)
			
			// Set the label to the field name. A span element is used, since
			// the checkbox itself is nested inside a label.
			var fieldName = getFieldRef(checkBox.properties.fieldID).name
			containerObj.find('span').text(fieldName)
			
			// Position the object withing the #layoutCanvas div
			$componentRow.append(containerObj)
			setElemDimensions(containerObj,checkBox.properties.geometry)
			
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(checkBox.checkBoxID,checkBox)
			
			// Callback for any specific initialization for either the form design or view mode 
			loadFormConfig.initCheckBoxFunc(checkBox)
		}
		
		function populateFormLayout(formLayout, parentLayoutSelector, compenentIDComponentMap) {
	
			var completedLayoutComponentIDs = {}
			for(var rowIndex = 0; rowIndex < formLayout.length; rowIndex++) {
	
				var currRowComponents = formLayout[rowIndex].componentIDs
				var $componentRow = createComponentRow()
				$(parentLayoutSelector).append($componentRow)
	
				for(var componentIndex = 0; componentIndex<currRowComponents.length; componentIndex++) {
					var componentID = currRowComponents[componentIndex]
					console.log("Form layout: row=" + rowIndex + " component ID=" + componentID)
					var initInfo = compenentIDComponentMap[componentID]
					console.log("Form layout: component info=" + JSON.stringify(initInfo.componentInfo))
					initInfo.initFunc($componentRow,initInfo.componentInfo)
					completedLayoutComponentIDs[componentID] = true
				}
	
			}
	
			// Layout any "orphans" which may are not, for whatever reason in the
			// list of rows and component IDs
			if(Object.keys(completedLayoutComponentIDs).length < Object.keys(compenentIDComponentMap).length) {
				console.log("populateFormLayout: Layout orphan components")
				var $orphanLayoutRow = createComponentRow()
				$(parentLayoutSelector).append($orphanLayoutRow)
				for(var componentID in compenentIDComponentMap) {
					if(completedLayoutComponentIDs[componentID] != true) {
						var initInfo = compenentIDComponentMap[componentID]
						console.log("populateFormLayout: Layout orphan component: " + componentID)
						initInfo.initFunc($orphanLayoutRow,initInfo.componentInfo)	
					}
				}	
			}

			var $placeholderRowForDrop = createComponentRow()
			$(parentLayoutSelector).append($placeholderRowForDrop)
	
		}
	
		for (var textBoxIter in formInfo.textBoxes) {
			
			var textBoxProps = formInfo.textBoxes[textBoxIter]
					
			console.log("Form layout: text box component info=" + JSON.stringify(textBoxProps))
			compenentIDComponentMap[textBoxProps.textBoxID] = {
				componentInfo: textBoxProps,
				initFunc: initTextBox
			}			

		} // for each text box
	
		for (var checkBoxIter in formInfo.checkBoxes) {
			var checkBoxProps = formInfo.checkBoxes[checkBoxIter]
			console.log("loadFormComponents: initializing check box: " + JSON.stringify(checkBoxProps))
			compenentIDComponentMap[checkBoxProps.checkBoxID] = {
				componentInfo: checkBoxProps,
				initFunc: initCheckBox
			}			
			
		}
		
		var formLayout = formInfo.form.properties.layout
		populateFormLayout(formLayout,loadFormConfig.formParentElemID,compenentIDComponentMap)
		
		
		
/*			
		for (var datePickerIter in formInfo.datePickers) {
			
			// Create an HTML block for the container
			var datePicker = formInfo.datePickers[datePickerIter]
			console.log("loadFormComponents: initializing date picker: " + JSON.stringify(datePicker))
			
			var containerHTML = datePickerContainerHTML(datePicker.datePickerID);
			var containerObj = $(containerHTML)
			
			// Set the label to the field name
			var fieldName = getFieldRef(datePicker.properties.fieldID).name
			containerObj.find('label').text(fieldName)
			
			// Position the object withing the #layoutCanvas div
			$(loadFormConfig.formParentElemID).append(containerObj)
			setElemGeometry(containerObj,datePicker.properties.geometry)
			
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(datePicker.datePickerID,datePicker)
			
			// Callback for any specific initialization for either the form design or view mode
			loadFormConfig.initDatePickerFunc(datePicker)
			

		} // for each date picker


		for (var htmlEditorIter in formInfo.htmlEditors) {
			
			// Create an HTML block for the container
			var htmlEditor = formInfo.htmlEditors[htmlEditorIter]
			console.log("loadFormComponents: initializing html editor: " + JSON.stringify(htmlEditor))
			
			var containerHTML = htmlEditorContainerHTML(htmlEditor.htmlEditorID);
			var containerObj = $(containerHTML)
			
			// Set the label to the field name
			var fieldName = getFieldRef(htmlEditor.properties.fieldID).name
			containerObj.find('label').text(fieldName)
			
			// Position the object withing the #layoutCanvas div
			$(loadFormConfig.formParentElemID).append(containerObj)
			setElemGeometry(containerObj,htmlEditor.properties.geometry)
			
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(htmlEditor.htmlEditorID,htmlEditor)
			
			// Callback for any specific initialization for either the form design or view mode
			loadFormConfig.initHtmlEditorFunc(htmlEditor)
			

		} // for each html editor

		for (var imageIter in formInfo.images) {
			
			// Create an HTML block for the container
			var image = formInfo.images[imageIter]
			console.log("loadFormComponents: initializing image editor: " + JSON.stringify(image))
			
			var containerHTML = imageContainerHTML(image.imageID);
			var containerObj = $(containerHTML)
			
			
			// Position the object withing the #layoutCanvas div
			$(loadFormConfig.formParentElemID).append(containerObj)
			setElemGeometry(containerObj,image.properties.geometry)
	
	
			// Set the label to the field name
			var labelID = imageContainerLabelIDFromContainerElemID(image.imageID)
			console.log("loadFormComponents: initializing label: id=" + labelID)
			var fieldName = getFieldRef(image.properties.fieldID).name
			$('#'+labelID).text(fieldName)
			
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(image.imageID,image)
			
			// Callback for any specific initialization for either the form design or view mode
			loadFormConfig.initImageFunc(image)
			

		} // for each html editor
	*/	
		
		loadFormConfig.doneLoadingFormDataFunc()
	})
	
	
	
}