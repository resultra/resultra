// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initAdminFieldSettings(databaseID) {
		
		
	function addFieldToAdminFieldList(fieldInfo) {
 
		var $fieldListRow = $('#adminFieldItemRowTemplate').clone()
		$fieldListRow.attr("id","")
	
		$fieldListRow.attr("data-fieldID",fieldInfo.fieldID)
		
		var $editFieldPropsButton = $fieldListRow.find(".editFieldPropsButton")
		$editFieldPropsButton.click(function(e) {
			e.preventDefault()
			$editFieldPropsButton.blur()
			
			var editPropsContentURL = '/admin/field/' + fieldInfo.fieldID
			setSettingsPageContent(editPropsContentURL,function() {
				initFieldPropsSettingsPageContent(fieldInfo.fieldID)
			})
		})
		

		var $fieldNameCol = $fieldListRow.find(".fieldNameCol")
		$fieldNameCol.text(fieldInfo.name)
		
		var $refNameCol = $fieldListRow.find(".fieldRefNameCol")
		$refNameCol.text(fieldInfo.refName)
		
		
		var fieldLabel = fieldTypeLabel(fieldInfo.type)
		var $typeCol = $fieldListRow.find(".fieldTypeCol")
		$typeCol.text(fieldLabel)

		var $calcFieldCol = $fieldListRow.find(".fieldCalcCol")
		var calcFieldLabel = "No"
		if(fieldInfo.isCalcField) {
			calcFieldLabel = "Yes"
		}
		$calcFieldCol.text(calcFieldLabel)

		
		console.log("field row info: " + $fieldListRow.html())
	
		$('#adminFieldListBody').append($fieldListRow)		
	}
		
	
	loadSortedFieldInfo(databaseID, [fieldTypeAll],function(sortedFields) {
		for (var fieldIndex in sortedFields) {
			var fieldInfo = sortedFields[fieldIndex]	
			addFieldToAdminFieldList(fieldInfo)		
		} // for each  field
		$('#adminFieldList').DataTable({
			destroy:true, // Destroy existing table before applying the options
			searching:false, // Hide the search box
			bInfo:false, // Hide the "Showing 1 of N Entries" below the footer
			paging:false,
			//scrollY: '500px',
			//scrollCollapse:true,
		
			aoColumnDefs: [
			          { bSortable: false, aTargets: [ 4 ] }
			       ]
		
		})
	})
	
	initButtonClickHandler('#adminNewFieldButton',function() {
		console.log("New field button clicked")
		openNewFieldDialog(databaseID)
	})
	
	
	
}