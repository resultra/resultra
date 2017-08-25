function loadSelectionProperties($selection,selectionRef) {
	console.log("loading selection properties")
		
	var elemPrefix = "selection_"
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentFormID: selectionRef.parentFormID,
			selectionID: selectionRef.selectionID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/selection/setLabelFormat", formatParams, function(updatedSelection) {
			setSelectionComponentLabel($selection,updatedSelection)
			setContainerComponentInfo($selection,updatedSelection,updatedSelection.selectionID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: selectionRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: selectionRef.parentFormID,
			selectionID: selectionRef.selectionID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/selection/setVisibility",params,function(updatedSelection) {
			setContainerComponentInfo($selection,updatedSelection,updatedSelection.selectionID)	
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: selectionRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)

	var permissionParams = {
		elemPrefix: elemPrefix,
		initialVal: selectionRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: selectionRef.parentFormID,
				selectionID: selectionRef.selectionID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/selection/setPermissions",params,function(updatedSelection) {
				setContainerComponentInfo($selection,updatedSelection,updatedSelection.selectionID)	
				initSelectionComponentClearValueButton($selection,updatedSelection)
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(permissionParams)

	function saveValueList(valueListID) {
		var setValueListParams = {
			parentFormID: selectionRef.parentFormID,
			selectionID: selectionRef.selectionID,
			valueListID: valueListID
		}
		jsonAPIRequest("frm/selection/setValueList", setValueListParams, function(updatedSelection) {
			setContainerComponentInfo($selection,updatedSelection,updatedSelection.selectionID)	
		})			
	}
	var valueListPropertyParams = {
		elemPrefix: elemPrefix,
		databaseID: designFormContext.databaseID,
		saveValueListCallback: saveValueList,
		defaultValueListID: selectionRef.properties.valueListID
	}
	initValueListSelectionPropertyPanel(valueListPropertyParams)



	var clearValueParams = {
		initialVal: selectionRef.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentFormID: selectionRef.parentFormID,
				selectionID: selectionRef.selectionID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("frm/selection/setClearValueSupported",formatParams,function(updatedSelection) {
				setContainerComponentInfo($selection,updatedSelection,updatedSelection.selectionID)	
				initSelectionComponentClearValueButton($selection,updatedSelection)
			})
		}
	}
	initClearValueProps(clearValueParams)


	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: selectionRef.parentFormID,
		componentID: selectionRef.selectionID,
		componentLabel: 'selection',
		$componentContainer: $selection
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#selectionProps')
		
	toggleFormulaEditorForField(selectionRef.properties.fieldID)
		
}