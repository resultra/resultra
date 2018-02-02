function initDashboardValueGroupingPropertyPanel(panelParams) {
		
	var $propertyForm = $(createPrefixedSelector(panelParams.elemPrefix,"ValueGroupingPropertiesForm"))
	
	var valGrouping = panelParams.valGroupingProps
	
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
			
	function initControlsForValGrouping(valueGroupingFieldsByID) {
		
		if (valGrouping !== null) {
			
			$bucketSize.val(valGrouping.groupValsByBucketWidth)	
			$bucketStart.val(valGrouping.bucketStart)	
			$bucketEnd.val(valGrouping.bucketEnd)
			$numberFormatSelection.val(valGrouping.numberFormat)
			$timeRangeSelection.val(valGrouping.timeRange)
			
			if (valGrouping.groupValsByFieldID !== undefined) {
			
				$groupByFieldOrTimeIncrementSelection.val(valGrouping.groupValsByFieldID)
				var existingFieldInfo = valueGroupingFieldsByID[valGrouping.groupValsByFieldID]
			
				populateDashboardValueGroupingSelection($groupBySelection,existingFieldInfo.type)
				$groupBySelection.val(valGrouping.groupValsBy)
			
				$timeGroupingInputs.hide()
				$fieldGroupingInputs.show()
				toggleNumberFormatForFieldType(existingFieldInfo.type)
				toggleBucketSizeForGrouping(valGrouping.groupValsBy)
			
			
			
			} else {
			
				$timeGroupingInputs.show()
				$fieldGroupingInputs.hide()
			
				$groupByFieldOrTimeIncrementSelection.val(valGrouping.groupValsByTimeIncrement)
				$timeRangeSelection.val(valGrouping.timeRange)
			
			}	
			
		} else {
				$timeGroupingInputs.hide()
				$fieldGroupingInputs.hide()
			
		}
		
		
		
	}
	
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

	function getRowGrouping() {
		if($propertyForm.valid()) {
			if(timeIncrementIsSelected()) {
				var newValGroupingParams = {
					groupValsByTimeIncrement: $groupByFieldOrTimeIncrementSelection.val(),
					timeRange: $timeRangeSelection.val(),
					includeBlank: $includeBlankCheckbox.prop("checked")
				}
				return newValGroupingParams
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
				return newValGroupingParams			
			} // field selected
		} else {
			return null
		}
	}

	loadSortedFieldInfo(panelParams.databaseID,[fieldTypeNumber,fieldTypeBool,fieldTypeTime,fieldTypeText],function(sortedFields) {
		
		var valueGroupingFieldsByID = createFieldsByIDMap(sortedFields)
		
		var $fieldOptGroup = $groupByFieldOrTimeIncrementSelection.find(".fieldSelectionOptGroup")
		populateSortedFieldSelectionOptGroup($fieldOptGroup,sortedFields)
	
		initControlsForValGrouping(valueGroupingFieldsByID)
		
		
		function saveGroupingIfValid() {
				
			var rowGrouping = getRowGrouping()
			if (rowGrouping !== null) {
					console.log("Saving new time increment grouping: " + JSON.stringify(rowGrouping))
					panelParams.saveValueGroupingFunc(rowGrouping)		
			}
			
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
		
		var initIncludeBlank = false
		if (valGrouping !== null) {
			initIncludeBlank = valGrouping.includeBlank
		}
		
		initCheckboxControlChangeHandler($includeBlankCheckbox,initIncludeBlank,function(includeBlank) {
			saveGroupingIfValid()			
		})

		initSelectControlChangeHandler($timeRangeSelection, function(timeRange) {
			saveGroupingIfValid()
		})
			
		$bucketInputs.blur(function() { saveGroupingIfValid() })
		
	})
	
	this.getRowGrouping = getRowGrouping
	
	
}


