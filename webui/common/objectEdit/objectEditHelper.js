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

// TODO - Verify a singular child objectID suffices as a unique object reference. 
// There can be multiple different types of objects within the same view. So, if child IDs 
// from different entity types are not guaranteed to be unique, the parent ID must also be included.
function setElemObjectRef(objectID, objectRef) {
	$('#'+objectID).data("objectRef",objectRef)
}


function getElemObjectRef(objectID) {
	
	var objectRef = $('#'+objectID).data("objectRef")
	assert(objectRef !== undefined, "getElemObjectRef: Can't get object element reference for object id = " + objectID)
	return objectRef
}