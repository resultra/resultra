$(document).ready(function() {
							
	var paletteConfig = {
		draggableItemHTML: function(placeholderID,paletteItemID) {
			var containerHTML = ''+
				'<div class="dashboardItemDesignContainer dashboardBarChartContainer draggable resizable" id="'+ placeholderID+'">' +
				'</div>';
				console.log("Container HTML: " + containerHTML)
			return containerHTML
		},
		
		dropComplete: function(droppedItemInfo) {
			console.log("Dashboard design pallete: drop item: " + JSON.stringify(droppedItemInfo))
			
			// At this point, the placholder div for the bar chart will have just been inserted. However, the DOM may 
			// not be completely updated at this point. To ensure this, a small delay is needed before
			// drawing the dummy bar charts. See http://goo.gl/IloNM for more.
			setTimeout(function() {drawDesignModeDummyBarChart(droppedItemInfo.placeholderID); }, 50);
			
			// "repackage" the dropped item paramaters for creating a new layout element. Also add the layoutID
			// to the parameters.
			var barChartParams = {
				parentDashboardID: dashboardID,
				geometry: droppedItemInfo.geometry,
				containerID: droppedItemInfo.placeholderID,
				};
			newBarChart(barChartParams)
			
		},
		
		dropDestSelector: "#dashboardCanvas",
		paletteSelector: "#dashboardPaletteSidebar",
	}
	initDesignPalette(paletteConfig)	
	
	initNewBarChartDialog(dashboardID)		
						
	// Initialize the page layout
	$('#designDashboardPage').layout({
		north: fixedUILayoutPaneParams(40),
		east: fixedUILayoutPaneParams(300),
		west: fixedUILayoutPaneParams(200),
		west__showOverflowOnHover:	true
	})	  
	  
});
