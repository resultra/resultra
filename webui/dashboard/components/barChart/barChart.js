

function barChartContainerHTML(barChartID) {
	
	// The actual chart is placed inside a "chartWrapper" div. The outer div is used by draggable and resizable to position 
	// and resize the bar chart within the dashboard canvas. If the chart is placed directly within the out div, there
	// is a conflict with the Google chart code disabling the resize behavor after the chart is refreshed.
	var containerHTML = ''+
	'<div class="dashboardItemDesignContainer dashboardBarChartContainer draggable resizable selectable" id="'+ barChartID+'">' +
		'<div id="' + barChartID+'_chart" class="dashboardChartWrapper"</div>'+
	'</div>';
	return containerHTML
}


function drawBarChart(barChartData) {
	
	var dataRows = [];
	for(var dataIndex in barChartData.dataRows) {
		var rowData = barChartData.dataRows[dataIndex]
		console.log("Adding row: label=" + rowData.label + " val=" + rowData.value)
		dataRows.push([rowData.label,rowData.value])
	}
	
	var dataTable = new google.visualization.DataTable();
	dataTable.addColumn('string',barChartData.xAxisTitle)
	dataTable.addColumn('number',barChartData.yAxisTitle)
	dataTable.addRows(dataRows)
	
	var barChartOptions = {
		title: barChartData.title,
		hAxis: {
			title: barChartData.xAxisTitle,
			minValue: 0
		},
		vAxis: {
			title: barChartData.yAxisTitle
		},
		legend: { position: 'none' },
		chartArea:{left:40,top:30,width:'85%',height:'75%'}
	};
	
	// Place the chart in an inner div - see the comment in barChartContainerHTML()
  	var chartContainerElem = document.getElementById(barChartData.barChartID+'_chart')
	var barChart = new google.visualization.ColumnChart(chartContainerElem);
	
	// Whenever new data is loaded into the bar chart and it is redrawn,
	// a new handler to redraw/refresh the chart is registered. This is needed
	// in cases where the user resizes the bar chart.
	$('#' + barChartData.barChartID).data("redrawFunc", function () {
		barChart.draw(dataTable, barChartOptions);
	})

	barChart.draw(dataTable, barChartOptions);
}

// Helper method for drawing the placholder bar chart when designing the dashboard.
function drawDesignModeDummyBarChart(placeholderID) {
		
	var dummyBarChartData = {
		barChartID: placeholderID,
		title:"Chart Title",
		xAxisTitle:"Grouped Field",
		yAxisTitle:"Summarized Field",
		dataRows:[
			{label:"A",value:1},
			{label:"B",value:2}]
	}

	// Draw just the same as a real bar chart, but feedit dummy data
   	drawBarChart(dummyBarChartData)
}

function initBarChartEditBehavior(barChartID)
{

	// While in edit mode, disable input on the container
	var barChartContainer = $('#'+barChartID)

	// TODO - This could be put into a common function, since these
	// properties should be the same for objects loaded with the page
	// and newly added objects.
	barChartContainer.draggable ({
		grid: [20, 20], // snap to a grid
		cursor: "move",
		containment: "parent",
		clone: "original",						
		stop: function(event, ui) {
				  var layoutPos = {
					  positionLeft: ui.position.left,
					  positionTop: ui.position.top,
				   };
		  
				  console.log("drag stop: id: " + event.target.id);
				  console.log("drag stop: new position: " + JSON.stringify(layoutPos));
				  
				  // TODO: send ajax request to reposition the container
  
		      } // stop function
	})
	  
	barChartContainer.resizable({
		aspectRatio: false,
		handles: 'e, w, n, s', // Only allow resizing horizontally
		minWidth: 300,
		minHeight: 300,
		grid: 20, // snap to grid during resize

		stop: function(event, ui) {  
				  
  				var resizedGeometry = { positionTop: ui.position.top,positionLeft: ui.position.left,
					sizeWidth: ui.size.width, sizeHeight: ui.size.height }
				
				barChartContainer.data("redrawFunc")()
				
				var barChartRef = barChartContainer.data("barChartRef")
					
				var resizeParams = {
					uniqueID: {
						parentDashboardID: barChartRef.parentDashboardID,
						barChartID: barChartRef.barChartID
					},
					geometry: resizedGeometry
				}	
						
			 	jsonAPIRequest("dashboard/barChart/setDimensions",resizeParams,function(updatedBarChartRef) {
			 		console.log("bar chart dimensions updated")
			 	})	
				  
			  } // stop function
	});
	
	
	initObjectSelectionBehavior(barChartContainer, "#dashboardCanvas",function(barChartID) {
		var barChartPropsArgs = {
			dashboardID: dashboardID,
			barChartID: barChartID,
			
			propertyUpdateComplete: function (updatedBarChartRef) {
				
				var updateContainer = $('#'+updatedBarChartRef.barChartID)
				updateContainer.data("barChartRef",updatedBarChartRef)
				
				var getDataParams = {
					parentDashboardID:updatedBarChartRef.parentDashboardID,
					barChartID:updatedBarChartRef.barChartID
				}
				jsonAPIRequest("dashboard/barChart/getData",getDataParams,function(updatedBarChartData) {
					console.log("Redrawing barchart after properties update")
					drawBarChart(updatedBarChartData) // redraw the chart
				})
			}
		}
		
		loadBarChartProperties(barChartPropsArgs)
		
	})
	
}



function initBarChartData(dashboardID,barChartData) {
	
	drawBarChart(barChartData)
	initBarChartEditBehavior(barChartData.barChart.barChartID)
	
	var barChartContainer = $('#'+barChartData.barChart.barChartID)
	
	barChartContainer.data("barChartRef",barChartData.barChart)
	
}


