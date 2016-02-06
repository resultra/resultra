
function jsonAPIRequest(apiName,requestData, successFunc)
{
	var jsonReqData = JSON.stringify(requestData)
			
	// TODO - In debug builds, the API logging could be enabled, but disabled in production
	console.log("JSON API Request: api name = " + apiName + " requestData =" + jsonReqData)
	
    $.ajax({
       url: '/api/'+ apiName,
		contentType : 'application/json',
       data: jsonReqData,
       error: function() {
		  var errMsg = "ERROR: API Request failed: api name = " + apiName + " requestData =" + jsonReqData
		  console.log(errMsg)
          alert(errMsg)
       },
       dataType: 'json',
       success: function(replyData) {
		  console.log("JSON API Request succeeded: api name = " + apiName + " replyData =" + JSON.stringify(replyData))
		  successFunc(replyData)
       },
       type: 'POST'
    });
	
}



// A placeholderID is a temporary ID to assign to the div. After saving the 
// new object via JSON call, it is replaced with a unique ID created by the server.
var placeholderNum = 1
function allocNextPlaceholderID()
{
	placeholderID = "placeholderContainerID" + placeholderNum.toString()
	placeholderNum = placeholderNum + 1
	return placeholderID
}

function fieldContainerHTML(id)
{
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer layoutField draggable resizable" id="'+id+'">' +
			'<div class="field">'+
				'<label>New Field</label>'+
				'<input type="text" name="symbol" class="layoutInput" placeholder="Enter">'+
			'</div>'+
		'</div>`';
	return containerHTML
}

function initContainerEditBehavior(container)
{
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
	  				positionTop: ui.position.top,positionLeft: ui.position.left,
	  				sizeWidth: ui.size.width, sizeHeight: ui.size.height
	  			};
		
			 	jsonAPIRequest("resizeLayoutContainer",layoutContainerParams,function(replyData) {})	
				  
			  } // stop function
	});
	
}


function newLayoutContainer(containerParams)
{
	var jsonReqData = JSON.stringify(containerParams)
			
	console.log("newLayoutContainer: params for new layout container:" + jsonReqData)
	
	jsonAPIRequest("newLayoutContainer",containerParams,function(replyData) {
          console.log("Done getting new ID:response=" + JSON.stringify(replyData));
		  // TODO - Define some kind of common "validateJSONResponse" function
		  // and possibly write errors back to a server log.
		  if(replyData.hasOwnProperty("layoutContainerID") && 
			  replyData.hasOwnProperty("placeholderID")) {
				  // Replace the placeholder ID with the permanent one generated via
				  // the API call.
				  var placeholderContainerDiv = document.getElementById(replyData.placeholderID);
				  placeholderContainerDiv.id = replyData.layoutContainerID;
			  	
			  }
			  else {
	              console.log("ERROR: Missing properties in newLayoutContainer response:response=" + JSON.stringify(replyData));
			  }
       })
		
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


function initCanvas()
{
	var jsonReqData = jsonAPIRequest("getLayoutContainers",{layoutID: layoutID},
		function(replyData) {
			  for(containerIter in replyData)
			  {
				container = replyData[containerIter]
			  	console.log("initializing container: id=" + JSON.stringify(container))
				var containerHTML = fieldContainerHTML(container.containerID);
				var containerObj = $(containerHTML)
				containerObj.find('input').prop('disabled',true);
				initContainerEditBehavior(containerObj)
				$("#layoutCanvas").append(containerObj)
				containerObj.css({top: container.positionTop, left: container.positionLeft, 
						width: container.sizeWidth, height: container.sizeHeight,
					position:"absolute"});
				} // for each container
		 })	
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
				positionTop: droppedObjGeom.top, positionLeft: droppedObjGeom.left,
				sizeWidth: droppedObjGeom.width,sizeHeight: droppedObjGeom.height};
			
			newLayoutContainer(layoutContainerParams)
						
        }
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
	  
	initCanvas()


});
