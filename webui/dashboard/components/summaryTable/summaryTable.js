function summaryTableTableElemID(summaryTableID) {
	return summaryTableID + "_table"
}

function summaryTableComponentHTML(summaryTableID) {
	
	var tableElemID = summaryTableTableElemID(summaryTableID)
	// The actual chart is placed inside a "chartWrapper" div. The outer div is used by draggable and resizable to position 
	// and resize the bar chart within the dashboard canvas. If the chart is placed directly within the out div, there
	// is a conflict with the Google chart code disabling the resize behavor after the chart is refreshed.
	var containerHTML = ''+
	'<div class="layoutContainer dashboardBarChartComponent" id="'+ summaryTableID+'">' +
		'<table class="table table-hover table-bordered" id="' + tableElemID+'"></table>'+
	'</div>';
	return containerHTML
}

function populateSummaryTableHeader($summaryTable,summaryTableData) {
	
	var $tableHeader = $("<thead></thead>")
	var $headerRow = $("<tr></tr>")
	
	
	$headerRow.append("<th>Row Header</th>")
	$headerRow.append("<th>Summary Col 1</th>")
	
	$tableHeader.append($headerRow)
	$summaryTable.append($tableHeader)
}

function populateSummaryTableRow($tableBody) {
	
	var $tableRow = $("<tr></tr>")
	$tableRow.append("<td><strong>Row Val<strong></td>")
	$tableRow.append("<td>Col Val</td>")	
	$tableBody.append($tableRow)
}


function populateSummaryTableRows($summaryTable,summaryTableData) {
	
	var $tableBody = $("<tbody></tbody>")
	
	populateSummaryTableRow($tableBody)
	populateSummaryTableRow($tableBody)
	populateSummaryTableRow($tableBody)
	
	$summaryTable.append($tableBody)
	
}


function initSummaryTableData(dashboardID,summaryTableData) {
	
//	drawBarChart(barChartData)
//	initBarChartEditBehavior(dashboardID, barChartData.barChart.barChartID)
	var tableElemID = summaryTableTableElemID(summaryTableData.summaryTable.summaryTableID)
	var $summaryTable = $('#'+tableElemID)


	var $tableTitle = $("<caption>Table title</caption>")
	$summaryTable.append($tableTitle)

	populateSummaryTableHeader($summaryTable,summaryTableData)
	populateSummaryTableRows($summaryTable,summaryTableData)
	
	setElemObjectRef(summaryTableData.summaryTableID,summaryTableData.summaryTable)
	
	$summaryTable.data("summaryTableRef",summaryTableData.summaryTable)
	
}