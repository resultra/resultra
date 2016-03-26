$(document).ready(function() {
				
	var paletteConfig = {
		draggableItemHTML: function(placeholderID,paletteItemID) {
			var containerHTML = ''+
				'<div class="ui-widget-content layoutContainer layoutField draggable resizable" id="'+placeholderID+'">' +
					'<div class="field">'+
						'<label>New Bar Chart</label>'+
						'<input type="text" name="symbol" class="layoutInput" placeholder="Enter">'+
					'</div>'+
				'</div>';
			return containerHTML
		},
		
		dropComplete: function(droppedItemInfo) {
			console.log("Dashboard design pallete: drop item: " + JSON.stringify(droppedItemInfo))
			
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
	
	initNewBarChartDialog()		
						
	// Initialize the page layout
	$('#designDashboardPage').layout({
		north: fixedUILayoutPaneParams(40),
		east: fixedUILayoutPaneParams(300),
		west: fixedUILayoutPaneParams(200),
		west__showOverflowOnHover:	true
	})	  
	  
});
