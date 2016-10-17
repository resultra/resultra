

function initFormComponentTimelinePane(elemPrefix) {
	
	var commentAddButtonSelector = createPrefixedSelector(elemPrefix,'TimelineCommentAddButton')
	var commentTextSelector = createPrefixedSelector(elemPrefix,'TimelineCommentText')

	$(commentTextSelector).val("")
	
	initButtonClickHandler(commentAddButtonSelector,function() {
		console.log("Add comment button clicked: comment= " + $(commentTextSelector).val())
		$(commentTextSelector).val("")
	})
}