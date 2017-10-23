

function loadDashboardView(pageLayout,databaseID, dashboardID) {


	hideSiblingsShowOne('#dashboardViewSidebarProps')
	hideSiblingsShowOne('#dashboardCanvas')
	

	var viewDashboardCanvasSelector = '#dashboardCanvas'
	
	pageLayout.clearCenterContentArea()
	pageLayout.hideFooterLayout()
	

	viewDashboardContext = { 
				dashboardID:dashboardID,
		 		databaseID: databaseID} 
				
				
	
	var getDashboardParams = { dashboardID: dashboardID }	

	jsonAPIRequest("dashboard/getProperties",getDashboardParams,function(dashboardInfo) {
		pageLayout.setCenterContentHeader(dashboardInfo.name)		
	})
				
				
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

		},
		initSummaryValComponent: function($summaryVal,summaryValRef) {

			var viewConfig = summaryValViewDashboardConfig(summaryValRef)

			console.log("Init gauge component")
			
			initDashboardComponentViewBehavior($summaryVal,summaryValRef.summaryValID,viewConfig)

		}

	}

	loadDashboardData(loadDashboardConfig)
}