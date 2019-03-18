// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


// Configure and set a form component's label, depending on its format.
function setFormComponentLabel($labelContainer,fieldID, labelFormat) {
    switch (labelFormat.labelType) {
    	case "none":
    		$labelContainer.hide()
    		break;
    	case "custom":
			if (labelFormat.customLabel.length > 0) {
	    		$labelContainer.text(labelFormat.customLabel)
	    		$labelContainer.show()				
			} else {
				$labelContainer.hide()
			}
    		break;
    	default:
    		var fieldName = getFieldRef(fieldID).name
    		$labelContainer.text(fieldName)
    		$labelContainer.show()
    		break;
    }
	
	
}