function initTableViewColsProperties(tableRef) {
	
	var $columnList = $('#tableColPropsColList')
	
	function saveUpdatedColumnOrder() {
		console.log("Saving updated columns: " + tableRef.tableID)
		
		var columns = []
		
		$columnList.find('.list-group-item').each(function() {
			var columnID = $(this).attr('data-column-id')
			columns.push(columnID)
		})
		
		console.log("savedUpdatedColumnOrder: " + JSON.stringify(columns))
		
		var setColsParams = {
			tableID: tableRef.tableID,
			orderedColumns: columns
		}
		jsonAPIRequest("tableView/setOrderedCols",setColsParams,function(replyStatus) {
			
		})
	}
	
    $columnList.sortable({
		placeholder: "ui-state-highlight",
		cursor:"move",
		update: function( event, ui ) {
			saveUpdatedColumnOrder(tableRef.tableID)
		}
    });
	
	
	loadFieldInfo(tableRef.parentDatabaseID,[fieldTypeAll],function(fieldsByID) {
		
		function populateTablePreview(tableInfo) {
			var $tablePreview = $('#tablePreview')
			
			var $tableHeader = $("<thead></thead>")
			var $headerRow = $("<tr></tr>")
	
			$.each(tableInfo.cols,function(index,colInfo) {
				var $header = $('<th></th>')
				
				var colWidths = tableInfo.table.properties.colWidths
				
				if(colWidths.hasOwnProperty(colInfo.columnID))
				{
					var colWidth = colWidths[colInfo.columnID]
					$header.css("width",colWidth + 'px')
				}
				
				$header.attr('data-col-id',colInfo.columnID)
				
				setFormComponentLabel($header,colInfo.properties.fieldID,
						colInfo.properties.labelFormat)
				
				$headerRow.append($header)
			})
		
			$tableHeader.append($headerRow)
			$tableHeader.find("th").css("background-color","lightGrey")
			
			$tablePreview.append($tableHeader)
			
			var componentContext = {
				databaseID: tableRef.parentDatabaseID,
				fieldsByID: fieldsByID
			}
			var recordProxy = {
				changeSetID: MainLineFullyCommittedChangeSetID,
				getRecordFunc: function() { return nil },
				updateRecordFunc: function(updatedRecordRef) {}
			}
				
			var $tableBody = $("<tbody></tbody>")
			var $previewRow= $("<tr></tr>")
			$.each(tableInfo.cols,function(index,colInfo) {
				var $column = $('<td></td>')				
				switch (colInfo.colType) {
				case 'numberInput':
					var $cellContainer = $(numberInputTableCellContainerHTML())
					configureNumberInputButtonSpinner($cellContainer,colInfo)
					$column.append($cellContainer)
					break
				case 'textInput':
					var $cellContainer = $(textBoxTableViewContainerHTML())
					function noOpValueList(newVal) {}
					configureTextBoxComponentValueListDropdown($cellContainer,colInfo,noOpValueList)
					$column.append($cellContainer)
					break					
				case 'datePicker':
					var $cellContainer = $(datePickerTableViewCellContainerHTML())
					initDatePickerFormComponentInput($cellContainer,colInfo)
					$column.append($cellContainer)
					break
				case 'userSelection':
					var $cellContainer = $(userSelectionTableCellContainerHTML())
					initDatePickerFormComponentInput($cellContainer,colInfo)
					$column.append($cellContainer)
					break
				case 'checkbox':
					var $cellContainer = $(checkBoxTableViewCellContainerHTML())
					initCheckBoxControl($cellContainer,colInfo)
					$column.append($cellContainer)
					break
				case 'rating':
					var $cellContainer = $(ratingTableCellContainerHTML())
					initRatingFormComponentControl($cellContainer,colInfo)
					$column.append($cellContainer)
					break
				case 'toggle':
					var $cellContainer = $(toggleTableCellContainerHTML())
					initToggleComponentControl($cellContainer,colInfo)
					$column.append($cellContainer)
					break
				default:
					console.log("Missing preview info for column: " + JSON.stringify(colInfo))
				}
							
				$previewRow.append($column)
			})
			$tableBody.append($previewRow)
			$tablePreview.append($tableBody)
			
			
			$tablePreview.find('th').resizable({
				handles:'e',
				stop: function(event,ui) {
					var widthsByID = {}
					$tablePreview.find('th').each(function() {
						var columnID = $(this).attr('data-col-id')
						var width = $(this).outerWidth()
						widthsByID[columnID] = width
					})
					console.log("Updated column widths: " + JSON.stringify(widthsByID))
					var widthParams = {
						tableID: tableRef.tableID,
						colWidths: widthsByID
					}
					jsonAPIRequest("tableView/setColWidths",widthParams,function(replyStatus) {
					})
				}
			})
			
		}
		
		
		function populateOneTableColInTableColList(tableCol) {
			var $colListItem = $('#tableColItemTemplate').clone()
			$colListItem.attr("id","")
			
			var fieldName = fieldsByID[tableCol.properties.fieldID].name
			$colListItem.find('label').text(fieldName)
			
			var editColLink = '/admin/tablecol/' + tableCol.columnID
			$colListItem.find('.editTableColButton').attr("href",editColLink)
			
			$colListItem.attr('data-column-id',tableCol.columnID)
			
			var $deleteColButton = $colListItem.find('.deleteTableColButton')
			
			initButtonControlClickHandler($deleteColButton,function() {
				openConfirmDeleteDialog("column",function() {
					console.log("column deletion confirmed")
					var deleteParams = {
						parentTableID: tableCol.parentTableID,
						columnID: tableCol.columnID
					}
					jsonAPIRequest("tableView/deleteColumn",deleteParams,function(replyStatus) {
						$colListItem.remove()
						console.log("Delete confirmed")
						saveUpdatedColumnOrder()
					})
					
				})				
			})
			
		
			$columnList.append($colListItem)
		}
		
		var params = { tableID: tableRef.tableID }	
		jsonAPIRequest("tableView/getTableDisplayInfo",params,function(tableInfo) {
			console.log("Loading table column properties: " + JSON.stringify(tableInfo))
		
			$columnList.empty()
			$.each(tableInfo.cols,function(colIndex,tableCol) {
				populateOneTableColInTableColList(tableCol)
			})
			
			populateTablePreview(tableInfo)
		})
		
	})
	
	initButtonClickHandler('#adminNewTableColButton',function() {
		console.log("New table column button clicked")
		openNewTableColDialog(tableRef)
	})
	
}