// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function initCheckBoxFormatProps(params) {
	
	var $colorSchemeSelection = $('#adminCheckboxComponentColorSchemeSelection')
	$colorSchemeSelection.val(params.initialColorScheme)
	initSelectControlChangeHandler($colorSchemeSelection,function(newColorScheme) {
		params.setColorScheme(newColorScheme)
	})

	initCheckboxChangeHandler('#adminCheckboxComponentStrikethrough', 
			params.initialStrikethrough, function (newVal) {
			params.setStrikethrough(newVal)	
	})
	
}

