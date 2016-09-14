
var newBarChartParams = {}
var barChartElemPrefix = "barChart_"

function initNewBarChartDialog(dashboardContext) {

	newBarChartParams.dashboardID = dashboardContext.dashboardID
	newBarChartParams.progressDivID = '#newBarChartProgress'		
}

function saveNewBarChart($dialog) {
	
	console.log("Saving new bar chart: dashboard ID = " + newBarChartParams.dashboardID )
			
	var saveNewBarChartParams = {
		dataSrcTableID: getWizardDialogPanelVals($dialog,dashboardComponentSelectTablePanelID),
		parentDashboardID: newBarChartParams.dashboardID,
		xAxisVals: getWizardDialogPanelVals($dialog,dashboardComponentValueGroupingPanelID), // xAxisVals
		xAxisSortValues: "asc",
		yAxisVals: getWizardDialogPanelVals($dialog,dashboardComponentValueSummaryPanelID), // yAxisVals
		geometry: newBarChartParams.geometry
	}
	
	
	console.log("saveNewBarChart: new bar chart params:  " + JSON.stringify(saveNewBarChartParams) )
	jsonAPIRequest("dashboard/barChart/new",saveNewBarChartParams,function(barChartRef) {
		console.log("saveNewBarChart: bar chart saved: new bar chart ID = " + barChartRef.barChartID)
		
		// Replace the placholder ID with the instantiated bar chart's unique ID. In the case
		// of a bar chart, 2 DOM elements are associated with the bar chart's ID. The first
		// is the overall/wrapper container, and the 2nd is a child div for the bar chart itself.
		// See the function barChartContainerHTML() to see how this is setup.
		 $('#'+newBarChartParams.placeholderID).attr("id",barChartRef.barChartID)
		 $('#'+newBarChartParams.placeholderID+"_chart").attr("id",barChartRef.barChartID+"_chart")

		$dialog.modal("hide")
		
		barChartDataParams = { 
			parentDashboardID: newBarChartParams.dashboardID,
			barChartID: barChartRef.barChartID
		}
		
		setTimeout(function() { // Wait for eventual consistency
			jsonAPIRequest("dashboard/barChart/getData",barChartDataParams,function(barChartData) {
				initBarChartData(newBarChartParams.dashboardID,barChartData)
			})		
		}, 2000);
		
	})
}



var barChartTablePanelConfig = createNewDashboardComponentSelectTablePanelConfig(barChartElemPrefix)
var barChartXAxisPanelConfig = createNewDashboardComponentValueGroupingPanelConfig(barChartElemPrefix)
var barChartYAxisPanelConfig = createNewDashboardComponentValueSummaryPanelConfig(barChartElemPrefix,saveNewBarChart)

function openNewBarChartDialog(barChartParams) {
		
	newBarChartParams.placeholderID = barChartParams.placeholderComponentID
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
		dialogDivID: '#newBarchartDialog',
		panels: [barChartTablePanelConfig,barChartXAxisPanelConfig, barChartYAxisPanelConfig],
		progressDivID: '#barChart_WizardDialogProgress',
		minBodyHeight:'350px'
	})
		
} // newBarChart




