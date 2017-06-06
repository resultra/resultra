function initAlternateFormsProperties(listInfo) {

	function populateOneFormCheckbox(formInfo,altFormsLookup) {

		var $propertyCell = $('#adminItemListAlternateFormListPropertyCell')

		var $formItemCheckboxContainer = $('#adminItemListAlternateFormCheckboxTemplate').clone()
		$formItemCheckboxContainer.attr("id","")
		$formItemCheckboxContainer.attr("data-formID",formInfo.formID)
		var $nameLabel = $formItemCheckboxContainer.find("span")
		$nameLabel.text(formInfo.name)

		var $formCheckbox = $formItemCheckboxContainer.find("input")

		if (altFormsLookup.hasID(formInfo.formID)) {
			$formCheckbox.prop("checked",true)
		} else {
			$formCheckbox.prop("checked",false)
		}

		$formCheckbox.change(function() {
				var formIsChecked = $formCheckbox.prop("checked")
				console.log("checkbox changed: " + formInfo.name + " - " + formIsChecked)

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
