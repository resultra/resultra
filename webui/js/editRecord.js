

var currRecordSet;

function RecordSet(recordRefData) {
	
	this.currRecordIter = 0
	this.recordRefs = recordRefData

	this.numRecords = function() {
		return this.recordRefs.length
	}
	
	this.currRecordRef = function() {
		if(this.numRecords() > 0) {
			return this.recordRefs[this.currRecordIter]
		}
		else {
			return null
		}
	}
	
	this.advanceToNextRecord = function() {
		var nextRecordIter = this.currRecordIter + 1
		if(nextRecordIter < this.recordRefs.length)
		{
			this.currRecordIter = nextRecordIter
			return true
		}
		else {
			return false
		}
	}

	this.advanceToPrevRecord = function() {
		var prevRecordIter = this.currRecordIter - 1
		if((prevRecordIter >= 0) && (prevRecordIter < this.recordRefs.length)) {
			this.currRecordIter = prevRecordIter
			return true
		}
		else {
			return false
		}
	}
	
	this.currRecordNum = function() {
		return this.currRecordIter + 1 // iterator is 0 based
	}
	
	this.recPageLabel = function () {
		return this.currRecordNum() + " of " + this.numRecords()
	}
	
}

function enableRecordPagingButtons(isEnabled)
{
	
	var isDisabled = true;
	if(isEnabled) { isDisabled=false }
	$('#prevRecordButton').prop("disabled",isDisabled)
	$('#nextRecordButton').prop("disabled",isDisabled)
}

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
		
		currRecordRef = currRecordSet.currRecordRef()
		if(currRecordRef != null) {
			var setRecordValParams = { recordID:currRecordRef.recordID, fieldID:fieldID, value:inputVal }
			jsonAPIRequest("setRecordFieldValue",setRecordValParams,function(replyData) {
				console.log("Set record value complete")
			}) // set record value
		}
		
		
	}) // focus out
} // initContainerRecordEntryBehavior

function loadCurrRecordIntoLayout()
{
	
	recordRef = currRecordSet.currRecordRef()
	if(recordRef != null)
	{
		console.log("Loading record into layout: fieldValues: " + JSON.stringify(recordRef.fieldValues))

		// Iterate through all the containers in the current layout (which may be a subset of the record's fields),
		// and populate the container's value with the field's value from the record.
		$(".layoutContainer").each(function() {
	
			var containerFieldID = $(this).data("fieldID")
	
			// If the value is not set for the current container, then don't try to 
			// retrieve the value from the record data.
			//
			// In other words, we are populating the "intersection" of field values in the record
			// with the fields shown by the layout's containers.
			if(recordRef.fieldValues.hasOwnProperty(containerFieldID)) {
		
				var fieldVal = recordRef.fieldValues[containerFieldID]

				console.log("Load value into container: " + $(this).attr("id") + " field ID:" + 
							containerFieldID + "  value:" + fieldVal)
		
				$(this).find('input').val(fieldVal)
			} // If record has a value for the current container's associated field ID.
			else
			{
				$(this).find('input').val("") // clear the value in the container
			}
		}) // for each container in the layout
	
		// Update footer to reflect where the current record is in list of currently loaded records
		$('#recordNumLabel').text(currRecordSet.recPageLabel())
		
	} // if current record != null
		
}


function loadRecords()
{
	
	var getRecordsParams = {} // TODO - will include sort & filter options
	jsonAPIRequest("getRecords",getRecordsParams,function(replyData) {
		
		currRecordSet = new RecordSet(replyData);
		if(currRecordSet.numRecords() > 0) {
			loadCurrRecordIntoLayout()		
		}
		
		// Enable the buttons to page through the records
		if(currRecordSet.numRecords() > 0) {
			enableRecordPagingButtons(true)
		}
				
	}) // getRecord
	
}


$(document).ready(function() {	
	 
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }
	
	// Initially disabled the buttons for paging through the records. They'll be 
	// enabled once the records are loaded.
	enableRecordPagingButtons(false)

	// Initialize the page layout
	$('#layoutPage').layout({
		inset: zeroPaddingInset,
		north: fixedUILayoutPaneParams(50),
		east: {
			size: 250,
			resizable:false,
			slidable: false,
			spacing_open:16,
			spacing_closed:16,
			togglerClass:			"toggler",
			togglerLength_open:	128,
			togglerLength_closed: 128,
			togglerAlign_closed: "middle",	// align to top of resizer
			togglerAlign_open: "middle"		// align to top of resizer
			
		}
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
	
	$('#nextRecordButton').click(function(e){
	         e.preventDefault();
			 if(currRecordSet.advanceToNextRecord()) {
			 	loadCurrRecordIntoLayout()
			 }
	});
	
	$('#prevRecordButton').click(function(e){
	         e.preventDefault();
			 if(currRecordSet.advanceToPrevRecord()) {
				 console.log("Advance to next record")
			 	loadCurrRecordIntoLayout()
			 }
			 
	});
		
	// Initialize the semantic ui dropdown menus
	$('.ui.dropdown').dropdown(); 
	  
	initCanvas(initContainerRecordEntryBehavior,initRecordEntryFieldInfo, loadRecords)

}); // document ready
