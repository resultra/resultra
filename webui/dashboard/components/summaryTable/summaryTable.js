function summaryTableComponentHTML(summaryTableID) {
	
	// The actual chart is placed inside a "chartWrapper" div. The outer div is used by draggable and resizable to position 
	// and resize the bar chart within the dashboard canvas. If the chart is placed directly within the out div, there
	// is a conflict with the Google chart code disabling the resize behavor after the chart is refreshed.
	var containerHTML = ''+
	'<div class="layoutContainer dashboardBarChartComponent" id="'+ summaryTableID+'">' +
		'<table class="table" id="' + summaryTableID+'_table"></table>'+
	'</div>';
	return containerHTML
}


function initSummaryTableData(dashboardID,summaryTableData) {
	
//	drawBarChart(barChartData)
//	initBarChartEditBehavior(dashboardID, barChartData.barChart.barChartID)
	
	var $summaryTable = $('#'+summaryTableData.summaryTable.summaryTableID)
	
	$summaryTable.data("summaryTableRef",summaryTableData.summaryTable)
	
}