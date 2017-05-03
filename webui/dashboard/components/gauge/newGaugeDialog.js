

function openNewDashboardGaugeDialog(gaugeParams) {
	
	var gaugeElemPrefix = "gauge_"
	
	var newGaugeParams = {}
	newGaugeParams.dashboardID = gaugeParams.dashboardContext.dashboardID
	newGaugeParams.progressDivID = '#newGaugeProgress'
	newGaugeParams.placeholderID = gaugeParams.placeholderComponentID
	newGaugeParams.geometry = gaugeParams.geometry
	newGaugeParams.gaugeCreated = false
	newGaugeParams.dialog = $('#newGaugeDialog')
	
	function saveNewGauge($dialog) {
	
		console.log("Saving new summary table: dashboard ID = " + newBarChartParams.dashboardID )
				
		var saveNewGaugeParams = {
			parentDashboardID: newGaugeParams.dashboardID,
			valSummary: getWizardDialogPanelVals($dialog,dashboardComponentValueSummaryPanelID),
			geometry: newGaugeParams.geometry
		}
	
		console.log("saveNewGauge: new gauge params:  " + JSON.stringify(saveNewGaugeParams) )
		jsonAPIRequest("dashboard/gauge/new",saveNewGaugeParams,function(gaugeRef) {
		
			console.log("saveNewGauge: gauge saved: new gauge ID = " + gaugeRef.gaugeID)
		
			  var newComponentSetupParams = {
				  parentDashboardID: newGaugeParams.dashboardID,
			  	  $container: gaugeParams.$componentContainer,
				  componentID: gaugeRef.gaugeID,
				  componentObjRef: gaugeRef,
				  designFormConfig: gaugeDashboardDesignConfig
			  }
			  setupNewlyCreatedDashboardComponentInfo(newComponentSetupParams)
			
			newGaugeParams.gaugeCreated = true
			$dialog.modal("hide")
		
			var gaugeDataParams = { 
				parentDashboardID: newGaugeParams.dashboardID,
				gaugeID: gaugeRef.gaugeID,
				filterRules: gaugeRef.properties.defaultFilterRules
			}
			jsonAPIRequest("dashboardController/getGaugeData",gaugeDataParams,function(gaugeData) {
				initGaugeData(newGaugeParams.dashboardID,gaugeParams.$componentContainer,gaugeData)
			})			
		})
	}
	
	var databaseID = gaugeParams.dashboardContext.databaseID
	
	var gaugeColPanelConfig = createNewDashboardComponentValueSummaryPanelConfig(gaugeElemPrefix,saveNewGauge,databaseID)
	
	openWizardDialog({
		closeFunc: function () {
  		  console.log("Close dialog")
  		  if(!newGaugeParams.gaugeCreated)
  		  {
  			  // If the the bar chart creation is not complete, remove the placeholder
  			  // from the canvas.
  			  $('#'+newGaugeParams.placeholderID).remove()
  		  }	
		},
		dialogDivID: '#newGaugeDialog',
		panels: [gaugeColPanelConfig],
		progressDivID: '#gauge_WizardDialogProgress',
		minBodyHeight:'350px'
	})
		
} // newBarChart