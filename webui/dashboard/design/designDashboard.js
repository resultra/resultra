var dashboardPaletteItemsEditConfig = {
	paletteItemBarChart: barChartDashboardDesignConfig,
	paletteItemSummaryTable: summaryTableDashboardDesignConfig
}

var dashboardDesignCanvasSelector = "#dashboardCanvas"


$(document).ready(function() {
							
	var paletteConfig = {
		draggableItemHTML: function(placeholderID,paletteItemID) {			
			return dashboardPaletteItemsEditConfig[paletteItemID].draggableHTMLFunc(placeholderID)
		},
		
		dropComplete: function(droppedItemInfo) {						
			
			var componentEditConfig = dashboardPaletteItemsEditConfig[droppedItemInfo.paletteItemID]
			
			setTimeout(function() {
				initObjectGridEditBehavior(droppedItemInfo.placeholderID,componentEditConfig)
				componentEditConfig.populatePlaceholderData(droppedItemInfo.placeholderID)
			}, 50);
			
			// "repackage" the dropped item paramaters for creating a new dashboard component, adding
			// the dashboard context to the parameter list.				
			var componentParams = {
				dashboardContext: designDashboardContext, 
				geometry: droppedItemInfo.geometry,
				placeholderComponentID: droppedItemInfo.placeholderID,
				finalizeLayoutIncludingNewComponentFunc: droppedItemInfo.finalizeLayoutIncludingNewComponentFunc
			};
				
			componentEditConfig.createNewComponentAfterDropFunc(componentParams)
			
		},
		
		dropDestSelector: "#dashboardCanvas",
		paletteSelector: "#dashboardPaletteSidebar",
	}
	
	initDesignPalette(paletteConfig)	
	
	initNewBarChartDialog(designDashboardContext)
						
	// Initialize the page layout
	$('#designDashboardPage').layout({
		north: fixedUILayoutPaneParams(40),
		east: fixedUILayoutPaneParams(300),
		west: fixedUILayoutPaneParams(200),
		west__showOverflowOnHover:	true
	})
	
	function initDashboardComponentDesignDashboardEditBehavior($component,componentID, designDashboardConfig) {
		console.log("initDashboardComponentDesignDashboardEditBehavior: component ID = " + componentID)
	
		initObjectGridEditBehavior(componentID,designDashboardConfig)
	
	
		initObjectSelectionBehavior($component, 
				dashboardDesignCanvasSelector,function(selectedComponentID) {
			console.log("dashboard design object selected: " + selectedComponentID)
			var selectedObjRef	= getElemObjectRef(selectedComponentID)
			designDashboardConfig.selectionFunc(selectedObjRef)
		})
		
		
	}
	
	google.charts.setOnLoadCallback(function() {
		var loadDashboardConfig = {
			dashboardContext: designDashboardContext,
			doneLoadingDashboardDataFunc: function() {
				initObjectCanvasSelectionBehavior(dashboardDesignCanvasSelector, function() {
					initDesignDashboardProperties(designDashboardContext.dashboardID)
					hideSiblingsShowOne('#dashboardProps')
				})
			},
			initBarChartComponent: function($barChart,barChartRef) {
				console.log("Init bar chart component")
				initDashboardComponentDesignDashboardEditBehavior($barChart,
						barChartRef.barChartID,barChartDashboardDesignConfig)
			},
			initSummaryTableComponent: function($summaryTable,summaryTableRef) {
				console.log("Init summary table component")
				initDashboardComponentDesignDashboardEditBehavior($summaryTable,
						summaryTableRef.summaryTableID,summaryTableDashboardDesignConfig)
			}
		}
		
		loadDashboardData(loadDashboardConfig)
	});
	  
});
