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

	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: ratingRef.parentFormID,
			ratingID: ratingRef.ratingID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/rating/setVisibility",params,function(updatedRating) {
			setContainerComponentInfo($rating,updatedRating,updatedRating.ratingID)
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: ratingRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)

	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: ratingRef.properties.readOnly,
		readOnlyPropertyChangedCallback: function(updatedReadOnlyVal) {
			var params = {
				parentFormID: ratingRef.parentFormID,
				ratingID: ratingRef.ratingID,
				readOnly: updatedReadOnlyVal
			}
			jsonAPIRequest("frm/rating/setReadOnly",params,function(updatedRating) {
				setContainerComponentInfo($rating,updatedRating,updatedRating.ratingID)
			})
		}
	}
	initFormComponentReadOnlyPropertyPanel(readOnlyParams)
	
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#ratingProps')
		
	toggleFormulaEditorForField(ratingRef.properties.fieldID)
	
}