

function openNewSummaryTableDialog(summaryTableParams) {
	
	var summaryTableElemPrefix = "summaryTable_"
	var newSummaryTableElemPrefix = "newSummaryTable_"
	
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
			rowGroupingVals: getWizardDialogPanelVals($dialog,dashboardComponentValueGroupingPanelID), // rowGrouping
			columnValSummaries: tableColSummaryParams,
			geometry: newSummaryTableParams.geometry
		}
	
		console.log("saveNewSummaryTable: new summary table params:  " + JSON.stringify(saveNewSummaryTableParams) )
		jsonAPIRequest("dashboard/summaryTable/new",saveNewSummaryTableParams,function(summaryTableRef) {
		
			console.log("saveNewBarChart: bar chart saved: new bar chart ID = " + summaryTableRef.summaryTableID)
		
			  var newComponentSetupParams = {
				  parentDashboardID: newSummaryTableParams.dashboardID,
			  	  $container: summaryTableParams.$componentContainer,
				  componentID: summaryTableRef.summaryTableID,
				  componentObjRef: summaryTableRef,
				  designFormConfig: summaryTableDashboardDesignConfig
			  }
			  setupNewlyCreatedDashboardComponentInfo(newComponentSetupParams)
			
			newSummaryTableParams.summaryTableCreated = true
			$dialog.modal("hide")
		
			var summaryTableDataParams = { 
				parentDashboardID: newSummaryTableParams.dashboardID,
				summaryTableID: summaryTableRef.summaryTableID,
				filterRules: summaryTableRef.properties.defaultFilterRules
			}
			jsonAPIRequest("dashboardController/getSummaryTableData",summaryTableDataParams,function(summaryTableData) {
				initSummaryTableData(newSummaryTableParams.dashboardID,summaryTableParams.$componentContainer,summaryTableData)
			})			
		})
	}
	
	var databaseID = summaryTableParams.dashboardContext.databaseID
	
	
	
	var summaryTableRowGroupingPanelConfig = createNewDashboardComponentValueGroupingPanelConfig(newSummaryTableElemPrefix,databaseID)
	
	
	var summaryTableColPanelConfig = createNewDashboardComponentValueSummaryPanelConfig(summaryTableElemPrefix,saveNewSummaryTable,databaseID)
	
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
		panels: [summaryTableRowGroupingPanelConfig, summaryTableColPanelConfig],
		progressDivID: '#summaryTable_WizardDialogProgress',
		minBodyHeight:'350px'
	})
		
} // newBarChart