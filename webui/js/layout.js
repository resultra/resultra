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
	
	var placeholderID = "placeholderFieldID"
	var newFieldDraggablePlaceholderHTML = ''+
		'<div class="ui-widget-content layoutContainer layoutField draggable resizable" id="'+placeholderID+'">' +
			'<div class="field">'+
				'<label>New Field</label>'+
				'<input type="text" name="symbol" class="layoutInput" placeholder="Enter">'+
			'</div>'+
		'</div>`';
	
	$(".newField").draggable({
		
		revert: 'invalid',
		cursor: "move",
		helper: function() { 
			 // Instead of dragging the icon in the gallery, drag a "placeholder image" of the
			 // field itself. This allows the user to see the proporitions of the field relative
			 // to the other elements already on the canvas.
			 var placeholder = $(newFieldDraggablePlaceholderHTML); 
			 // While in layout mode, disable entry into the fields
			 // Interestingly, there's not CSS for this.
			 placeholder.find('input').prop('disabled',true);
			 return placeholder
		 },
		stack: "div", // important to keep the draggable above the other elements
		
		stop: function(event, ui) {
				  var layoutPos = {
					  x: ui.position.left,
					  y: ui.position.top,
				  };
				  
				  console.log("drag stop: id: " + event.target.id);
				  console.log("drag stop: new position: " + JSON.stringify(layoutPos));
      	  
		      } // stop function
	});
	
    $("#layoutCanvas").droppable({
        accept: ".newField",
        activeClass: "ui-state-highlight",
        drop: function( event, ui ) {
 			var theClone = $(ui.helper).clone()
			
			// After the drag operation is complete, the draggable and resizable
			// properties need to be initialized.
			// TODO - This could be put into a common function, since these
			// properties should be the same for objects loaded with the page
			// and newly added objects.
			theClone.draggable ({
				cursor: "move",
				containment: "parent",
				clone: "original",				
				stop: function(event, ui) {
						  var layoutPos = {
							  x: ui.position.left,
							  y: ui.position.top,
						  };
				  
						  console.log("drag stop: id: " + event.target.id);
						  console.log("drag stop: new position: " + JSON.stringify(layoutPos));
      	  
				      } // stop function
			})
			theClone.resizable({
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
			// Get the relative position of the field being dragged
			// and set it within the canvas area
			var canvasRelTop = ui.offset.top-$(this).offset().top
			var canvasRelLeft = ui.offset.left-$(this).offset().left
			theClone.css({top: canvasRelTop, left: canvasRelLeft, position:"absolute"});
			$(this).append(theClone);
		    console.log("drag stop: relTop=" + canvasRelTop + " relLeft=" + canvasRelLeft);
        }
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
	    east__paneSelector:   "#propertiesSidebar",
		west__paneSelector: "#gallerySidebar"
	  });


});
