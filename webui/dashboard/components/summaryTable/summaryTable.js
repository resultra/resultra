
function summaryTableComponentHTML(summaryTableID) {
	
	var containerHTML = ''+
	'<div class="layoutContainer dashboardSummaryTableComponent">' +
		'<div class="summaryTableTitle"></div>'+
		'<div class="tableContainer">' + 
			'<table class="table table-hover table-bordered display"></table>'+
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
	$tableHeader.find("th").css("background-color","lightGrey")
	
	$summaryTable.append($tableHeader)
}


function createOneSummaryTableRow(dataRow,numberFormats) {
	var $tableRow = $("<tr></tr>")
	$tableRow.append("<td><strong>" + dataRow.groupLabel + "<strong></td>")
	
	for(var summaryValIndex = 0; summaryValIndex < dataRow.summaryVals.length; summaryValIndex++) {
		var summaryVal = dataRow.summaryVals[summaryValIndex]
		var colNumberFormat = numberFormats[summaryValIndex]
		var formattedVal = formatNumberValue(colNumberFormat,summaryVal)
		$tableRow.append("<td>" + formattedVal + "</td>")	
	}
	return $tableRow
}

function populateSummaryTableFooter($summaryTable,summaryTableData) {
	var $tableFooter = $("<tfoot></tfoot>")
	
	var numberFormats = summaryTableData.groupedSummarizedVals.summaryNumberFormats
	
	var $footerRow = createOneSummaryTableRow(summaryTableData.groupedSummarizedVals.overallDataRow,numberFormats)
	$footerRow.find("td").css("background-color","lightGrey")
	$tableFooter.append($footerRow)
	
	$summaryTable.append($tableFooter)
	
}

function populateSummaryTableRow($tableBody,dataRow,numberFormats) {
	var $tableRow = createOneSummaryTableRow(dataRow,numberFormats)
	$tableBody.append($tableRow)
}

function populateSummaryTableRows($summaryTable,summaryTableData) {
	
	var $tableBody = $("<tbody></tbody>")
	
	var numberFormats = summaryTableData.groupedSummarizedVals.summaryNumberFormats
	
	var dataRows = summaryTableData.groupedSummarizedVals.groupedDataRows
	for(var dataRowIndex = 0; dataRowIndex < dataRows.length; dataRowIndex++) {
		populateSummaryTableRow($tableBody,dataRows[dataRowIndex],numberFormats)
	}	
	$summaryTable.append($tableBody)
	
}


function initSummaryTableData(dashboardID,$summaryTable, summaryTableData) {
	
	var $tableElem = $summaryTable.find(".table")
	$tableElem.empty()


	var tableTitle = summaryTableData.summaryTable.properties.title
	var $tableTitleDiv = $summaryTable.find(".summaryTableTitle")
	if (tableTitle !== null && tableTitle.length >0) {
		var $tableTitleLabel = $("<label>" + tableTitle + "</label>")
		$tableTitleDiv.append($tableTitleLabel)	
	}

	populateSummaryTableHeader($tableElem,summaryTableData)
	populateSummaryTableFooter($tableElem,summaryTableData)
		
	populateSummaryTableRows($tableElem,summaryTableData)
		
	$summaryTable.data("summaryTableRef",summaryTableData.summaryTable)
	
	var dataTable = $tableElem.DataTable({
		destroy:true, // Destroy existing table before applying the options
		searching:false, // Hide the search box
		bInfo:false, // Hide the "Showing 1 of N Entries" below the footer
		paging:false,
		scrollY: '100px',
		scrollCollapse:true,
	})
	
	var $tableContainer = $summaryTable.find(".tableContainer")
	var $scrollHead = $tableContainer.find(".dataTables_scrollHead")
	var $scrollFoot = $tableContainer.find(".dataTables_scrollFoot")
	var $scrollBody = $tableContainer.find(".dataTables_scrollBody")
	
	// Set the color of the entire header and footer to match the color of
	// of the individual header and footer cells. Otherwise, the scroll bar
	// on the RHS of the table stands out.
	$scrollFoot.css("background-color","lightGrey")
	$scrollHead.css("background-color","lightGrey")
	
	
	var scrollBodyHeight = $tableContainer.outerHeight() -
			$scrollHead.outerHeight() - $scrollFoot.outerHeight()
	var scrollBodyHeightPx = scrollBodyHeight + 'px'
	
	$tableElem.closest('.dataTables_scrollBody').css('max-height', scrollBodyHeightPx);
	dataTable.draw()
	
}