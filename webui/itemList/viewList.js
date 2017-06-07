

// TODO - Need to move currGlobalVals out of this file and/or figure something else out for retrieval
// and access to global values.
var currGlobalVals

var listItemController
var tableViewController
var itemListLayout


function initItemListView(listInfo) {
	
		
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


	function updateViewConfig(viewOptions) {
		console.log("Updating item list view configuration: " + JSON.stringify(viewOptions))
		if(viewOptions.formID !== undefined) {
			itemListLayout.showFooterLayout()
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
			itemListLayout.hideFooterLayout()
			
		}
	}

	var itemListViewConfig = {
		setViewCallback: updateViewConfig,
		databaseID: viewListContext.databaseID,
		initialView: listInfo.properties.defaultView,
		alternateViews: listInfo.properties.alternateViews
	}
	initItemListViewSelection(itemListViewConfig)
	
	// Perform an initial update of the view, based upon the default view.
	updateViewConfig(listInfo.properties.defaultView)
	

}


$(document).ready(function() {	
	 
	function resizeListView() {
		console.log("Resizing list view")
		if (tableViewController !== undefined) {
			tableViewController.refresh()
		}
	}
	 
	 
	itemListLayout = new ItemListLayout(resizeListView)
				
	initUserDropdownMenu()
	
	var tocConfig = {
		databaseID: viewListContext.databaseID,
		newItemFormButtonFunc: openSubmitFormDialog
	}	
	initDatabaseTOC(tocConfig)
	
	hideSiblingsShowOne('#listViewProps')
	
	var $formViewContainer = $('#formViewContainer')
	var $tableViewContainer = $('#tableViewContainer')
	
	initFieldInfo(viewListContext.databaseID, function() {
		var getListParams = { listID: viewListContext.listID }
		jsonAPIRequest("itemList/get",getListParams,function(listInfo) {
			listItemController = new ListItemController($formViewContainer)
			tableViewController = new ItemListTableViewController($tableViewContainer,viewListContext.databaseID)
			initItemListView(listInfo)		
		})	
		
	})
				
}); // document ready
