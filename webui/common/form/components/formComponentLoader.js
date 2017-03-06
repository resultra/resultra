function getFormComponentContext(formContext, contextLoadCompleteCallback) {
	var contextPartsRemaining = 4;
	var context = {}
	
	context.formID = formContext.formID
	context.databaseID = formContext.databaseID
	
	function completeOneContextPart() {
		contextPartsRemaining -= 1
		if(contextPartsRemaining <= 0) {
			contextLoadCompleteCallback(context)
		}
	}
	
	loadFieldInfo(context.databaseID, [fieldTypeAll],function(fieldsByID) {
		context.fieldsByID = fieldsByID
		completeOneContextPart()
	})
	
	initFieldInfo( formContext.databaseID, function () {
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

function populateOneFormLayoutWithComponents(loadFormConfig, componentContext) {
	
	var compenentIDComponentMap = {}	

	function initHeaderLayout($componentRow,header) {
		// Create an HTML block for the container
		console.log("loadFormComponents: initializing header: " + JSON.stringify(header))
	
		var containerHTML = formHeaderContainerHTML(header.headerID);
		var $containerObj = $(containerHTML)
		$containerObj.find(".formHeader").text(header.properties.label)
				
		$componentRow.append($containerObj)
		
		setElemFixedWidthFlexibleHeight($containerObj,header.properties.geometry.sizeWidth)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($containerObj,header,header.headerID)
	
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initHeaderFunc($containerObj,componentContext,header)
		
	}


	function initCaptionLayout($componentRow,caption) {
		// Create an HTML block for the container
		console.log("loadFormComponents: initializing caption: " + JSON.stringify(caption))
	
		var containerHTML = formCaptionContainerHTML(caption.captionID);
		var $containerObj = $(containerHTML)
		$containerObj.find(".formCaption").text(caption.properties.label)
				
		$componentRow.append($containerObj)
		
		setElemDimensions($containerObj,caption.properties.geometry)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($containerObj,caption,caption.captionID)
	
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initCaptionFunc($containerObj,componentContext,caption)
		
	}


	function initFormButtonLayout($componentRow,formButton) {
		// Create an HTML block for the container
		console.log("loadFormComponents: initializing form button: " + JSON.stringify(formButton))
	
		var containerHTML = formButtonContainerHTML(formButton.buttonID);
		var $containerObj = $(containerHTML)
		
		
		jsonAPIRequest("frm/getFormInfo", { formID: formButton.properties.linkedFormID }, function(formInfo) {
			$containerObj.find(".formButton").text(formInfo.form.name)		
		})
		
				
		$componentRow.append($containerObj)
		
		setElemDimensions($containerObj,formButton.properties.geometry)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($containerObj,formButton,formButton.buttonID)
	
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initFormButtonFunc(componentContext,$containerObj,formButton)
		
	}


	function initTextBoxLayout($componentRow,textBox) {
		// Create an HTML block for the container
		console.log("loadFormComponents: initializing text box: " + JSON.stringify(textBox))
	
		var containerHTML = textBoxContainerHTML(textBox.textBoxID);
		var containerObj = $(containerHTML)
		
		function setTextBoxLabel($textBoxContainer,label) {
			$textBoxContainer.find('label').text(label)
		}
	
		var fieldName = getFieldRef(textBox.properties.fieldID).name
		setTextBoxLabel(containerObj,fieldName)
	
		$componentRow.append(containerObj)
		
		setElemDimensions(containerObj,textBox.properties.geometry)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo(containerObj,textBox,textBox.textBoxID)
	
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initTextBoxFunc(componentContext,containerObj,textBox)
		
	}

	function initSelectionLayout($componentRow,selection) {
		// Create an HTML block for the container
		console.log("loadFormComponents: initializing selection: " + JSON.stringify(selection))
	
		var containerHTML = selectionContainerHTML(selection.selectionID);
		var containerObj = $(containerHTML)
		
		function setSelectionLabel($selectionContainer,label) {
			$selectionContainer.find('label').text(label)
		}
	
		var fieldName = getFieldRef(selection.properties.fieldID).name
		setSelectionLabel(containerObj,fieldName)
	
		$componentRow.append(containerObj)
		
		setElemDimensions(containerObj,selection.properties.geometry)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo(containerObj,selection,selection.selectionID)
	
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initSelectionFunc(componentContext,containerObj,selection)
		
	}


	function initCommentLayout($componentRow,comment) {
		// Create an HTML block for the container
		
		var containerHTML = commentContainerHTML(comment.commentID);
		var containerObj = $(containerHTML)
		
		var fieldID = comment.properties.fieldID			
		var componentLabel = getFieldRef(fieldID).name		
		containerObj.find('label').text(componentLabel)
		
		// Position the object withing the #layoutCanvas div
		$componentRow.append(containerObj)
		setElemDimensions(containerObj,comment.properties.geometry)
		
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo(containerObj,comment,comment.commentID)
		
		// Callback for any specific initialization for either the form design or view mode 
		loadFormConfig.initCommentFunc(componentContext,containerObj,comment)
	}



	function initProgressLayout($componentRow,progress) {
		// Create an HTML block for the container
		
		var containerHTML = progressContainerHTML();
		var $progressContainer = $(containerHTML)
				
		var componentLabel = getFieldRef(progress.properties.fieldID).name		
		$progressContainer.find('label').text(componentLabel)
		
		// Position the object withing the #layoutCanvas div
		$componentRow.append($progressContainer)
		setElemDimensions($progressContainer,progress.properties.geometry)
		
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($progressContainer,progress,progress.progressID)
		
		// Callback for any specific initialization for either the form design or view mode 
		loadFormConfig.initProgressFunc(componentContext,$progressContainer,progress)
	}

	function initCheckBoxLayout($componentRow,checkBox) {
		// Create an HTML block for the container
		
		var containerHTML = checkBoxContainerHTML(checkBox.checkBoxID);
		var $checkboxContainer = $(containerHTML)
				
		var componentLabel = getFieldRef(checkBox.properties.fieldID).name		
		$checkboxContainer.find('label').text(componentLabel)
		var checkboxColorSchemeClass = "checkbox-"+checkBox.properties.colorScheme
		$checkboxContainer.addClass(checkboxColorSchemeClass)
		
		// Position the object withing the #layoutCanvas div
		$componentRow.append($checkboxContainer)
		
		setElemFixedWidthFlexibleHeight($checkboxContainer,checkBox.properties.geometry.sizeWidth)
				
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($checkboxContainer,checkBox,checkBox.checkBoxID)
		
		// Callback for any specific initialization for either the form design or view mode 
		loadFormConfig.initCheckBoxFunc(componentContext,$checkboxContainer,checkBox)
	}


	function initRatingLayout($componentRow,rating) {
		// Create an HTML block for the container
		
		var containerHTML = ratingContainerHTML(rating.ratingID);
		
		var $ratingContainer = $(containerHTML)
		
		var componentLabel = getFieldRef(rating.properties.fieldID).name		
		$ratingContainer.find('label').text(componentLabel)
		
		// Position the object withing the #layoutCanvas div
		$componentRow.append($ratingContainer)
		setElemDimensions($ratingContainer,rating.properties.geometry)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($ratingContainer,rating,rating.ratingID)
		
		// Callback for any specific initialization for either the form design or view mode 
		loadFormConfig.initRatingFunc(componentContext,$ratingContainer,rating)
	}


	function initUserSelectionLayout($componentRow,userSelection) {
		// Create an HTML block for the container
		
		var containerHTML = userSelectionContainerHTML(userSelection.userSelectionID);
		
		var containerObj = $(containerHTML)
				
		var componentLabel = getFieldRef(userSelection.properties.fieldID).name
		containerObj.find('label').text(componentLabel)
		
		// Position the object withing the #layoutCanvas div
		$componentRow.append(containerObj)
		setElemDimensions(containerObj,userSelection.properties.geometry)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo(containerObj,userSelection,userSelection.userSelectionID)
		
		// Callback for any specific initialization for either the form design or view mode 
		loadFormConfig.initUserSelectionFunc(componentContext,containerObj,userSelection)
	}


	
	function initDatePickerLayout($componentRow,datePicker) {
		// Create an HTML block for the container			
		var containerHTML = datePickerContainerHTML(datePicker.datePickerID);
		var containerObj = $(containerHTML)
		
		var componentLabel = getFieldRef(datePicker.properties.fieldID).name
		containerObj.find('label').text(componentLabel)
		
		// Position the object withing the #layoutCanvas div
		$componentRow.append(containerObj)
		setElemDimensions(containerObj,datePicker.properties.geometry)
		
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo(containerObj,datePicker,datePicker.datePickerID)
					
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initDatePickerFunc(componentContext,containerObj,datePicker)
		
	}

	function initHtmlEditorLayout($componentRow,htmlEditor) {
		
		var containerHTML = htmlEditorContainerHTML(htmlEditor.htmlEditorID);
		var containerObj = $(containerHTML)
		
		// Set the label to the field name		
		var componentLabel = getFieldRef(htmlEditor.properties.fieldID).name			
		containerObj.find('label').text(componentLabel)
		
		// Position the object withing the #layoutCanvas div
		$componentRow.append(containerObj)
		setElemDimensions(containerObj,htmlEditor.properties.geometry)
		
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo(containerObj,htmlEditor,htmlEditor.htmlEditorID)
		
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
		var componentLabel = getFieldRef(image.properties.fieldID).name		
		var $imageLabel = imageContainerLabelImageComponentContainer(containerObj)
		$imageLabel.text(componentLabel)
		
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo(containerObj,image,image.imageID)
		
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

	for (var progressIter in formInfo.progressIndicators) {
		var progressProps = formInfo.progressIndicators[progressIter]
		console.log("loadFormComponents: initializing progress indicator: " + JSON.stringify(progressProps))
		compenentIDComponentMap[progressProps.progressID] = {
			componentInfo: progressProps,
			initFunc: initProgressLayout
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

	for (var commentIter in formInfo.comments) {
		var commentProps = formInfo.comments[commentIter]
		console.log("loadFormComponents: initializing comment component: " + JSON.stringify(commentProps))
		compenentIDComponentMap[commentProps.commentID] = {
			componentInfo: commentProps,
			initFunc: initCommentLayout
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

	for (var buttonIter in formInfo.formButtons) {
		var buttonProps = formInfo.formButtons[buttonIter]
		console.log("loadFormComponents: initializing form button: " + JSON.stringify(buttonProps))
		compenentIDComponentMap[buttonProps.buttonID] = {
			componentInfo: buttonProps,
			initFunc: initFormButtonLayout
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

	for (var captionIter in formInfo.captions) {
		var captionProps = formInfo.captions[captionIter]
		console.log("loadFormComponents: initializing caption: " + JSON.stringify(captionProps))

		compenentIDComponentMap[captionProps.captionID] = {
			componentInfo: captionProps,
			initFunc: initCaptionLayout
		}			

	}		


	
	var formLayout = formInfo.form.properties.layout
	populateComponentLayout(formLayout,loadFormConfig.$parentFormLayout,compenentIDComponentMap)
	
}


function loadFormComponentsIntoSingleLayout(loadFormConfig, doneLoadingFormDataFunc) {
	
	getFormComponentContext(loadFormConfig.formContext, function(componentContext) {											
		populateOneFormLayoutWithComponents(loadFormConfig,componentContext)		
		doneLoadingFormDataFunc()
	})
	
}