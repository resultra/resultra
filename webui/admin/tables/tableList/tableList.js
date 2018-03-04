

function initAdminTableListSettings(databaseID) {
	
	var $tableList = $('#adminTableList')
	
	function addTableListItem(tableRef) {
		var $tableListItem = $('#tableListItemTemplate').clone()
		$tableListItem.attr("id","")
		
		var $tableName = $tableListItem.find("label")
		$tableName.text(tableRef.name)
		
		
		var $editTableButton = $tableListItem.find(".editTablePropsButton")
		setPageContentButtonClickHandler($editTableButton,
			'/admin/table/' + tableRef.tableID, function() {
				initTablePropsAdminSettingsPageContent(tableRef)
		})
		
		
		$tableList.append($tableListItem)
	}
	
	
	var getTableParams = { 
		parentDatabaseID: databaseID 
	}
	jsonAPIRequest("tableView/list",getTableParams,function(tableRefs) {
		
		$tableList.empty()
		
		$.each(tableRefs,function(index,tableRef) {
			addTableListItem(tableRef)			
		})
		
	})	
	
	initButtonClickHandler('#adminNewTableButton',function() {
		console.log("New table button clicked")
			
		openNewTableDialog(databaseID)
	})
	
	
}