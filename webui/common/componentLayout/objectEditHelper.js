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
	
	var objectRef = $('#'+objectID).data("objectRef")
	assert(objectRef !== undefined, "getElemObjectRef: Can't get object element reference for object id = " + objectID)
	return objectRef
}