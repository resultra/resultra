function loadDashboardData()
{
	// Load the dashboard data
	var loadBarChartDataParams = { dashboardID: dashboardID }
	jsonAPIRequest("getDashboardData",loadBarChartDataParams,function(dashboardData) {
		
		for (var barChartDataIndex in dashboardData.barChartsData) {
			var barChartData = dashboardData.barChartsData[barChartDataIndex]
			console.log ("Loading bar chart: id = " + barChartData.barChartID)
			
			var barChartHTML = barChartContainerHTML(barChartData.barChartID);
			var barChartElem = $(barChartHTML)
			
			$("#dashboardCanvas").append(barChartElem)
			setElemGeometry(barChartElem,barChartData.barChartRef.geometry)
			
			initBarChartData(dashboardID,barChartData);			
			
		}
		
		// jQuery UI selectable and draggable conflict with one another for click handling, so there is specialized
		// click handling for the selection and deselection of individual dashboard elements. When a click is made
		// on the canvas, all the items are deselected.
		$( "#dashboardCanvas").click(function(e) {
			console.log("click on dashboard canvas")
	        $( "#dashboardCanvas > div" ).removeClass("ui-selected");
			
			// Toggle to the overall dashboard properties, hiding the other property panels
			hideSiblingsShowOne('#dashboardProps')
		})
		$( "#dashboardProps" ).accordion();	
						
	})
	
}

$(document).ready(function() {
							
	var paletteConfig = {
		draggableItemHTML: function(placeholderID,paletteItemID) {
			var containerHTML = barChartContainerHTML(placeholderID);
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
	
	google.charts.setOnLoadCallback(loadDashboardData);
	  
});
