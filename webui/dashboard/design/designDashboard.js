var dashboardPaletteItemsEditConfig = {
	paletteItemBarChart: barChartDashboardDesignConfig,
	paletteItemSummaryTable: summaryTableDashboardDesignConfig,
	paletteItemHeader: headerDashboardDesignConfig,
	paletteItemGauge: gaugeDashboardDesignConfig,
	paletteItemSummaryVal: summaryValDashboardDesignConfig
}


$(document).ready(function() {
	
	initUserDropdownMenu()
	initAlertHeader(designDashboardContext.databaseID)
	
							
	var paletteConfig = {
		draggableItemHTML: function(placeholderID,paletteItemID) {			
			return dashboardPaletteItemsEditConfig[paletteItemID].draggableHTMLFunc(placeholderID)
		},
		
		initDummyDragAndDropComponentContainer: function(paletteItemID, $paletteItemContainer) {
			// No-op - If a palette item needs to initialize the dragged item after it's
			// been inserted into the DOM, then code for that would go here.			
		},
		
		dropComplete: function(droppedItemInfo) {						
			
			var componentEditConfig = dashboardPaletteItemsEditConfig[droppedItemInfo.paletteItemID]
			
			var $componentContainer = droppedItemInfo.droppedElem
			
			setTimeout(function() {
				// TODO - need to pass "layoutDesignConfig" parameter to initObjectGridEditBehavior
				initObjectGridEditBehavior(droppedItemInfo.droppedElem,componentEditConfig)
				componentEditConfig.populatePlaceholderData($componentContainer)
			}, 50);
			
			// "repackage" the dropped item paramaters for creating a new dashboard component, adding
			// the dashboard context to the parameter list.				
			var componentParams = {
				dashboardContext: designDashboardContext, 
				geometry: droppedItemInfo.geometry,
				$componentContainer: $componentContainer,
				placeholderComponentID: droppedItemInfo.placeholderID,
				finalizeLayoutIncludingNewComponentFunc: droppedItemInfo.finalizeLayoutIncludingNewComponentFunc
			};
				
			componentEditConfig.createNewComponentAfterDropFunc(componentParams)
			
		},
		
		dropDestSelector: "#dashboardCanvas",
		paletteSelector: "#dashboardPaletteSidebar",
	}
	
	var designDashboardPaletteLayoutConfig = {
		parentLayoutSelector: dashboardDesignCanvasSelector,
		saveLayoutFunc: function() {} // no-op layout gets saved after component is finalized
	}
	
	initDesignPalette(paletteConfig,designDashboardPaletteLayoutConfig)	
	
	initNewBarChartDialog(designDashboardContext)
						
	// Initialize the page layout
	$('#designDashboardPage').layout({
		north: fixedUILayoutPaneParams(40),
		east: fixedUILayoutPaneParams(300),
		west: fixedUILayoutPaneParams(200),
		west__showOverflowOnHover:	true
	})
						
	var layoutDesignConfig = createDashboardLayoutDesignConfig()
	
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
					barChartRef.barChartID,barChartDashboardDesignConfig,layoutDesignConfig)
		},
		initSummaryTableComponent: function($summaryTable,summaryTableRef) {
			console.log("Init summary table component")
			initDashboardComponentDesignDashboardEditBehavior($summaryTable,
					summaryTableRef.summaryTableID,summaryTableDashboardDesignConfig,layoutDesignConfig)
		},
		initHeaderComponent: function($header,headerRef) {
			console.log("Init header component")
			initDashboardComponentDesignDashboardEditBehavior($header,
					headerRef.headerID,headerDashboardDesignConfig,layoutDesignConfig)
		},
		initGaugeComponent: function($gauge,gaugeRef) {
			console.log("Init gauge component")
			initDashboardComponentDesignDashboardEditBehavior($gauge,
					gaugeRef.gaugeID,gaugeDashboardDesignConfig,layoutDesignConfig)
		},
		initSummaryValComponent: function($summaryVal,summaryValRef) {
			console.log("Init summary value")
			initDashboardComponentDesignDashboardEditBehavior($summaryVal,
					summaryValRef.summaryValID,summaryValDashboardDesignConfig,layoutDesignConfig)
		}
	}
	
	loadDashboardData(loadDashboardConfig)
	
	// When first loading the dashboard in design mode, show the properties for the dashboard as a whole.
	initDesignDashboardProperties(designDashboardContext.dashboardID)
	hideSiblingsShowOne('#dashboardProps')
	  
});
