

function barChartContainerHTML() {
	
	// The actual chart is placed inside a "chartWrapper" div. The outer div is used by draggable and resizable to position 
	// and resize the bar chart within the dashboard canvas. If the chart is placed directly within the out div, there
	// is a conflict with the Google chart code disabling the resize behavor after the chart is refreshed.
	var containerHTML = ''+
	'<div class="layoutContainer dashboardBarChartComponent">' +
		'<div class="row">' +
			'<div class="col-sm-8">' +
				'<label class="barChartTitle">' + 'New Bar Chart' + '</label>' +
			'</div>' +
			'<div class="col-sm-4 componentHeaderButtons">' +
				componentHelpPopupButtonHTML() +
			'</div>' +
		'</div>' +	
		'<canvas class="dashboardChartWrapper"</canvas>'+
	'</div>';
	return containerHTML
}


function drawBarChart($barChart, barChartData) {
	
	
	var $chart = $barChart.find(".dashboardChartWrapper")
	
	var $title = $barChart.find(".barChartTitle")
	$title.text(barChartData.title)
	
	var summarizedVals = barChartData.groupedSummarizedVals
	
	var numberFormat = summarizedVals.summaryNumberFormats[0]
	
	
	var chartLabels = []
	var chartData = []
	for(var dataIndex in summarizedVals.groupedDataRows) {
		var rowData = summarizedVals.groupedDataRows[dataIndex]
		console.log("Adding row: label=" + rowData.label + " val=" + rowData.value)
		chartLabels.push(rowData.groupLabel)
		var summaryVal = rowData.summaryVals[0]
		chartData.push(summaryVal)
	}

	var myChart = new Chart($chart, {
	  type: 'bar',
	  data: {
	    labels: chartLabels,
	    datasets: [{
	      label: summarizedVals.groupingLabel,
	      data: chartData
	    }]
	  },
	  options: {
		  legend: {
		              display: false
		           },
           tooltips: {
              enabled: false
           },
		  title: {
		              display: false
		          },
		  scales: {
		    yAxes: [{
		      scaleLabel: {
		        display: true,
		        labelString: summarizedVals.summaryLabels[0]
		      },
			  ticks: {
			      // Use custom labels on the Y Axis
				  callback: function(value, index, values) {
					  return formatNumberValue(numberFormat,value)
				  }
			  	} // ticks
			  
		    }],
		    xAxes: [{
		      scaleLabel: {
		        display: true,
		        labelString: summarizedVals.groupingLabel
		      }
		    }]
		  }
	  }
	});
	
}

// Helper method for drawing the placholder bar chart when designing the dashboard.
function drawDesignModeDummyBarChart($barChart) {
		
	var dummyBarChartData = {
		"groupedSummarizedVals": {
			"groupedDataRows": [
				{
					"groupLabel": "A",
					"summaryVals": [
						3
					]
				},
				{
					"groupLabel": "B",
					"summaryVals": [
						2
					]
				}
			],
			"overallDataRow": {
				"groupLabel": "Overall",
				"summaryVals": [
					2
				]
			},
			"groupingLabel": "X Axis",
			"summaryLabels": [
				"Y Axis"
			],
			"summaryNumberFormats": [
				"general"
			]
		}
	}
		

	// Draw just the same as a real bar chart, but feedit dummy data
   	drawBarChart($barChart,dummyBarChartData)
}


function initBarChartData(dashboardID,$barChart, barChartData) {
	
	drawBarChart($barChart, barChartData)
}


