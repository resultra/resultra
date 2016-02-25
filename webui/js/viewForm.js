

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
	
	this.jumpToRecord = function(recordID) {
		// Use an integer iterator, since the value of the matching record
		// will be used to update this.currRecordIter. Otherwise, a string
		// will be assigned to this.currRecordIter.
		for (recIter = 0; recIter < this.recordRefs.length; recIter++) {
			if(this.recordRefs[recIter].recordID == recordID) {
				this.currRecordIter = recIter
				return true
			}
		}
		return false // TODO - some type of better error handling/assertion checking needed here
	}
	
	this.updateRecordRef = function(updatedRecordRef) {
		for (recIter in this.recordRefs) {
			if(this.recordRefs[recIter].recordID == updatedRecordRef.recordID) {
				this.recordRefs[recIter] = updatedRecordRef
				return
			}
		}
	}
	
	this.appendNewRecord = function(recordRef) {
		this.recordRefs.push(recordRef)
	}
	
	this.currRecordNum = function() {
		return this.currRecordIter + 1 // iterator is 0 based
	}
		
	this.recPageLabel = function () {
		var recNumLabel = (this.currRecordIter + 1).toString()
		var totalRecsLabel = this.numRecords().toString()		
		var recLabel = recNumLabel + " of " + totalRecsLabel
		return recLabel
	}
	
}

function enableRecordButtons(isEnabled)
{
	
	var isDisabled = true;
	if(isEnabled) { isDisabled=false }
	$('#prevRecordButton').prop("disabled",isDisabled)
	$('#nextRecordButton').prop("disabled",isDisabled)
	$('#newRecordButton').prop("disabled",isDisabled)
}

function enableNewRecordButton()
{
	$('#newRecordButton').prop("disabled",false)
}

function initRecordEntryFieldInfo(fieldRef)
{
	// TODO - If the field is a calculated field, disable editing on the field.
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
		var fieldType = container.data("fieldType")
		console.log("container focus out:" 
		    + " containerID: " + containerID
			+ " ,fieldID: " + fieldID
		    + " ,fieldType: " + fieldType
			+ " , inputval:" + inputVal)
		
		currRecordRef = currRecordSet.currRecordRef()
		if(currRecordRef != null) {
			
			// Only update the value if it has changed. Sometimes a user may focus on or tab
			// through a field but not change it. In this case we don't need to update the record.
			if(currRecordRef.fieldValues[fieldID] != inputVal) {
				
				if(fieldType == "text") {
					currRecordRef.fieldValues[fieldID] = inputVal
					var setRecordValParams = { recordID:currRecordRef.recordID, fieldID:fieldID, value:inputVal }
					jsonAPIRequest("setTextFieldValue",setRecordValParams,function(replyData) {
						// After updating the record, the local cache of records in currentRecordSet will
						// be out of date. So after updating the record on the server, the locally cached
						// version of the record also needs to be updated.
						currRecordSet.updateRecordRef(replyData)
					}) // set record's text field value
					
				} else if (fieldType == "number") {
					var numberVal = Number(inputVal)
					if(!isNaN(numberVal)) {
						console.log("Change number val: "
							+ "fieldID: " + fieldID
						    + " ,number = " + numberVal)
						currRecordRef.fieldValues[fieldID] = numberVal
						var setRecordValParams = { recordID:currRecordRef.recordID, fieldID:fieldID, value:numberVal }
						jsonAPIRequest("setNumberFieldValue",setRecordValParams,function(replyData) {
							// After updating the record, the local cache of records in currentRecordSet will
							// be out of date. So after updating the record on the server, the locally cached
							// version of the record also needs to be updated.
							currRecordSet.updateRecordRef(replyData)
						}) // set record's number field value
					}
					
				}
				
			
			} // if input value is different than currently cached value
			
			
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
			enableRecordButtons(true)
		}
		else {
			enableNewRecordButton() // just enable the "New Record" button
		}
				
	}) // getRecord
	
}

function createNewRecord() {
	var newRecordsParams = {}
	jsonAPIRequest("newRecord",newRecordsParams,function(newRecordRef) {
		currRecordSet.appendNewRecord(newRecordRef);
		currRecordSet.jumpToRecord(newRecordRef.recordID)
		loadCurrRecordIntoLayout()
	}) // getRecord
	
}

function initRecordButtonsBehavior()
{
	// Initially disabled the buttons for paging through the records. They'll be 
	// enabled once the records are loaded.
	enableRecordButtons(false)
	
	
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
	
	$('#newRecordButton').click(function(e){ createNewRecord() });
	
}

function initUILayoutPanes()
{
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }
	

	// Initialize the page layout
	$('#layoutPage').layout({
		inset: zeroPaddingInset,
		north: fixedUILayoutPaneParams(40),
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
	
}


$(document).ready(function() {	
	 
	initUILayoutPanes()
		
	// Initialize the semantic ui dropdown menus
	$('.ui.dropdown').dropdown()
	
	initRecordButtonsBehavior()
	  
	initCanvas(initContainerRecordEntryBehavior,initRecordEntryFieldInfo, loadRecords)

}); // document ready
