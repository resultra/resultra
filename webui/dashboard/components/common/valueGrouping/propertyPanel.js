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
	validator.resetForm()	

	
	function toggleBucketSizeForGrouping(grouping) {
		if (grouping == "bucket") {
			$bucketInputGroups.show()
		} else {
			$bucketInputGroups.hide()		
		}
	}


	loadFieldInfo(panelParams.databaseID,[fieldTypeAll],function(valueGroupingFieldsByID) {
		populateFieldSelectionControlMenu(valueGroupingFieldsByID,$fieldSelection)
	
		// Initialize the controls to the existing values
		$fieldSelection.val(panelParams.valGroupingProps.groupValsByFieldID)
		
		var existingFieldInfo = valueGroupingFieldsByID[panelParams.valGroupingProps.groupValsByFieldID]
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
					bucketEnd: convertStringToNumber($bucketEnd.val())
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
			}
			saveGroupingIfValid() 
		})

		initSelectControlChangeHandler($groupBySelection, function(grouping) {		
			toggleBucketSizeForGrouping(grouping)
			saveGroupingIfValid() 
		})
			
		$bucketInputs.blur(function() { saveGroupingIfValid() })
		
	})
	
	
}


