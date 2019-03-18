// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function colorClassByColorScheme(colorScheme) {
	var classLookup = {
		default:"",
		info:"condFormat-info",
		primary: "condFormat-primary",
		success: "condFormat-success",
		warning: "condFormat-warning",
		danger:"condFormat-danger"
	}

	var bgClass = classLookup[colorScheme]
	if (bgClass === undefined) {
		bgClass = null
	}
	return bgClass
}


function removeConditionalFormatClasses($container) {
	$container.removeClass("condFormat-info condFormat-primary condFormat-success condFormat-warning condFormat-danger")
}
