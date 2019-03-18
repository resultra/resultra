// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function summaryTableTableElem() {
	return '<table class="table table-hover table-bordered display"></table>'
}

function summaryTableComponentHTML(summaryTableID) {
	
	var containerHTML = ''+
	'<div class="layoutContainer dashboardSummaryTableComponent">' +
		'<div class="row summaryTableHeader">' +
			'<div class="col-sm-10">' +
				'<div class="summaryTableTitle"></div>'+
			'</div>' +
			'<div class="col-sm-2 summaryTableButtons">' +
				componentHelpPopupButtonHTML() +
			'</div>' +
		'</div>' +
		'<div class="tableContainer">' + 
			summaryTableTableElem() +
		'</div>' +
	'</div>';
	return containerHTML
}

function populateSummaryTableHeader($summaryTable,summaryTableData) {
	
	var $tableHeader = $("<thead></thead>")
	var $headerRow = $("<tr></tr>")
	
	
	// A minimum width is needed on the header. This will be respected by the DataTables
	// plugin to ensure the formating of the summary table looks sensible and doesn't wrap
	// the grouping label unnecessarily.
	var $groupingHeader = $("<th>" + summaryTableData.groupedSummarizedVals.groupingLabel + "</th>")
	$groupingHeader.css("min-width","100px")
	
	$headerRow.append($groupingHeader)
	
	var summaryLabels = summaryTableData.groupedSummarizedVals.summaryLabels
	
	for(var summaryLabelIndex = 0; summaryLabelIndex < summaryLabels.length; summaryLabelIndex++) {
		$headerRow.append('<th>' + summaryLabels[summaryLabelIndex] + '</th>')
	}
	
	
	$tableHeader.append($headerRow)
	$tableHeader.find("th").css("background-color","lightGrey")
	
	$summaryTable.append($tableHeader)
}


function createOneSummaryTableRow(dataRow,numberFormats,dataRowIndex) {
	var $tableRow = $("<tr></tr>")
	
	
	var $groupTableCell = $("<td></td>")
	$groupTableCell.attr("data-order",dataRowIndex)
	$groupTableCell.append('<strong>' + dataRow.groupLabel + '</strong>')
	
	$tableRow.append($groupTableCell)
	console.log("Initializing summary table grouping label: " + $tableRow.html())
	
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
	
	var $footerRow = createOneSummaryTableRow(summaryTableData.groupedSummarizedVals.overallDataRow,numberFormats,null)
	$footerRow.find("td").css("background-color","lightGrey")
	$tableFooter.append($footerRow)
	
	$summaryTable.append($tableFooter)
	
}

function populateSummaryTableRow($tableBody,dataRow,numberFormats,dataRowIndex) {
	var $tableRow = createOneSummaryTableRow(dataRow,numberFormats,dataRowIndex)
	$tableBody.append($tableRow)
}

function populateSummaryTableRows($summaryTable,summaryTableData) {
	
	var $tableBody = $("<tbody></tbody>")
	
	var numberFormats = summaryTableData.groupedSummarizedVals.summaryNumberFormats
	
	var dataRows = summaryTableData.groupedSummarizedVals.groupedDataRows
	for(var dataRowIndex = 0; dataRowIndex < dataRows.length; dataRowIndex++) {
		populateSummaryTableRow($tableBody,dataRows[dataRowIndex],numberFormats,dataRowIndex)
	}	
	$summaryTable.append($tableBody)
	
}


function initSummaryTableData(dashboardID,$summaryTable, summaryTableData) {
	
	var $tableContainer = $summaryTable.find(".tableContainer")
	var $summaryTableHeader = $summaryTable.find(".summaryTableHeader")
	
	$tableContainer.empty()
	$tableContainer.append(summaryTableTableElem())
	
	var $tableElem = $summaryTable.find(".table")
	$tableElem.empty()


	var tableTitle = summaryTableData.summaryTable.properties.title
	var $tableTitleDiv = $summaryTable.find(".summaryTableTitle")
	$tableTitleDiv.empty()
	if (tableTitle !== null && tableTitle.length >0) {
		var $tableTitleLabel = $("<label>" + tableTitle + "</label>")
		$tableTitleDiv.append($tableTitleLabel)	
	}
	
	// Dynamically compute the summary table component header's height,
	// then set the top of the table container to be just below this height.
	// This computation needs to happen after setting the label.
	var headerHeightPx = $summaryTableHeader.outerHeight(true) + 'px'
	console.log("Summary table height: " + headerHeightPx)
	$tableContainer.css("top",headerHeightPx)
	

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
	
	$scrollBody.css('max-height', scrollBodyHeightPx);
	dataTable.draw() // force redraw
	
	// $tableContainer wraps the table as a whole. By preventing clicks from propagating higher than
	// $tableContainer, clicking on the table itself will not cause the overall dashboard component
	// to be selected, which would be distracting. The table's header can be clicked on to select
	// the table as a whole.
	$tableContainer.click(function (event){
		event.stopPropagation();
   	 	//   ... your code here
		return false;
	});
	
	
}