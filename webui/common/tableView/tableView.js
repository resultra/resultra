function initItemListTableView($tableContainer, databaseID, tableID,initDoneCallback) {
	
	function TableViewRecordProxy(initialRecord,$cell) {
		
		var currRecord = initialRecord
		
		function getCurrentRecord() {
			return currRecord
		}
		
		function updateAllTableRowCells(currRecord) {
			// Get the parent row and update all cells in the same row as the given cell.
			var $parentRow = $cell.closest('tr')
			$parentRow.find('.layoutContainer').each(function() {
				var $cellContainer = $(this)
				var viewConfig = $cellContainer.data("viewFormConfig")
				viewConfig.loadRecord($cellContainer,currRecord)
			})
		
		}
		
		function updateCurrentRecord(updatedRecordRef) {
			currRecord = updatedRecordRef
			updateAllTableRowCells(currRecord)
		}
		
		
		this.changeSetID = MainLineFullyCommittedChangeSetID
		this.getRecordFunc = getCurrentRecord
		this.updateRecordFunc = updateCurrentRecord
	}
	
	function createTableViewColDef(colInfo,fieldsByID,renderCellHTMLFunc,initContainerFunc) {
		var fieldID = colInfo.properties.fieldID
		var colDef = {
			data:'fieldValues.' + fieldID,
			defaultContent:'', // used when there is null or undefined data
			createdCell: function( cell, cellData, rowData, rowIndex, colIndex ) {
				
				var $cellContainer = $(cell).find('.layoutContainer')
				
				var recordProxy = new TableViewRecordProxy(rowData,$(cell))
								
				var componentContext = {
					databaseID: databaseID,
					fieldsByID: fieldsByID
				}
				
				initContainerFunc(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext)
								
				var viewConfig = $cellContainer.data("viewFormConfig")
				viewConfig.loadRecord($cellContainer,recordProxy.getRecordFunc())
			},
			render: function(data, type, row, meta) {
				if (type==='display') {
					return renderCellHTMLFunc()
				} else {
					return data
				}
			}
		}
		return colDef
	}
	
	
	function createNumberInputColDef(colInfo,fieldsByID) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
			setContainerComponentInfo($cellContainer,colInfo,colInfo.numberInputID)
			initNumberInputTableRecordEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,numberInputTableCellContainerHTML,initContainer)
	}

	function createTextInputColDef(colInfo,fieldsByID) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
				setContainerComponentInfo($cellContainer,colInfo,colInfo.textInputID)
				initTextBoxRecordEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,textBoxTableViewContainerHTML,initContainer)
	}

	function createDateInputColDef(colInfo,fieldsByID) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
			setContainerComponentInfo($cellContainer,colInfo,colInfo.datePickerID)
			initTableViewDatePickerEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,datePickerTableViewCellContainerHTML,initContainer)
	}

	
	function createColDef(colInfo,fieldsByID) {
		switch (colInfo.colType) {
		case 'numberInput':
			return createNumberInputColDef(colInfo,fieldsByID)
		case 'textInput':
			return createTextInputColDef(colInfo,fieldsByID)
		case 'datePicker':
			return createDateInputColDef(colInfo,fieldsByID)
		default:
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
		
		var $tableElem = $('<table class="table table-hover table-bordered display tableView"></table>')
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