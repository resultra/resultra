

function checkBoxElemIDFromContainerElemID(checkBoxElemID) {
	return 	checkBoxElemID + '_checkbox'
}

function checkBoxContainerHTML(elementID)
{	
	var checkboxID = checkBoxElemIDFromContainerElemID(elementID)
	var checkboxInput = elementID + '_checkboxInput'
	
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer checkBoxFormContainer draggable resizable" id="'+elementID+'">' +
			'<div class="field">'+
			      '<div class="ui checkbox" id="' + checkboxID + '">' +
			        '<input type="checkbox" name="' + checkboxInput + '" tabindex="0">' +
			        '<label>Check box label.</label>' +
			      '</div>' +
			'</div>'+
		'</div>';
		
	return containerHTML
}
