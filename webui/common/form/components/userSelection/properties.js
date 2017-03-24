function loadUserSelectionProperties($userSelection,userSelectionRef) {
	console.log("Loading user selection properties")
	
	var elemPrefix = "userSelection_"
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for user selection")
		var formatParams = {
			parentFormID: userSelectionRef.parentFormID,
			userSelectionID: userSelectionRef.userSelectionID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/userSelection/setLabelFormat", formatParams, function(updatedUserSelection) {
			setUserSelectionComponentLabel($userSelection,updatedUserSelection)
			setContainerComponentInfo($userSelection,updatedUserSelection,updatedUserSelection.userSelectionID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: userSelectionRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: userSelectionRef.parentFormID,
			userSelectionID: userSelectionRef.userSelectionID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/userSelection/setVisibility",params,function(updatedUserSelection) {
			setContainerComponentInfo($userSelection,updatedUserSelection,updatedUserSelection.userSelectionID)
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: userSelectionRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
		
	var permissionParams = {
		elemPrefix: elemPrefix,
		initialVal: userSelectionRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: userSelectionRef.parentFormID,
				userSelectionID: userSelectionRef.userSelectionID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/userSelection/setPermissions",params,function(updatedUserSelection) {
				setContainerComponentInfo($userSelection,updatedUserSelection,updatedUserSelection.userSelectionID)
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(permissionParams)


	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#userSelectionProps')
		
	toggleFormulaEditorForField(userSelectionRef.properties.fieldID)
	
}