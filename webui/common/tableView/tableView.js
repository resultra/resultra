
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


	function createUserSelectionColDef(colInfo,fieldsByID,percColWidths) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
			setContainerComponentInfo($cellContainer,colInfo,colInfo.numberInputID)
			initUserSelectionTableRecordEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				userSelectionTableCellContainerHTML,initContainer,percColWidths)
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
		case 'userSelection':
			return createUserSelectionColDef(colInfo,fieldsByID,percColWidths)
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
			paging:true, // pagination must be enabled for pageResize plug-in
			pageResize:true, // enable plug-in for vertical page resizing
			lengthChange:true, // needed for pageResize plug-in
			deferRender:true, // only create elements when required (needed with paging)
			columns:dataCols
		})
		
		
		// The order.dt event can be triggered when a table is redrawn. However, we only
		// want to propagate the order.dt event when the end-user clicks on a table 
		// header to re-sort the table.
		var drawInProgress = true
		$tableElem.on('draw.dt',function() {
			drawInProgress = false
		})
		
		// If the table is reordered by the end-user, then synchronize the sort order with the side-bar's
		// sorting preferences.
		$tableElem.on('order.dt',function(e,settings) {
			if (!drawInProgress) {
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
				params.resortCallback(sortRules)	
			}
		})
		
		
	
		var $scrollHead = params.$tableContainer.find(".dataTables_scrollHead")
		$scrollHead.css("background-color","lightGrey")
		
		function resizeToContainerHeight() {
			drawInProgress = true
			dataTable.draw() // force redraw
		}
		resizeToContainerHeight()
		
		function setSortOrder(sortRules) {
			var dataTableSortRules = []
			$.each(sortRules,function(index,sortRule) {
				var colIndex = 0
				$.each(tableInfo.cols,function(index,colInfo) {
					if(sortRule.fieldID === colInfo.properties.fieldID) {
						var dataTableSortRule = [colIndex,sortRule.direction]
						dataTableSortRules.push(dataTableSortRule)
					}
					colIndex++
				})
			})
			// If there is a 1-1 match between each sort rule and a column in the table,
			// then setting the order will accurately reflect the given sort order from the side-bar.
			// However, if there are sort fields in the sidebar which are not columns, it's not possible
			// to map the side-bars sort rules accurately onto the table; in this case the table is 
			// shown unordered. 
			if(dataTableSortRules.length === sortRules.length) {
				console.log("Setting table view sort rules: " + JSON.stringify(sortRules) + 
						" -> " + JSON.stringify(dataTableSortRules))
				dataTable.order(dataTableSortRules)
			} else {
				dataTable.order([])
			}
		}
		
		function updateData(recordData,sortRules) {
			drawInProgress = true
			dataTable.clear()
			setSortOrder(sortRules)
			dataTable.rows.add(recordData)
			dataTable.draw()
		}
		
		
		var tableContext = {
			resizeTable: resizeToContainerHeight,
			dataTable: dataTable,
			updateData: updateData,
			setSortOrder: setSortOrder
		}
		
		params.initDoneCallback(tableContext)
		
	}
	
	getTableInfo(populateTable)
	
	
}