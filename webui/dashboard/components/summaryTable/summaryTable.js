
function summaryTableComponentHTML(summaryTableID) {
	
	var containerHTML = ''+
	'<div class="layoutContainer dashboardSummaryTableComponent">' +
		'<div class="tableContainer">' + 
			'<table class="table table-hover table-bordered"></table>'+
		'</div>' +
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


function initSummaryTableData(dashboardID,$summaryTable, summaryTableData) {
	
	
	var $tableElem = $summaryTable.find(".table")
	$tableElem.empty()


	var tableTitle = summaryTableData.summaryTable.properties.title
	if (tableTitle !== null && tableTitle.length >0) {
		var $tableTitle = $("<caption>" + tableTitle + "</caption>")
		$tableElem.append($tableTitle)	
	}

	populateSummaryTableHeader($tableElem,summaryTableData)
	populateSummaryTableRows($tableElem,summaryTableData)
		
	$summaryTable.data("summaryTableRef",summaryTableData.summaryTable)
	
}