// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function initItemListViewSelection(config) {
	
	var $viewSelection = $('#itemListViewSelection')
	var $pageSizeSelection = $('#itemListPageSizeSelection')
	var $formPageSizeSelectionFormGroup = $('#formPageSizeSelectionFormGroup')
	
	var alternateViewLookup = null
	if(config.alternateViews !== undefined) {
		alternateViewLookup = createAlternateViewLookupTable(config.alternateViews)
	}
	function showView(viewID) {
		if (config.initialView !== undefined) {
			if (config.initialView.formID === viewID) {
				return true
			} else if (config.initialView.tableID === viewID) {
				return true
			}
		}
		if (alternateViewLookup ===null) {
			return true
		} else {
			if(alternateViewLookup.hasOwnProperty(viewID)) {
				return true
			} else {
				return false
			}
		}
	}
	
	function populateViewSelection() {
		function populateTableViewList(doneCallback) {
			var $tableOptGroup = $('#itemListTableSelectionOptGroup')
			$tableOptGroup.empty()
			var getTableParams = { parentDatabaseID: config.databaseID }
			jsonAPIRequest("tableView/list",getTableParams,function(tableRefs) {
				$tableOptGroup.empty()
				var numTablesShown = 0
				$.each(tableRefs,function(index,tableRef) {
					if(showView(tableRef.tableID)) {
						numTablesShown++
						var $tableItem = $(selectOptionHTML(tableRef.tableID,tableRef.name))
						$tableItem.attr('data-view-type','table')
						$tableOptGroup.append($tableItem)						
					}
				})
				if (numTablesShown === 0) {
					$tableOptGroup.hide()
				}
				doneCallback()
			})
		}
		function populateFormList(doneCallback) {
			var listParams =  { parentDatabaseID: config.databaseID }
			jsonAPIRequest("frm/list",listParams,function(formsInfo) {
				var $formOptGroup = $('#itemListFormSelectionOptGroup')
				$formOptGroup.empty()
				var numFormsShown = 0
				$.each(formsInfo,function(index,formInfo) {
					if(showView(formInfo.formID)) {
						numFormsShown++
						var $formItem = $(selectOptionHTML(formInfo.formID,formInfo.name))
						$formItem.attr('data-view-type','form')
						$formOptGroup.append($formItem)	
					}
				})
				if (numFormsShown === 0) {
					$formOptGroup.hide()
				}
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
						$formPageSizeSelectionFormGroup.show()
					} else if (config.initialView.tableID !== undefined) {
						$viewSelection.val(config.initialView.tableID)					
						$formPageSizeSelectionFormGroup.hide()
					}
				} else {
					$formPageSizeSelectionFormGroup.hide()
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
				$formPageSizeSelectionFormGroup.show()
				var viewParams = {
					formID: selectedID,
					pageSize: convertStringToNumber($pageSizeSelection.val())
				}
				config.setViewCallback(viewParams)
			} else { // viewerType === 'table'
				$formPageSizeSelectionFormGroup.hide()
				var viewParams = {
					tableID: selectedID,
					pageSize: 0
				}
				config.setViewCallback(viewParams)
			}
		}
		
	}
	
	initSelectControlChangeHandler($viewSelection, function(selectedID) {
		if(alternateViewLookup != null) {
			var altView = alternateViewLookup[selectedID]
			if (altView !== undefined) {
				$pageSizeSelection.val(altView.pageSize)
			}
		}
		setView() 
	})	
	initNumberSelectionChangeHandler($pageSizeSelection, function(pageSize) { setView() })
	

} // initItemListFormProperties

