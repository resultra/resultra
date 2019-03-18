// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initComponentHelpPopupPropertyPanel(params) {
	
	var propsSelector = '#' + params.elemPrefix + "ComponentHelpPopupProps"
	var $props = $(propsSelector)
	
	var $msgInput = $props.find(".helpPopupInput")
	
	$msgInput.html(params.initialMsg)
	
	$msgInput.dblclick(function() {
		if (!inlineCKEditorEnabled($msgInput)) {
			
			var editor = enableInlineCKEditor($msgInput)
			$msgInput.focus()
			
			editor.on('blur', function(event) {
				var popupMsg = editor.getData();
				
				params.setMsg(popupMsg)
							
				disableInlineCKEditor($msgInput,editor)
			})
			
		}
	})
			
}