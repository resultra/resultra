

var currRecordSet;


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


function loadCurrRecordIntoLayout()
{
	recordRef = currRecordSet.currRecordRef()
	if(recordRef != null)
	{
		console.log("Loading record into layout: record field values: " + JSON.stringify(recordRef.fieldValues))

		// Iterate through all the containers in the current layout (which may be a subset of the record's fields),
		// and populate the container's value with the field's value from the record.
		$(".layoutContainer").each(function() {
			
			// Each type of form object needs to set a "viewFormConfig" object on it's DOM element. The loadRecord()
			// function is called on each of these objects to perform per form object record initialization.
			var viewFormConfig = $(this).data("viewFormConfig")
			viewFormConfig.loadRecord($(this),recordRef)
	
		}) // for each container in the layout
	
		// Update footer to reflect where the current record is in list of currently loaded records
		$('#recordNumLabel').text(currRecordSet.recPageLabel())
		
	} // if current record != null
}


function reloadRecords(reloadParams) {
	
	jsonAPIRequest("recordRead/getFilteredSortedRecordValues",reloadParams,function(replyData) {
		
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

function reloadSortedAndFilterRecords()
{
	
	var selectedFilterIDs = getSelectedFilterPanelFilterIDs()
	var sortRules = getSortPaneSortRules()
	
	var getFilteredRecordsParams = { 
		tableID: viewFormContext.tableID,
		filterIDs: selectedFilterIDs,
		sortRules: sortRules}
	
	reloadRecords(getFilteredRecordsParams)
}

function createNewRecord() {
	var newRecordsParams = {parentTableID:tableID}
	jsonAPIRequest("recordUpdate/newRecord",newRecordsParams,function(newRecordRef) {
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
		north__showOverflowOnHover:	true,
		south__showOverflowOnHover:	true
	})
	
	$('#eastFilterSortPane').layout({
		inset: zeroPaddingInset,
		center__size:.6,
		south__size:.4,
	})
	
}

function initAfterViewFormComponentsAlreadyLoaded() {
	initRecordFilterPanel(tableID,reloadSortedAndFilterRecords)
	
	initFormSortRecordsPane(reloadSortedAndFilterRecords)
	
	// Initially all the records are loaded. This can be refined through the filter panel (see above).
	// TODO - At some point, there should be default filters for forms.
	var getRecordsParams = {
		tableID:tableID,
		filterIDs:[],
		sortRules:[]} 
	reloadRecords(getRecordsParams)
}


$(document).ready(function() {	
	 
	initUILayoutPanes()
			
	initRecordButtonsBehavior()
	
	initFieldInfo( function () {
		loadFormComponents({
			formParentElemID: "#layoutCanvas",
			initTextBoxFunc: function(textBoxObjectRef) {			
				initTextBoxRecordEditBehavior(textBoxObjectRef)
			},
			initCheckBoxFunc: function(checkBoxObjectRef) {
				console.log("Init check box in view form")
				initCheckBoxRecordEditBehavior(checkBoxObjectRef)
			},
			initDatePickerFunc: function(datePickerObjectRef) {
				console.log("Init date picker in view form")
				initDatePickerRecordEditBehavior(datePickerObjectRef)
			},
			initHtmlEditorFunc: function(htmlEditorObjectRef) {
				console.log("Init html editor in view form")
				initHtmlEditorRecordEditBehavior(htmlEditorObjectRef)
			},
			initImageFunc: function(imageObjectRef) {
				console.log("Init image in view form")
				initImageRecordEditBehavior(imageObjectRef)
			},
			doneLoadingFormDataFunc: initAfterViewFormComponentsAlreadyLoaded	
		}); 
	})
	

}); // document ready
