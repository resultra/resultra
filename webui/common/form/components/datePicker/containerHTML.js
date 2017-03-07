

function datePickerInputFromContainer($datePickerContainer) {
	return 	$datePickerContainer.find(".datePickerComponentInput")
}


function datePickerContainerHTML(elementID)
{	
	
	var containerHTML = ''+
	'<div class="layoutContainer datePickerContainer">' +
		'<div class="form-group">'+
			'<label>New Field</label>'+
			'<div class="datePickerInputContainer">' + 
				'<input type="text" name="symbol"  class="form-control datePickerComponentInput" placeholder="Select a date">' +
			'</div>'+
		'</div>'+
	'</div>';
	
	return containerHTML
}
