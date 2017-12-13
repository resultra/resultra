
var dashboardComponentValueGroupingPanelID = "dashboardComponentValueGrouping"

function createNewDashboardComponentValueGroupingPanelConfig(elemPrefix,databaseID) {
	
	var panelSelector = "#" + elemPrefix + "DashboardComponentValueGroupingPanel"
	var groupedFieldSelection = createPrefixedTemplElemInfo(elemPrefix,"NewComponentGroupedFieldSelection")
	var groupBySelection = createPrefixedTemplElemInfo(elemPrefix,"NewComponentGroupBySelection")
	var bucketSizeInput = createPrefixedTemplElemInfo(elemPrefix,"NewComponentBucketSizeInput")
	var $bucketSizeFormComponents = $(panelSelector).find(".valGroupingBucketSizeFormComponents")
	
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
		
	function getPanelValues() {
		var valGrouping = {
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
					$bucketSizeFormComponents.show()
					validateWithBucketSize = true
				}
				else {
					$bucketSizeFormComponents.hide()
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
			$bucketSizeFormComponents.hide()
			validateWithBucketSize = true
			
			revalidateNonEmptyFormFieldOnChange(groupedFieldSelection.selector)
			revalidateNonEmptyFormFieldOnChange(groupBySelection.selector)
			revalidateNonEmptyFormFieldOnChange(bucketSizeInput.selector)		

		}, // init panel
		transitionIntoPanel: function ($dialog) { 
			
			setWizardDialogButtonSet("newDashboardComponentValueGroupingButtons")				
				
			loadSortedFieldInfo(databaseID,[fieldTypeAll],function(sortedGroupingFields) {
				
				var valueGroupingFieldsByID = createFieldsByIDMap(sortedGroupingFields)
				
				var  $groupedFieldSelection = $(groupedFieldSelection.selector)
				
				populateSortedFieldSelectionMenu($groupedFieldSelection,sortedGroupingFields)
				$(groupBySelection.selector).attr("disabled",true)
					
				 $groupedFieldSelection.unbind("change")				
				 $groupedFieldSelection.change(function(){
					var fieldID =  $groupedFieldSelection.val()
			        console.log("select field: " + fieldID )
					if(fieldID in valueGroupingFieldsByID) {
						fieldInfo = valueGroupingFieldsByID[fieldID]			
			        	console.log("select field: field ID = " + fieldID  + " name = " + fieldInfo.name + " type = " + fieldInfo.type)
						populateDashboardValueGroupingSelection($(groupBySelection.selector),fieldInfo.type)
						$(groupBySelection.selector).attr("disabled",false)
					}
			    });
				
				
				
			}) // loadFieldInfo
			
				
		} // transitionIntoPanel
		
	}
	
	return dashboardComponentValueGroupingPanelConfig
	
} // createNewDashboardComponentValueGroupingPanelConfig

