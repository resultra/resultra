

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
	
	var globalParams = { parentDatabaseID: viewFormContext.databaseID }
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
		listID: viewFormContext.listID
	}
	
	
	
	jsonAPIRequest("itemList/get",getListParams,function(listInfo) {
		
		var filterPanelElemPrefix = "form_"
		
			
		function reloadSortedAndFilterRecords()
		{
			var filterRules = getRecordFilterRuleListRules(filterPanelElemPrefix)		
			var sortRules = getSortPaneSortRules()
	
			var getFilteredRecordsParams = { 
				tableID: viewFormContext.tableID,
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
			tableID: listInfo.parentTableID,
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
			tableID: listInfo.parentTableID,
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
	
	initDatabaseTOC(viewFormContext.databaseID)
	
	var viewFormCanvasSelector = '#layoutCanvas'
	
	function initFormComponentViewBehavior($component,componentID, selectionFunc) {	
		initObjectSelectionBehavior($component, 
				viewFormCanvasSelector,function(selectedComponentID) {
			console.log("Form view object selected: " + selectedComponentID)
			var selectedObjRef	= getElemObjectRef(selectedComponentID)
			selectionFunc(selectedObjRef)
		})
	}
	
	hideSiblingsShowOne('#formViewProps')
	initObjectCanvasSelectionBehavior(viewFormCanvasSelector, function() {
		hideSiblingsShowOne('#formViewProps')
	})

	
	loadFormComponents({
		formParentElemID: viewFormCanvasSelector,
		formContext: viewFormContext,
		initTextBoxFunc: function(componentContext,$textBox,textBoxObjectRef) {			
			initTextBoxRecordEditBehavior(componentContext,textBoxObjectRef)
			initFormComponentViewBehavior($textBox,
					textBoxObjectRef.textBoxID,initTextBoxViewProperties)
		},
		initSelectionFunc: function(componentContext,$selection,selectionObjectRef) {			
			initSelectionRecordEditBehavior(componentContext,selectionObjectRef)
			initFormComponentViewBehavior($selection,
					selectionObjectRef.selectionID,initSelectionViewProperties)
		},
		initCheckBoxFunc: function(componentContext,$checkBox,checkBoxObjectRef) {
			console.log("Init check box in view form")
			initCheckBoxRecordEditBehavior(componentContext,checkBoxObjectRef)
			initFormComponentViewBehavior($checkBox,
					checkBoxObjectRef.checkBoxID,initCheckBoxViewProperties)
		},
		initRatingFunc: function(componentContext,$rating,ratingObjectRef) {
			console.log("Init rating in view form")
			initRatingRecordEditBehavior(componentContext,ratingObjectRef)
			initFormComponentViewBehavior($rating,
					ratingObjectRef.ratingID,initRatingViewProperties)
		},
		initUserSelectionFunc: function(componentContext,$userSelection,userSelectionObjectRef) {
			console.log("Init user selection in view form")
			initUserSelectionRecordEditBehavior(componentContext,userSelectionObjectRef)
			initFormComponentViewBehavior($userSelection,
					userSelectionObjectRef.userSelectionID,initUserSelectionViewProperties)
		},
		initDatePickerFunc: function(componentContext,$datePicker,datePickerObjectRef) {
			console.log("Init date picker in view form")
			initDatePickerRecordEditBehavior(componentContext,datePickerObjectRef)
			initFormComponentViewBehavior($datePicker,
					datePickerObjectRef.datePickerID,initDatePickerViewProperties)
		},
		initHtmlEditorFunc: function(componentContext,$htmlEditor,htmlEditorObjectRef) {
			console.log("Init html editor in view form")
			initHtmlEditorRecordEditBehavior(componentContext,htmlEditorObjectRef)
			initFormComponentViewBehavior($htmlEditor,
					htmlEditorObjectRef.htmlEditorID,initHTMLEditorViewProperties)
		},
		initImageFunc: function(componentContext,$image,imageObjectRef) {
			console.log("Init image in view form")
			initImageRecordEditBehavior(componentContext,imageObjectRef)
			initFormComponentViewBehavior($image,
					imageObjectRef.imageID,initImageViewProperties)
		},
		initHeaderFunc: function(componentContext,headerObjectRef) {
			console.log("Init header in view form")
			initHeaderRecordEditBehavior(componentContext,headerObjectRef)
		},
		doneLoadingFormDataFunc: initAfterViewFormComponentsAlreadyLoaded	
	}); 	
		
}); // document ready
