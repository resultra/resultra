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

function setContainerObjectRef($container,objectRef) {
	$container.data("objectRef",objectRef)
}

function setContainerComponentInfo($container,objectRef,componentID) {
	$container.data("objectRef",objectRef)
	$container.attr("data-componentID",componentID)
}

function getContainerComponentID($container) {
	return $container.attr("data-componentID")
}

function getContainerObjectRef($container) {
	var objectRef = $container.data("objectRef")
	assert(objectRef !== undefined, "getElemObjectRef: Can't get object element reference")
	return objectRef	
}

function setElemObjectRef(objectID, objectRef) {
	$('#'+objectID).data("objectRef",objectRef)
}


function getElemObjectRef(objectID) {
	
	var objectRef = $('#'+objectID).data("objectRef")
	assert(objectRef !== undefined, "getElemObjectRef: Can't get object element reference for object id = " + objectID)
	return objectRef
}