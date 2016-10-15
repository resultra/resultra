
function initHeaderTextProperties(headerRef) {
	
//	$('#headerTextNameInput').val(headerRef.properties.label)
	
	var $headerLabelForm = $('#headerTextPropertyPanelForm')
	
	var remoteValidationParams = {
		url: '/api/generic/stringValidation/validateItemLabel',
		data: {
			label: function() { return $('#headerTextNameInput').val(); }
		}
	}
	
	var validationSettings = createInlineFormValidationSettings({
		rules: {
			headerTextNameInput: {
				minlength: 3,
				required: true,
				remote:  remoteValidationParams
			}
		},
	})	
	var validator = $headerLabelForm.validate(validationSettings)
	
	
	initInlineInputValidationOnBlur(validator,'#headerTextNameInput',
			remoteValidationParams, function(validatedName) {
				
				console.log("Header label changed: " + validatedName)
				
				var setLabelParams = {
					label: validatedName
				}
				
//				jsonAPIRequest("frm/setName",{formID:formInfo.formID,newFormName:validatedName},function(formInfo) {
//					console.log("Done changing header label: " + validatedName)
//				})		
	})
		

	validator.resetForm()
	
}


function loadFormHeaderProperties(headerRef) {
	console.log("Loading header properties")
	
	initHeaderTextProperties(headerRef)
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#formHeaderProps')
		
	closeFormulaEditor()
	
}