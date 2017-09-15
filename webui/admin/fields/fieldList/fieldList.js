function initAdminFieldSettings(databaseID) {
		
		
	function addFieldToAdminFieldList(fieldInfo) {
 
		var $fieldListRow = $('#adminFieldItemRowTemplate').clone()
		$fieldListRow.attr("id","")
	
		$fieldListRow.attr("data-fieldID",fieldInfo.fieldID)
		
		var $editFieldPropsButton = $fieldListRow.find(".editFieldPropsButton")
		var editPropsLink = '/admin/field/' + fieldInfo.fieldID
		$editFieldPropsButton.attr("href",editPropsLink)

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