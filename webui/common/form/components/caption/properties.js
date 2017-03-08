

function loadFormCaptionProperties($caption,captionRef) {
	console.log("Loading caption properties")
	
	function initColorSchemeProperties() {
		var $schemeSelection = $('#adminCaptionComponentColorSchemeSelection')
		$schemeSelection.val(captionRef.properties.colorScheme)
		initSelectControlChangeHandler($schemeSelection,function(newScheme) {
		
			var sizeParams = {
				parentFormID: captionRef.parentFormID,
				captionID: captionRef.captionID,
				colorScheme: newScheme
			}
			jsonAPIRequest("frm/caption/setColorScheme",sizeParams,function(updatedCaption) {
				setContainerComponentInfo($caption,updatedCaption,updatedCaption.captionID)	
				setFormCaptionColorScheme($caption,updatedCaption.properties.colorScheme)
			})
		
		})
		
	}
	initColorSchemeProperties()
	
	
	
		
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#formCaptionProps')
		
	closeFormulaEditor()
	
}