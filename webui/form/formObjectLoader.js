function loadFormObjects(loadFormConfig) {
	
	jsonAPIRequest("frm/getFormInfo", { uniqueID: { objectID: layoutID} }, function(formInfo) {
						
		for (textBoxIter in formInfo.textBoxes) {
			
			// Create an HTML block for the container
			textBox = formInfo.textBoxes[textBoxIter]
			console.log("initializing text box: id=" + JSON.stringify(textBox))
			
			// TODO - textBoxContainerHTMl is specific to text boxes only. Need to use a callback
			// to create the right HTML for the containers.
			var containerHTML = textBoxContainerHTML(textBox.uniqueID.objectID);
			var containerObj = $(containerHTML)
			
			// Set the label to the field name
			containerObj.find('label').text(textBox.fieldRef.fieldInfo.name)
			
			// Position the object withing the #layoutCanvas div
			$(loadFormConfig.formParentElemID).append(containerObj)
			setElemGeometry(containerObj,textBox.geometry)
			
			 // Store the newly created object reference in the DOM element. This is needed for follow-on
			 // property setting, resizing, etc.
			setElemObjectRef(textBox.uniqueID.objectID,textBox)
			
			// Callback for any specific initialization for either 
			loadFormConfig.initTextBoxFunc(textBox)
			

		} // for each text box
		
		loadFormConfig.doneLoadingFormDataFunc()
	})
	
	
	
}