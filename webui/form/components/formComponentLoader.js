


function loadFormComponents(loadFormConfig) {
	
	function createComponentRow() {
		var rowHTML = '<div class="componentRow">' +
		  '</div>'
		
		function saveUpdatedFormComponentLayout(parentComponentLayoutSelector, formID) {
			console.log("saveUpdatedFormComponentLayout: saving updated layout " 
				+ parentComponentLayoutSelector + ", form id = " + formID)
			
			componentRows = getComponentLayout(parentComponentLayoutSelector)
			
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

		function initTextBoxLayout($componentRow,textBox) {
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

		function initCheckBoxLayout($componentRow,checkBox) {
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
		
		function initDatePickerLayout($componentRow,datePicker) {
			// Create an HTML block for the container			
			var containerHTML = datePickerContainerHTML(datePicker.datePickerID);
			var containerObj = $(containerHTML)
			
			// Set the label to the field name
			var fieldName = getFieldRef(datePicker.properties.fieldID).name
			containerObj.find('label').text(fieldName)
			
			// Position the object withing the #layoutCanvas div
			$componentRow.append(containerObj)
			setElemDimensions(containerObj,datePicker.properties.geometry)
			
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(datePicker.datePickerID,datePicker)
			
			// Callback for any specific initialization for either the form design or view mode
			loadFormConfig.initDatePickerFunc(datePicker)
			
		}

		function initHtmlEditorLayout($componentRow,htmlEditor) {
			
			var containerHTML = htmlEditorContainerHTML(htmlEditor.htmlEditorID);
			var containerObj = $(containerHTML)
			
			// Set the label to the field name
			var fieldName = getFieldRef(htmlEditor.properties.fieldID).name
			containerObj.find('label').text(fieldName)
			
			// Position the object withing the #layoutCanvas div
			$componentRow.append(containerObj)
			setElemDimensions(containerObj,htmlEditor.properties.geometry)
			
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(htmlEditor.htmlEditorID,htmlEditor)
			
			// Callback for any specific initialization for either the form design or view mode
			loadFormConfig.initHtmlEditorFunc(htmlEditor)
		}
		
		function initImageEditorLayout($componentRow,image) {
			var containerHTML = imageContainerHTML(image.imageID);
			var containerObj = $(containerHTML)
			
			// Position the object withing the #layoutCanvas div
			$componentRow.append(containerObj)
			setElemDimensions(containerObj,image.properties.geometry)
	
	
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
	
		// Iterate through each type of component and populate a map/dictonary with 
		// the component ID and method to initialize the compnent. This map is then
		// referenced when populating the layout to put each component in the right place
		// on the layout.
		for (var textBoxIter in formInfo.textBoxes) {			
			var textBoxProps = formInfo.textBoxes[textBoxIter]			
			console.log("Form layout: text box component info=" + JSON.stringify(textBoxProps))
			compenentIDComponentMap[textBoxProps.textBoxID] = {
				componentInfo: textBoxProps,
				initFunc: initTextBoxLayout
			}			

		} // for each text box
	
		for (var checkBoxIter in formInfo.checkBoxes) {
			var checkBoxProps = formInfo.checkBoxes[checkBoxIter]
			console.log("loadFormComponents: initializing check box: " + JSON.stringify(checkBoxProps))
			compenentIDComponentMap[checkBoxProps.checkBoxID] = {
				componentInfo: checkBoxProps,
				initFunc: initCheckBoxLayout
			}		
		}
		
		for (var datePickerIter in formInfo.datePickers) {
			var datePickerProps = formInfo.datePickers[datePickerIter]
			console.log("loadFormComponents: initializing date picker: " + JSON.stringify(datePickerProps))
			compenentIDComponentMap[datePickerProps.datePickerID] = {
				componentInfo: datePickerProps,
				initFunc: initDatePickerLayout
			}			
			
		}

		for (var htmlEditorIter in formInfo.htmlEditors) {
			var htmlEditorProps = formInfo.htmlEditors[htmlEditorIter]
			console.log("loadFormComponents: initializing html editor: " + JSON.stringify(htmlEditorProps))

			compenentIDComponentMap[htmlEditorProps.htmlEditorID] = {
				componentInfo: htmlEditorProps,
				initFunc: initHtmlEditorLayout
			}			
		}
	
		for (var imageIter in formInfo.images) {
			var imageProps = formInfo.images[imageIter]
			console.log("loadFormComponents: initializing image editor: " + JSON.stringify(imageProps))

			compenentIDComponentMap[imageProps.imageID] = {
				componentInfo: imageProps,
				initFunc: initImageEditorLayout
			}			

		}
		
		var formLayout = formInfo.form.properties.layout
		populateFormLayout(formLayout,loadFormConfig.formParentElemID,compenentIDComponentMap)
				
		loadFormConfig.doneLoadingFormDataFunc()
	})
	
	
	
}