function loadRatingProperties($rating,ratingRef) {
	console.log("Loading rating properties")
	
	function initIconProps() {
		var $iconSelection = $('#adminRatingComponentIconSelection')
		$iconSelection.val(ratingRef.properties.icon)
		initSelectControlChangeHandler($iconSelection,function(newIcon) {
		
			var iconParams = {
				parentFormID: ratingRef.parentFormID,
				ratingID: ratingRef.ratingID,
				icon: newIcon
			}
			jsonAPIRequest("frm/rating/setIcon",iconParams,function(updatedRating) {
				setContainerComponentInfo($rating,updatedRating,updatedRating.ratingID)
				reInitRatingFormComponentControl($rating,updatedRating)
			})
		
		})
		
	}
	
	initRatingTooltipProperties($rating,ratingRef)
	initIconProps()


	var elemPrefix = "rating_"
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentFormID: ratingRef.parentFormID,
			ratingID: ratingRef.ratingID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/rating/setLabelFormat", formatParams, function(updatedRating) {
			setRatingComponentLabel($rating,updatedRating)
			setContainerComponentInfo($rating,updatedRating,updatedRating.ratingID)
		})	
	}
	
	
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: ratingRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)


	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#ratingProps')
		
	toggleFormulaEditorForField(ratingRef.properties.fieldID)
	
}