

function navigateToTablePropsPage(pageContext,tableRef) {
	var contentURL = '/admin/table/' + tableRef.tableID
	setSettingsPageContent(contentURL,function() {
		initTablePropsAdminSettingsPageContent(pageContext,tableRef)	
	})
	var offPageURL = "/admin/table/tableProps/offPageContent"
	setSettingsPageOffPageContent(offPageURL,function() {})
	
}

function initAdminTableListSettings(pageContext) {
	
	var $tableList = $('#adminTableList')
	
	function addTableListItem(tableRef) {
		var $tableListItem = $('#tableListItemTemplate').clone()
		$tableListItem.attr("id","")
		
		var $tableName = $tableListItem.find("label")
		$tableName.text(tableRef.name)
		
		
		var $editTableButton = $tableListItem.find(".editTablePropsButton")
		$editTableButton.click(function(e) {
			e.preventDefault()
			$editTableButton.blur()
			navigateToTablePropsPage(pageContext,tableRef)
		})		
		$tableList.append($tableListItem)
	}
	
	
	var getTableParams = { 
		parentDatabaseID: pageContext.databaseID 
	}
	jsonAPIRequest("tableView/list",getTableParams,function(tableRefs) {
		
		$tableList.empty()
		
		$.each(tableRefs,function(index,tableRef) {
			addTableListItem(tableRef)			
		})
		
	})	
	
	initButtonClickHandler('#adminNewTableButton',function() {
		console.log("New table button clicked")
			
		openNewTableDialog(pageContext)
	})
	
	
}