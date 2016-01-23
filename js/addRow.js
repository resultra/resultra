$(document).ready(function() {
	
	var addRowValidationRules = {
      fields: {
		symbol : 'empty',
        qty     : 'integer[1..10000000]',
        price   : 'number'
      }
  	 // Additional validation fields would go here
	};
	
		
	function submitAJAXForm() {
		var formData = JSON.stringify($('#addRowForm').form('get values'));
		$.ajax({
			url: '/addRow',
			data: formData,
			dataType: 'json',
			type: 'POST',
			async: false,
			success: function(responseData) { // callback method for further manipulations   
				//	TODO - Add success handling (e.g. clear form or redirect)     
			},
			error: function(responseData) { // if error occured
					alert("Failure submitting form")
			}
		});
		return false;		
	}


	$( "#addRowForm" ).form(addRowValidationRules);      
		  
	$( "#addRowButton" ).click(function(event){
	    event.preventDefault();
		// Override the default form handling to perform AJAX/JSON
		// based form handling instead.
		if($( "#addRowForm" ).form('validate form'))
		{
			submitAJAXForm();
		}
	   	return false;

	});

}); // document ready