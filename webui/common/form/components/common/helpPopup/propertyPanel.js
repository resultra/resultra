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