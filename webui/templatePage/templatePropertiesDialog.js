function openTemplatePropertiesDialog(templateInfo) {
	
	function initActiveTemplateProperty(templateInfo) {
			
		initCheckboxChangeHandler('#activeTemplatePropIsActive', 
			templateInfo.isActive, function(isActive) {
				var setActiveParams = {
					databaseID:templateInfo.databaseID,
					isActive:isActive
				}
				jsonAPIRequest("database/setActive",setActiveParams,function(dbInfo) {
				})
		})
	
		
			
	}
	
	function initNameProperty(templateInfo) {
	
		$('#databasePropsNameInput').val(templateInfo.name)
	
		var $nameForm = $('#templateNamePropertyForm')
		var $nameInput = $('#templatePropsNameInput')
		$nameInput.val(templateInfo.name)
		
		var remoteValidationParams = {
			url: '/api/database/validateDatabaseName',
			data: {
				databaseID: function() { return templateInfo.databaseID },
				databaseName: function() { return $nameInput.val() }
			}	
		}
	
		var validationSettings = createInlineFormValidationSettings({
			rules: {
				templatePropsNameInput: {
					minlength: 3,
					required: true,
					remote: remoteValidationParams
				} // newRoleNameInput
			}
		})	
	
	
		var validator = $nameForm.validate(validationSettings)
	
		initInlineInputValidationOnBlur(validator,'#templatePropsNameInput',
			remoteValidationParams, function(validatedName) {		
				var setNameParams = {
					databaseID:templateInfo.databaseID,
					newName:validatedName
				}
				jsonAPIRequest("database/setName",setNameParams,function(dbInfo) {
					console.log("Done changing database name: " + validatedName)
				})
		})	

		validator.resetForm()
	
	}
	
	
	function initTemplateDescriptionProperty(templateInfo) {
		
		var $descInput = $('#templateDescription')
	
		function setTrackerDescription(description) {
			var setDescParams = {
				databaseID:templateInfo.databaseID,
				description:description
			}
			jsonAPIRequest("database/setDescription",setDescParams,function(dbInfo) {
			})
		
		}
	
		$descInput.html(templateInfo.description)
	
		$descInput.dblclick(function() {
			if (!inlineCKEditorEnabled($descInput)) {
			
				var editor = enableInlineCKEditor($descInput)
				$descInput.focus()
			
				editor.on('blur', function(event) {
					var popupMsg = editor.getData();
				
					setTrackerDescription(popupMsg)
							
					disableInlineCKEditor($descInput,editor)
				})
			
			}
		})
			
	}
	
	
	var getDBInfoParams = { databaseID: templateInfo.databaseID }
	jsonAPIRequest("database/getInfo",getDBInfoParams,function(templateInfo) {
		
		initTemplateDescriptionProperty(templateInfo.databaseInfo)
		initNameProperty(templateInfo.databaseInfo)
		initActiveTemplateProperty(templateInfo.databaseInfo)
		
		var $dialog = $('#templatePropertiesDialog')
		$dialog.modal('show')
		
	})
	
	
	
}