
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
				renderCellHTMLFunc,initContainerFunc) {
					
		function columnSortType(colInfo,fieldsByID) {	
			function dataColSortType(colInfo,fieldsByID) {
				var fieldID = colInfo.properties.fieldID
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
			if (colInfo.colType === 'button') {
				return 'string'
			} else {
				return dataColSortType(colInfo,fieldsByID)
			}
		}
		
		var colType = columnSortType(colInfo,fieldsByID)
		
		function columnDataKey(colInfo) {
			if (colInfo.colType === 'button') {
				return null
			} else {
				var fieldID = colInfo.properties.fieldID
				return 'fieldValues.' + fieldID
			}
		}
								
		var colDef = {
			data:columnDataKey(colInfo),
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
		
		return colDef
	}
	
	
	function createNumberInputColDef(colInfo,fieldsByID) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
			setContainerComponentInfo($cellContainer,colInfo,colInfo.numberInputID)
			initNumberInputTableRecordEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				numberInputTableCellContainerHTML,initContainer)
	}


	function createUserSelectionColDef(colInfo,fieldsByID) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
			setContainerComponentInfo($cellContainer,colInfo,colInfo.numberInputID)
			initUserSelectionTableRecordEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				userSelectionTableCellContainerHTML,initContainer)
	}


	function createTextInputColDef(colInfo,fieldsByID) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
				setContainerComponentInfo($cellContainer,colInfo,colInfo.textInputID)
				initTextBoxRecordEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				textBoxTableViewContainerHTML,initContainer)
	}

	function createNoteColDef(colInfo,fieldsByID) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
				setContainerComponentInfo($cellContainer,colInfo,colInfo.noteID)
				initNoteEditorTableCellEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				noteEditorTableViewCellContainerHTML,initContainer)
	}

	function createCommentColDef(colInfo,fieldsByID) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
				setContainerComponentInfo($cellContainer,colInfo,colInfo.noteID)
				initCommentBoxTableViewRecordEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				commentBoxTableViewContainerHTML,initContainer)
	}


	function createAttachmentColDef(colInfo,fieldsByID) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
				setContainerComponentInfo($cellContainer,colInfo,colInfo.noteID)
				initAttachmentTableViewRecordEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				attachmentTableViewContainerHTML,initContainer)
	}
	
	function createDateInputColDef(colInfo,fieldsByID) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
			setContainerComponentInfo($cellContainer,colInfo,colInfo.datePickerID)
			initTableViewDatePickerEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				datePickerTableViewCellContainerHTML,initContainer)
	}

	function createCheckboxColDef(colInfo,fieldsByID) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
			setContainerComponentInfo($cellContainer,colInfo,colInfo.datePickerID)
			initTableViewCheckboxEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				checkBoxTableViewCellContainerHTML,initContainer)
	}

	function createFormButtonColDef(colInfo,fieldsByID) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
			setContainerComponentInfo($cellContainer,colInfo,colInfo.datePickerID)
			
			setFormButtonSize($cellContainer,colInfo.properties.size)
			setFormButtonColorScheme($cellContainer,colInfo.properties.colorScheme)
			setFormButtonLabel($cellContainer,colInfo)
			
			// The loadFormViewComponents and loadRecordIntoFormLayout functions
			// need to be passed to initFormButtonRecordEditBehavior in order
			// to avoid a cyclical package dependency.
			var defaultValSrc = "col="+colInfo.columnID
			
			initFormButtonRecordEditBehavior($cellContainer,componentContext,recordProxy, colInfo,defaultValSrc,
					loadFormViewComponents,loadRecordIntoFormLayout)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				formButtonContainerHTML,initContainer)
	}



	function createToggleColDef(colInfo,fieldsByID) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
			setContainerComponentInfo($cellContainer,colInfo,colInfo.toggleID)
			
			initToggleTableCellRecordEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				toggleTableCellContainerHTML,initContainer)
	}


	function createRatingColDef(colInfo,fieldsByID) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
			setContainerComponentInfo($cellContainer,colInfo,colInfo.ratingID)
			initRatingTableCellRecordEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,
				ratingTableCellContainerHTML,initContainer)
	}
	
	
	function createProgressColDef(colInfo,fieldsByID) {
		
		function initContainer(colInfo, $cellContainer, fieldsByID,recordProxy,componentContext) {
			setContainerComponentInfo($cellContainer,colInfo,colInfo.progressID)
			initProgressRecordEditBehavior($cellContainer,componentContext,recordProxy, colInfo)
		}
		return createTableViewColDef(colInfo,fieldsByID,progressTableCellContainerHTML,initContainer)
	}

	
	function createColDef(colInfo,fieldsByID) {
		switch (colInfo.colType) {
		case 'numberInput':
			return createNumberInputColDef(colInfo,fieldsByID)
		case 'textInput':
			return createTextInputColDef(colInfo,fieldsByID)
		case 'datePicker':
			return createDateInputColDef(colInfo,fieldsByID)
		case 'checkbox':
			return createCheckboxColDef(colInfo,fieldsByID)
		case 'rating':
			return createRatingColDef(colInfo,fieldsByID)
		case 'toggle':
			return createToggleColDef(colInfo,fieldsByID)
		case 'userSelection':
			return createUserSelectionColDef(colInfo,fieldsByID)
		case 'note':
			return createNoteColDef(colInfo,fieldsByID)
		case 'comment':
			return createCommentColDef(colInfo,fieldsByID)
		case 'attachment':
			return createAttachmentColDef(colInfo,fieldsByID)
		case 'button':
			return createFormButtonColDef(colInfo,fieldsByID)
		case 'progress':
			return createProgressColDef(colInfo,fieldsByID)
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
				
		
		function tableHeader() {
	
			var $tableHeader = $("<thead></thead>")
			var $headerRow = $("<tr></tr>")
	
			$.each(tableInfo.cols,function(index,colInfo) {
				var $header = $('<th></th>')
								
				if (colInfo.colType !== 'button') {
					setFormComponentLabel($header,colInfo.properties.fieldID,
							colInfo.properties.labelFormat)
						
				} else {
					setFormButtonHeader($header,colInfo)
				}
				// TODO - Add support for other column types
				if(colInfo.properties.helpPopupMsg !== undefined) {
					$header.append(componentHelpPopupButtonHTML())
					initComponentHelpPopupButton($header,colInfo,"auto bottom")
				}					
				
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
			var colDataDef = createColDef(colInfo,fieldsByID)
			dataCols.push(colDataDef)
		})
		
		var dataTable = $tableElem.DataTable({
			destroy:true, // Destroy existing table before applying the options
			searching:false, // Hide the search box
			paging:true, // pagination must be enabled for pageResize plug-in
			pageLength:1,
			lengthChange:true,  // needed for pageResize plug-in
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
			
			// To dynamically resize, the header and footer need to subtracted from
			// the parent container's height, then the number of rows needs to be recalculated
			// and set to the page length. 
			//
			// The approach below basically resizes the table by setting the page length
			// to a number of elements which will fit within the table's parent container.
			// Another approach was to use the 'scrollY' option for DataTables, but this
			// caused the table header and content to be split up, so that the header
			// wouldn't horizontally scroll with the table body.
			var headerHeight = params.$tableContainer.find("thead").height()
			var contentHeight = params.$tableContainer.find(".dataTables_scroll").height()
			var overallHeight = $tableElem.height()
			var containerHeight = params.$tableContainer.height()
			
			// There's no single div to surround all footer elements. However, the following
			// number "rounds up" the expected height for any footer elements.
			var footerHeight = 40
			
			
			var scrollHeight = containerHeight - footerHeight - headerHeight
			
			var rowHeight = 40
			
			// Dynamically set the page length, based upon the current scroll height. We want 
			// the number of rows to fill the visible page, but not cause scroll bars to enable.
			// So, the page length is adjusted downwards by 1 if there is minimal remaining space
			// below the table rows which fully fit within the scroll body.
			var rowsRemainder = scrollHeight % rowHeight
			var pageLen = Math.floor(scrollHeight / rowHeight)
			if (rowsRemainder <= 10) { pageLen-- }
			dataTable.page.len(pageLen)
			
			
			dataTable.draw() // force redraw
		}
		
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