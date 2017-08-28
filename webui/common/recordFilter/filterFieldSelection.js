function initFilterFieldSelection(params) {
	
	var $fieldSelectionList = $('#' + params.elemPrefix + 'AdminFilterFieldSelection')
	
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
		
		var fieldEnabled = false
		initCheckboxControlChangeHandler($selectionCheckbox,fieldEnabled,function(isEnabled) {
			
			var checkedFields = []
			$fieldSelectionList.find("input").each(function() {
				if($(this).prop("checked")===true) {
					var fieldID = $(this).attr('data-fieldID')
					checkedFields.push(fieldID)
				}
			})
			console.log("checked fields: " + JSON.stringify(checkedFields))
//			params.setRolesCallback(checkedRoles)
			
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