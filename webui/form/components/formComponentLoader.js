


function loadFormComponents(loadFormConfig) {
	
	jsonAPIRequest("frm/getFormInfo", { formID: loadFormConfig.formID }, function(formInfo) {
													
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
		
		function saveUpdatedFormComponentLayout(updatedLayout) {
			console.log("saveUpdatedFormComponentLayout: component layout = " + JSON.stringify(updatedLayout))		
			var setLayoutParams = {
				formID: loadFormConfig.formID,
				layout: updatedLayout
			}
			jsonAPIRequest("frm/setLayout", setLayoutParams, function(formInfo) {
			})
		}		
		
		
		var formLayout = formInfo.form.properties.layout
		populateComponentLayout(formLayout,loadFormConfig.formParentElemID,
				compenentIDComponentMap,saveUpdatedFormComponentLayout)
				
		loadFormConfig.doneLoadingFormDataFunc()
	})
	
	
	
}