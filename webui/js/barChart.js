

function drawBarChart(barChartData) {
	
	var dataRows = [];
	for(var dataIndex in barChartData.dataRows) {
		var rowData = barChartData.dataRows[dataIndex]
		console.log("Adding row: " + rowData.label + " " + rowData.value)
		dataRows.push([rowData.label,rowData.value])
	}
	
	var dataTable = new google.visualization.DataTable();
	dataTable.addColumn('string',barChartData.xAxisTitle)
	dataTable.addColumn('number',barChartData.yAxisTitle)
	dataTable.addRows(dataRows)
	
	console.log("Drawing dummy bar chart: " + placeholderID)

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
	
  	var chartContainerElem = document.getElementById(barChartData.barChartID)
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
				barChartRef.geometry = resizedGeometry
		
			 	jsonAPIRequest("updateBarChartProps",barChartRef,function(updatedBarChartRef) {
			 		
			 	})	
				  
			  } // stop function
	});
	
}

function initBarChart(barChartRef,barChartData) {
	
	drawBarChart(barChartData)
	initBarChartEditBehavior(barChartRef.barChartID)

	var barChartContainer = $('#'+barChartRef.barChartID)
	barChartContainer.data("barChartRef",barChartRef)
	
}


