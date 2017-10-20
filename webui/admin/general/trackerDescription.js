function initTrackerDescriptionPropertyPanel(trackerDatabaseInfo) {
	
	var $props = $('#adminGeneralTrackerDescription')
	
	var $descInput = $props.find(".adminGeneralTrackerDescriptionInput")
	
	function setTrackerDescription(description) {
		var setDescParams = {
			databaseID:trackerDatabaseInfo.databaseID,
			description:description
		}
		jsonAPIRequest("database/setDescription",setDescParams,function(dbInfo) {
		})
		
	}
	
	$descInput.html(trackerDatabaseInfo.description)
	
	$descInput.dblclick(function() {
		if (!inlineCKEditorEnabled($descInput)) {
			
			var editor = enableInlineCKEditor($descInput)
			$descInput.focus()
			
			editor.on('blur', function(event) {
				var popupMsg = editor.getData();
				
				setTrackerDescription(popupMsg)
							
				disableInlineCKEditor($descInput,editor)
			})
			
		}
	})
			
}