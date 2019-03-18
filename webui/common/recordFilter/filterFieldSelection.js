// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initFilterFieldSelection(params) {
	
	var $fieldSelectionList = $('#' + params.elemPrefix + 'AdminFilterFieldSelection')
	
	var defaultFilterFieldsLookup = new IDLookupTable(params.defaultFilterFields)
	
	function addFieldSelection(fieldInfo) {
			
		var fieldSelectionHTML = '' +
				'<div class="list-group-item adminFilterFieldSelectionCheckbox" id="adminFilterFieldSelectionCheckboxTemplate">' +
	      			'<div class="checkbox">'+
	        			'<input type="checkbox">'+
						'<label><span>Field name goes here.</span></label>' +
	      			'</div>'+
	  			'</div>'		

		var $fieldSelection = $(fieldSelectionHTML)

		
		$fieldSelection.find("span").text(fieldInfo.name)
		
		var inputID = params.elemPrefix + '_' + fieldInfo.fieldID
		
		var $selectionCheckbox = $fieldSelection.find("input")
		$selectionCheckbox.attr("data-fieldID",fieldInfo.fieldID)
		$selectionCheckbox.attr("id",inputID)
		$fieldSelection.find("label").attr("for",inputID)
		
		var fieldEnabled = defaultFilterFieldsLookup.hasID(fieldInfo.fieldID)
		initCheckboxControlChangeHandler($selectionCheckbox,fieldEnabled,function(isEnabled) {
			
			var checkedFields = []
			$fieldSelectionList.find("input").each(function() {
				if($(this).prop("checked")===true) {
					var fieldID = $(this).attr('data-fieldID')
					checkedFields.push(fieldID)
				}
			})
			console.log("checked fields: " + JSON.stringify(checkedFields))
			params.setFilterFieldsCallback(checkedFields)
		})
		
		$fieldSelectionList.append($fieldSelection)
		
		
	}
	
	loadSortedFieldInfo(params.databaseID, [fieldTypeNumber,fieldTypeText,fieldTypeTime],function(sortedFields) {
		for (var fieldIndex in sortedFields) {
			var fieldInfo = sortedFields[fieldIndex]	
			addFieldSelection(fieldInfo)		
		} // for each  field
	})
	
	
	
}