function openNewFieldDialog(databaseID) {
	
	var $newFieldDialogForm = $('#newFieldDialogForm')
	var $newFieldDialog = $('#newFieldDialog')
	
	var $isCalcField = $('#newFieldIsCalcField')
	var $fieldName = $newFieldDialogForm.find("input[name=newFieldNameInput]")
	var $fieldRefName = $newFieldDialogForm.find("input[name=fieldRefName]")
	var $fieldType = $newFieldDialogForm.find("select[name=newFieldTypeSelection]")
	
	var validator = $newFieldDialogForm.validate({
		rules: {
			newFieldNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/field/validateNewFieldName',
					data: {
						databaseID: databaseID,
						fieldName: function() { return $fieldName.val(); }
					}
				} // remote
			}, // newListNameInput
			fieldRefName: { 
				minlength: 3,
				required:true, 
				remote: {
					url: '/api/field/validateNewFieldRefName',
					data: {
						databaseID: databaseID,
						fieldRefName: function() { return $fieldRefName.val(); }
					}
				} // remote
			},
			newFieldTypeSelection: { required: true }
		},
		messages: {
			fieldRefName: {
				required: "Formula reference name is required"
			},
			newFieldNameInput: {
				required: "Field name is required"
			},
			newFieldTypeSelection: {
				required: "Select a field type"
			}
		}
	})

	resetFormValidationFeedback($newFieldDialogForm)
	$fieldType.val("")
	$fieldName.val("")
	$fieldRefName.val("")
	$isCalcField.prop("checked",false)
	validator.resetForm()
		
	// Set a default reference name based upon the field name the user inputs.
	// However, after the field reference name has been manually edited, don't
	// change it to the default reference name anymore.
	var fieldRefNameManuallyEdited = false
	$fieldName.on('input',function() {
		if (!fieldRefNameManuallyEdited) {
			var fieldName = $fieldName.val()
			var defaultFieldRefName = fieldName.replace(/[^0-9a-zA-Z]/g,"")
			$fieldRefName.val(defaultFieldRefName)
			// Immediately trigger validation of the field reference name, based
			// upon the default value.
			$fieldRefName.valid()
		}
	})
	$fieldRefName.change(function() {
		fieldRefNameManuallyEdited = true
	})
	
	
	initButtonClickHandler('#newFieldSaveButton',function() {
		console.log("New field save button clicked")
		if($newFieldDialogForm.valid()) {
			
			var isCalcField = $isCalcField.prop("checked")
			var newFieldParams = {
				parentDatabaseID: databaseID,
				name: $fieldName.val(),
				refName: $fieldRefName.val(),
				isCalcField: $isCalcField.prop("checked"),
				type: $fieldType.val()
			}
			if(isCalcField) {
				jsonAPIRequest("calcField/new",newFieldParams,function(newField) {
					console.log("new field created: " + JSON.stringify(newField))
					$newFieldDialog.modal('hide')
					navigateToURL('/admin/field/'+newField.fieldID)
				})
				
			} else {
				console.log("creating new field: params= " + JSON.stringify(newFieldParams))
				jsonAPIRequest("field/new",newFieldParams,function(newField) {
					console.log("new field created: " + JSON.stringify(newField))
					$newFieldDialog.modal('hide')
					navigateToURL('/admin/field/'+newField.fieldID)
				})
				
			}
			

		}
	})
	
	$newFieldDialog.modal('show')
	
}