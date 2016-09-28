function createElemRect($elem) {
	var offset = $elem.offset()
	
	var rect = {
		top: offset.top,
		left: offset.left,
		width: $elem.outerWidth(),
		height: $elem.outerHeight()
	}
	return rect
}

function outerTopInsetRect(mainRect, insetX, height) {
	
	var outerRect = {
		top: mainRect.top - height,
		left: mainRect.left + insetX,
		width: mainRect.width - insetX * 2,
		height: height
	}
	return outerRect
}

function outerBottomInsetRect(mainRect, insetX, height) {
	
	var outerRect = {
		top: mainRect.top + mainRect.height,
		left: mainRect.left + insetX,
		width: mainRect.width - insetX * 2,
		height: height
	}
	return outerRect
}

function outerRightInsetRect(mainRect, insetY, width) {
	
	var outerRect = {
		top: mainRect.top + insetY,
		left: mainRect.left + mainRect.width,
		width: width,
		height: mainRect.height - insetY*2
	}
	return outerRect
}

function outerLeftInsetRect(mainRect, insetY, width) {
	
	var outerRect = {
		top: mainRect.top + insetY,
		left: mainRect.left - width,
		width: width,
		height: mainRect.height - insetY*2
	}
	return outerRect
}

function hitTestLayoutRect(position, rect) {
	
	var rectBottom = rect.top + rect.height
	var rectRight = rect.left + rect.width
	
	if( (position.top >= rect.top) && (position.top <= rectBottom) &&
		(position.left >= rect.left) && (position.left <= rectRight)) {
			return true
	} else {
		return false
	}
}

function hitExistingPlaceholder(selector,currMouseOffset) {
	var placeholderFound = false
	$(selector).each(function() {
		if (placeholderFound) { return } // short-circuit loop if placeholder already found
		var placeholderRect = createElemRect($(this))
		if(hitTestLayoutRect(currMouseOffset,placeholderRect)) {
			console.log("found existing placeholder: " + JSON.stringify(placeholderRect))
			placeholderFound = true
		}	
	})
	return placeholderFound
}

