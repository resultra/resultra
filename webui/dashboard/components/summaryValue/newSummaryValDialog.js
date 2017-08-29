

function openNewDashboardSummaryValDialog(summaryValParams) {
	
	var elemPrefix = "summaryVal_"
	
	var newSummaryValParams = {}
	newSummaryValParams.dashboardID = summaryValParams.dashboardContext.dashboardID
	newSummaryValParams.progressDivID = '#newGaugeProgress'
	newSummaryValParams.placeholderID = summaryValParams.placeholderComponentID
	newSummaryValParams.geometry = summaryValParams.geometry
	newSummaryValParams.summaryValCreated = false
	newSummaryValParams.dialog = $('#newSummaryValDialog')
	
	function saveNewSummaryVal($dialog) {
					
		var saveNewSummaryValParams = {
			parentDashboardID: newSummaryValParams.dashboardID,
			valSummary: getWizardDialogPanelVals($dialog,dashboardComponentValueSummaryPanelID),
			geometry: newSummaryValParams.geometry
		}
	
		jsonAPIRequest("dashboard/summaryVal/new",saveNewSummaryValParams,function(summaryValRef) {
				
			  var newComponentSetupParams = {
				  parentDashboardID: newSummaryValParams.dashboardID,
			  	  $container: summaryValParams.$componentContainer,
				  componentID: summaryValRef.summaryValID,
				  componentObjRef: summaryValRef,
				  designFormConfig: summaryValDashboardDesignConfig
			  }
			  setupNewlyCreatedDashboardComponentInfo(newComponentSetupParams)
			
			newSummaryValParams.summaryValCreated = true
			$dialog.modal("hide")
		
			var summaryValDataParams = { 
				parentDashboardID: newSummaryValParams.dashboardID,
				summaryValID: summaryValRef.summaryValID,
				filterRules: summaryValRef.properties.defaultFilterRules
			}
			jsonAPIRequest("dashboardController/getSummaryValData",summaryValDataParams,function(summaryValData) {
				initSummaryValData(newGaugeParams.dashboardID,summaryValParams.$componentContainer,summaryValData)
			})			
		})
	}
	
	var databaseID = summaryValParams.dashboardContext.databaseID
	
	var summaryValColPanelConfig = createNewDashboardComponentValueSummaryPanelConfig(elemPrefix,saveNewSummaryVal,databaseID)
	
	openWizardDialog({
		closeFunc: function () {
  		  console.log("Close dialog")
  		  if(!newSummaryValParams.summaryValCreated)
  		  {
  			  // If the the bar chart creation is not complete, remove the placeholder
  			  // from the canvas.
  			  $('#'+newSummaryValParams.placeholderID).remove()
  		  }	
		},
		dialogDivID: '#newSummaryValDialog',
		panels: [summaryValColPanelConfig],
		progressDivID: '#summaryVal_WizardDialogProgress',
		minBodyHeight:'350px'
	})
		
} // newBarChart