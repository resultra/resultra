$(document).ready(function() {

	$(".layoutContainer").draggable({
		
		grid: [20, 20], // snap to a grid
		cursor: "move",
		containment: "parent", // don't allow the draggable item to move outsize its parent
		
		stop: function(event, ui) {
				  var layoutPos = {
					  x: ui.position.left,
					  y: ui.position.top,
				  };
				  
				  console.log("drag stop: id: " + event.target.id);
				  console.log("drag stop: new position: " + JSON.stringify(layoutPos));
      
	  			/* TODO - Update on server via something like the following:
		          // Save the new position via an AJAX call to the server
		          $.ajax({
		              type: "POST",
		              url: "/",
		              data: JSON.stringify(layoutPos)
		            }).done(function( msg ) {
		              alert( "Data Saved: " + msg );
		            }); 
				  */
	  
		      } // stop function
				  
		
	});

	$(".layoutContainer").resizable({
		aspectRatio: false,
		handles: 'e, w', // Only allow resizing horizontally
		maxWidth: 600,
		minWidth: 100,
		grid: 20, // snap to grid during resize
		
		stop: function(event, ui) {
				  
				  var layoutDim = {
					  width: ui.size.width,
					  height: ui.size.height
				  };
     
				  console.log("resize stop: id: " + event.target.id);
		          console.log("resize stop: new dimensions:" + JSON.stringify(layoutDim));
			  } // 
	});

	// While in layout mode, disable entry into the fields
	$('input[type=text]').prop('disabled', true);
	
	// Set the initial positions of the page elements. 
	// TODO - The list of layout objects and their positions
	// needs to come from the server.
	$("#layoutCanvas").css({position: 'relative'});
	$("#widget01").css({top: 100, left: 100, width: 160, height: 75, position:'absolute'});
	$("#widget02").css({top: 200, left: 100, width: 160, height: 75, position:'absolute'});
	$("#widget03").css({top: 300, left: 100, width: 160, height: 75, position:'absolute'});
	
	$('.layoutPageDiv').layout({
	    center__paneSelector: "#layoutCanvas",
	    east__paneSelector:   "#propertiesSidebar"
	  });

});
