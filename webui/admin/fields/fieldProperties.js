function openFieldPropertiesDialog(fieldInfo) {
	
	console.log("Edit field properties: ")
	
	var $fieldPropertiesDialog = $("#adminFieldPropertiesDialog")
	
	 var $nameInput = $('#fieldPropertiesFieldNameInput')
	 $nameInput.val(fieldInfo.name)
	
	var $fieldType = $('#fieldPropertiesFieldType')
	$fieldType.text(fieldTypeLabel(fieldInfo.type))
	
	var $refNameInput = $('#fieldPropertiesFieldRefNameInput')
	$refNameInput.val(fieldInfo.refName)
	
	var $isCalFieldCheckbox = $('#fieldPropertiesFieldIsCalcField')
	$isCalFieldCheckbox.prop("checked",fieldInfo.isCalcField)	
	
	$fieldPropertiesDialog.modal("show")
}