function initDashboardValueGroupingPropertyPanel(panelParams) {
	
	var groupedFieldSelectionElemInfo = createPrefixedTemplElemInfo(panelParams.elemPrefix,"GroupedFieldSelection")
	var groupBySelectionElemInfo = createPrefixedTemplElemInfo(panelParams.elemPrefix,"GroupBySelection")
	var bucketSizeInputElemInfo = createPrefixedTemplElemInfo(panelParams.elemPrefix,"BucketSizeInput")
	var bucketSizeFormGroupSelector = createPrefixedSelector(panelParams.elemPrefix,"BucketSizeFormGroup")
	var saveChangesButtonElemInfo = createPrefixedTemplElemInfo(panelParams.elemPrefix,"ValGroupingPropertiesSaveChangesButton")
	
	var $propertyForm = $(createPrefixedSelector(panelParams.elemPrefix,"ValueGroupingPropertiesForm"))
	
	
	var validationRules = {}	
	validationRules[bucketSizeInputElemInfo.id] = { 
		positiveNumber: {
			depends: function(element) {
				return groupBySelectionElemInfo.val() == "bucket"
			}
		} 
	}
	validationRules[groupBySelectionElemInfo.id] = { required: true }
	validationRules[groupedFieldSelectionElemInfo.id] = { required: true }
	
	var validationSettings = createInlineFormValidationSettings({rules: validationRules })	
	var validator = $propertyForm.validate(validationSettings)
		
	validator.resetForm()	

	
	function toggleBucketSizeForGrouping(grouping) {
		if (grouping == "bucket") {
			$(bucketSizeFormGroupSelector).show()
		} else {
			$(bucketSizeFormGroupSelector).hide()			
		}
	}


	loadFieldInfo(panelParams.databaseID,[fieldTypeAll],function(valueGroupingFieldsByID) {
		populateFieldSelectionMenu(valueGroupingFieldsByID,groupedFieldSelectionElemInfo.selector)
	
		// Initialize the controls to the existing values
		$(groupedFieldSelectionElemInfo.selector).val(panelParams.valGroupingProps.groupValsByFieldID)
		var existingFieldInfo = valueGroupingFieldsByID[panelParams.valGroupingProps.groupValsByFieldID]
		populateDashboardValueGroupingSelection($(groupBySelectionElemInfo.selector),existingFieldInfo.type)
		
		$(groupBySelectionElemInfo.selector).val(panelParams.valGroupingProps.groupValsBy)
		disableButton(saveChangesButtonElemInfo.selector)
		toggleBucketSizeForGrouping(panelParams.valGroupingProps.groupValsBy)
		
		initSelectionChangedHandler(groupedFieldSelectionElemInfo.selector, function(fieldID) {
			if(fieldID in valueGroupingFieldsByID) {
				fieldInfo = valueGroupingFieldsByID[fieldID]			
		    	console.log("select field: field ID = " + fieldID  + " name = " + fieldInfo.name + " type = " + fieldInfo.type)
				populateDashboardValueGroupingSelection($(groupBySelectionElemInfo.selector),fieldInfo.type)
				$(groupBySelectionElemInfo.selector).attr("disabled",false)
				
				// Disable the Save button until a "Group Values By" selection is also made
				disableButton(saveChangesButtonElemInfo.selector)
			}
			
		})

		initSelectionChangedHandler(groupBySelectionElemInfo.selector, function(grouping) {		
			enableButton(saveChangesButtonElemInfo.selector)
			toggleBucketSizeForGrouping(grouping)
		})

		
		initButtonClickHandler(saveChangesButtonElemInfo.selector, function() {
			if($propertyForm.valid()) {
				var newValGroupingParams = {
					groupValsByFieldID: $(groupedFieldSelectionElemInfo.selector).val(),
					groupValsBy: $(groupBySelectionElemInfo.selector).val(),
					groupValsByBucketWidth: Number($(bucketSizeInputElemInfo.selector).val())
				}
				console.log("Saving new value grouping: " + JSON.stringify(newValGroupingParams))
				panelParams.saveValueGroupingFunc(newValGroupingParams)
				disableButton(saveChangesButtonElemInfo.selector)
				
			}
		})
		
	})
	
	
}


