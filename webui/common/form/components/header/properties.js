
function loadFormHeaderProperties($header,headerRef) {
	
	
	function initHeaderTextProperties() {
	
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
	
	function initHeaderSizeProperties() {
		var $sizeSelection = $('#adminHeaderComponentSizeSelection')
		$sizeSelection.val(headerRef.properties.headerSize)
		initSelectControlChangeHandler($sizeSelection,function(newSize) {
		
			var sizeParams = {
				parentFormID: headerRef.parentFormID,
				headerID: headerRef.headerID,
				size: newSize
			}
			console.log("Setting new header size: " + JSON.stringify(sizeParams))
		
			jsonAPIRequest("frm/header/setSize",sizeParams,function(updatedHeader) {
				setContainerComponentInfo($header,updatedHeader,updatedHeader.headerID)	
				setHeaderFormComponentHeaderSize($header,updatedHeader.properties.headerSize)			
			})
		
		})
		
	}
	
	function initHeaderUnderlinedProperties() {
		initCheckboxChangeHandler('#adminHeaderComponentUnderline', 
					headerRef.properties.underlined, function (newVal) {
				
			var underlinedParams = {
				parentFormID: headerRef.parentFormID,
				headerID: headerRef.headerID,
				underlined: newVal
			}

			jsonAPIRequest("frm/header/setUnderlined",underlinedParams,function(updatedHeader) {
				setContainerComponentInfo($header,updatedHeader,updatedHeader.checkBoxID)
				setHeaderFormComponentUnderlined($header,updatedHeader.properties.underlined)		
			})
		})
		
	}
	
	var elemPrefix = "header_"
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
				parentFormID: headerRef.parentFormID,
				headerID: headerRef.headerID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/header/setVisibility",params,function(updatedHeader) {
			setContainerComponentInfo($header,updatedHeader,updatedHeader.headerID)	
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: headerRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
	
	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: headerRef.parentFormID,
		componentID: headerRef.headerID,
		componentLabel: 'header',
		$componentContainer: $header
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	console.log("Loading header properties")
	
	initHeaderTextProperties()
	initHeaderSizeProperties()
	initHeaderUnderlinedProperties()
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#formHeaderProps')
		
	closeFormulaEditor()
	
}