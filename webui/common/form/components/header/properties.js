
function initHeaderTextProperties($header,headerRef) {
	
	$('#headerTextNameInput').val(headerRef.properties.label)
	
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
				parentFormID: headerRef.parentFormID,
				headerID: headerRef.headerID,
				label: validatedName
			}
			
			jsonAPIRequest("frm/header/setLabel",setLabelParams,function(updatedHeader) {
				console.log("Done changing header label: " + validatedName)
				
				setContainerComponentInfo($header,updatedHeader,updatedHeader.headerID)				
				$header.find(".formHeader").text(updatedHeader.properties.label)
					
			})		
	})
		

	validator.resetForm()
	
}


function loadFormHeaderProperties($header,headerRef) {
	console.log("Loading header properties")
	
	initHeaderTextProperties($header,headerRef)
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#formHeaderProps')
		
	closeFormulaEditor()
	
}