

function initRecordEntryFieldInfo(fieldRef)
{
	// TBD - While entering records, is there any initialization to do for the fields?
}


function initContainerRecordEntryBehavior(container)
{

	// TODOS:
	// - Setup the ability for events to be triggered when value changes
	// - Set tab order of the container vs the others
	// - Disable editing if the field is calculated
	// - Setup validation
	// - Set the default value
	
	// While in edit mode, disable input on the container
	container.focusout(function () {
		var inputVal = container.find("input").val()
		
		var containerID = container.attr("id")
		var fieldID = container.data("fieldID")
		console.log("container focus out:" 
		    + " containerID: " + containerID
			+ " ,fieldID: " + fieldID
			+ " , inputval:" + inputVal)
		
		var setRecordValParams = { recordID:recordID, fieldID:fieldID, value:inputVal }
		jsonAPIRequest("setRecordFieldValue",setRecordValParams,function(replyData) {
			console.log("Set record value complete")
		}) // set record value
		
	}) // focus out
} // initContainerRecordEntryBehavior

function loadRecordIntoLayout()
{
	jsonAPIRequest("getRecord",{recordID:recordID},function(replyData) {
		
		console.log("Loading record into layout: fieldValues: " + JSON.stringify(replyData.fieldValues))
		
		$(".layoutContainer").each(function() {
			
			var containerFieldID = $(this).data("fieldID")
			
			if(replyData.fieldValues.hasOwnProperty(containerFieldID)) {
				var fieldVal = replyData.fieldValues[containerFieldID]

			console.log("Load value into container: " + $(this).attr("id") + " field ID:" + containerFieldID + "  value:" + fieldVal)
				
				
				$(this).find('input').val(fieldVal)
			}
			
		})
		
	}) // getRecord
	
}


$(document).ready(function() {	
	 
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }

	// Initialize the page layout
	$('#layoutPage').layout({
		north: fixedUILayoutPaneParams(50),
		east: fixedUILayoutPaneParams(250)
	})
	
	$('#recordsPane').layout({
		north: fixedUILayoutPaneAutoSizeToFitContentsParams(),
		south: fixedUILayoutPaneAutoSizeToFitContentsParams(),
		north__showOverflowOnHover:	true
	})
	
	
	$('#eastFilterSortPane').layout({
		inset: zeroPaddingInset,
		center__size:.6,
		south__size:.4,
	})
	
	// Initialize the semantic ui dropdown menus
	$('.ui.dropdown').dropdown(); 
	  
	initCanvas(initContainerRecordEntryBehavior,initRecordEntryFieldInfo, loadRecordIntoLayout)


}); // document ready
