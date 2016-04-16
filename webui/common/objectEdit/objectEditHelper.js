function elemResizeConstraints(minWidth, maxWidth, minHeight, maxHeight) {
	assert(minWidth > 0)
	assert(minHeight > 0)
	assert(maxWidth >= minWidth)
	assert(maxHeight >= minHeight)
	return {
		minWidth: minWidth,
		minHeight: minHeight,
		maxWidth: maxWidth,
		maxHeight: maxHeight,
	}
}

function setElemObjectRef(objectID, objectRef) {
	$('#'+objectID).data("objectRef",objectRef)
}

function getElemObjectRef(objectID) {
	return $('#'+objectID).data("objectRef")
}