

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