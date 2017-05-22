function initItemListTableView($tableContainer, databaseID, tableID) {
	
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
		
	}
	
	getTableInfo(populateTable)
	
	
}