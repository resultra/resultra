

function datePickerElemIDFromContainerElemID(datePickerElemID) {
	return 	datePickerElemID + '_datePicker'
}

function datePickerInputIDFromContainerElemID(datePickerElemID) {
	return 	datePickerElemID + '_datePickerInput'
}


function datePickerContainerHTML(elementID)
{	
	var datePickerID = datePickerElemIDFromContainerElemID(elementID)
	var datePickerInputID = datePickerInputIDFromContainerElemID(elementID)
	
	var containerHTML = ''+
	'<div class="ui-widget-content layoutContainer datePickerContainer  draggable resizable" id="'+elementID+'">' +
		'<div class="field">'+
			'<label>New Field</label>'+
			'<input type="text" name="symbol" id="' + datePickerInputID + '" class="layoutInput" placeholder="Select a Date">'+
		'</div>'+
	'</div>';
	
	
/* Container for Bootstrap datetime	(not yet integrated)
	var containerHTML = ''+
		'<div class="container  datePickerContainer draggable resizable" id="' + elementID + '">' + 
    		'<div class="row">' +
        		'<div class="col-sm-12"">' +
					'<div class="form-group">' +
                		'<div class="input-group date" id="' + datePickerID + '">' +
                    		'<input type="text" class="form-control" />' +
                    		'<span class="input-group-addon">' +
                        		'<span class="glyphicon glyphicon-calendar"></span>' +
                    		'</span>' +
                		'</div>' + // date picker
            		'</div>' + // form group
        		'</div>' + // column
    		'</div>' + // row
	'</div>'; // container
*/		
	return containerHTML
}
