function initItemListTableView($tableContainer, databaseID, tableID,initDoneCallback) {
	
	function createColDef(colInfo,fieldsByID) {
		if(colInfo.colType === 'numberInput') {
			var fieldID = colInfo.properties.fieldID
			var colDef = {
				data:'fieldValues.' + fieldID,
				defaultContent:'', // used when there is null or undefined data
				createdCell: function( cell, cellData, rowData, rowIndex, colIndex ) {
					
		
					var $numberInputContainer = $(cell).find('.layoutContainer')
					
					var currRecord = rowData
					function getCurrentRecord() {
						return currRecord
					}
	
					function updateCurrentRecord(updatedRecordRef) {
						currRecord = updatedRecordRef
						// Get the parent row and update all cells in the row to the updated record.
						var $parentRow = $(cell).closest('tr')
						$parentRow.find('.layoutContainer').each(function() {
							var $cellContainer = $(this)
							var viewConfig = $cellContainer.data("viewFormConfig")
							viewConfig.loadRecord($cellContainer,currRecord)
						})
					}
			
					var recordProxy = {
						changeSetID: MainLineFullyCommittedChangeSetID,
						getRecordFunc: getCurrentRecord,
						updateRecordFunc: updateCurrentRecord
					}
					
					
					var componentContext = {
						databaseID: databaseID,
						fieldsByID: fieldsByID
					}
					
					var $numberInputContainer = $(cell).find('.layoutContainer')
					setContainerComponentInfo($numberInputContainer,colInfo,colInfo.numberInputID)
					initNumberInputTableRecordEditBehavior($numberInputContainer,componentContext,recordProxy, colInfo)
					var viewConfig = $numberInputContainer.data("viewFormConfig")
					viewConfig.loadRecord($numberInputContainer,currRecord)
				},
				render: function(data, type, row, meta) {
					if (type==='display') {
						return numberInputTableCellContainerHTML()
					} else if (type==='filter') {
						return data
					} else {
						return data
					}
				}
			}
			return colDef
		} else {
			var colDef = {
				data:'fieldValues.' + colInfo.properties.fieldID,
				defaultContent:'' // used when there is null or undefined data
			}
			return colDef
		}
	}
	
	function getTableInfo(tableInfoCallback) {
		
		var numTableInfoRemaining = 2
		var tableInfo
		var fieldsByID
		
		function tableInfoReceived() {
			numTableInfoRemaining--
			if(numTableInfoRemaining <= 0) {
				tableInfoCallback(tableInfo,fieldsByID)
			}
		}
		
		var tableInfoParams = { tableID: tableID }
		jsonAPIRequest("tableView/getTableDisplayInfo",tableInfoParams,function(info) {
			tableInfo = info
			tableInfoReceived()
		})
		
		loadFieldInfo(databaseID,[fieldTypeAll],function(retrievedFieldsByID) {
			fieldsByID = retrievedFieldsByID
			tableInfoReceived()
		})
		
	}
	
	function populateTable(tableInfo,fieldsByID) {
		function tableHeader() {
	
			var $tableHeader = $("<thead></thead>")
			var $headerRow = $("<tr></tr>")
	
			$.each(tableInfo.cols,function(index,colInfo) {
				var $header = $('<th></th>')
				var fieldName = fieldsByID[colInfo.properties.fieldID].name
				$header.text(fieldName)
				$headerRow.append($header)
			})
		
			$tableHeader.append($headerRow)
			$tableHeader.find("th").css("background-color","lightGrey")
	
			return $tableHeader
		}
		
		$tableContainer.empty()
		
		var $tableElem = $('<table class="table table-hover table-bordered display"></table>')
		$tableElem.append(tableHeader())
		$tableContainer.append($tableElem)
		
		var dataCols = []
		$.each(tableInfo.cols,function(index,colInfo) {
			var colDataDef = createColDef(colInfo,fieldsByID)
			dataCols.push(colDataDef)
		})
		
		var dataTable = $tableElem.DataTable({
			destroy:true, // Destroy existing table before applying the options
			searching:false, // Hide the search box
			bInfo:false, // Hide the "Showing 1 of N Entries" below the footer
			paging:false,
			scrollY: '100px',
			scrollCollapse:true,
			columns:dataCols
		})
	
		var $scrollHead = $tableContainer.find(".dataTables_scrollHead")
// TODO - incorporate footer into the table.
//		var $scrollFoot = $tableContainer.find(".dataTables_scrollFoot")
		var $scrollBody = $tableContainer.find(".dataTables_scrollBody")
	
		// Set the color of the entire header and footer to match the color of
		// of the individual header and footer cells. Otherwise, the scroll bar
		// on the RHS of the table stands out.
//		$scrollFoot.css("background-color","lightGrey")
		$scrollHead.css("background-color","lightGrey")
		
		var scrollBodyHeight = $tableContainer.outerHeight() -
				$scrollHead.outerHeight() // - $scrollFoot.outerHeight()
		var scrollBodyHeightPx = scrollBodyHeight + 'px'
	
		$scrollBody.css('max-height', scrollBodyHeightPx);
		dataTable.draw() // force redraw
		
		
		initDoneCallback(dataTable)
		
	}
	
	getTableInfo(populateTable)
	
	
}