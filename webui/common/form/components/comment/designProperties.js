function loadCommentComponentProperties($comment,commentRef) {
	console.log("Loading comment component properties")
	
	var elemPrefix = "comment_"
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for comment box")
		var formatParams = {
			parentFormID: commentRef.parentFormID,
			commentID: commentRef.commentID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/comment/setLabelFormat", formatParams, function(updatedComment) {
			setCommentComponentLabel($comment,updatedComment)
			setContainerComponentInfo($comment,updatedComment,commentRef.commentID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: commentRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)


	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#commentComponentProps')
	
	toggleFormulaEditorForField(commentRef.properties.fieldID)
	
}