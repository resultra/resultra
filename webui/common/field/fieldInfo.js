var fieldTypeNumber = "number"
var fieldTypeText = "text"
var fieldTypeLongText = "longText"
var fieldTypeAttachment = "file"
var fieldTypeBool = "bool"
var fieldTypeTime = "time"
var fieldTypeUser = "user"
var fieldTypeAll = "all"
var fieldTypeComment = "comment"

function fieldTypeLabel(fieldType) {
	switch (fieldType) {
	case fieldTypeNumber: return "Number"
	case fieldTypeText: return "Text"
	case fieldTypeTime: return "Date and/or Time"
	case fieldTypeUser: return "User"
	case fieldTypeBool: return "True or False (Boolean)"
	case fieldTypeAttachment: return "File"
	case fieldTypeLongText: return "Long Text"
	case fieldTypeComment: return "Comment"
	default: return "Unknown field type"
	}
}



var fieldInfoFieldsByID

// initFieldInfo loads the field information once, then it is available globally as a service to other parts of client
// code which also references field information.
// TODO - Some type of polling/push mechanism is needed to keep this table up to date.
function initFieldInfo(databaseID, fieldInitDoneCallbackFunc) {
	loadFieldInfo(databaseID, [fieldTypeAll],function(fieldsByID) {
		fieldInfoFieldsByID = fieldsByID;
		fieldInitDoneCallbackFunc()
	})
}


function getFieldRef(fieldID) {
	
	var fieldRef = fieldInfoFieldsByID[fieldID]
	assert(fieldRef !== undefined, "No field information found for field ID = " + fieldID)
	
	return fieldRef
	
}

function getFieldsByID() {
	
	assert(fieldInfoFieldsByID !== undefined, "getFieldsByID: field info not initialized")
	
	return fieldInfoFieldsByID
}

function createFieldTypesFilterInfo(fieldTypes) {
	var doLoadFieldByType = {}
	var loadAllFieldTypes = false
	var specificFieldTypes = []
	for(var fieldTypeIndex = 0; fieldTypeIndex != fieldTypes.length; fieldTypeIndex++) {
		var fieldType = fieldTypes[fieldTypeIndex]
		if(fieldType == fieldTypeAll)
		{
			loadAllFieldTypes = true
			console.log("loadFieldInfo: loading field info for all field types")
		}
		else {
			specificFieldTypes.push(fieldType)
			doLoadFieldByType[fieldType] = true
		}
	}
	var filterInfo = {
		loadAllFieldTypes: loadAllFieldTypes, 
		doLoadFieldByType: doLoadFieldByType,
		fieldTypes: specificFieldTypes
	}
	
	return filterInfo
}


function getFilteredFieldsByID(fieldTypes) {
	var unfilteredFieldsByID = getFieldsByID()
	var filterInfo = createFieldTypesFilterInfo(fieldTypes)
	var filteredFields = {}
	for (var fieldID in unfilteredFieldsByID) {
		var fieldRef = unfilteredFieldsByID[fieldID]
		if(filterInfo.loadAllFieldTypes || filterInfo.doLoadFieldByType[fieldRef.type]==true) {
			filteredFields[fieldID] = fieldRef
		}
	}
	return filteredFields
}

// Use for populating menus with field selections
function selectFieldHTML(fieldID, fieldName) {
 	return selectFieldOptionHTML = '<option value="' +
 		fieldID + '">' + fieldName + '</option>'
}

function populateFieldSelectionControlMenu(fieldsByID, $menuSelector) {
	$menuSelector.empty()
	$menuSelector.append(defaultSelectOptionPromptHTML("Select a Field"))
	
	for (var fieldID in fieldsByID) {
	  if (fieldsByID.hasOwnProperty(fieldID)) {
		  var fieldInfo = fieldsByID[fieldID]	  	
		  $menuSelector.append(selectFieldHTML(fieldID, fieldInfo.name))		
	  }
	}
}

function populateSortedFieldSelectionMenu($menuSelector,sortedFields) {
	$menuSelector.empty()
	$menuSelector.append(defaultSelectOptionPromptHTML("Select a Field"))
	for(var fieldIndex in sortedFields) {
		var fieldInfo = sortedFields[fieldIndex]
		$menuSelector.append(selectFieldHTML(fieldInfo.fieldID, fieldInfo.name))		
	}
}

function createFieldsByIDMap(fieldList) {
	var fieldsByID = {}
	for (var fieldIndex in fieldList) {
		var fieldInfo = fieldList[fieldIndex]
		fieldsByID[fieldInfo.fieldID] = fieldInfo
	}
	
	return fieldsByID
}


function loadFieldInfo(parentDatabaseID,fieldTypes,fieldInfoCallback) {
	
	// Map of field ID's to the fieldRef object: see initialization below
	var fieldsByID = {}
	
	var filterInfo = createFieldTypesFilterInfo(fieldTypes)
		
	function processFieldInfo(fieldsByType) {
		
		if(filterInfo.loadAllFieldTypes || filterInfo.doLoadFieldByType[fieldTypeText]==true) {
			var textFields = fieldsByType.textFields
			for (textFieldIter in textFields) {		
				var textField = textFields[textFieldIter]			
				console.log("Text field: " + textField.name)
				fieldsByID[textField.fieldID] = textField
			} // for each text field
		}
		
		if(filterInfo.loadAllFieldTypes || filterInfo.doLoadFieldByType[fieldTypeNumber]==true) {
			var numberFields = fieldsByType.numberFields
			for (numberFieldIter in numberFields) {		
				var numberField = numberFields[numberFieldIter]
				console.log("Number field: " + numberField.name)
				fieldsByID[numberField.fieldID] = numberField
			} // for each number field
		}

		if(filterInfo.loadAllFieldTypes || filterInfo.doLoadFieldByType[fieldTypeUser]==true) {
			var userFields = fieldsByType.userFields
			for (userFieldIter in userFields) {		
				var userField = userFields[userFieldIter]
				console.log("User field: " + userField.name)
				fieldsByID[userField.fieldID] = userField
			} // for each number field
		}


		if(filterInfo.loadAllFieldTypes || filterInfo.doLoadFieldByType[fieldTypeBool]==true) {
			var boolFields = fieldsByType.boolFields
			for (boolFieldIter in boolFields) {
				var boolField = boolFields[boolFieldIter]
				console.log("Bool field: " + boolField.name)
				fieldsByID[boolField.fieldID] = boolField
			} // for each bool field
		}
		
		if(filterInfo.loadAllFieldTypes || filterInfo.doLoadFieldByType[fieldTypeTime]==true) {
			var timeFields = fieldsByType.timeFields
			for (timeFieldIter in timeFields) {
				var timeField = timeFields[timeFieldIter]
				console.log("Time field: " + timeField.name)
				fieldsByID[timeField.fieldID] = timeField
			} // for each bool field
		}
		
		if(filterInfo.loadAllFieldTypes || filterInfo.doLoadFieldByType[fieldTypeLongText]==true) {
			var longTextFields = fieldsByType.longTextFields
			for (longTextFieldIter in longTextFields) {		
				var longTextField = longTextFields[longTextFieldIter]			
				console.log("Long Text field: " + longTextField.name)
				fieldsByID[longTextField.fieldID] = longTextField
			} // for each long text field
		}
	
		if(filterInfo.loadAllFieldTypes || filterInfo.doLoadFieldByType[fieldTypeAttachment]==true) {
			var fileFields = fieldsByType.fileFields
			for (fileFieldIter in fileFields) {		
				var fileField = fileFields[fileFieldIter]			
				console.log("file field: " + fileField.name)
				fieldsByID[fileField.fieldID] = fileField
			} // for each file field
		}

		if(filterInfo.loadAllFieldTypes || filterInfo.doLoadFieldByType[fieldTypeComment]==true) {
			var commentFields = fieldsByType.commentFields
			for (commentFieldIter in commentFields) {		
				var commentField = commentFields[commentFieldIter]			
				console.log("comment field: " + commentField.name)
				fieldsByID[commentField.fieldID] = commentField
			} // for each file field
		}
	
		
		fieldInfoCallback(fieldsByID)
	}
	
	jsonAPIRequest("field/getListByType", {parentDatabaseID:parentDatabaseID},processFieldInfo)
	
} // loadFieldInfo

function loadSortedFieldInfo(parentDatabaseID,fieldTypes,sortedFieldListCallback) {
	
	
	var filterInfo = createFieldTypesFilterInfo(fieldTypes)
	
	if (filterInfo.loadAllFieldTypes) {
		var getFieldParams = {
			parentDatabaseID: parentDatabaseID
		}
		
		jsonAPIRequest("field/getAllSortedFields",getFieldParams,function(sortedFields) {
			sortedFieldListCallback(sortedFields)
		})
	} else {
		var getFieldParams = {
			parentDatabaseID: parentDatabaseID,
			fieldTypes: filterInfo.fieldTypes			
		}
		jsonAPIRequest("field/getSortedListByType",getFieldParams,function(sortedFields) {
			sortedFieldListCallback(sortedFields)
		})
	}
	
	
	
}

