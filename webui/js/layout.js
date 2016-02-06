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
	
	var placeholderID = "placeholderContainerID"
	var placeholderNum = 1
	
	$(".newField").draggable({
		
		revert: 'invalid',
		cursor: "move",
		helper: function() { 
			 // Instead of dragging the icon in the gallery, drag a "placeholder image" of the
			 // field itself. This allows the user to see the proporitions of the field relative
			 // to the other elements already on the canvas.
			//
			// The placeholderID is a temporary ID to assign to the div. After saving the 
			// new object via JSON call, it is replaced with a unique ID created by the server.
			placeholderID = "placeholderContainerID" + placeholderNum.toString()
			placeholderNum = placeholderNum + 1
			var newFieldDraggablePlaceholderHTML = ''+
				'<div class="ui-widget-content layoutContainer layoutField draggable resizable" id="'+placeholderID+'">' +
					'<div class="field">'+
						'<label>New Field</label>'+
						'<input type="text" name="symbol" class="layoutInput" placeholder="Enter">'+
					'</div>'+
				'</div>`';
			 var placeholder = $(newFieldDraggablePlaceholderHTML); 
			 // While in layout mode, disable entry into the fields
			 // Interestingly, there's not CSS for this.
			 placeholder.find('input').prop('disabled',true);
			 
	//		 console.log("Start drag and drop: placeholder: ID=" + placeholder.attr("id"))
			 
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
							  positionLeft: ui.position.left,
							  positionTop: ui.position.top,
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
						  
			  			var layoutContainerParams = JSON.stringify({
			  				parentLayoutID: layoutID,
			  				containerID: event.target.id,
			  				positionTop: ui.position.top,
			  				positionLeft: ui.position.left,
			  				sizeWidth: ui.size.width,
			  				sizeHeight: ui.size.height
			  			});
					  console.log("resize stop: id: " + event.target.id);
			          console.log("resize stop: new dimensions:" + JSON.stringify(layoutContainerParams));
						console.log("Sending params to resize layout container:" + layoutContainerParams)
			
				        $.ajax({
				           url: '/api/resizeLayoutContainer',
							contentType : 'application/json',
				           data: layoutContainerParams,
				           error: function() {
				              alert("ERROR: Couldn't resizelayout field")
				           },
				           dataType: 'json',
				           success: function(data) {
				              console.log("Done resizing new ID: AJAX response=" + JSON.stringify(data));
				           },
				           type: 'POST'
				        });
						  
					  } // 
			});
			// Get the relative position of the field being dragged
			// and set it within the canvas area
			var canvasRelTop = ui.offset.top-$(this).offset().top
			var canvasRelLeft = ui.offset.left-$(this).offset().left
			theClone.css({top: canvasRelTop, left: canvasRelLeft, position:"absolute"});
			$(this).append(theClone);
			var placeholderID = $(theClone).attr("id");
		    console.log("End Drag and drop: placeholder ID=" + placeholderID + " relTop=" + canvasRelTop + " relLeft=" + canvasRelLeft);
			
			var layoutContainerParams = JSON.stringify({
				parentLayoutID: layoutID,
				containerID: placeholderID,
				positionTop: canvasRelTop,
				positionLeft: canvasRelLeft,
				sizeWidth: $(theClone).width(),
				sizeHeight: $(theClone).height()
			});
			
			console.log("Sending params for new layout container:" + layoutContainerParams)
			
	        $.ajax({
	           url: '/api/newLayoutContainer',
				contentType : 'application/json',
	           data: layoutContainerParams,
	           error: function() {
	              alert("ERROR: Couldn't get new ID for layout field")
	           },
	           dataType: 'json',
	           success: function(data) {
	              console.log("Done getting new ID:response=" + JSON.stringify(data));
				  // TODO - Define some kind of common "validateJSONResponse" function
				  // and possibly write errors back to a server log.
				  if(data.hasOwnProperty("layoutContainerID") && 
					  data.hasOwnProperty("placeholderID")) {
						  // 
			              console.log("New ID for layout field:" + data.layoutContainerID);
						  // Replace the placeholder ID with the permanent one generated via
						  // the API call.
						  var placeholderContainerDiv = document.getElementById(data.placeholderID);
						  placeholderContainerDiv.id = data.layoutContainerID;
					  	
					  }
					  else {
			              console.log("ERROR: Missing properties in newLayoutContainer response:response=" + JSON.stringify(data));
					  }
	           },
	           type: 'POST'
	        });
			
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
	
	// Set the initial positions of the page elements. 
	// TODO - The list of layout objects and their positions
	// needs to come from the server.
	$("#layoutCanvas").css({position: 'relative'});
	
	$('.layoutPageDiv').layout({
	    center__paneSelector: "#layoutCanvas",
	    east__paneSelector:   "#propertiesSidebar",
		west__paneSelector: "#gallerySidebar"
	  });


});
