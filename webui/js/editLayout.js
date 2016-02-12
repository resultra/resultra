

function initLayoutEditFieldInfo(fieldRef)
{
	// Add the ability to select the field from the new layout container
	// dialog.
 	var selectFieldOptionHTML = '<option value="' +
 		fieldRef.fieldID + '">' +
 		fieldRef.fieldInfo.name + '</option>'
 	$("#textBoxFieldSelection").append(selectFieldOptionHTML)	
}


function initContainerEditBehavior(container)
{

	// While in edit mode, disable input on the container
	container.find('input').prop('disabled',true);


	// TODO - This could be put into a common function, since these
	// properties should be the same for objects loaded with the page
	// and newly added objects.
	container.draggable ({
		grid: [20, 20], // snap to a grid
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
				  
				  // TODO: send ajax request to reposition the container
  
		      } // stop function
	})
	container.resizable({
		aspectRatio: false,
		handles: 'e, w', // Only allow resizing horizontally
		maxWidth: 600,
		minWidth: 100,
		grid: 20, // snap to grid during resize

		stop: function(event, ui) {  
				  
	  			var layoutContainerParams = {
	  				parentLayoutID: layoutID,containerID: event.target.id,
	  				geometry: { positionTop: ui.position.top,positionLeft: ui.position.left,
						sizeWidth: ui.size.width, sizeHeight: ui.size.height }
	  			};
		
			 	jsonAPIRequest("resizeLayoutContainer",layoutContainerParams,function(replyData) {})	
				  
			  } // stop function
	});
	
}


function newLayoutContainer(containerParams)
{
	var placeholderID = containerParams.containerID
	var containerCreated = false
	
	$( "#newTextBox" ).form({
    	fields: {
	        textBoxFieldSelection: {
	          identifier: 'textBoxFieldSelection',
	          rules: [
	            {
	              type   : 'empty',
	              prompt : 'Please select a field'
	            }
	          ]
	        }
     	},
		inline : true,
  	})
	
	function saveNewTextBox()
	{
		
		// Complete the parameters to create the new layout container by including the 
		// fieldID selected in the dialog box.
		var fieldID = $( "#newTextBox" ).form('get value','textBoxFieldSelection')
		console.log("saveNewTextBox: Selected field ID: " + fieldID)
		containerParams["fieldID"] = fieldID

		jsonAPIRequest("newLayoutContainer",containerParams,function(replyData) {
	          console.log("Done getting new ID:response=" + JSON.stringify(replyData));
			  // TODO - Define some kind of common "validateJSONResponse" function
			  // and possibly write errors back to a server log.
  
			  if(replyData.hasOwnProperty("layoutContainerID") && 
				  replyData.hasOwnProperty("placeholderID")) {
					  // Replace the placeholder ID with the permanent one generated via
					  // the API call.
					  $('#'+placeholderID).find('label').text(fieldsByID[fieldID].name)
					  $('#'+placeholderID).attr("id",replyData.layoutContainerID)
	 				  
					  containerCreated = true
					  dialog.dialog("close")
				  }
				  else {
		              console.log("ERROR: Missing properties in newLayoutContainer response:response=" + JSON.stringify(replyData));
					  dialog.dialog("close")
				  }
	       }) // newLayoutContainer API request
		
	}
	
    dialog = $( "#newTextBox" ).dialog({
      autoOpen: false,
      height: 325, width: 300,
      modal: true,
      buttons: {
        "Create Text Box": function() {
			if($( "#newTextBox" ).form('validate form')) {
				saveNewTextBox()
			} // if validate form
         }, // Create Text Box function
        Cancel: function() {
          dialog.dialog( "close" );
        }
      },
      close: function() {
		  console.log("Close dialog")
		  if(!containerCreated)
		  {
			  // If the the text box creation is not complete, remove the placeholder
			  // from the canvas.
			  $('#'+placeholderID).remove()
		  }
      }
    });
 
    form = dialog.find( "form" ).on( "submit", function( event ) {
      	event.preventDefault();
		if($( "#newTextBox" ).form('validate form')) {
			saveNewTextBox()
		}
    });
	
	$('#newTextBox').form('clear') // clear any previous entries
	$( "#newTextBox" ).dialog("open")
	
}

function droppedObjGeometry(dropDest,droppedObj,ui)
{
	// TODO - Support snapping of the top and left
	var relTop = ui.offset.top-$(dropDest).offset().top
	var relLeft = ui.offset.left-$(dropDest).offset().left
	var objWidth = droppedObj.width()
	var objHeight = droppedObj.height()
		
	return { top: relTop, left: relLeft, width: objWidth, height: objHeight }
}


$(document).ready(function() {
		
	
	$(".newField").draggable({
		
		revert: 'invalid',
		cursor: "move",
		helper: function() { 
			 // Instead of dragging the icon in the gallery, drag a "placeholder image" of the
			 // field itself. This allows the user to see the proporitions of the field relative
			 // to the other elements already on the canvas.
			var newFieldDraggablePlaceholderHTML = fieldContainerHTML(allocNextPlaceholderID());
			var placeholder = $(newFieldDraggablePlaceholderHTML);
			
			 // While in layout mode, disable entry into the fields
			 // Interestingly, there's not CSS for this.
			placeholder.find('input').prop('disabled',true);
			 			 
			 return placeholder
		 },
		stack: "div" // important to keep the draggable above the other elements
		
	});
	
    $("#layoutCanvas").droppable({
        accept: ".newField",
        activeClass: "ui-state-highlight",
        drop: function( event, ui ) {
 			var theClone = $(ui.helper).clone()
			
			// After the drag operation is complete, the draggable and resizable
			// properties need to be initialized.
			initContainerEditBehavior(theClone)
					
			$(this).append(theClone);

			// IMPORTANT - Don't call droppedObjGeometry until after theClone has been appended
			// Otherwise the width and height will be returned as 0.
			var placeholderID = $(theClone).attr("id");
			var droppedObjGeom = droppedObjGeometry(this,theClone,ui)
			
			theClone.css({top: droppedObjGeom.top, left: droppedObjGeom.left, position:"absolute"});
			
		    console.log("End Drag and drop: placeholder ID=" + placeholderID + 
				" drop loc =" + JSON.stringify(droppedObjGeom));
			
			var layoutContainerParams = {
				parentLayoutID: layoutID, containerID: placeholderID,
				geometry: {positionTop: droppedObjGeom.top, positionLeft: droppedObjGeom.left,
				sizeWidth: droppedObjGeom.width,sizeHeight: droppedObjGeom.height}
				};
			
			newLayoutContainer(layoutContainerParams)
						
        }
    }); // #layoutCanvas droppable
	
	
	// Initialize the newTextBox dialog with the minimum parameters. This is necessary
	// to hide the dialog from view when the document is initially loaded. The
	// dialog is fully re-initialized just prior to it being opened.
    $( "#newTextBox" ).dialog({ autoOpen: false })
 
		
	// Set the initial positions of the page elements. 
	// TODO - The list of layout objects and their positions
	// needs to come from the server.
	$("#layoutCanvas").css({position: 'relative'});
	
	$('.layoutPageDiv').layout({
	    center__paneSelector: "#layoutCanvas",
	    east__paneSelector:   "#propertiesSidebar",
		west__paneSelector: "#gallerySidebar"
	  });
	  
	  
	  function initCanvasComplete() {} // noop
	initCanvas(initContainerEditBehavior,initLayoutEditFieldInfo,initCanvasComplete)


});
