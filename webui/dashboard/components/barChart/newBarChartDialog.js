
var newBarChartParams = {}
var barChartElemPrefix = "barChart_"

function initNewBarChartDialog(dashboardContext) {

	newBarChartParams.dashboardID = dashboardContext.dashboardID
	newBarChartParams.progressDivID = '#newBarChartProgress'		
}


function openNewBarChartDialog(barChartParams) {
	
	function saveNewBarChart($dialog) {
	
		console.log("Saving new bar chart: dashboard ID = " + newBarChartParams.dashboardID )
			
		var saveNewBarChartParams = {
			parentDashboardID: newBarChartParams.dashboardID,
			xAxisVals: getWizardDialogPanelVals($dialog,dashboardComponentValueGroupingPanelID), // xAxisVals
			xAxisSortValues: "asc",
			yAxisVals: getWizardDialogPanelVals($dialog,dashboardComponentValueSummaryPanelID), // yAxisVals
			geometry: newBarChartParams.geometry
		}
		
	
	
		console.log("saveNewBarChart: new bar chart params:  " + JSON.stringify(saveNewBarChartParams) )
		jsonAPIRequest("dashboard/barChart/new",saveNewBarChartParams,function(barChartRef) {
			console.log("saveNewBarChart: bar chart saved: new bar chart ID = " + barChartRef.barChartID)
			
			setContainerComponentInfo(barChartParams.$componentContainer,barChartRef,barChartRef.barChartID)
			
			$dialog.modal("hide")
		
			var barChartDataParams = { 
				parentDashboardID: newBarChartParams.dashboardID,
				barChartID: barChartRef.barChartID,
				filterRules: barChartRef.properties.defaultFilterRules
			}
		
			setTimeout(function() {
				jsonAPIRequest("dashboardController/getBarChartData",barChartDataParams,function(barChartData) {
					initBarChartData(newBarChartParams.dashboardID,barChartParams.$componentContainer, barChartData)
				})		
			}, 2000);
		
		})
	}
	
	//
		
	newBarChartParams.placeholderID = barChartParams.placeholderComponentID
	newBarChartParams.geometry = barChartParams.geometry
	newBarChartParams.barChartCreated = false
	newBarChartParams.dialog = $('#newBarchartDialog')
	
	var databaseID = barChartParams.dashboardContext.databaseID
	
	var barChartXAxisPanelConfig = createNewDashboardComponentValueGroupingPanelConfig(barChartElemPrefix,databaseID)
	var barChartYAxisPanelConfig = createNewDashboardComponentValueSummaryPanelConfig(barChartElemPrefix,saveNewBarChart,databaseID)
	

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
		panels: [barChartXAxisPanelConfig, barChartYAxisPanelConfig],
		progressDivID: '#barChart_WizardDialogProgress',
		minBodyHeight:'350px'
	})
		
} // newBarChart




