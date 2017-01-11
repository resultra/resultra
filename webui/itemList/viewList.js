

var currRecordSet;
var currGlobalVals;


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

var viewFormCanvasSelector = '#layoutCanvas'

function loadCurrRecordIntoLayout()
{
	recordRef = currRecordSet.currRecordRef()
	if(recordRef != null)
	{

		loadRecordIntoFormLayout(viewFormCanvasSelector,recordRef)
	
		// Update footer to reflect where the current record is in list of currently loaded records
		$('#recordNumLabel').text(currRecordSet.recPageLabel())
		
		// If the record changed, and one of the form components is already loaded, it needs to be 
		// re-selected so the sidebar can be re-initialized with any settings specific to this 
		// record.
		reselectCurrentObjectSelection()
		
	} // if current record != null
}

function loadFormData(reloadRecordParams, formDataCallback) {
	var numDataSetsRemainingToLoad = 2
	
	var formData =  {}
	
	function oneDataSetLoaded() {
		numDataSetsRemainingToLoad -= 1
		if(numDataSetsRemainingToLoad <= 0) {
			formDataCallback(formData)
		}
	}
	
	jsonAPIRequest("recordRead/getFilteredSortedRecordValues",reloadRecordParams,function(recordsData) {
		formData.recordData = recordsData
		oneDataSetLoaded()
	})
	
	var globalParams = { parentDatabaseID: viewListContext.databaseID }
	jsonAPIRequest("global/getValues",globalParams,function(globalVals) {
		formData.globalVals = globalVals
		oneDataSetLoaded()
	})
	
}


function reloadRecords(reloadParams) {
	
	
	loadFormData(reloadParams,function(formData) {
		currGlobalVals = formData.globalVals	
		currRecordSet = new RecordSet(formData.recordData);
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
		
	})	
}


function createNewRecord() {
	var newRecordsParams = {parentDatabaseID:viewListContext.databaseID}
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
	var mainLayout = $('#layoutPage').layout({
		inset: zeroPaddingInset,
		north: fixedUILayoutPaneParams(40),
		east: {
			size: 300,
			resizable:false,
			slidable: false,
			spacing_open:16,
			spacing_closed:16,
			togglerClass:			"toggler",
			togglerLength_open:	128,
			togglerLength_closed: 128,
			togglerAlign_closed: "middle",	// align to top of resizer
			togglerAlign_open: "middle"		// align to top of resizer
			
		},
		west: {
			size: 250,
			resizable:false,
			slidable: false,
			spacing_open:4,
			spacing_closed:4,
			initClosed:true // panel is initially closed	
		}
	})
	
	$('#recordsPane').layout({
		north: fixedUILayoutPaneAutoSizeToFitContentsParams(),
		south: fixedUILayoutPaneAutoSizeToFitContentsParams(),
		north__showOverflowOnHover:	true,
		south__showOverflowOnHover:	true
	})
			
	initButtonClickHandler("#viewTableOfContentsMenuButton", function() {
		console.log("TOC button clicked")
		mainLayout.toggle("west")
	})
}



function initAfterViewFormComponentsAlreadyLoaded() {
	
	var getListParams = {
		listID: viewListContext.listID
	}
	
	
	
	jsonAPIRequest("itemList/get",getListParams,function(listInfo) {
		
		var filterPanelElemPrefix = "form_"
		
			
		function reloadSortedAndFilterRecords()
		{
			var filterRules = getRecordFilterRuleListRules(filterPanelElemPrefix)		
			var sortRules = getSortPaneSortRules()
	
			var getFilteredRecordsParams = { 
				databaseID: viewListContext.databaseID,
				filterRules: filterRules,
				sortRules: sortRules}
	
			reloadRecords(getFilteredRecordsParams)
		}
		
		var panelInitRemaining = 2
		function decrementRemainingPanelInitCount() {
			panelInitRemaining--
			if(panelInitRemaining <= 0) {
				reloadSortedAndFilterRecords()	
			}
		}
		
		var filterPropertyPanelParams = {
			elemPrefix: filterPanelElemPrefix,
			databaseID: viewListContext.databaseID,
			defaultFilterRules: listInfo.properties.defaultFilterRules,
			initDone: decrementRemainingPanelInitCount,
			updateFilterRules: function (updatedFilterRules) {
				console.log("View form: filters changed - updating filtering")
				reloadSortedAndFilterRecords()
			},
			refilterWithCurrentFilterRules: function() {
				reloadSortedAndFilterRecords()
			}
		}
		initRecordFilterViewPanel(filterPropertyPanelParams)
						
		var recordSortPaneParams = {
			defaultSortRules: listInfo.properties.defaultRecordSortRules,
			databaseID: viewListContext.databaseID,
			resortFunc: reloadSortedAndFilterRecords,
			initDoneFunc: decrementRemainingPanelInitCount,
			saveUpdatedSortRulesFunc: function(sortRules) {} // no-op
		}
		initSortRecordsPane(recordSortPaneParams)

	})
}


$(document).ready(function() {	
	 
	initUILayoutPanes()
			
	initRecordButtonsBehavior()
	
	initUserDropdownMenu()
	
	var tocConfig = {
		databaseID: viewListContext.databaseID,
		newItemFormButtonFunc: openSubmitFormDialog
	}
	
	initDatabaseTOC(tocConfig)
	
	
		
	hideSiblingsShowOne('#listViewProps')
	initObjectCanvasSelectionBehavior(viewFormCanvasSelector, function() {
		hideSiblingsShowOne('#listViewProps')
	})
	
	function getCurrentRecord() {
		return  currRecordSet.currRecordRef()
	}
	function updateCurrentRecord(updatedRecordRef) {
		currRecordSet.updateRecordRef(updatedRecordRef)
		loadCurrRecordIntoLayout()
	}


	loadFormViewComponents(viewFormCanvasSelector,viewListContext,
		getCurrentRecord,updateCurrentRecord,
		initAfterViewFormComponentsAlreadyLoaded)
			
}); // document ready
