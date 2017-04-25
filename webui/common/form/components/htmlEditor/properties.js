
function loadHtmlEditorProperties($editor, htmlEditorRef) {
	console.log("Loading html editor properties")
	
	var elemPrefix = "htmlEditor_"
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for html editor")
		var formatParams = {
			parentFormID: htmlEditorRef.parentFormID,
			htmlEditorID: htmlEditorRef.htmlEditorID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/htmlEditor/setLabelFormat", formatParams, function(updatedEditor) {
			setEditorComponentLabel($editor,updatedEditor)
			setContainerComponentInfo($editor,updatedEditor,updatedEditor.htmlEditorID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: htmlEditorRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: htmlEditorRef.parentFormID,
			htmlEditorID: htmlEditorRef.htmlEditorID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/htmlEditor/setVisibility",params,function(updatedEditor) {
			setContainerComponentInfo($editor,updatedEditor,updatedEditor.htmlEditorID)	
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: htmlEditorRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
	
	var permissionParams = {
		elemPrefix: elemPrefix,
		initialVal: htmlEditorRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: htmlEditorRef.parentFormID,
				htmlEditorID: htmlEditorRef.htmlEditorID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/htmlEditor/setPermissions",params,function(updatedEditor) {
				setContainerComponentInfo($editor,updatedEditor,updatedEditor.htmlEditorID)	
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(permissionParams)
	
	
	initCheckboxChangeHandler('#adminHtmlEditorComponentValidationRequired', 
				htmlEditorRef.properties.validation.valueRequired, function (newVal) {
			
		var validationProps = {
			valueRequired: newVal
		}		
			
		var validationParams = {
			parentFormID: htmlEditorRef.parentFormID,
			htmlEditorID: htmlEditorRef.htmlEditorID,
			validation: validationProps
		}
		console.log("Setting new validation settings: " + JSON.stringify(validationParams))

		jsonAPIRequest("frm/htmlEditor/setValidation",validationParams,function(updatedEditor) {
				setContainerComponentInfo($editor,updatedEditor,updatedEditor.htmlEditorID)	
		})
	})
	
	
	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: htmlEditorRef.parentFormID,
		componentID: htmlEditorRef.htmlEditorID,
		componentLabel: 'editor',
		$componentContainer: $editor
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	

	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#htmlEditorProps')
		
	toggleFormulaEditorForField(htmlEditorRef.properties.fieldID)
	
}