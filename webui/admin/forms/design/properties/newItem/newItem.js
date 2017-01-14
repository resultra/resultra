


function initFormPropertiesNewItem(formInfo) {
	
	var $addNewCheckbox = $('#formNewItemPropertyAddNew')
		
	$addNewCheckbox.prop("checked",formInfo.properties.addNewItemFromForm)
	
	$addNewCheckbox.click(function() {
		
		var doAddNewItemsFromThisForm = $addNewCheckbox.prop("checked")
		
		var setPropertyParams = {
			formID:formInfo.formID,
			addNewItemFromForm: doAddNewItemsFromThisForm
		}
		
		jsonAPIRequest("frm/setAddNewFromForm",setPropertyParams,function(formInfo) {
			console.log("Done changing property for adding new items: " + JSON.stringify(setPropertyParams))
		})		
		
		
	})
	
	
		
}