

function initRecordEntryFieldInfo(fieldRef)
{
	// TBD - While entering records, is there any initialization to do for the fields?
}


function initContainerRecordEntryBehavior(container)
{

	// TODO - Setup the ability for events to be triggered when value changes
	
}


$(document).ready(function() {	
	 
	// Set the initial positions of the page elements. 
	$("#layoutCanvas").css({position: 'relative'});
	
	$('.layoutPageDiv').layout({
	    center__paneSelector: "#layoutCanvas",
	    east__paneSelector:   "#propertiesSidebar",
		west__paneSelector: "#gallerySidebar"
	  });
	  
	initCanvas(initContainerRecordEntryBehavior,initRecordEntryFieldInfo)


}); // document ready
