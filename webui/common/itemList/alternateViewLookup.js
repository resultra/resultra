// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function createAlternateViewLookupTable(altViews) {
	var altViewLookup = {}
	$.each(altViews,function(index,altView) {
		if (altView.formID != null) {
			altViewLookup[altView.formID] = altView
		} else if (altView.tableID != null) {
			altViewLookup[altView.tableID] = altView
		}
	})
	return altViewLookup
}
