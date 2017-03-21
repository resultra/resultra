

function barChartContainerHTML() {
	
	// The actual chart is placed inside a "chartWrapper" div. The outer div is used by draggable and resizable to position 
	// and resize the bar chart within the dashboard canvas. If the chart is placed directly within the out div, there
	// is a conflict with the Google chart code disabling the resize behavor after the chart is refreshed.
	var containerHTML = ''+
	'<div class="layoutContainer dashboardBarChartComponent">' +
		'<div class="dashboardChartWrapper"</div>'+
	'</div>';
	return containerHTML
}


function drawBarChart($barChart, barChartData) {
	
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
	
	var $chartWrapper = $barChart.find(".dashboardChartWrapper")
  	var chartContainerElem = $chartWrapper.get(0)
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
function drawDesignModeDummyBarChart($barChart) {
		
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
   	drawBarChart($barChart,dummyBarChartData)
}


function initBarChartData(dashboardID,$barChart, barChartData) {
	
	drawBarChart($barChart, barChartData)
}


