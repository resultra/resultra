

function checkBoxContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer layoutField draggable resizable" id="'+elementID+'">' +
			'<div class="field">'+
				'<label>New Field</label>'+
				'<input type="text" name="symbol" class="layoutInput" placeholder="Enter">'+
			'</div>'+
		'</div>';
	return containerHTML
}
