function initDashboardUILayoutPanes()
{
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }
	
	// Initialize the page layout
	var mainLayout = $('#layoutPage').layout({
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
	
	$('#dashboardPane').layout({
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


$(document).ready(function() {	
	 
	initDashboardUILayoutPanes()
			
	initUserDropdownMenu()
	
	var viewDashboardCanvasSelector = '#dashboardCanvas'
	
	initDatabaseTOC(viewDashboardContext.databaseID)
	
	google.charts.setOnLoadCallback(function() {
		
		function initDashboardComponentViewBehavior($component,componentID, viewDashboardConfig) {	
			initObjectSelectionBehavior($component, 
					viewDashboardCanvasSelector,function(selectedComponentID) {
				console.log("dashboard view object selected: " + selectedComponentID)
				var selectedObjRef	= getElemObjectRef(selectedComponentID)
				viewDashboardConfig.selectionFunc(selectedObjRef)
			})
		}
			
		var loadDashboardConfig = {
			dashboardContext: viewDashboardContext,
			doneLoadingDashboardDataFunc: function() {
				
				initObjectCanvasSelectionBehavior(viewDashboardCanvasSelector, function() {
	//				initViewDashboardProperties(viewDashboardContext.dashboardID)
					hideSiblingsShowOne('#dashboardViewProps')
				})
				
				
			},
			initBarChartComponent: function($barChart,barChartRef) {
				
				var barChartViewConfig = barChartViewDashboardConfig()
				
				console.log("Init bar chart component")
				initDashboardComponentViewBehavior($barChart,
						barChartRef.barChartID,barChartViewConfig)
			},
			initSummaryTableComponent: function($summaryTable,summaryTableRef) {
				
				var summaryTableViewConfig = summaryTableViewDashboardConfig()
				
				console.log("Init summary table component")
				initDashboardComponentViewBehavior($summaryTable,
						summaryTableRef.summaryTableID,summaryTableViewConfig)
				
			}
		}
		
		loadDashboardData(loadDashboardConfig)
	})
		

}); // document ready