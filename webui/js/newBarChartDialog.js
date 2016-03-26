
var barChartProgressDivID = '#newBarChartProgress'

var xAxisPanelDivID = "#newBarChartDlgXAxisPanel"

var newBarChartFieldsByID;

var barChartXAxisPanelConfig = {
	divID: xAxisPanelDivID,
	progressPerc:80,
	dlgButtons: { 
		"Next" : function() { 
			if($( xAxisPanelDivID ).form('validate form')) {
				transitionToNextWizardDlgPanel(this,barChartProgressDivID,
						barChartXAxisPanelConfig,barChartYAxisPanelConfig)
			} // if validate panel's form
		},
		"Cancel" : function() { $(this).dialog('close'); },
 	}, // dialog buttons
	initPanel: function() {
		
		function setValidationRulesWithoutBucketSize() {
			$( xAxisPanelDivID ).form({
		    	fields: {
			        xAxisFieldSelection: nonEmptyFieldValidation('Please enter a field name'),
			        xAxisGroupBySelection: nonEmptyFieldValidation('Select a grouping'),
			        xAxisSortSelection: nonEmptyFieldValidation('Choose a sort order')
		     	},
		  	})
			
		}
		
		function setValidationRulesWithBucketSize() {
			$( xAxisPanelDivID ).form({
		    	fields: {
			        xAxisFieldSelection: nonEmptyFieldValidation('Please enter a field name'),
			        xAxisGroupBySelection: nonEmptyFieldValidation('Select a grouping'),
			        xAxisSortSelection: nonEmptyFieldValidation('Choose a sort order'),
					xAxisBucketSizeInput: validPositiveNumberFieldValidation()
		     	},
		  	})
			
		}
		
		var xAxisFieldSelectionID = '#xAxisFieldSelection'
		var xAxisSelectGroupByID = '#xAxisGroupBySelection'
		
		// If the showOnFocus setting is not used, the dropdown will open automatically when the 
		// dialog is opened. 
		$(xAxisFieldSelectionID).dropdown({showOnFocus:false})
		$(xAxisSelectGroupByID).dropdown()
		$('#xAxisSortSelection').dropdown()
		
		function populateXAxisGroupSelection(fieldType) {
			$(xAxisSelectGroupByID).empty()
			$(xAxisSelectGroupByID).append(emptyOptionHTML("Select a grouping"))
			if(fieldType == fieldTypeNumber) {
				$(xAxisSelectGroupByID).append(selectOptionHTML("none","Don't group values"))
				$(xAxisSelectGroupByID).append(selectOptionHTML("bucket","Bucket values"))
			}
			else if (fieldType == fieldTypeText) {
				$(xAxisSelectGroupByID).append(selectOptionHTML("none","Don't group values"))
			}
			else {
				console.log("unrecocognized field type: " + fieldType)
			}
		}
		
		$(xAxisFieldSelectionID).change(function(){
			var fieldID = $(xAxisPanelDivID).form('get value','xAxisFieldSelection')
	        console.log("select field: " + fieldID )
			if(fieldID in newBarChartFieldsByID) {
				fieldInfo = newBarChartFieldsByID[fieldID]			
	        	console.log("select field: field ID = " + fieldID  + " name = " + fieldInfo.name + " type = " + fieldInfo.type)
				populateXAxisGroupSelection(fieldInfo.type)
				$(xAxisSelectGroupByID).removeClass("disabled")
			}
	    });
		
		$(xAxisSelectGroupByID).change(function() {
			groupBy = $(xAxisPanelDivID).form('get value','xAxisGroupBySelection')
			if(groupBy == "bucket") {
				$("#xAxisBucketSize").show()
				setValidationRulesWithBucketSize()
			}
			else {
				$("#xAxisBucketSize").hide()
				setValidationRulesWithoutBucketSize()
			}
		});
		
		// The field for entering a bucket size is initially hidden. It is only shown if
		// the group by parameter is set to use a bucket.
		$("#xAxisBucketSize").hide()
		setValidationRulesWithoutBucketSize()
		
		// Grouping selection is initially disabled until a field is selected
		$(xAxisSelectGroupByID).empty()
		$(xAxisSelectGroupByID).addClass("disabled")
	},	
}


var yAxisPanelDivID = "#newBarChartDlgYAxisPanel"
var barChartYAxisPanelConfig = {
	divID: yAxisPanelDivID,
	progressPerc:80,
	dlgButtons: { 
		"Previous": function() {
			transitionToPrevWizardDlgPanel(this,barChartProgressDivID,
				barChartYAxisPanelConfig,barChartXAxisPanelConfig)	
		 },
		"Done" : function() { 
			if($( yAxisPanelDivID ).form('validate form')) {
				$(this).dialog('close');
			} // if validate panel's form
		},
		"Cancel" : function() { $(this).dialog('close'); },
 	}, // dialog buttons
	
	
	initPanel: function() {
		$( yAxisPanelDivID ).form({
	    	fields: {
		        yAxisFieldSelection: nonEmptyFieldValidation('Select a field'),
		        yAxisSummarySelection: nonEmptyFieldValidation('Choose how to summarize values'),
	     	},
	  	})
		
		var yAxisFieldSelectionID = '#yAxisFieldSelection'
		var yAxisSummarySelectionID = '#yAxisSummarySelection'
		
		$(yAxisFieldSelectionID).dropdown()
		$(yAxisSummarySelectionID).dropdown()
		
		function populateYAxisSummarySelection(fieldType) {
			$(yAxisSummarySelectionID).empty()
			$(yAxisSummarySelectionID).append(emptyOptionHTML("Choose how to summarize values"))
			if(fieldType == fieldTypeNumber) {
				$(yAxisSummarySelectionID).append(selectOptionHTML("count","Count of values"))
				$(yAxisSummarySelectionID).append(selectOptionHTML("sum","Sum of values"))
				$(yAxisSummarySelectionID).append(selectOptionHTML("average","Average of values"))
			}
			else if (fieldType == fieldTypeText) {
				$(yAxisSummarySelectionID).append(selectOptionHTML("count","Count of values"))
			}
			else {
				console.log("unrecocognized field type: " + fieldType)
			}
		}
		
		
		$(yAxisFieldSelectionID).change(function(){
			var fieldID = $(yAxisPanelDivID).form('get value','yAxisFieldSelection')
	        console.log("select field: " + fieldID )
			if(fieldID in newBarChartFieldsByID) {
				fieldInfo = newBarChartFieldsByID[fieldID]			
	        	console.log("select field: field ID = " + fieldID  + " name = " + fieldInfo.name + " type = " + fieldInfo.type)
				
				populateYAxisSummarySelection(fieldInfo.type)
				$(yAxisSummarySelectionID).removeClass("disabled")
			}
	    });
		
		$(yAxisSummarySelectionID).empty()
		$(yAxisSummarySelectionID).addClass("disabled")
		
		
	},	
}


function newBarChart(barChartParams) {
		
	var barChartCreated = false
	var placholderID = barChartParams.containerID
	
	
	openWizardDialog({
		closeFunc: function () {
  		  console.log("Close dialog")
  		  if(!barChartCreated)
  		  {
  			  // If the the bar chart creation is not complete, remove the placeholder
  			  // from the canvas.
  			  $('#'+placeholderID).remove()
  		  }	
		},
		width: 550, height: 500,
		dialogDivID: '#newBarchartDialog',
		panels: [barChartXAxisPanelConfig, barChartYAxisPanelConfig],
		progressDivID: '#newBarChartProgress',
	})
		
					
	
} // newBarChart


function initNewBarChartDialog() {
	initWizardDialog('#newBarchartDialog')
	
	loadFieldInfo(function(fieldsByID) {
		populateFieldSelectionMenu(fieldsByID,'#xAxisFieldSelection')
		populateFieldSelectionMenu(fieldsByID,'#yAxisFieldSelection')
		newBarChartFieldsByID = fieldsByID
	})
	

}
