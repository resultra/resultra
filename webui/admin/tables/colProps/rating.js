function initRatingColPropertiesImpl(ratingCol) {
	
	setColPropsHeader(ratingCol)
	
	var elemPrefix = "rating_"
	hideSiblingsShowOne("#ratingColProps")
		
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: ratingCol.parentTableID,
			ratingID: ratingCol.ratingID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/rating/setLabelFormat", formatParams, function(updateRating) {
			setColPropsHeader(updateRating)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: ratingCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: ratingCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: ratingCol.parentTableID,
				ratingID: ratingCol.ratingID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/rating/setPermissions",params,function(updatedRating) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	
	
}

function initRatingColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		ratingID: columnID
	}
	jsonAPIRequest("tableView/rating/get", getColParams, function(ratingCol) { 
		initRatingColPropertiesImpl(ratingCol)
	})
	
	
	
}