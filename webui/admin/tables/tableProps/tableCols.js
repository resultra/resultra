function registerTablePropsLoader(pageContext,tableRef) {
	const tableSettingsLinkID = 'table-' + tableRef.tableID
	var contentURL = '/admin/table/' + tableRef.tableID
	registerPageContentLoader(tableSettingsLinkID, contentURL, function() {
		initTablePropsAdminSettingsPageContent(pageContext,tableRef)	
		var offPageURL = "/admin/table/tableProps/offPageContent"
		setSettingsPageOffPageContent(offPageURL,function() {})
	})
}

function getDisplayedTableColumnOrder() {
	var $columnList = $('#tableColPropsColList')
	var columns = []
	
	$columnList.find('.list-group-item').each(function() {
		var columnID = $(this).attr('data-column-id')
		columns.push(columnID)
	})
	console.log("savedUpdatedColumnOrder: " + JSON.stringify(columns))
	return columns
}

function saveUpdatedTableColumnOrder(tableRef) {
	console.log("Saving updated columns: " + tableRef.tableID)
	
	var columns = getDisplayedTableColumnOrder()
	
	
	var setColsParams = {
		tableID: tableRef.tableID,
		orderedColumns: columns
	}
	jsonAPIRequest("tableView/setOrderedCols",setColsParams,function(replyStatus) {
		
	})
	
}

function appendNewTableColToTableColumnOrder(tableRef,newColumnID) {
	var columns = getDisplayedTableColumnOrder()
	columns.push(newColumnID)
	var setColsParams = {
		tableID: tableRef.tableID,
		orderedColumns: columns
	}
	jsonAPIRequest("tableView/setOrderedCols",setColsParams,function(replyStatus) {
	})
	
}


function initTableViewColsProperties(pageContext,tableRef) {
	
	var $columnList = $('#tableColPropsColList')
		
    $columnList.sortable({
		placeholder: "ui-state-highlight",
		cursor:"move",
		update: function( event, ui ) {
			saveUpdatedTableColumnOrder(tableRef)
		}
    });
	
	loadFieldInfo(tableRef.parentDatabaseID,[fieldTypeAll],function(fieldsByID) {
		
		function populateOneTableColInTableColList(tableCol) {
			var $colListItem = $('#tableColItemTemplate').clone()
			$colListItem.attr("id","")
			
			if (tableCol.colType !== 'button') {
				var fieldName = fieldsByID[tableCol.properties.fieldID].name
				$colListItem.find('label').text(fieldName)
			} else {
				$colListItem.find('label').text("Button: open form")
			}
			
			
			var editColContentURL = '/admin/tablecol/' + tableCol.columnID
			var $tableColButton = $colListItem.find('.editTableColButton')
			setPageContentButtonClickHandler($tableColButton,editColContentURL,function() {
					initTableColPropsPageConent(pageContext,tableCol)
					registerTablePropsLoader(pageContext,tableRef)
			})			
			
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
						saveUpdatedTableColumnOrder(tableRef)
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
		
		})
		
	})
	
	initButtonClickHandler('#adminNewTableColButton',function() {
		console.log("New table column button clicked")
		openNewTableColDialog(pageContext,tableRef)
	})
	
	initButtonClickHandler('#adminNewFormButtonColButton',function() {
		console.log("New table column button clicked")
		openNewButtonTableColDialog(pageContext,tableRef)
	})
	
	
}