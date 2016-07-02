
var dashboardComponentValueGroupingPanelID = "dashboardComponentValueGrouping"

var valueGroupingDialogPanelTableID;

function createNewDashboardComponentValueGroupingPanelConfig(elemPrefix) {
	
	var panelSelector = "#" + elemPrefix + "DashboardComponentValueGroupingPanel"
	var groupedFieldSelection = createPrefixedTemplElemInfo(elemPrefix,"GroupedFieldSelection")
	var groupBySelection = createPrefixedTemplElemInfo(elemPrefix,"GroupBySelection")
	var bucketSizeInput = createPrefixedTemplElemInfo(elemPrefix,"BucketSizeInput")
	
	var validateWithBucketSize = false
	
	
	function validateValueGroupingForm() {
		
		var validationResults = true
		
		// Any one of the fields not passing validation makes the whole validation fail
		if(!validateNonEmptyFormField(groupedFieldSelection.selector)) { validationResults = false }
		if(!validateNonEmptyFormField(groupBySelection.selector)) { validationResults = false }	
		if (validateWithBucketSize) {
			if(!validateNonEmptyFormField(bucketSizeInput.selector)) { validationResults = false }
			if(!validateNumberFormField(bucketSizeInput.selector))	{ validationResults = false }
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
	
	function getPanelValues() {
		var valGrouping = {
			fieldParentTableID: valueGroupingDialogPanelTableID,
			fieldID: groupedFieldSelection.val(),
			groupValsBy: groupBySelection.val(),
			groupByValBucketWidth: Number(bucketSizeInput.val())
		}
		return valGrouping
	}
	
	var dashboardComponentValueGroupingPanelConfig = {
		divID: panelSelector,
		panelID: dashboardComponentValueGroupingPanelID,
		progressPerc:40,
		getPanelVals:getPanelValues,
		initPanel: function($dialog) {
				
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
			
			var nextButtonSelector = '#' + elemPrefix + 'NewDashboardComponentValueGroupingNextButton'
			initButtonClickHandler(nextButtonSelector,function() {
				if(validateValueGroupingForm()) {				
					transitionToNextWizardDlgPanelByID($dialog,dashboardComponentValueSummaryPanelID)
				} // if validate panel's form
			})

			var prevButtonSelector = '#' + elemPrefix + 'NewDashboardComponentValueGroupingPrevButton'
			initButtonClickHandler(prevButtonSelector,function() {
				transitionToPrevWizardDlgPanelByPanelID($dialog,dashboardComponentSelectTablePanelID)	
			})

		
			// The field for entering a bucket size is initially hidden. It is only shown if
			// the group by parameter is set to use a bucket.
			$(bucketSizeInput.selector).addClass("hidden")
			validateWithBucketSize = true
			
			revalidateNonEmptyFormFieldOnChange(groupedFieldSelection.selector)
			revalidateNonEmptyFormFieldOnChange(groupBySelection.selector)
			revalidateNonEmptyFormFieldOnChange(bucketSizeInput.selector)		

		}, // init panel
		transitionIntoPanel: function ($dialog) { 
			
			setWizardDialogButtonSet("newDashboardComponentValueGroupingButtons")
			
			var selectedTableID = getWizardDialogPanelVals($dialog,dashboardComponentSelectTablePanelID)
			console.log("Transitioning into value grouping panel: selected table ID = " + selectedTableID)
			valueGroupingDialogPanelTableID = 	selectedTableID
				
				
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

