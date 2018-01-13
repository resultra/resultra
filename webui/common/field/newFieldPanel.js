function NewFieldPanel(databaseID,$newFieldForm) {
	
	var $isCalcField = $('#newFieldIsCalcField')
	var $fieldName = $newFieldForm.find("input[name=newFieldNameInput]")
	var $fieldRefName = $newFieldForm.find("input[name=fieldRefName]")
	var $fieldType = $newFieldForm.find("select[name=newFieldTypeSelection]")
	var $calcFieldFormGroup = $('#newFieldCalcFieldFormGroup')
	
	var validator = $newFieldForm.validate({
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

	resetFormValidationFeedback($newFieldForm)
	$fieldType.val("")
	$fieldName.val("")
	$fieldRefName.val("")
	$isCalcField.prop("checked",false)
	$calcFieldFormGroup.hide()
	validator.resetForm()
	
	function validateNewFieldParams() {
		return $newFieldForm.valid()
	}
	
		
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
	
	initSelectControlChangeHandler($fieldType,function(fieldType) {
		if(fieldTypeSupportsCalcField(fieldType)) {
			$calcFieldFormGroup.show()			
		} else {
			$calcFieldFormGroup.hide()
		}
	})
	
	
	var newFieldParams =  function() {
		
		var fieldType =  $fieldType.val()
		
		var isCalcField = false
		if(fieldTypeSupportsCalcField(fieldType)) {
			isCalcField = $isCalcField.prop("checked")
		}
		
		var newFieldParams = {
			parentDatabaseID: databaseID,
			name: $fieldName.val(),
			refName: $fieldRefName.val(),
			isCalcField: isCalcField,
			type: fieldType
		}
		return newFieldParams
	}
	
	function createNewField(newFieldCallback) {
		if($newFieldForm.valid()) {
			
			var isCalcField = $isCalcField.prop("checked")
			var params = newFieldParams()
			if(isCalcField) {
				jsonAPIRequest("calcField/new",params,function(newField) {
					console.log("new field created: " + JSON.stringify(newField))
					newFieldCallback(newField)
				})
				
			} else {
				console.log("creating new field: params= " + JSON.stringify(newFieldParams))
				jsonAPIRequest("field/new",params,function(newField) {
					console.log("new field created: " + JSON.stringify(newField))
					newFieldCallback(newField)
				})
				
			}
		} else {
			newFieldCreatedCallback(null)
		}
	}
	
	this.createNewField = createNewField
	this.validateNewFieldParams =  validateNewFieldParams
	this.newFieldParams = newFieldParams
	
}