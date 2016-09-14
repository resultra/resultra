function loadSummaryTableProperties(propArgs) {
	
	var $summaryTable = $('#'+propArgs.summaryTableID)
	var summaryTableRef = $summaryTable.data("summaryTableRef")
	var summaryTableElemPrefix = "summaryTable_"
	
	
	var titlePropertyPanelParams = {
		dashboardID: propArgs.dashboardID,
		title: summaryTableRef.properties.title,
		setTitleFunc: function(newTitle) {
			
			var setTitleParams = {
				parentDashboardID:propArgs.dashboardID,
				summaryTableID: summaryTableRef.summaryTableID,
				newTitle:newTitle
			}
			jsonAPIRequest("dashboard/summaryTable/setTitle",setTitleParams,function(updatedSummaryTable) {
					$summaryTable.data("summaryTableRef",updatedSummaryTable)
			})
			
		}
	}
	initDashboardComponentTitlePropertyPanel(summaryTableElemPrefix,titlePropertyPanelParams)
	
}