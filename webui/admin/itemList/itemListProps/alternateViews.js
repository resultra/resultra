function initAlternateFormsProperties(listInfo) {

	function populateOneFormCheckbox(formInfo,altFormsLookup) {

		var $propertyCell = $('#adminItemListAlternateFormListPropertyCell')
		

		var $formItemCheckboxContainer = $('#adminItemListAlternateFormCheckboxTemplate').clone()
		$formItemCheckboxContainer.attr("id","")
		$formItemCheckboxContainer.attr("data-formID",formInfo.formID)
		var $nameLabel = $formItemCheckboxContainer.find("span")
		$nameLabel.text(formInfo.name)

		var $formCheckbox = $formItemCheckboxContainer.find("input")
		var $itemsPerPageFormGroup = $formItemCheckboxContainer.find('.itemsPerPageFormGroup')
		var $itemsPerPageFormSelection = $formItemCheckboxContainer.find('.itemsPerPageSelection')

		if (altFormsLookup.hasID(formInfo.formID)) {
			$formCheckbox.prop("checked",true)
			$itemsPerPageFormGroup.show()
		} else {
			$formCheckbox.prop("checked",false)
			$itemsPerPageFormGroup.hide()				
		}
		
		function updateAlternateForms() {
			var alternateForms = []
			$propertyCell.find(".alternateFormCheckboxContainer").each(function() {
				var formID = $(this).attr("data-formID")
				var $checkbox = $(this).find("input")
				var isChecked = $checkbox.prop("checked")
				if (isChecked) {
					alternateForms.push(formID)
				}
				console.log("form checkbox: " + formID + " " + isChecked)
			})
			var altFormsParams = {
				listID:listInfo.listID,
				alternateForms: alternateForms
			}
			jsonAPIRequest("itemList/setAlternateForms",altFormsParams,function(updatedListInfo) {
			})
			
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

	var formListParams =  { parentDatabaseID: listInfo.parentDatabaseID }
	jsonAPIRequest("frm/list",formListParams,function(formsInfo) {
		var altFormsLookup = new IDLookupTable(listInfo.properties.alternateForms)
		$.each(formsInfo, function(index, formInfo) {
			populateOneFormCheckbox(formInfo,altFormsLookup)
		})
	})

}
