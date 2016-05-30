
var dashboardComponentValueGroupingPanelID = "dashboardComponentValueGrouping"

function createNewDashboardComponentValueGroupingPanelConfig(elemPrefix) {
	
	var panelSelector = "#" + elemPrefix + "DashboardComponentValueGroupingPanel"
	var groupedFieldSelection = createPrefixedTemplElemInfo(elemPrefix,"GroupedFieldSelection")
	var groupBySelection = createPrefixedTemplElemInfo(elemPrefix,"GroupBySelection")
	var bucketSizeInput = createPrefixedTemplElemInfo(elemPrefix,"BucketSizeInput")
	var sortSelection = createPrefixedTemplElemInfo(elemPrefix,"SortSelection")
	
	var validateWithBucketSize = false
	
	
	function validateValueGroupingForm() {
		
		var validationResults = true
		
		// Any one of the fields not passing validation makes the whole validation fail
		if(!validateNonEmptyFormField(groupedFieldSelection.selector)) { validationResults = false }
		if(!validateNonEmptyFormField(groupBySelection.selector)) { validationResults = false }
		if(!validateNonEmptyFormField(sortSelection.selector)) { validationResults = false }
		
		if (validateWithBucketSize) {
			if(!validateNonEmptyFormField(bucketSizeInput.selector)) { validationResults = false }
		}
		
		return validationResults
	}
	
	function populateValueGroupingSelection(fieldType) {
		$(groupBySelection.selector).empty()
		$(groupBySelection.selector).append(defaultSelectOptionPromptHTML("Select a grouping"))
		if(fieldType == fieldTypeNumber) {
			$(groupBySelection.selector).append(selectOptionHTML("none","Don't group values"))
			$(groupBySelection.selector).append(selectOptionHTML("bucket","Bucket values"))
		}
		else if (fieldType == fieldTypeText) {
			$(groupBySelection.selector).append(selectOptionHTML("none","Don't group values"))
		}
		else {
			console.log("unrecocognized field type: " + fieldType)
		}
	}
	
	var dashboardComponentValueGroupingPanelConfig = {
		divID: panelSelector,
		panelID: dashboardComponentValueGroupingPanelID,
		progressPerc:40,
		dlgButtons: { 
			"Previous": function() {
				transitionToPrevWizardDlgPanelByPanelID(this,dashboardComponentSelectTablePanelID)	
			 },
			"Next" : function() { 
				if(validateValueGroupingForm()) {
					console.log("Value grouping panel validated")
					
					var panelData = {
						valGrouping: {
							fieldID: groupedFieldSelection.val(),
							groupValsBy: groupBySelection.val(),
							groupByValBucketWidth: bucketSizeInput.val()
						},
						sortOrder: sortSelection.val()
					}
					setWizardDialogPanelData($(this),elemPrefix,dashboardComponentValueGroupingPanelID,panelData)
					
				//	transitionToNextWizardDlgPanel(this,barChartYAxisPanelConfig)
				} // if validate panel's form
			},
			"Cancel" : function() { $(this).dialog('close'); },
	 	}, // dialog buttons
		initPanel: function() {
				
			$(groupBySelection.selector).change(function() {
				var groupBy = groupBySelection.val()
			    console.log(groupBySelection.id)
				console.log("Value grouping changed: " + groupBy)
				if(groupBy == "bucket") {
					$(bucketSizeInput.selector).removeClass("hidden")
					validateWithBucketSize = true
				}
				else {
					$(bucketSizeInput.selector).addClass("hidden")
					validateWithBucketSize = false
				}
			});
		
			// The field for entering a bucket size is initially hidden. It is only shown if
			// the group by parameter is set to use a bucket.
			$(bucketSizeInput.selector).addClass("hidden")
			validateWithBucketSize = true
			
			revalidateNonEmptyFormFieldOnChange(groupedFieldSelection.selector)
			revalidateNonEmptyFormFieldOnChange(groupBySelection.selector)
			revalidateNonEmptyFormFieldOnChange(sortSelection.selector)
			revalidateNonEmptyFormFieldOnChange(bucketSizeInput.selector)		
		
			return {}
		}, // init panel
		transitionIntoPanel: function ($dialog) { 
			
			var selectedTableID = getWizardDialogPanelData($dialog,
					elemPrefix,dashboardComponentSelectTablePanelID)
			console.log("Transitioning into value grouping panel: selected table ID = " + selectedTableID)
			loadFieldInfo(selectedTableID,[fieldTypeAll],function(valueGroupingFieldsByID) {
				
				populateFieldSelectionMenu(valueGroupingFieldsByID,groupedFieldSelection.selector)
				$(groupBySelection.selector).attr("disabled",true)
					
				$(groupedFieldSelection.selector).unbind("change")				
				$(groupedFieldSelection.selector).change(function(){
					var fieldID = $(groupedFieldSelection.selector).val()
			        console.log("select field: " + fieldID )
					if(fieldID in valueGroupingFieldsByID) {
						fieldInfo = valueGroupingFieldsByID[fieldID]			
			        	console.log("select field: field ID = " + fieldID  + " name = " + fieldInfo.name + " type = " + fieldInfo.type)
						populateValueGroupingSelection(fieldInfo.type)
						$(groupBySelection.selector).attr("disabled",false)
					}
			    });
				
			}) // loadFieldInfo
			
				
		} // transitionIntoPanel
		
	}
	
	return dashboardComponentValueGroupingPanelConfig
	
} // createNewDashboardComponentValueGroupingPanelConfig

