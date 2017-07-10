function initNoteColPropertiesImpl(noteCol) {
	
	setColPropsHeader(noteCol)
	
	var elemPrefix = "note_"
	hideSiblingsShowOne("#noteColProps")
	
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: noteCol.parentTableID,
			noteID: noteCol.noteID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/note/setLabelFormat", formatParams, function(updateNoteCol) {
			setColPropsHeader(updateNoteCol)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: noteCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: noteCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: noteCol.parentTableID,
				noteID: noteCol.noteID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/note/setPermissions",params,function(updateNoteCol) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
}


function initNoteColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		noteID: columnID
	}
	jsonAPIRequest("tableView/note/get", getColParams, function(noteCol) { 
		initNoteColPropertiesImpl(noteCol)
	})
}