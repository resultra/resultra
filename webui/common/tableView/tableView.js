
function initItemListTableView(params) {
	
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
	
	function createTableViewColDef(colInfo,fieldsByID,
				renderCellHTMLFunc,initContainerFunc,percColWidths) {
		var fieldID = colInfo.properties.fieldID
					
		function columnSortType(fieldID,fieldsByID) {
			var fieldInfo = fieldsByID[fieldID]
			switch(fieldInfo.type) {
			case fieldTypeNumber:
				return 'custom-num'
			case fieldTypeText:
				return 'string'
			case fieldTypeBool:
				return 'custom-bool'
			case fieldTypeTime:
				return 'date'
			default:
				return 'string'
			}							
		}
		
		var colType = columnSortType(fieldID,fieldsByID)
								
		var colDef = {
			data:'fieldValues.' + fieldID,
			defaultContent:'', // used when there is null or undefined data
			createdCell: function( cell, cellData, rowData, rowIndex, colIndex ) {
				
				var $cellContainer = $(cell).find('.layoutContainer')
				
				var recordProxy = new TableViewRecordProxy(rowData,$(cell))
								
				var componentContext = {
					databaseID: params.databaseID,
					fieldsByID: fieldsByID
				}
				
				initContainerFunc(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext)
								
				var viewConfig = $cellContainer.data("viewFormConfig")
				viewConfig.loadRecord($cellContainer,recordProxy.getRecordFunc())
			},
			type: colType,
			render: function(data, type, row, meta) {
				if (type==='display') {
					return renderCellHTMLFunc()
				} else {
					return data
				}
			}
		}
		
		if(percColWidths.hasOwnProperty(colInfo.columnID)) {
			colDef.width = percColWidths[colInfo.columnID]
		}
		return colDef
	}
	
	
	function createNumberInputColDef(colInfo,fieldsByID,percColWidths) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
			setContainerComponentInfo($cellContainer,colInfo,colInfo.numberInputID)
			initNumberInputTableRecordEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				numberInputTableCellContainerHTML,initContainer,percColWidths)
	}

	function createTextInputColDef(colInfo,fieldsByID,percColWidths) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
				setContainerComponentInfo($cellContainer,colInfo,colInfo.textInputID)
				initTextBoxRecordEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				textBoxTableViewContainerHTML,initContainer,percColWidths)
	}

	function createDateInputColDef(colInfo,fieldsByID,percColWidths) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
			setContainerComponentInfo($cellContainer,colInfo,colInfo.datePickerID)
			initTableViewDatePickerEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				datePickerTableViewCellContainerHTML,initContainer,percColWidths)
	}

	function createCheckboxColDef(colInfo,fieldsByID,percColWidths) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
			setContainerComponentInfo($cellContainer,colInfo,colInfo.datePickerID)
			initTableViewCheckboxEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				checkBoxTableViewCellContainerHTML,initContainer,percColWidths)
	}

	function createToggleColDef(colInfo,fieldsByID,percColWidths) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
			setContainerComponentInfo($cellContainer,colInfo,colInfo.toggleID)
			initToggleTableCellRecordEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				toggleTableCellContainerHTML,initContainer,percColWidths)
	}


	function createRatingColDef(colInfo,fieldsByID,percColWidths) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
			setContainerComponentInfo($cellContainer,colInfo,colInfo.ratingID)
			initRatingTableCellRecordEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				ratingTableCellContainerHTML,initContainer,percColWidths)
	}

	
	function createColDef(colInfo,fieldsByID,percColWidths) {
		switch (colInfo.colType) {
		case 'numberInput':
			return createNumberInputColDef(colInfo,fieldsByID,percColWidths)
		case 'textInput':
			return createTextInputColDef(colInfo,fieldsByID,percColWidths)
		case 'datePicker':
			return createDateInputColDef(colInfo,fieldsByID,percColWidths)
		case 'checkbox':
			return createCheckboxColDef(colInfo,fieldsByID,percColWidths)
		case 'rating':
			return createRatingColDef(colInfo,fieldsByID,percColWidths)
		case 'toggle':
			return createToggleColDef(colInfo,fieldsByID,percColWidths)
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
		
		var tableInfoParams = { tableID: params.tableID }
		jsonAPIRequest("tableView/getTableDisplayInfo",tableInfoParams,function(info) {
			tableInfo = info
			tableInfoReceived()
		})
		
		loadFieldInfo(params.databaseID,[fieldTypeAll],function(retrievedFieldsByID) {
			fieldsByID = retrievedFieldsByID
			tableInfoReceived()
		})
		
	}
	
	function populateTable(tableInfo,fieldsByID) {
		
		// When displaying the table, use percentage widths instead of pixel widths. This allows
		// the column widths to scale up naturally as the table view is expanded.
		var percColWidths = {}
		var overallWidth = 0.0
		$.each(tableInfo.table.properties.colWidths,function(colID,width) {
			overallWidth += width
		})
		$.each(tableInfo.table.properties.colWidths,function(colID,width) {
			percColWidths[colID] = (width/overallWidth * 100).toFixed(2) + '%'
		})
		
		
		function tableHeader() {
	
			var $tableHeader = $("<thead></thead>")
			var $headerRow = $("<tr></tr>")
	
			$.each(tableInfo.cols,function(index,colInfo) {
				var $header = $('<th></th>')
				
				if(percColWidths.hasOwnProperty(colInfo.columnID)) {
					$header.css("width",percColWidths[colInfo.columnID])
				}
				
				setFormComponentLabel($header,colInfo.properties.fieldID,
						colInfo.properties.labelFormat)
				
				$headerRow.append($header)
			})
		
			$tableHeader.append($headerRow)
			$tableHeader.find("th").css("background-color","lightGrey")
	
			return $tableHeader
		}
		
		params.$tableContainer.empty()
		
		var $tableElem = $('<table class="table table-hover table-bordered display tableView"></table>')
		$tableElem.append(tableHeader())
		params.$tableContainer.append($tableElem)
		
		var dataCols = []
		$.each(tableInfo.cols,function(index,colInfo) {
			var colDataDef = createColDef(colInfo,fieldsByID,percColWidths)
			dataCols.push(colDataDef)
		})
		
		var dataTable = $tableElem.DataTable({
			destroy:true, // Destroy existing table before applying the options
			searching:false, // Hide the search box
			bInfo:false, // Hide the "Showing 1 of N Entries" below the footer
			paging:false,
			scrollY: '100px',
// TODO - Evaluation the use of horizontal scrolling through the scrollX option
// and possible other options. Currently the headers don't scroll with the body of
// table when horizontal scrolling is enabled.
//			scrollX:true,
			scrollCollapse:true,
			columns:dataCols
		})
		
		// If the table is reordered by the end-user, then synchronize the sort order with the side-bar's
		// sorting preferences.
		$tableElem.on('order.dt',function() {
			var order = dataTable.order();
			console.log("Table reordered: " + JSON.stringify(order))
			var sortRules = []
			$.each(order,function(index,orderInfo) {
				var colIndex = orderInfo[0]
				var sortDirection = orderInfo[1]
				var tableColInfo = tableInfo.cols[colIndex]
				var fieldID = tableColInfo.properties.fieldID
				sortRules.push({
					direction:sortDirection,
					fieldID: fieldID
				})
				
			})
			console.log("Table reordered: " + JSON.stringify(sortRules))
			
		})
		
	
		var $scrollHead = params.$tableContainer.find(".dataTables_scrollHead")
// TODO - incorporate footer into the table.
//		var $scrollFoot = params.$tableContainer.find(".dataTables_scrollFoot")
		var $scrollBody = params.$tableContainer.find(".dataTables_scrollBody")
	
		// Set the color of the entire header and footer to match the color of
		// of the individual header and footer cells. Otherwise, the scroll bar
		// on the RHS of the table stands out.
//		$scrollFoot.css("background-color","lightGrey")
		$scrollHead.css("background-color","lightGrey")
		
		function resizeToContainerHeight() {
			var scrollBodyHeight = params.$tableContainer.outerHeight() -
					$scrollHead.outerHeight() // TODO: after adding footer, also subtract footer height: - $scrollFoot.outerHeight()
			var scrollBodyHeightPx = scrollBodyHeight + 'px'
	
			$scrollBody.css('max-height', scrollBodyHeightPx);
			dataTable.draw() // force redraw
		}
		resizeToContainerHeight()
		
		params.initDoneCallback(dataTable,resizeToContainerHeight)
		
	}
	
	getTableInfo(populateTable)
	
	
}