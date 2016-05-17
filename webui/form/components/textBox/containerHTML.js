function textBoxContainerHTML(elementID)
{
	var elementInputID = elementID + "_input"
	
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer layoutField draggable resizable" id="'+elementID+'">' +
			'<div class="form-group">'+
				'<label for="' + elementInputID + '">New Field</label>'+
				'<input type="text" name="symbol" class="layoutInput form-control" placeholder="Enter" id="' + elementInputID + '">'+
			'</div>'+
		'</div>';
	return containerHTML
}