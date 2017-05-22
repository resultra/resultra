
function initItemListViewSelection(config) {
	
	var $viewSelection = $('#itemListViewSelection')
	var $pageSizeSelection = $('#itemListPageSizeSelection')
	
	
	function populateViewSelection() {
		function populateTableViewList(doneCallback) {
			var $tableOptGroup = $('#itemListTableSelectionOptGroup')
			var getTableParams = { parentDatabaseID: config.databaseID }
			jsonAPIRequest("tableView/list",getTableParams,function(tableRefs) {
				$tableOptGroup.empty()
				$.each(tableRefs,function(index,tableRef) {
					var $tableItem = $(selectOptionHTML(tableRef.tableID,tableRef.name))
					$tableItem.attr('data-view-type','table')
					$tableOptGroup.append($tableItem)	
				})
				doneCallback()
			})
		}
		function populateFormList(doneCallback) {
			var listParams =  { parentDatabaseID: config.databaseID }
			jsonAPIRequest("frm/list",listParams,function(formsInfo) {
				var $formOptGroup = $('#itemListFormSelectionOptGroup')
				$.each(formsInfo,function(index,formInfo) {
					var $formItem = $(selectOptionHTML(formInfo.formID,formInfo.name))
					$formItem.attr('data-view-type','form')
					$formOptGroup.append($formItem)
				})
				doneCallback()
			})
	
		}
		var numOptGroupsRemaining = 2
		function donePopulatingOptGroup() {
			numOptGroupsRemaining--
			if(numOptGroupsRemaining<=0) {
				if (config.initialView !== undefined) {
					if(config.initialView.formID !== undefined) {
						$viewSelection.val(config.initialView.formID)
						$pageSizeSelection.val(config.initialView.pageSize)
					} else if (config.initialView.tableID !== undefined) {
						$viewSelection.val(config.initialView.tableID)					
					}
				}		
			}
		}
		populateFormList(donePopulatingOptGroup)
		populateTableViewList(donePopulatingOptGroup)
	}
	populateViewSelection()
	
	function setView() {
		var $selectedFormOrTable = $('#itemListViewSelection option:selected')
		var selectedID = $selectedFormOrTable.val()
		
		if(selectedID.length > 0) {
			var viewerType = $selectedFormOrTable.attr('data-view-type')
			console.log("Selected form or table: " + 
					$selectedFormOrTable.text() + ' type = ' + viewerType)
			if(viewerType === 'form') {
				var viewParams = {
					formID: selectedID,
					pageSize: convertStringToNumber($pageSizeSelection.val())
				}
				config.setViewCallback(viewParams)
			} else { // viewerType === 'table'
				var viewParams = {
					tableID: selectedID,
					pageSize: 0
				}
				config.setViewCallback(viewParams)
			}
		}
		
	}
	
	initSelectControlChangeHandler($viewSelection, function(selectedID) { setView() })	
	initNumberSelectionChangeHandler($pageSizeSelection, function(pageSize) { setView() })
	

} // initItemListFormProperties

