$(document).ready(function() {
				
	var paletteConfig = {
		draggableItemHTML: function(placeholderID) {
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
		},
		
		dropDestSelector: "#dashboardCanvas",
		paletteSelector: "#dashboardPaletteSidebar",
	}
	initDesignPalette(paletteConfig)			
						
	// Initialize the page layout
	$('#designDashboardPage').layout({
		north: fixedUILayoutPaneParams(40),
		east: fixedUILayoutPaneParams(300),
		west: fixedUILayoutPaneParams(200),
		west__showOverflowOnHover:	true
	})	  
	  
});
