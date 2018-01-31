function initDashboardValueGroupingPropertyPanel(panelParams) {
		
	var $propertyForm = $(createPrefixedSelector(panelParams.elemPrefix,"ValueGroupingPropertiesForm"))
	
	// Main drop-down for either time increments or fields
	var $groupByFieldOrTimeIncrementSelection = $propertyForm.find("select[name=groupedFieldOrTimeIntervalSelection]")
	
	
	// Field grouping inputs
	var $fieldGroupingInputs = $propertyForm.find(".fieldGroupingInput")
	var $bucketStart = $propertyForm.find("input[name=bucketStart]")	
	var $bucketEnd = $propertyForm.find("input[name=bucketEnd]")
	var $bucketSize = $propertyForm.find("input[name=bucketSize]")
	var $selectionInputs = $propertyForm.find(".groupBySelectionInput")
	var $bucketInputs = $propertyForm.find(".bucketInput")
	var $bucketInputGroups = $propertyForm.find(".groupBucketInputs")
	var $groupBySelection = $propertyForm.find("select[name=groupBySelection]")
	var $numberFormatSelection = $propertyForm.find("select[name=numberFormat]")
	
	
	// Time grouping inputs
	var $timeGroupingInputs = $propertyForm.find(".timeGroupingInput")
	var $timeRangeSelection = $propertyForm.find("select[name=timeRangeSelection]")
	
	var $includeBlankCheckbox = $propertyForm.find(".includeBlankValuesCheckbox")
	
	function timeIncrementIsSelected() {
		var $selected = $groupByFieldOrTimeIncrementSelection.find(":selected")
		if ($selected.hasClass("timeIncrementSelection")) {
			return true
		} else {
			return false
		}
	}
		
	//var validationRules = {}
	var validationRules = {
		bucketSize: { 
			positiveNumber: {
				depends: function(element) { return $groupBySelection.val() === "bucket" }
			}	
		},
		timeRangeSelection: {
			required: {
				depends: function(element) { return timeIncrementIsSelected() }
			}
		},
		groupedFieldOrTimeIntervalSelection: { 
			required: true
		},
		groupBySelection: { 			
			required: {
				depends: function(element) { return !timeIncrementIsSelected() }
			} 
		}
	}
	var validationSettings = createInlineFormValidationSettings({rules: validationRules })	
	var validator = $propertyForm.validate(validationSettings)
		
	var valGrouping = panelParams.valGroupingProps
		
	$bucketSize.val(valGrouping.groupValsByBucketWidth)	
	$bucketStart.val(valGrouping.bucketStart)	
	$bucketEnd.val(valGrouping.bucketEnd)
	$numberFormatSelection.val(valGrouping.numberFormat)
	validator.resetForm()	

	
	var $numberFormatGroup =  $propertyForm.find(".numberFieldFormatGroup")
	
	function toggleBucketSizeForGrouping(grouping) {
		if (grouping == "bucket") {
			$bucketInputGroups.show()
		} else {
			$bucketInputGroups.hide()		
		}
	}
	
	function toggleNumberFormatForFieldType(fieldType) {
		if (fieldType == fieldTypeNumber) {
			$numberFormatGroup.show()
		} else {
			$numberFormatGroup.hide()
		}	
	}
	
	function toggleBucketSizeForFieldType(fieldType) {
		if (fieldType == fieldTypeNumber) {
			$bucketInputGroups.show()
		} else {
			$bucketInputGroups.hide()		
		}	
	}


	loadSortedFieldInfo(panelParams.databaseID,[fieldTypeNumber,fieldTypeBool,fieldTypeTime,fieldTypeText],function(sortedFields) {
		
		var valueGroupingFieldsByID = createFieldsByIDMap(sortedFields)
		
		var $fieldOptGroup = $groupByFieldOrTimeIncrementSelection.find(".fieldSelectionOptGroup")
		populateSortedFieldSelectionOptGroup($fieldOptGroup,sortedFields)
	
		// Initialize the controls to the existing values
		$groupByFieldOrTimeIncrementSelection.val(panelParams.valGroupingProps.groupValsByFieldID)
		
		var existingFieldInfo = valueGroupingFieldsByID[panelParams.valGroupingProps.groupValsByFieldID]
		toggleNumberFormatForFieldType(existingFieldInfo.type)
		
		populateDashboardValueGroupingSelection($groupBySelection,existingFieldInfo.type)
		
		$groupBySelection.val(panelParams.valGroupingProps.groupValsBy)
		toggleBucketSizeForGrouping(panelParams.valGroupingProps.groupValsBy)
		
		function saveGroupingIfValid() {
			
			if($propertyForm.valid()) {
				if(timeIncrementIsSelected()) {
					var newValGroupingParams = {
						groupValsByTime: $groupByFieldOrTimeIncrementSelection.val(),
						timeRange: $timeRangeSelection.val(),
						includeBlank: $includeBlankCheckbox.prop("checked")
					}
					console.log("Saving new time increment grouping: " + JSON.stringify(newValGroupingParams))
				
				} else {
					var newValGroupingParams = {
						groupValsByFieldID: $groupByFieldOrTimeIncrementSelection.val(),
						groupValsBy: $groupBySelection.val(),
						groupValsByBucketWidth: convertStringToNumber($bucketSize.val()),
						bucketStart: convertStringToNumber($bucketStart.val()),
						bucketEnd: convertStringToNumber($bucketEnd.val()),
						numberFormat: $numberFormatSelection.val(),
						includeBlank: $includeBlankCheckbox.prop("checked")
					}
					console.log("Saving new field value grouping: " + JSON.stringify(newValGroupingParams))
		// TODO - Re-enable after finishing this form.
		//			panelParams.saveValueGroupingFunc(newValGroupingParams)
				
				} // field selected
			} // if form valid
			
		}
		
		// Overall selection of either a time-based increment or a field selection.
		initSelectControlChangeHandler($groupByFieldOrTimeIncrementSelection, function(fieldIDOrTimeIncrement) {
			if(timeIncrementIsSelected()) {
				var timeIncrement = fieldIDOrTimeIncrement
				$timeGroupingInputs.show()
				$fieldGroupingInputs.hide()
				console.log("selected time increment: " + timeIncrement)
			} else {
				$timeGroupingInputs.hide()
				$fieldGroupingInputs.show()
				var fieldID = fieldIDOrTimeIncrement
				if(fieldID in valueGroupingFieldsByID) {
					fieldInfo = valueGroupingFieldsByID[fieldID]			
			    	console.log("selected field: field ID = " + fieldID  + " name = " + fieldInfo.name + " type = " + fieldInfo.type)
					populateDashboardValueGroupingSelection($groupBySelection,fieldInfo.type)
					$groupBySelection.attr("disabled",false)		
					toggleNumberFormatForFieldType(fieldInfo.type)
					$bucketInputGroups.hide() // initially hide bucket inputs when selecting a new field
				}
			}
			saveGroupingIfValid() 		
		})
		
		
		// Changing any of the "sub controls" (for field value or time increment) triggers saving of the grouping if setting
		// that control makes the grouping valid.
		initSelectControlChangeHandler($numberFormatSelection, function(fieldID) {
			saveGroupingIfValid()
		})
		
		initSelectControlChangeHandler($groupBySelection, function(grouping) {		
			toggleBucketSizeForGrouping(grouping)
			saveGroupingIfValid() 
		})
		
		initCheckboxControlChangeHandler($includeBlankCheckbox,valGrouping.includeBlank,function(includeBlank) {
			saveGroupingIfValid()			
		})

		initSelectControlChangeHandler($timeRangeSelection, function(timeRange) {
			saveGroupingIfValid()
		})
			
		$bucketInputs.blur(function() { saveGroupingIfValid() })
		
	})
	
	
}


