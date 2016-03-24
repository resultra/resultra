
var barChartProgressDivID = '#newBarChartProgress'

var barChartXAxisPanelConfig = {
	divID: "#newBarChartDlgXAxisPanel",
	progressPerc:80,
	dlgButtons: { 
		"Next" : function() { 
			if($( "#newBarChartDlgXAxisPanel" ).form('validate form')) {
				transitionToNextWizardDlgPanel(this,barChartProgressDivID,
						barChartXAxisPanelConfig,barChartYAxisPanelConfig)
			} // if validate panel's form
		},
		"Cancel" : function() { $(this).dialog('close'); },
 	}, // dialog buttons
	initPanel: function() {
		$('#xAxisFieldSelection').dropdown()
		$('#xAxisGroupBySelection').dropdown()
		$('#xAxisSortSelection').dropdown()
	},	
}


var barChartYAxisPanelConfig = {
	divID: "#newBarChartDlgYAxisPanel",
	progressPerc:80,
	dlgButtons: { 
		"Previous": function() {
			transitionToPrevWizardDlgPanel(this,barChartProgressDivID,
				barChartYAxisPanelConfig,barChartXAxisPanelConfig)	
		 },
		"Done" : function() { 
			if($( "#newBarChartDlgYAxisPanel" ).form('validate form')) {
				$(this).dialog('close');
			} // if validate panel's form
		},
		"Cancel" : function() { $(this).dialog('close'); },
 	}, // dialog buttons
	initPanel: function() {
		$('#yAxisFieldSelection').dropdown()
		$('#yAxisSummarySelection').dropdown()
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
}
