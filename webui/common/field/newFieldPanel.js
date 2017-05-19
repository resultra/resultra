function NewFieldPanel(databaseID,$newFieldForm) {
	
	var $isCalcField = $('#newFieldIsCalcField')
	var $fieldName = $newFieldForm.find("input[name=newFieldNameInput]")
	var $fieldRefName = $newFieldForm.find("input[name=fieldRefName]")
	var $fieldType = $newFieldForm.find("select[name=newFieldTypeSelection]")
	
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
	
	var newFieldParams =  function() {
		var newFieldParams = {
			parentDatabaseID: databaseID,
			name: $fieldName.val(),
			refName: $fieldRefName.val(),
			isCalcField: $isCalcField.prop("checked"),
			type: $fieldType.val()
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