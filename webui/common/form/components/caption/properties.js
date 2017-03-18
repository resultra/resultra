

function loadFormCaptionProperties($caption,captionRef) {
	console.log("Loading caption properties")
	
	var elemPrefix = "caption_"
	
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
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: captionRef.parentFormID,
			captionID: captionRef.captionID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/caption/setVisibility",params,function(updatedCaption) {
			setContainerComponentInfo($caption,updatedCaption,updatedCaption.captionID)	
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: captionRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
	
		
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#formCaptionProps')
		
	closeFormulaEditor()
	
}