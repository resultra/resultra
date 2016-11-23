

function openNewSummaryTableDialog(summaryTableParams) {
	
	var summaryTableElemPrefix = "summaryTable_"
	
	var newSummaryTableParams = {}
	newSummaryTableParams.dashboardID = summaryTableParams.dashboardContext.dashboardID
	newSummaryTableParams.progressDivID = '#newSummaryTableProgress'
	newSummaryTableParams.placeholderID = summaryTableParams.placeholderComponentID
	newSummaryTableParams.geometry = summaryTableParams.geometry
	newSummaryTableParams.summaryTableCreated = false
	newSummaryTableParams.dialog = $('#newSummaryTableDialog')
	
	function saveNewSummaryTable($dialog) {
	
		console.log("Saving new summary table: dashboard ID = " + newBarChartParams.dashboardID )
	
		// The summary table is created with a list of columns.
		// TODO - support adding multiple cols from dialog
		var tableColSummaryParams = [
			getWizardDialogPanelVals($dialog,dashboardComponentValueSummaryPanelID)
		]
			
		var saveNewSummaryTableParams = {
			parentDashboardID: newSummaryTableParams.dashboardID,
			dataSrcTableID: getWizardDialogPanelVals($dialog,dashboardComponentSelectTablePanelID),
			rowGroupingVals: getWizardDialogPanelVals($dialog,dashboardComponentValueGroupingPanelID), // rowGrouping
			columnValSummaries: tableColSummaryParams,
			geometry: newSummaryTableParams.geometry
		}
	
		console.log("saveNewSummaryTable: new summary table params:  " + JSON.stringify(saveNewSummaryTableParams) )
		jsonAPIRequest("dashboard/summaryTable/new",saveNewSummaryTableParams,function(summaryTableRef) {
		
			console.log("saveNewBarChart: bar chart saved: new bar chart ID = " + summaryTableRef.summaryTableID)
		
			// Replace the placholder ID with the instantiated bar chart's unique ID. In the case
			// of a bar chart, 2 DOM elements are associated with the bar chart's ID. The first
			// is the overall/wrapper container, and the 2nd is a child div for the bar chart itself.
			// See the function barChartContainerHTML() to see how this is setup.
			 $('#'+newBarChartParams.placeholderID).attr("id",summaryTableRef.summaryTableID)
			 $('#'+newBarChartParams.placeholderID+"_table").attr("id",summaryTableRef.summaryTableID+"_table")

			newSummaryTableParams.summaryTableCreated = true
			$dialog.modal("hide")
		
			var summaryTableDataParams = { 
				parentDashboardID: newSummaryTableParams.dashboardID,
				summaryTableID: summaryTableRef.summaryTableID,
				filterRules: []
			}
			jsonAPIRequest("dashboardController/getSummaryTableData",summaryTableDataParams,function(summaryTableData) {
				initSummaryTableData(newSummaryTableParams.dashboardID,summaryTableData)
			})			
		})
	}
	
	var summaryTableTablePanelConfig = createNewDashboardComponentSelectTablePanelConfig(summaryTableElemPrefix)
	var summaryTableRowGroupingPanelConfig = createNewDashboardComponentValueGroupingPanelConfig(summaryTableElemPrefix)
	var summaryTableColPanelConfig = createNewDashboardComponentValueSummaryPanelConfig(summaryTableElemPrefix,saveNewSummaryTable)
	
	openWizardDialog({
		closeFunc: function () {
  		  console.log("Close dialog")
  		  if(!newSummaryTableParams.summaryTableCreated)
  		  {
  			  // If the the bar chart creation is not complete, remove the placeholder
  			  // from the canvas.
  			  $('#'+newSummaryTableParams.placeholderID).remove()
  		  }	
		},
		dialogDivID: '#newSummaryTableDialog',
		panels: [summaryTableTablePanelConfig,summaryTableRowGroupingPanelConfig, summaryTableColPanelConfig],
		progressDivID: '#summaryTable_WizardDialogProgress',
		minBodyHeight:'350px'
	})
		
} // newBarChart