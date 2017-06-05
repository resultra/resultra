

// TODO - Need to move currGlobalVals out of this file and/or figure something else out for retrieval
// and access to global values.
var currGlobalVals

var listItemController
var tableViewController

function initUILayoutPanes(resizeCallback)
{
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }
	
	// Initialize the page layout
	var mainLayout = $('#layoutPage').layout({
		inset: zeroPaddingInset,
		north: fixedUILayoutPaneParams(40),
		onopen_end: function(pane, $pane, paneState, paneOptions) {
			if (pane === 'west' || pane === 'east') {
				resizeCallback()				
			}
		},
		onclose_end: function(pane, $pane, paneState, paneOptions) {
			if (pane === 'west' || pane === 'east') {
				resizeCallback()				
			}
		},
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
		onresize_end: function(pane, $pane, paneState, paneOptions) {
			if(pane === 'center'){
				// only propagate the resize event for the center/content pane
				console.log("resize triggered")
				resizeCallback()
			}
			
		},
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

function initAfterViewFormComponentsAlreadyLoaded(listInfo) {
	
		
	var filterPanelElemPrefix = "form_"
	var $formLayoutContainer = $('#formViewContainer')
	
	var $tableViewLayoutContainer = $('#tableViewContainer')
	$tableViewLayoutContainer.hide()
	
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
	
	
				
	function reloadSortedAndFilterRecords()
	{
		var filterRules = getRecordFilterRuleListRules(filterPanelElemPrefix)		
		var sortRules = getSortPaneSortRules()

		var getFilteredRecordsParams = { 
			databaseID: viewListContext.databaseID,
			preFilterRules: listInfo.properties.preFilterRules,
			filterRules: filterRules,		
			sortRules: sortRules}
			
		loadFormData(getFilteredRecordsParams,function(formData) {
			currGlobalVals = formData.globalVals	
			tableViewController.setRecordData(formData.recordData)
			listItemController.setRecordData(formData.recordData)
		})

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

	function updatePageSize(newPageSize) {
		listItemController.setPageSize(newPageSize)
	}
	
	function updateForm(newFormID) {
		listItemController.setForm(newFormID)
	}

	function updateViewConfig(viewOptions) {
		console.log("Updating item list view configuration: " + JSON.stringify(viewOptions))
		if(viewOptions.formID !== undefined) {
			listItemController.setFormAndPageSize(viewOptions.formID,viewOptions.pageSize)
			$formLayoutContainer.show()
//			$tableViewLayoutContainer.hide()
			$tableViewLayoutContainer.css("display","none")
		} else {
//			$tableViewLayoutContainer.show()
			$tableViewLayoutContainer.css("display","")
			$formLayoutContainer.hide()
	// TODO - Clear the form layout container
	//		$formLayoutContainer.empty()
			tableViewController.setTable(viewOptions.tableID)
			
		}
	}
	initItemListDisplayConfigPanel(listInfo,updatePageSize,updateForm)

	var limitSelectionToFormIDs = listInfo.properties.alternateForms.slice(0)
	limitSelectionToFormIDs.push(listInfo.formID)
	var itemListViewConfig = {
		setViewCallback: updateViewConfig,
		databaseID: viewListContext.databaseID,
		initialView: listInfo.properties.defaultView
	}
	initItemListViewSelection(itemListViewConfig)
	

}


$(document).ready(function() {	
	 
	function resizeListView() {
		console.log("Resizing list view")
		if (tableViewController !== undefined) {
			tableViewController.refresh()
		}
	}
	 
	 
	initUILayoutPanes(resizeListView)
				
	initUserDropdownMenu()
	
	var tocConfig = {
		databaseID: viewListContext.databaseID,
		newItemFormButtonFunc: openSubmitFormDialog
	}
	
	initDatabaseTOC(tocConfig)
	
	hideSiblingsShowOne('#listViewProps')
	
	var getListParams = {
		listID: viewListContext.listID
	}
	
	var $formViewContainer = $('#formViewContainer')
	var $tableViewContainer = $('#tableViewContainer')
	
	jsonAPIRequest("itemList/get",getListParams,function(listInfo) {
		listItemController = new ListItemController($formViewContainer,listInfo.properties.defaultPageSize)
		tableViewController = new ItemListTableViewController($tableViewContainer,viewListContext.databaseID)
		
		
		var defaultListContext = {
			databaseID: listInfo.parentDatabaseID,
			formID: listInfo.formID,
			listID: listInfo.listID
		}
		listItemController.populateListViewWithListItemContainers(defaultListContext,function() {
			initAfterViewFormComponentsAlreadyLoaded(listInfo)
		})
		
	})
	
	
	
				
}); // document ready
