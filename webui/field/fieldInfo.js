var fieldTypeNumber = "number"
var fieldTypeText = "text"
var fieldTypeBool = "bool"


// Use for populating menus with field selections
function selectFieldHTML(fieldID, fieldName) {
 	return selectFieldOptionHTML = '<option value="' +
 		fieldID + '">' + fieldName + '</option>'
}

function populateFieldSelectionMenu(fieldsByID, menuSelector) {
	$(menuSelector).empty()
	$(menuSelector).append(emptyOptionHTML("Select a Field"))
	$.each(fieldsByID, function(fieldID, fieldInfo) {
		$(menuSelector).append(selectFieldHTML(fieldID, fieldInfo.name))		
	})
}

function loadFieldInfo(fieldInfoCallback) {
	
	// Map of field ID's to the fieldRef object: see initialization below
	var fieldsByID = {}
	
	function processFieldInfo(fieldsByType) {
		var textFields = fieldsByType.textFields
		for (textFieldIter in textFields) {
			
			var textField = textFields[textFieldIter]
			
			console.log("Text field: " + textField.fieldInfo.name)

			// Populate a map/dictionary of field IDs to the field information.
			// This is needed when creating new layout elements (text boxes, etc.),
			// so the fields information can be used after creation of the layout
			// element.
			fieldsByID[textField.fieldID] = textField.fieldInfo


		} // for each text field
		
		var numberFields = fieldsByType.numberFields
		for (numberFieldIter in numberFields) {
			
			var numberField = numberFields[numberFieldIter]

			console.log("Number field: " + numberField.fieldInfo.name)

			
			// Populate a map/dictionary of field IDs to the field information.
			// This is needed when creating new layout elements (text boxes, etc.),
			// so the fields information can be used after creation of the layout
			// element.
			fieldsByID[numberField.fieldID] = numberField.fieldInfo

		} // for each number field
		
		fieldInfoCallback(fieldsByID)
	}
	
	jsonAPIRequest("getFieldsByType", {},processFieldInfo)
	
} // loadFieldInfo

