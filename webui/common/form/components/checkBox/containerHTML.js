

function checkBoxElemIDFromContainerElemID(checkBoxElemID) {
	return 	checkBoxElemID + '_checkbox'
}

function checkBoxContainerHTML(elementID)
{	
	var checkboxID = checkBoxElemIDFromContainerElemID(elementID)
	
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer checkBoxFormContainer draggable resizable" id="'+elementID+'">' +
			'<div class="checkbox">' +
				'<label>' + 
				  		'<input type="checkbox" id="' + checkboxID + '"></input><span>Checkbox Label</span> ' +
				'</label>' +
			'</div>' +
		'</div><';
				
	console.log ("Checkbox HTML: " + containerHTML)
		
	return containerHTML
}
