function loadFormComponents(loadFormConfig) {
	
	// TODO: formID is accessing a global. Pass it into this function instead.
	jsonAPIRequest("frm/getFormInfo", { formID: formID }, function(formInfo) {
												
		for (var textBoxIter in formInfo.textBoxes) {
			
			// Create an HTML block for the container
			var textBox = formInfo.textBoxes[textBoxIter]
			console.log("loadFormComponents: initializing text box: " + JSON.stringify(textBox))
			
			var containerHTML = textBoxContainerHTML(textBox.textBoxID);
			var containerObj = $(containerHTML)
			
			// Set the label to the field name
			containerObj.find('label').text(textBox.fieldRef.fieldInfo.name)
			
			// Position the object withing the #layoutCanvas div
			$(loadFormConfig.formParentElemID).append(containerObj)
			setElemGeometry(containerObj,textBox.geometry)
			
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(textBox.textBoxID,textBox)
			
			// Callback for any specific initialization for either the form design or view mode
			loadFormConfig.initTextBoxFunc(textBox)
			

		} // for each text box
		
		
		for (var checkBoxIter in formInfo.checkBoxes) {
			
			// Create an HTML block for the container
			var checkBox = formInfo.checkBoxes[checkBoxIter]
			console.log("loadFormComponents: initializing check box: " + JSON.stringify(checkBox))
			
			var containerHTML = checkBoxContainerHTML(checkBox.checkBoxID);
			var containerObj = $(containerHTML)
			
			// Set the label to the field name. A span element is used, since
			// the checkbox itself is nested inside a label.
			containerObj.find('span').text(checkBox.fieldRef.fieldInfo.name)
			
			// Position the object withing the #layoutCanvas div
			$(loadFormConfig.formParentElemID).append(containerObj)
			setElemGeometry(containerObj,checkBox.geometry)
			
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(checkBox.checkBoxID,checkBox)
			
			// Callback for any specific initialization for either the form design or view mode 
			loadFormConfig.initCheckBoxFunc(checkBox)
			

		} // for each text box
	
		for (var datePickerIter in formInfo.datePickers) {
			
			// Create an HTML block for the container
			var datePicker = formInfo.datePickers[datePickerIter]
			console.log("loadFormComponents: initializing date picker: " + JSON.stringify(datePicker))
			
			var containerHTML = datePickerContainerHTML(datePicker.datePickerID);
			var containerObj = $(containerHTML)
			
			// Set the label to the field name
			containerObj.find('label').text(datePicker.fieldRef.fieldInfo.name)
			
			// Position the object withing the #layoutCanvas div
			$(loadFormConfig.formParentElemID).append(containerObj)
			setElemGeometry(containerObj,datePicker.geometry)
			
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
			containerObj.find('label').text(htmlEditor.fieldRef.fieldInfo.name)
			
			// Position the object withing the #layoutCanvas div
			$(loadFormConfig.formParentElemID).append(containerObj)
			setElemGeometry(containerObj,htmlEditor.geometry)
			
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
			setElemGeometry(containerObj,image.geometry)
	
	
			// Set the label to the field name
			var labelID = imageContainerLabelIDFromContainerElemID(image.imageID)
			console.log("loadFormComponents: initializing label: id=" + labelID)
			$('#'+labelID).text(image.fieldRef.fieldInfo.name)
			
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(image.imageID,image)
			
			// Callback for any specific initialization for either the form design or view mode
			loadFormConfig.initImageFunc(image)
			

		} // for each html editor
		
		
		loadFormConfig.doneLoadingFormDataFunc()
	})
	
	
	
}