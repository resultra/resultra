function getFormComponentContext(formContext, contextLoadCompleteCallback) {
	var contextPartsRemaining = 4;
	var context = {}
	
	context.formID = formContext.formID
	context.databaseID = formContext.databaseID
	context.tableID = formContext.tableID
	
	function completeOneContextPart() {
		contextPartsRemaining -= 1
		if(contextPartsRemaining <= 0) {
			contextLoadCompleteCallback(context)
		}
	}
	
	loadFieldInfo(context.tableID, [fieldTypeAll],function(fieldsByID) {
		context.fieldsByID = fieldsByID
		completeOneContextPart()
	})
	
	initFieldInfo( function () {
		completeOneContextPart()
	})
	
	getGlobalInfoIndexedByID(context.databaseID,function(globalsByID) {
		context.globalsByID = globalsByID
		completeOneContextPart()
	})
	
	jsonAPIRequest("frm/getFormInfo", { formID: context.formID }, function(formInfo) {
		context.formInfo = formInfo
		completeOneContextPart()
	})
	
}


function loadFormComponents(loadFormConfig) {
	
	getFormComponentContext(loadFormConfig.formContext, function(componentContext) {
													
		var compenentIDComponentMap = {}	

		function initHeaderLayout($componentRow,header) {
			// Create an HTML block for the container
			console.log("loadFormComponents: initializing header: " + JSON.stringify(header))
		
			var containerHTML = formHeaderContainerHTML(header.headerID);
			var $containerObj = $(containerHTML)
			$containerObj.find(".formHeader").text(header.properties.label)
					
			$componentRow.append($containerObj)
			
			setElemDimensions($containerObj,header.properties.geometry)
		
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(header.headerID,header)
		
			// Callback for any specific initialization for either the form design or view mode
			loadFormConfig.initHeaderFunc(componentContext,header)
			
		}


		function initTextBoxLayout($componentRow,textBox) {
			// Create an HTML block for the container
			console.log("loadFormComponents: initializing text box: " + JSON.stringify(textBox))
		
			var containerHTML = textBoxContainerHTML(textBox.textBoxID);
			var containerObj = $(containerHTML)
			
			function setTextBoxLabel($textBoxContainer,label) {
				$textBoxContainer.find('label').text(label)
			}
		
			var componentLink = textBox.properties.componentLink
		
			if(componentLink.linkedValType == linkedComponentValTypeField) {
				// text box is linked to a field type
				// Set the label to the field name
				var fieldName = getFieldRef(componentLink.fieldID).name
				setTextBoxLabel(containerObj,fieldName)
			} else {
				// text box is linked to a global
				var globalInfo = componentContext.globalsByID[componentLink.globalID]
				var globalName = globalInfo.name
				setTextBoxLabel(containerObj,globalName)
			}
		
			$componentRow.append(containerObj)
			
			setElemDimensions(containerObj,textBox.properties.geometry)
		
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(textBox.textBoxID,textBox)
		
			// Callback for any specific initialization for either the form design or view mode
			loadFormConfig.initTextBoxFunc(componentContext,containerObj,textBox)
			
		}

		function initSelectionLayout($componentRow,selection) {
			// Create an HTML block for the container
			console.log("loadFormComponents: initializing text box: " + JSON.stringify(selection))
		
			var containerHTML = selectionContainerHTML(selection.selectionID);
			var containerObj = $(containerHTML)
			
			function setSelectionLabel($selectionContainer,label) {
				$selectionContainer.find('label').text(label)
			}
		
			var componentLink = selection.properties.componentLink
		
			if(componentLink.linkedValType == linkedComponentValTypeField) {
				// text box is linked to a field type
				// Set the label to the field name
				var fieldName = getFieldRef(componentLink.fieldID).name
				setSelectionLabel(containerObj,fieldName)
			} else {
				// text box is linked to a global
				var globalInfo = componentContext.globalsByID[componentLink.globalID]
				var globalName = globalInfo.name
				setSelectionLabel(containerObj,globalName)
			}
		
			$componentRow.append(containerObj)
			
			setElemDimensions(containerObj,selection.properties.geometry)
		
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(selection.selectionID,selection)
		
			// Callback for any specific initialization for either the form design or view mode
			loadFormConfig.initSelectionFunc(componentContext,containerObj,selection)
			
		}



		function initCheckBoxLayout($componentRow,checkBox) {
			// Create an HTML block for the container
			
			var containerHTML = checkBoxContainerHTML(checkBox.checkBoxID);
			var containerObj = $(containerHTML)
			
			var componentLink = checkBox.properties.componentLink
			
			var componentLabel
			if(componentLink.linkedValType == linkedComponentValTypeField) {
				// Set the label to the field name. A span element is used, since
				// the checkbox itself is nested inside a label.
				componentLabel = getFieldRef(componentLink.fieldID).name
			} else {
				var globalInfo = componentContext.globalsByID[componentLink.globalID]
				componentLabel = globalInfo.name
			}
			
			containerObj.find('span').text(componentLabel)
			
			// Position the object withing the #layoutCanvas div
			$componentRow.append(containerObj)
			setElemDimensions(containerObj,checkBox.properties.geometry)
			
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(checkBox.checkBoxID,checkBox)
			
			// Callback for any specific initialization for either the form design or view mode 
			loadFormConfig.initCheckBoxFunc(componentContext,containerObj,checkBox)
		}

		function initRatingLayout($componentRow,rating) {
			// Create an HTML block for the container
			
			var containerHTML = ratingContainerHTML(rating.ratingID);
			
			var containerObj = $(containerHTML)
			
			var componentLink = rating.properties.componentLink
			
			var componentLabel
			if(componentLink.linkedValType == linkedComponentValTypeField) {
				// Set the label to the field name. A span element is used, since
				// the checkbox itself is nested inside a label.
				componentLabel = getFieldRef(componentLink.fieldID).name
			} else {
				var globalInfo = componentContext.globalsByID[componentLink.globalID]
				componentLabel = globalInfo.name
			}
			
			containerObj.find('label').text(componentLabel)
			
			// Position the object withing the #layoutCanvas div
			$componentRow.append(containerObj)
			setElemDimensions(containerObj,rating.properties.geometry)
		
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(rating.ratingID,rating)
			
			// Callback for any specific initialization for either the form design or view mode 
			loadFormConfig.initRatingFunc(componentContext,containerObj,rating)
		}


		function initUserSelectionLayout($componentRow,userSelection) {
			// Create an HTML block for the container
			
			var containerHTML = userSelectionContainerHTML(userSelection.userSelectionID);
			
			var containerObj = $(containerHTML)
			
			var componentLink = userSelection.properties.componentLink
			
			var componentLabel
			if(componentLink.linkedValType == linkedComponentValTypeField) {
				// Set the label to the field name. A span element is used, since
				// the checkbox itself is nested inside a label.
				componentLabel = getFieldRef(componentLink.fieldID).name
			} else {
				var globalInfo = componentContext.globalsByID[componentLink.globalID]
				componentLabel = globalInfo.name
			}
			
			containerObj.find('label').text(componentLabel)
			
			// Position the object withing the #layoutCanvas div
			$componentRow.append(containerObj)
			setElemDimensions(containerObj,userSelection.properties.geometry)
		
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(userSelection.userSelectionID,userSelection)
			
			// Callback for any specific initialization for either the form design or view mode 
			loadFormConfig.initUserSelectionFunc(componentContext,containerObj,userSelection)
		}


		
		function initDatePickerLayout($componentRow,datePicker) {
			// Create an HTML block for the container			
			var containerHTML = datePickerContainerHTML(datePicker.datePickerID);
			var containerObj = $(containerHTML)
			
			var componentLink = datePicker.properties.componentLink
			
			var componentLabel
			if(componentLink.linkedValType == linkedComponentValTypeField) {
				// Set the label to the field name.
				componentLabel = getFieldRef(componentLink.fieldID).name
			} else {
				var globalInfo = componentContext.globalsByID[componentLink.globalID]
				componentLabel = globalInfo.name
			}
			containerObj.find('label').text(componentLabel)
			
			// Position the object withing the #layoutCanvas div
			$componentRow.append(containerObj)
			setElemDimensions(containerObj,datePicker.properties.geometry)
			
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(datePicker.datePickerID,datePicker)
			
			// Callback for any specific initialization for either the form design or view mode
			loadFormConfig.initDatePickerFunc(componentContext,containerObj,datePicker)
			
		}

		function initHtmlEditorLayout($componentRow,htmlEditor) {
			
			var containerHTML = htmlEditorContainerHTML(htmlEditor.htmlEditorID);
			var containerObj = $(containerHTML)
			
			// Set the label to the field name
			var componentLink = htmlEditor.properties.componentLink
			
			var componentLabel
			if(componentLink.linkedValType == linkedComponentValTypeField) {
				// Set the label to the field name.
				componentLabel = getFieldRef(componentLink.fieldID).name
			} else {
				var globalInfo = componentContext.globalsByID[componentLink.globalID]
				componentLabel = globalInfo.name
			}
				
			containerObj.find('label').text(componentLabel)
			
			// Position the object withing the #layoutCanvas div
			$componentRow.append(containerObj)
			setElemDimensions(containerObj,htmlEditor.properties.geometry)
			
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(htmlEditor.htmlEditorID,htmlEditor)
			
			// Callback for any specific initialization for either the form design or view mode
			loadFormConfig.initHtmlEditorFunc(componentContext,containerObj,htmlEditor)
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
			
			var componentLink = image.properties.componentLink
			
			var componentLabel
			if(componentLink.linkedValType == linkedComponentValTypeField) {
				componentLabel = getFieldRef(componentLink.fieldID).name
			} else {
				var globalInfo = componentContext.globalsByID[componentLink.globalID]
				componentLabel = globalInfo.name
			}
			
			$('#'+labelID).text(componentLabel)
			
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(image.imageID,image)
			
			// Callback for any specific initialization for either the form design or view mode
			loadFormConfig.initImageFunc(componentContext,containerObj,image)
			
		}

		var formInfo = componentContext.formInfo
		
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


		for (var selectionIter in formInfo.selections) {			
			var selectionProps = formInfo.selections[selectionIter]			
			console.log("Form layout: selection component info=" + JSON.stringify(selectionProps))
			compenentIDComponentMap[selectionProps.selectionID] = {
				componentInfo: selectionProps,
				initFunc: initSelectionLayout
			}			

		} // for each selection

	
		for (var checkBoxIter in formInfo.checkBoxes) {
			var checkBoxProps = formInfo.checkBoxes[checkBoxIter]
			console.log("loadFormComponents: initializing check box: " + JSON.stringify(checkBoxProps))
			compenentIDComponentMap[checkBoxProps.checkBoxID] = {
				componentInfo: checkBoxProps,
				initFunc: initCheckBoxLayout
			}		
		}

		for (var ratingIter in formInfo.ratings) {
			var ratingProps = formInfo.ratings[ratingIter]
			console.log("loadFormComponents: initializing rating: " + JSON.stringify(ratingProps))
			compenentIDComponentMap[ratingProps.ratingID] = {
				componentInfo: ratingProps,
				initFunc: initRatingLayout
			}		
		}


		for (var userSelIter in formInfo.userSelections) {
			var userSelProps = formInfo.userSelections[userSelIter]
			console.log("loadFormComponents: initializing user selection: " + JSON.stringify(userSelProps))
			compenentIDComponentMap[userSelProps.userSelectionID] = {
				componentInfo: userSelProps,
				initFunc: initUserSelectionLayout
			}		
		}


		for (var headerIter in formInfo.headers) {
			var headerProps = formInfo.headers[headerIter]
			console.log("loadFormComponents: initializing header: " + JSON.stringify(headerProps))
			compenentIDComponentMap[headerProps.headerID] = {
				componentInfo: headerProps,
				initFunc: initHeaderLayout
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
		populateComponentLayout(formLayout,loadFormConfig.formParentElemID,compenentIDComponentMap)
				
		loadFormConfig.doneLoadingFormDataFunc()
	})
	
	
	
}