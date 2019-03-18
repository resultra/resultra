// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initAlternateFormsProperties(listInfo) {
	
	var $propertyCell = $('#adminItemListAlternateFormListPropertyCell')
	var $alternateTablesContainer = $('#adminItemListAlternateTableList')
	
	function updateAlternateForms() {
		var alternateViews = []
		$propertyCell.find(".alternateFormCheckboxContainer").each(function() {
			var formID = $(this).attr("data-formID")
			var $checkbox = $(this).find("input")
			var $pageSizeSelection = $(this).find('.itemsPerPageSelection')
			var pageSize = Number($pageSizeSelection.val())
			var isChecked = $checkbox.prop("checked")
			if (isChecked) {
				var altView = {
					formID: formID,
					pageSize: pageSize
				}
				alternateViews.push(altView)
			}
			console.log("form checkbox: " + formID + " " + isChecked)
		})
		$alternateTablesContainer.find(".alternateTableCheckboxContainer").each(function() {
			var tableID = $(this).attr("data-tableID")
			var $checkbox = $(this).find("input")
			var isChecked = $checkbox.prop("checked")
			if(isChecked) {
				var altView = {
					tableID: tableID,
					pageSize: 1
				}
				alternateViews.push(altView)				
			}
		})
		var altViewsParams = {
			listID:listInfo.listID,
			alternateViews: alternateViews
		}
		jsonAPIRequest("itemList/setAlternateViews",altViewsParams,function(updatedListInfo) {
		})
		
	}

	function populateOneFormCheckbox(formInfo,altViewsLookup) {

		
		var $formItemCheckboxContainer = $('#adminItemListAlternateFormCheckboxTemplate').clone()
		$formItemCheckboxContainer.attr("id","")
		$formItemCheckboxContainer.attr("data-formID",formInfo.formID)
		var $nameLabel = $formItemCheckboxContainer.find("span")
		$nameLabel.text(formInfo.name)

		var $formCheckbox = $formItemCheckboxContainer.find("input")
		var $itemsPerPageFormGroup = $formItemCheckboxContainer.find('.itemsPerPageFormGroup')
		var $itemsPerPageFormSelection = $formItemCheckboxContainer.find('.itemsPerPageSelection')

		if (altViewsLookup.hasOwnProperty(formInfo.formID)) {
			var altView = altViewsLookup[formInfo.formID]
			$formCheckbox.prop("checked",true)
			$itemsPerPageFormSelection.val(altView.pageSize)
			$itemsPerPageFormGroup.show()
		} else {
			$formCheckbox.prop("checked",false)
			$itemsPerPageFormGroup.hide()				
		}
		

		$formCheckbox.change(function() {
			var formIsChecked = $formCheckbox.prop("checked")
			console.log("checkbox changed: " + formInfo.name + " - " + formIsChecked)
			
			if(formIsChecked) {
				$itemsPerPageFormGroup.show()
				$itemsPerPageFormSelection.val("1")
			} else {
				$itemsPerPageFormGroup.hide()				
			}
			updateAlternateForms()
		})
		
		$itemsPerPageFormSelection.change(function() {
			updateAlternateForms()
		})
		

		$propertyCell.append($formItemCheckboxContainer)

	}
	
	function populateOneTableCheckbox(tableRef,altViewLookup) {
		
		var $tableCheckboxItem = $('#adminItemListAlternateTableCheckboxTemplate').clone()
		$tableCheckboxItem.attr("id","")
		
		$tableCheckboxItem.attr("data-tableID",tableRef.tableID)
		
		var $nameLabel = $tableCheckboxItem.find("span")
		$nameLabel.text(tableRef.name)
		
		var $tableCheckbox = $tableCheckboxItem.find("input")
		if (altViewLookup.hasOwnProperty(tableRef.tableID)) {
			$tableCheckbox.prop("checked",true)
		} else {
			$tableCheckbox.prop("checked",false)
		}
		
		$tableCheckbox.change(function() {
			updateAlternateForms()		
		})
		
		$alternateTablesContainer.append($tableCheckboxItem)
		
	}

	var altViewLookup = createAlternateViewLookupTable(listInfo.properties.alternateViews)

	var formListParams =  { parentDatabaseID: listInfo.parentDatabaseID }
	jsonAPIRequest("frm/list",formListParams,function(formsInfo) {
		$.each(formsInfo, function(index, formInfo) {
			populateOneFormCheckbox(formInfo,altViewLookup)
		})
	})
	
	var getTableParams = { parentDatabaseID: listInfo.parentDatabaseID }
	jsonAPIRequest("tableView/list",getTableParams,function(tableRefs) {
		$.each(tableRefs,function(index,tableRef) {
			populateOneTableCheckbox(tableRef,altViewLookup)
		})
	})
	

}
