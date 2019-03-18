// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initTrackerDescriptionPropertyPanel(trackerDatabaseInfo) {
	
	var $props = $('#adminGeneralTrackerDescription')
	
	var $descInput = $props
	
	function setTrackerDescription(description) {
		var setDescParams = {
			databaseID:trackerDatabaseInfo.databaseID,
			description:description
		}
		jsonAPIRequest("database/setDescription",setDescParams,function(dbInfo) {
		})
		
	}
	
	$descInput.html(formatInlineContentHTMLDisplay(trackerDatabaseInfo.description))
	
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