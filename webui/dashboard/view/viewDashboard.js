function initDashboardUILayoutPanes()
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
			spacing_open:4,
			spacing_closed:4,
			initClosed:true // panel is initially closed				
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
		north__showOverflowOnHover:	true,
		south__showOverflowOnHover:	true
	})
			
	initButtonClickHandler("#viewTableOfContentsMenuButton", function() {
		console.log("TOC button clicked")
		mainLayout.toggle("west")
	})
	
	var pageLayout = {
		closePropertyPanel: function() {
			mainLayout.close("east")
		},
		openPropertyPanel: function() {
			mainLayout.open("east")
		}
	}
	
	return pageLayout
	
}


$(document).ready(function() {	
	 
	var pageLayout = initDashboardUILayoutPanes()
			
	initUserDropdownMenu()
	
	var viewDashboardCanvasSelector = '#dashboardCanvas'
	
	var tocConfig = {
		databaseID: viewDashboardContext.databaseID,
		newItemFormButtonFunc: openSubmitFormDialog
	}
	
	initDatabaseTOC(tocConfig)
	
		
	function initDashboardComponentViewBehavior($component,componentID, viewDashboardConfig) {
		
		var $parentDashboardCanvas = $(viewDashboardCanvasSelector)	
		initObjectSelectionBehavior($component, 
				$parentDashboardCanvas,function(selectedComponentID) {
			console.log("dashboard view object selected: " + selectedComponentID)
			var selectedObjRef	= getContainerObjectRef($component)
			viewDashboardConfig.selectionFunc($component,selectedObjRef)
			pageLayout.openPropertyPanel()
		})
	}
		
	var loadDashboardConfig = {
		dashboardContext: viewDashboardContext,
		doneLoadingDashboardDataFunc: function() {
			
			initObjectCanvasSelectionBehavior(viewDashboardCanvasSelector, function() {
				pageLayout.closePropertyPanel()
				hideSiblingsShowOne('#dashboardViewProps')
			})
			
			
		},
		initBarChartComponent: function($barChart,barChartRef) {
			
			var barChartViewConfig = barChartViewDashboardConfig(barChartRef)
			
			console.log("Init bar chart component")
			
			initDashboardComponentViewBehavior($barChart,
					barChartRef.barChartID,barChartViewConfig)
		},
		initSummaryTableComponent: function($summaryTable,summaryTableRef) {
			
			var summaryTableViewConfig = summaryTableViewDashboardConfig(summaryTableRef)
			
			console.log("Init summary table component")
						
			initDashboardComponentViewBehavior($summaryTable,
					summaryTableRef.summaryTableID,summaryTableViewConfig)
			
		},
		initHeaderComponent: function($header,headerRef) {
			
			var viewConfig = headerViewDashboardConfig(headerRef)
			
			console.log("Init header component")
						
			initDashboardComponentViewBehavior($header,
					headerRef.headerID,viewConfig)
			
		},
		initGaugeComponent: function($gauge,gaugeRef) {
			
			var viewConfig = gaugeViewDashboardConfig(gaugeRef)
			
			console.log("Init gauge component")
						
			initDashboardComponentViewBehavior($gauge,gaugeRef.gaugeID,viewConfig)
			
		}
	}
	
	loadDashboardData(loadDashboardConfig)
		

}); // document ready