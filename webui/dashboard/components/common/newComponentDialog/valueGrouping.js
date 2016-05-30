
var dashboardComponentValueGroupingPanelID = "dashboardComponentValueGrouping"

function createNewDashboardComponentValueGroupingPanelConfig(elemPrefix) {
	
	var panelSelector = "#" + elemPrefix + "DashboardComponentValueGroupingPanel"
	var groupedFieldSelection = createPrefixedTemplElemInfo(elemPrefix,"GroupedFieldSelection")
	var groupBySelection = createPrefixedTemplElemInfo(elemPrefix,"GroupBySelection")
	var bucketSizeInput = createPrefixedTemplElemInfo(elemPrefix,"BucketSizeInput")
	var sortSelection = createPrefixedTemplElemInfo(elemPrefix,"SortSelection")
	
	var validateWithBucketSize = false
	
	function validateValueGroupingForm() {
		return true
	}
	
	function populateValueGroupingSelection(fieldType) {
		$(groupBySelection.selector).empty()
		$(groupBySelection.selector).append(emptyOptionHTML("Select a grouping"))
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
					transitionToNextWizardDlgPanel(this,barChartYAxisPanelConfig)
				} // if validate panel's form
			},
			"Cancel" : function() { $(this).dialog('close'); },
	 	}, // dialog buttons
		initPanel: function() {
		
			function setValidationRulesWithoutBucketSize() {
				// TODO - Add back validation for xAxisFieldSelection,
				// xAxisSortSelection, and xAxisGroupBySelection
			}
		
			function setValidationRulesWithBucketSize() {
				// TODO - Add back validation for xAxisFieldSelection,
				// xAxisSortSelection, and xAxisGroupBySelection
			
			}
			
			$(groupedFieldSelection.selector).change(function(){
				var fieldID = $(groupedFieldSelection.selector).val()
		        console.log("select field: " + fieldID )
				if(fieldID in newBarChartParams.fieldsByID) {
					fieldInfo = newBarChartParams.fieldsByID[fieldID]			
		        	console.log("select field: field ID = " + fieldID  + " name = " + fieldInfo.name + " type = " + fieldInfo.type)
					populateValueGroupingSelection(fieldInfo.type)
					$(groupBySelection.selector).removeClass("disabled")
				}
		    });
		
			$(groupBySelection.selector).change(function() {
				groupBy = $(groupBySelection.selector).val()
				if(groupBy == "bucket") {
					$(bucketSizeInput.selector).show()
					validateWithBucketSize = false
				}
				else {
					$(bucketSizeInput.selector).hide()
					validateWithBucketSize = true
				}
			});
		
			// The field for entering a bucket size is initially hidden. It is only shown if
			// the group by parameter is set to use a bucket.
			$(bucketSizeInput.selector).hide()
			setValidationRulesWithoutBucketSize()
		
			// Grouping selection is initially disabled until a field is selected
			$(groupBySelection.selector).empty()
			$(groupBySelection.selector).addClass("disabled")
		
			return {}
		}, // init panel
	}
	
	return dashboardComponentValueGroupingPanelConfig
	
} // createNewDashboardComponentValueGroupingPanelConfig

