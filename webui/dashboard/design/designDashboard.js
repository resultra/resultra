var dashboardPaletteItemsEditConfig = {
	paletteItemBarChart: barChartDashboardDesignConfig,
	paletteItemSummaryTable: summaryTableDashboardDesignConfig
}

var dashboardDesignCanvasSelector = "#dashboardCanvas"


function createDashboardLayoutDesignConfig() {
	// TODO - Move to function for dragging items
	function saveUpdatedDashboardComponentLayout(updatedLayout) {
		console.log("saveUpdatedDashboardComponentLayout: component layout = " + JSON.stringify(updatedLayout))		
		var setLayoutParams = {
			dashboardID: designDashboardContext.dashboardID,
			layout: updatedLayout
		}
		jsonAPIRequest("dashboard/setLayout", setLayoutParams, function(dashboardInfo) {})
		
	}
	
	var layoutDesignConfig = {
		parentLayoutSelector: dashboardDesignCanvasSelector,
		saveLayoutFunc: saveUpdatedDashboardComponentLayout
	}
	return layoutDesignConfig
}

$(document).ready(function() {
	
	initUserDropdownMenu()
	
							
	var paletteConfig = {
		draggableItemHTML: function(placeholderID,paletteItemID) {			
			return dashboardPaletteItemsEditConfig[paletteItemID].draggableHTMLFunc(placeholderID)
		},
		
		startPaletteDrag: function(placeholderID,paletteItemID,$paletteItemContainer) {
			// No-op - If a palette item needs to initialize the dragged item after it's
			// been inserted into the DOM, then code for that would go here.			
		},
		
		dropComplete: function(droppedItemInfo) {						
			
			var componentEditConfig = dashboardPaletteItemsEditConfig[droppedItemInfo.paletteItemID]
			
			setTimeout(function() {
				// TODO - need to pass "layoutDesignConfig" parameter to initObjectGridEditBehavior
				initObjectGridEditBehavior(droppedItemInfo.droppedElem,componentEditConfig)
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
	
	function initDashboardComponentDesignDashboardEditBehavior($component,componentID, designDashboardConfig,layoutDesignConfig) {
		console.log("initDashboardComponentDesignDashboardEditBehavior: component ID = " + componentID)
	
		initObjectGridEditBehavior($component,designDashboardConfig,layoutDesignConfig)
	
		// When in design mode, dashboard components need to have dotted lines as their border.
		$component.addClass("layoutDesignContainer")
		
		var $parentDashboardCanvas = $(dashboardDesignCanvasSelector)
		initObjectSelectionBehavior($component, 
				$parentDashboardCanvas,function(selectedComponentID) {
			console.log("dashboard design object selected: " + selectedComponentID)
			var selectedObjRef	= getElemObjectRef(selectedComponentID)
			designDashboardConfig.selectionFunc($component,selectedObjRef)
		})
		
		
	}
	
	google.charts.setOnLoadCallback(function() {
				
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
			}
		}
		
		loadDashboardData(loadDashboardConfig)
	});
	  
});
