function initDashboardValueGroupingPropertyPanel(panelParams) {
		
	var $propertyForm = $(createPrefixedSelector(panelParams.elemPrefix,"ValueGroupingPropertiesForm"))
	var $bucketStart = $propertyForm.find("input[name=bucketStart]")	
	var $bucketEnd = $propertyForm.find("input[name=bucketEnd]")
	var $bucketSize = $propertyForm.find("input[name=bucketSize]")
	var $selectionInputs = $propertyForm.find(".groupBySelectionInput")
	var $bucketInputs = $propertyForm.find(".bucketInput")
	var $bucketInputGroups = $propertyForm.find(".groupBucketInputs")
	var $fieldSelection = $propertyForm.find("select[name=groupedFieldSelection]")
	var $groupBySelection = $propertyForm.find("select[name=groupBySelection]")
	var $numberFormatSelection = $propertyForm.find("select[name=numberFormat]")
	
	//var validationRules = {}
	var validationRules = {
		bucketSize: { 
			positiveNumber: {
				depends: function(element) { return $groupBySelection.val() === "bucket" }
			}	
		},
		groupedFieldSelection: { required:true },
		groupBySelection: { required: true }
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


	loadSortedFieldInfo(panelParams.databaseID,[fieldTypeAll],function(sortedFields) {
		
		var valueGroupingFieldsByID = createFieldsByIDMap(sortedFields)
		
		populateSortedFieldSelectionMenu($fieldSelection,sortedFields)
	
		// Initialize the controls to the existing values
		$fieldSelection.val(panelParams.valGroupingProps.groupValsByFieldID)
		
		var existingFieldInfo = valueGroupingFieldsByID[panelParams.valGroupingProps.groupValsByFieldID]
		toggleNumberFormatForFieldType(existingFieldInfo.type)
		
		populateDashboardValueGroupingSelection($groupBySelection,existingFieldInfo.type)
		
		$groupBySelection.val(panelParams.valGroupingProps.groupValsBy)
		toggleBucketSizeForGrouping(panelParams.valGroupingProps.groupValsBy)
		
		function saveGroupingIfValid() {
			if($propertyForm.valid()) {
				var newValGroupingParams = {
					groupValsByFieldID: $fieldSelection.val(),
					groupValsBy: $groupBySelection.val(),
					groupValsByBucketWidth: convertStringToNumber($bucketSize.val()),
					bucketStart: convertStringToNumber($bucketStart.val()),
					bucketEnd: convertStringToNumber($bucketEnd.val()),
					numberFormat: $numberFormatSelection.val()
				}
				console.log("Saving new value grouping: " + JSON.stringify(newValGroupingParams))
				panelParams.saveValueGroupingFunc(newValGroupingParams)
			}
		}
		
		initSelectControlChangeHandler($fieldSelection, function(fieldID) {
			if(fieldID in valueGroupingFieldsByID) {
				fieldInfo = valueGroupingFieldsByID[fieldID]			
		    	console.log("select field: field ID = " + fieldID  + " name = " + fieldInfo.name + " type = " + fieldInfo.type)
				populateDashboardValueGroupingSelection($groupBySelection,fieldInfo.type)
				$groupBySelection.attr("disabled",false)		
				toggleNumberFormatForFieldType(fieldInfo.type)
			}
			saveGroupingIfValid() 
		})
		
		initSelectControlChangeHandler($numberFormatSelection, function(fieldID) {
			saveGroupingIfValid()
		})
		

		initSelectControlChangeHandler($groupBySelection, function(grouping) {		
			toggleBucketSizeForGrouping(grouping)
			saveGroupingIfValid() 
		})
			
		$bucketInputs.blur(function() { saveGroupingIfValid() })
		
	})
	
	
}


