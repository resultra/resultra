

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
		legend: { position: 'none' }
	};

   	var chartContainerElem = document.getElementById(barChartData.barChartID)
	var barChart = new google.visualization.ColumnChart(chartContainerElem);
	  
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

