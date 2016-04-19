
var newBarChartParams = {}


function initNewBarChartDialog(dashboardID) {

	newBarChartParams.dashboardID = dashboardID
	newBarChartParams.progressDivID = '#newBarChartProgress'	
	
	initWizardDialog('#newBarchartDialog')
	
	loadFieldInfo(function(fieldsByID) {
		populateFieldSelectionMenu(fieldsByID,'#xAxisFieldSelection')
		populateFieldSelectionMenu(fieldsByID,'#yAxisFieldSelection')
		newBarChartParams.fieldsByID = fieldsByID
	})
}

function saveNewBarChart() {
	console.log("Saving new bar chart: dashboard ID = " + newBarChartParams.dashboardID )
	

	
	var formID = '#newBarchartDialog'
		
	var saveNewBarChartParams = {
		parentDashboardID: newBarChartParams.dashboardID,
		xAxisVals: {
			fieldID: getFormStringValue(formID,'xAxisFieldSelection'),
			groupValsBy: getFormStringValue(formID,'xAxisGroupBySelection'),
			groupByValBucketWidth: getFormFloatValue(formID,'xAxisBucketSizeInput'),		
		}, // xAxisVals
		xAxisSortValues: getFormStringValue(formID,'xAxisSortSelection'),
		yAxisVals: {
			fieldID: getFormStringValue(formID,'yAxisFieldSelection'),
			summarizeValsWith: getFormStringValue(formID,'yAxisSummarySelection')
		}, // yAxisVals
		geometry: newBarChartParams.geometry
	}
	
	
	console.log("saveNewBarChart: new bar chart params:  " + JSON.stringify(saveNewBarChartParams) )
	jsonAPIRequest("newBarChart",saveNewBarChartParams,function(barChartRef) {
		console.log("saveNewBarChart: bar chart saved: new bar chart ID = " + barChartRef.barChartID)
		
		// Replace the placholder ID with the instantiated bar chart's unique ID.
		 $('#'+newBarChartParams.placeholderID).attr("id",barChartRef.barChartID)

		newBarChartParams.barChartCreated = true;		
		newBarChartParams.dialog.dialog("close")
		
		barChartDataParams = { 
			parentDashboardID: newBarChartParams.dashboardID,
			barChartID: barChartRef.barChartID
		}
		jsonAPIRequest("getBarChartData",barChartDataParams,function(barChartData) {
			initBarChartData(barChartData)
		})
	})
}

var barChartXAxisPanelConfig = {
	divID: "#newBarChartDlgXAxisPanel",
	panelID: "barChartXAxis",
	progressPerc:80,
	dlgButtons: { 
		"Next" : function() { 
			if($( "#newBarChartDlgXAxisPanel" ).form('validate form')) {
				transitionToNextWizardDlgPanel(this,newBarChartParams.progressDivID,
						barChartXAxisPanelConfig,barChartYAxisPanelConfig)
			} // if validate panel's form
		},
		"Cancel" : function() { $(this).dialog('close'); },
 	}, // dialog buttons
	initPanel: function() {
		
		function setValidationRulesWithoutBucketSize() {
			$( "#newBarChartDlgXAxisPanel" ).form({
		    	fields: {
			        xAxisFieldSelection: nonEmptyFieldValidation('Please enter a field name'),
			        xAxisGroupBySelection: nonEmptyFieldValidation('Select a grouping'),
			        xAxisSortSelection: nonEmptyFieldValidation('Choose a sort order')
		     	},
		  	})		
		}
		
		function setValidationRulesWithBucketSize() {
			$( "#newBarChartDlgXAxisPanel" ).form({
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
			var fieldID = $("#newBarChartDlgXAxisPanel").form('get value','xAxisFieldSelection')
	        console.log("select field: " + fieldID )
			if(fieldID in newBarChartParams.fieldsByID) {
				fieldInfo = newBarChartParams.fieldsByID[fieldID]			
	        	console.log("select field: field ID = " + fieldID  + " name = " + fieldInfo.name + " type = " + fieldInfo.type)
				populateXAxisGroupSelection(fieldInfo.type)
				$(xAxisSelectGroupByID).removeClass("disabled")
			}
	    });
		
		$(xAxisSelectGroupByID).change(function() {
			groupBy = $("#newBarChartDlgXAxisPanel").form('get value','xAxisGroupBySelection')
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


var barChartYAxisPanelConfig = {
	divID: "#newBarChartDlgYAxisPanel",
	panelID: "barChartYAxis",
	progressPerc:80,
	dlgButtons: { 
		"Previous": function() {
			transitionToPrevWizardDlgPanel(this,newBarChartParams.progressDivID,
				barChartYAxisPanelConfig,barChartXAxisPanelConfig)	
		 },
		"Done" : function() { 
			if($( "#newBarChartDlgYAxisPanel" ).form('validate form')) {
				saveNewBarChart()
			} // if validate panel's form
		},
		"Cancel" : function() { $(this).dialog('close'); },
 	}, // dialog buttons
	
	
	initPanel: function() {
		$( "#newBarChartDlgYAxisPanel" ).form({
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
			var fieldID = $("#newBarChartDlgYAxisPanel").form('get value','yAxisFieldSelection')
	        console.log("select field: " + fieldID )
			if(fieldID in newBarChartParams.fieldsByID) {
				fieldInfo = newBarChartParams.fieldsByID[fieldID]			
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
		
	newBarChartParams.placeholderID = barChartParams.containerID
	newBarChartParams.geometry = barChartParams.geometry
	newBarChartParams.barChartCreated = false
	newBarChartParams.dialog = $('#newBarchartDialog')

	openWizardDialog({
		closeFunc: function () {
  		  console.log("Close dialog")
  		  if(!newBarChartParams.barChartCreated)
  		  {
  			  // If the the bar chart creation is not complete, remove the placeholder
  			  // from the canvas.
  			  $('#'+newBarChartParams.placeholderID).remove()
  		  }	
		},
		width: 550, height: 500,
		dialogDivID: '#newBarchartDialog',
		panels: [barChartXAxisPanelConfig, barChartYAxisPanelConfig],
		progressDivID: '#newBarChartProgress',
	})
		
} // newBarChart




