

// Helper method for drawing the placholder bar chart when designing the dashboard.
function drawDesignModeDummyBarChart(placeholderID) {
	
	console.log("Drawing dummy bar chart: " + placeholderID)

	 var dummyData = google.visualization.arrayToDataTable([
		['Grouped Values', 'Summarized Values', ],
		['A', 1],
		['B', 2.5],
	]);

	var barChartOptions = {
		title: 'Chart Title',
		hAxis: {
			title: 'Grouped Values',
			minValue: 0
		},
		vAxis: {
			title: 'Summarized Values'
		}
	};

   	var chartContainerElem = document.getElementById(placeholderID)
	var barChart = new google.visualization.ColumnChart(chartContainerElem);
	  
	barChart.draw(dummyData, barChartOptions);
}



function drawBarChart(barChartID, barChartData) {
	
	var dataRows = [];
	for(var dataIndex in barChartData.dataRows) {
		var rowData = barChartData.dataRows[dataIndex]
		console.log("Adding row: " + rowData.label + " " + rowData.value)
		dataRows.push([rowData.label,rowData.value])
	}
	
	var dataTable = new google.visualization.DataTable();
	dataTable.addColumn('string',"Grouped Values")
	dataTable.addColumn('number',"Summarized Values")
	dataTable.addRows(dataRows)
	
	console.log("Drawing dummy bar chart: " + placeholderID)

	var barChartOptions = {
		title: 'Chart Title',
		hAxis: {
			title: 'Grouped Values',
			minValue: 0
		},
		vAxis: {
			title: 'Summarized Values'
		}
	};

   	var chartContainerElem = document.getElementById(barChartID)
	var barChart = new google.visualization.ColumnChart(chartContainerElem);
	  
	barChart.draw(dataTable, barChartOptions);
}
