var fieldTypeNumber = "number"
var fieldTypeText = "text"
var fieldTypeLongText = "longText"
var fieldTypeFile = "file"
var fieldTypeBool = "bool"
var fieldTypeTime = "time"
var fieldTypeAll = "all"



// Use for populating menus with field selections
function selectFieldHTML(fieldID, fieldName) {
 	return selectFieldOptionHTML = '<option value="' +
 		fieldID + '">' + fieldName + '</option>'
}

function populateFieldSelectionMenu(fieldsByID, menuSelector) {
	$(menuSelector).empty()
	$(menuSelector).append(defaultSelectOptionPromptHTML("Select a Field"))
	$.each(fieldsByID, function(fieldID, fieldInfo) {
		$(menuSelector).append(selectFieldHTML(fieldID, fieldInfo.name))		
	})
}

function loadFieldInfo(parentTableID,fieldTypes,fieldInfoCallback) {
	
	// Map of field ID's to the fieldRef object: see initialization below
	var fieldsByID = {}
	
	var doLoadFieldByType = {}
	var loadAllFieldTypes = false
	for(var fieldTypeIndex = 0; fieldTypeIndex != fieldTypes.length; fieldTypeIndex++) {
		var fieldType = fieldTypes[fieldTypeIndex]
		if(fieldType == fieldTypeAll)
		{
			loadAllFieldTypes = true
			console.log("loadFieldInfo: loading field info for all field types")
		}
		else {
			doLoadFieldByType[fieldType] = true
		}
	}
	
	function processFieldInfo(fieldsByType) {
		
		if(loadAllFieldTypes || doLoadFieldByType[fieldTypeText]==true) {
			var textFields = fieldsByType.textFields
			for (textFieldIter in textFields) {		
				var textField = textFields[textFieldIter]			
				console.log("Text field: " + textField.fieldInfo.name)
				fieldsByID[textField.fieldID] = textField.fieldInfo
			} // for each text field
		}
		
		if(loadAllFieldTypes || doLoadFieldByType[fieldTypeNumber]==true) {
			var numberFields = fieldsByType.numberFields
			for (numberFieldIter in numberFields) {		
				var numberField = numberFields[numberFieldIter]
				console.log("Number field: " + numberField.fieldInfo.name)
				fieldsByID[numberField.fieldID] = numberField.fieldInfo
			} // for each number field
		}

		if(loadAllFieldTypes || doLoadFieldByType[fieldTypeBool]==true) {
			var boolFields = fieldsByType.boolFields
			for (boolFieldIter in boolFields) {
				var boolField = boolFields[boolFieldIter]
				console.log("Bool field: " + boolField.fieldInfo.name)
				fieldsByID[boolField.fieldID] = boolField.fieldInfo
			} // for each bool field
		}
		
		if(loadAllFieldTypes || doLoadFieldByType[fieldTypeTime]==true) {
			var timeFields = fieldsByType.timeFields
			for (timeFieldIter in timeFields) {
				var timeField = timeFields[timeFieldIter]
				console.log("Time field: " + timeField.fieldInfo.name)
				fieldsByID[timeField.fieldID] = timeField.fieldInfo
			} // for each bool field
		}
		
		if(loadAllFieldTypes || doLoadFieldByType[fieldTypeLongText]==true) {
			var longTextFields = fieldsByType.longTextFields
			for (longTextFieldIter in longTextFields) {		
				var longTextField = longTextFields[longTextFieldIter]			
				console.log("Long Text field: " + longTextField.fieldInfo.name)
				fieldsByID[longTextField.fieldID] = longTextField.fieldInfo
			} // for each long text field
		}
	
		if(loadAllFieldTypes || doLoadFieldByType[fieldTypeFile]==true) {
			var fileFields = fieldsByType.fileFields
			for (fileFieldIter in fileFields) {		
				var fileField = fileFields[fileFieldIter]			
				console.log("file field: " + fileField.fieldInfo.name)
				fieldsByID[fileField.fieldID] = fileField.fieldInfo
			} // for each file field
		}
	
		
		fieldInfoCallback(fieldsByID)
	}
	
	assert(parentTableID !== undefined)
	assert(parentTableID.length > 0)
	jsonAPIRequest("field/getListByType", {parentTableID:parentTableID},processFieldInfo)
	
} // loadFieldInfo

