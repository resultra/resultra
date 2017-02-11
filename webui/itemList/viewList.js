

// TODO - Need to move currGlobalVals out of this file and/or figure something else out for retrieval
// and access to global values.
var currGlobalVals

var listItemController

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
				preFilterRules: listInfo.properties.preFilterRules,
				filterRules: filterRules,		
				sortRules: sortRules}
	
			listItemController.reloadRecords(getFilteredRecordsParams)
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
				
	initUserDropdownMenu()
	
	var tocConfig = {
		databaseID: viewListContext.databaseID,
		newItemFormButtonFunc: openSubmitFormDialog
	}
	
	initDatabaseTOC(tocConfig)
	
	hideSiblingsShowOne('#listViewProps')
	
	listItemController = new ListItemController()
	listItemController.populateListViewWithListItemContainers(initAfterViewFormComponentsAlreadyLoaded)
				
}); // document ready
