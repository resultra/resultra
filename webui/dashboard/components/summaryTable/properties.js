function loadSummaryTableProperties(propArgs) {
	
	var $summaryTable = $('#'+propArgs.summaryTableID)
	
	
	var summaryTableRef = getElemObjectRef(propArgs.summaryTableID)
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
	
	
	var columnsPropertyPanelParams = {
		listElemPrefix: summaryTableElemPrefix,
		dataSrcTableID: summaryTableRef.properties.dataSrcTableID,
		initialColumnValSummaries: summaryTableRef.properties.columnValSummaries,
		setColumnsFunc: function(newColumns) {
			console.log("Setting summary table column properties: " + JSON.stringify(newColumns))
			var setColumnParams = {
				parentDashboardID:propArgs.dashboardID,
				summaryTableID: summaryTableRef.summaryTableID,
				columnValSummaries: newColumns }
			jsonAPIRequest("dashboard/summaryTable/setColumns",setColumnParams,function(updatedSummaryTable) {
					$summaryTable.data("summaryTableRef",updatedSummaryTable)
			})
		}
	}
	initDashboardComponentSummaryColsPropertyPanel(columnsPropertyPanelParams)
	
	// Toggle to the summary properties, hiding the other property panels
	hideSiblingsShowOne('#summaryTableProps')
	
}