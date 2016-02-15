

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

// To ensure the popup menus in the header aren't hidden below the center pane, some
// onmouseover behavior is configured for the header/north pane. For this purpose, this 
// global is initialized below then referenced directly within the div tag of the 
// header. 
var recordLayout;

$(document).ready(function() {	
	 
	// Initialize the page layout
	$('#layoutPage').layout({
		north: fixedUILayoutPaneParams(50),
		east: fixedUILayoutPaneParams(200)
	})
	
	recordLayout = $('#recordsPane').layout({
		north: fixedUILayoutPaneAutoSizeToFitContentsParams(),
		south: fixedUILayoutPaneAutoSizeToFitContentsParams(),
		north__showOverflowOnHover:	true,
		center_showOverflowOnHover:false
	})
	
	// Initialize the semantic ui dropdown menus
	$('.ui.dropdown').dropdown(); 
	  
	initCanvas(initContainerRecordEntryBehavior,initRecordEntryFieldInfo, loadRecordIntoLayout)


}); // document ready
