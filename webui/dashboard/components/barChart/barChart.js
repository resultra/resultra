

function barChartContainerHTML() {
	
	// The actual chart is placed inside a "chartWrapper" div. The outer div is used by draggable and resizable to position 
	// and resize the bar chart within the dashboard canvas. If the chart is placed directly within the out div, there
	// is a conflict with the Google chart code disabling the resize behavor after the chart is refreshed.
	var containerHTML = ''+
	'<div class="layoutContainer dashboardBarChartComponent">' +
		'<canvas class="dashboardChartWrapper"</canvas>'+
	'</div>';
	return containerHTML
}


function drawBarChart($barChart, barChartData) {
	
	
	var $chart = $barChart.find(".dashboardChartWrapper")
	
	var chartLabels = []
	var chartData = []
	for(var dataIndex in barChartData.dataRows) {
		var rowData = barChartData.dataRows[dataIndex]
		console.log("Adding row: label=" + rowData.label + " val=" + rowData.value)
		chartLabels.push(rowData.label)
		chartData.push(rowData.value)
	}

	var myChart = new Chart($chart, {
	  type: 'bar',
	  data: {
	    labels: chartLabels,
	    datasets: [{
	      label: barChartData.xAxisTitle,
	      data: chartData
	    }]
	  },
	  options: {
		  title: {
		              display: true,
		              text: barChartData.title
		          },
		  scales: {
		    yAxes: [{
		      scaleLabel: {
		        display: true,
		        labelString: barChartData.yAxisTitle
		      }
		    }],
		    xAxes: [{
		      scaleLabel: {
		        display: true,
		        labelString: barChartData.xAxisTitle
		      }
		    }]
		  }
	  }
	});
	
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


