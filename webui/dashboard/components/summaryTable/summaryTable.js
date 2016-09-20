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
	
	
	$headerRow.append("<th>" + summaryTableData.groupedSummarizedVals.groupingLabel + "</th>")
	
	var summaryLabels = summaryTableData.groupedSummarizedVals.summaryLabels
	
	for(var summaryLabelIndex = 0; summaryLabelIndex < summaryLabels.length; summaryLabelIndex++) {
		$headerRow.append('<th>' + summaryLabels[summaryLabelIndex] + '</th>')
	}
	
	$tableHeader.append($headerRow)
	$summaryTable.append($tableHeader)
}

function populateSummaryTableRow($tableBody,dataRow) {
	
	var $tableRow = $("<tr></tr>")
	$tableRow.append("<td><strong>" + dataRow.groupLabel + "<strong></td>")
	
	for(var summaryValIndex = 0; summaryValIndex < dataRow.summaryVals.length; summaryValIndex++) {
		var summaryVal = dataRow.summaryVals[summaryValIndex]
		$tableRow.append("<td>" + summaryVal + "</td>")	
	}
	
	$tableBody.append($tableRow)
}


function populateSummaryTableRows($summaryTable,summaryTableData) {
	
	var $tableBody = $("<tbody></tbody>")
	
	var dataRows = summaryTableData.groupedSummarizedVals.groupedDataRows
	for(var dataRowIndex = 0; dataRowIndex < dataRows.length; dataRowIndex++) {
		populateSummaryTableRow($tableBody,dataRows[dataRowIndex])
	}
		
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