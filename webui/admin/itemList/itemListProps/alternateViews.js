function initAlternateFormsProperties(listInfo) {

	function populateOneFormCheckbox(formInfo,altViewsLookup) {

		var $propertyCell = $('#adminItemListAlternateFormListPropertyCell')
		
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
			var altViewsParams = {
				listID:listInfo.listID,
				alternateViews: alternateViews
			}
			jsonAPIRequest("itemList/setAlternateViews",altViewsParams,function(updatedListInfo) {
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
		var altViewLookup = createAlternateViewLookupTable(listInfo.properties.alternateViews)
		$.each(formsInfo, function(index, formInfo) {
			populateOneFormCheckbox(formInfo,altViewLookup)
		})
	})

}
