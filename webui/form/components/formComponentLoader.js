function loadFormComponents(loadFormConfig) {
	
	jsonAPIRequest("frm/getFormInfo", { formID: layoutID }, function(formInfo) {
												
		for (var textBoxIter in formInfo.textBoxes) {
			
			// Create an HTML block for the container
			var textBox = formInfo.textBoxes[textBoxIter]
			console.log("loadFormComponents: initializing text box: " + JSON.stringify(textBox))
			
			// TODO - textBoxContainerHTMl is specific to text boxes only. Need to use a callback
			// to create the right HTML for the containers.
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
			
			// TODO - textBoxContainerHTMl is specific to text boxes only. Need to use a callback
			// to create the right HTML for the containers.
			var containerHTML = checkBoxContainerHTML(checkBox.checkBoxID);
			var containerObj = $(containerHTML)
			
			// Set the label to the field name
			containerObj.find('label').text(checkBox.fieldRef.fieldInfo.name)
			
			// Position the object withing the #layoutCanvas div
			$(loadFormConfig.formParentElemID).append(containerObj)
			setElemGeometry(containerObj,checkBox.geometry)
			
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(checkBox.checkBoxID,checkBox)
			
			// Callback for any specific initialization for either the form design or view mode 
			loadFormConfig.initCheckBoxFunc(checkBox)
			

		} // for each text box
		
		
		loadFormConfig.doneLoadingFormDataFunc()
	})
	
	
	
}