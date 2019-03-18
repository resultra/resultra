// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function openNewDashboardSummaryValDialog(summaryValParams) {
	
	var elemPrefix = "summaryVal_"
	
	var newSummaryValParams = {}
	newSummaryValParams.dashboardID = summaryValParams.dashboardContext.dashboardID
	newSummaryValParams.progressDivID = '#newGaugeProgress'
	newSummaryValParams.placeholderID = summaryValParams.placeholderComponentID
	newSummaryValParams.geometry = summaryValParams.geometry
	newSummaryValParams.summaryValCreated = false
	newSummaryValParams.dialog = $('#newSummaryValDialog')
	
	var summaryValCreated = false
	
	function saveNewSummaryVal($dialog) {
		
		if (summaryValCreated === true) {
			return
		}
		summaryValCreated = true // only allow creation of a single summaryVal from the dialog
					
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
				initSummaryValData(newSummaryValParams.dashboardID,summaryValParams.$componentContainer,summaryValData)
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