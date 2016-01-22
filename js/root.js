

$(document).ready(function(){
	
    $('#action-button').click(function() {
       $.ajax({
          url: '/pageinfo',
          data: { format: 'json' },
          error: function() {
             $('#info').html('<p>An error has occurred</p>');
          },
          dataType: 'json',
          success: function(data) {
             var $title = $('<h3>').text(data.title);
             $('#info').append($title)
          },
          type: 'GET'
       });
    }); // action buttion click

	// Dynamically initialize the data table
   $.ajax({
      url: '/dataTable',
      data: {
         format: 'json'
      },
      error: function() {
         $('#info').html('<p>An error has occurred</p>');
      },
      dataType: 'json',
      success: function(data) {
		 $('#datatable').DataTable( {
		       columns: data.columns,
		       data: data.data
		     }); // Data Table
      },
      type: 'GET'
   }); // initialize data table
	
			
}); // document ready

