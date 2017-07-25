function initCommentColPropertiesImpl(commentCol) {
	
	setColPropsHeader(commentCol)
	
	var elemPrefix = "comment_"
	hideSiblingsShowOne("#commentColProps")
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: commentCol.parentTableID,
			commentID: commentCol.commentID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/comment/setLabelFormat", formatParams, function(updateCol) {
			setColPropsHeader(updateCol)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: commentCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: commentCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: commentCol.parentTableID,
				commentID: commentCol.commentID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/comment/setPermissions",params,function(updateCol) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)


	var helpPopupParams = {
		initialMsg: commentCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: commentCol.parentTableID,
				commentID: commentCol.commentID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/comment/setHelpPopupMsg",params,function(updateCol) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
}


function initCommentColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		commentID: columnID
	}
	jsonAPIRequest("tableView/comment/get", getColParams, function(commentCol) { 
		initCommentColPropertiesImpl(commentCol)
	})
}