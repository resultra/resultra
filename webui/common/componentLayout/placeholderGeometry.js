// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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

function findPlaceholderUnderMousePosition(selector,currMouseOffset) {
	var foundPlaceholder = null
	$(selector).each(function() {
		if (foundPlaceholder !== null) { return } // short-circuit loop if placeholder already found
		var placeholderRect = createElemRect($(this))
		if(hitTestLayoutRect(currMouseOffset,placeholderRect)) {
			console.log("found existing placeholder: " + JSON.stringify(placeholderRect))
			foundPlaceholder = $(this)
		}	
	})
	return foundPlaceholder
	
}

function hitExistingPlaceholder(selector,currMouseOffset) {
	var placeholderFound = false
	
	var $placeholderFound = findPlaceholderUnderMousePosition(selector,currMouseOffset)
	if($placeholderFound !== null) {
		placeholderFound = true
	}	
	return placeholderFound
}

