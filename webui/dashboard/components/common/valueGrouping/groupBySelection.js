// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function populateDashboardValueGroupingSelection($selection,fieldType) {
	$selection.empty()
	$selection.append(defaultSelectOptionPromptHTML("Select a grouping"))
	if(fieldType == fieldTypeNumber) {
		$selection.append(selectOptionHTML("none","Don't group values"))
		$selection.append(selectOptionHTML("bucket","Bucket values"))
	}
	else if (fieldType === fieldTypeText) {
		$selection.append(selectOptionHTML("none","Don't group values"))
	}
	else if (fieldType === fieldTypeBool) {
		$selection.append(selectOptionHTML("none","Group into true and false"))
	}
	else if (fieldType === fieldTypeTime) {
		$selection.append(selectOptionHTML("none","Don't group values"))
		$selection.append(selectOptionHTML("day","Day"))	
		$selection.append(selectOptionHTML("week","Week"))	
		$selection.append(selectOptionHTML("monthYear","Month and year"))	
	}
	else {
		console.log("unrecocognized field type: " + fieldType)
	}
}
