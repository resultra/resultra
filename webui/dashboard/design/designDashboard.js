var dashboardPaletteItemsEditConfig = {
	paletteItemBarChart: barChartDashboardDesignConfig,
	paletteItemSummaryTable: summaryTableDashboardDesignConfig,
	paletteItemHeader: headerDashboardDesignConfig,
	paletteItemGauge: gaugeDashboardDesignConfig,
	paletteItemSummaryVal: summaryValDashboardDesignConfig
}


// TODO - designDashboardContext is used as a global by dashboard components in the dashboard designer.
// The code needs to be enhanced to eliminate dependencies on this global.
var designDashboardContext

function initDesignDashboardPageContent(pageContext,dashboardInfo) {
		
	designDashboardContext = { 
		dashboardID: dashboardInfo.dashboardID,
		dashboardName: dashboardInfo.name,
		databaseID: dashboardInfo.parentDatabaseID,
		isSingleUserWorkspace: pageContext.isSingleUserWorkspace }
		
	function showDashboardProperties() {
		// When first loading the dashboard in design mode, show the properties for the dashboard as a whole.
		initDesignDashboardProperties(designDashboardContext.dashboardID)
		hideSiblingsShowOne('#dashboardProps')
	}
		
	
		
	var layoutDesignConfig = createDashboardLayoutDesignConfig()
	
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
				initObjectGridEditBehavior(droppedItemInfo.droppedElem,componentEditConfig,layoutDesignConfig)
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
	$('#designDashboardLayoutContent').layout({
		north: fixedUILayoutPaneAutoSizeToFitContentsParams(),
		south: fixedUILayoutPaneAutoSizeToFitContentsParams(),
		// Important: The 'showOverflowOnHover' options give a higher
		// z-index to sidebars and other panels with controls, etc. Otherwise
		// popups and other controlls will not be shown on top of the rest
		// of the layout.
		north__showOverflowOnHover:	true,
		south__showOverflowOnHover:	true 
	})
						
	
	var loadDashboardConfig = {
		dashboardContext: designDashboardContext,
		doneLoadingDashboardDataFunc: function() {
			initObjectCanvasSelectionBehavior(dashboardDesignCanvasSelector, function() {
				showDashboardProperties()
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
	
	
	showDashboardProperties()
	
	  
}



function navigateToDashboardDesignerPageContent(pageContext,dashboardInfo) {
	console.log("navigating to form designer")
	
	function initDashboardDesignerContent(initDoneCallback) {
		
		var contentSectionsRemaining = 3
		function processOneSection() {
			contentSectionsRemaining--
			if (contentSectionsRemaining <=0) {
				initDoneCallback()
			}
		}
		
		const sidebarContentURL = '/admin/dashboard/designDashboardSidebarContent/' + dashboardInfo.dashboardID
		setRHSSidebarContent(sidebarContentURL, function() {
			processOneSection()
		})

		const settingsPageURL = '/admin/dashboard/designDashboardMainContent/' + dashboardInfo.dashboardID
		setSettingsPageContent(settingsPageURL,function() {
			processOneSection()
		})
		
		const offPageContentURL = '/admin/dashboard/designDashboardOffpageContent/' + dashboardInfo.dashboardID
		setMainWindowOffPageContent(offPageContentURL,function() {
			processOneSection()
		})
		
	}
	
	initDashboardDesignerContent(function() {
		theMainWindowLayout.showRHSSidebar()
		theMainWindowLayout.openRHSSidebar()
		initDesignDashboardPageContent(pageContext,dashboardInfo)
	})		
	
	
}

