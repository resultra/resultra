// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function loadDashboardView(pageLayout,databaseID, dashboardID) {


	hideSiblingsShowOne('#dashboardViewSidebarProps')
	hideSiblingsShowOne('#dashboardCanvas')
	
	// Initially hide the RHS sidebar. The sidebar is only shown when individual components are selected.
	theMainWindowLayout.hideRHSSidebar()
	

	var viewDashboardCanvasSelector = '#dashboardCanvas'
		
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
			theMainWindowLayout.openRHSSidebar()
		})
	}		

	var loadDashboardConfig = {
		dashboardContext: viewDashboardContext,
		doneLoadingDashboardDataFunc: function() {

			initObjectCanvasSelectionBehavior(viewDashboardCanvasSelector, function() {
				// Hide the RHS sidebar when the main canvas of the dashboard is selected.
				theMainWindowLayout.hideRHSSidebar()
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