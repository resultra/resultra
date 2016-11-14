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
					setElemObjectRef(updatedSummaryTable.summaryTableID,updatedSummaryTable)
			})
			
		}
	}
	initDashboardComponentTitlePropertyPanel(summaryTableElemPrefix,titlePropertyPanelParams)
	
	
	var rowGroupingPropertyPanelParams = {
		elemPrefix: summaryTableElemPrefix,
		tableID: summaryTableRef.properties.dataSrcTableID,
		valGroupingProps: summaryTableRef.properties.rowGroupingVals,
		saveValueGroupingFunc: function(newValueGroupingParams) {
			var setRowGroupingParams = {
				parentDashboardID:propArgs.dashboardID,
				summaryTableID: summaryTableRef.summaryTableID,
				rowValueGrouping:newValueGroupingParams
			}
			jsonAPIRequest("dashboard/summaryTable/setRowValueGrouping",setRowGroupingParams,function(updatedSummaryTable) {
					setElemObjectRef(updatedSummaryTable.summaryTableID,updatedSummaryTable)
			})
		}
		
	}
	initDashboardValueGroupingPropertyPanel(rowGroupingPropertyPanelParams)
	
	
	var filterPropertyPanelParams = {
		elemPrefix: summaryTableElemPrefix,
		tableID: summaryTableRef.properties.dataSrcTableID,
		/* TODO - restore a property panel with a callback and params like the following: 
			var params = {
				parentDashboardID: propArgs.dashboardID,
				summaryTableID: summaryTableRef.summaryTableID,
				defaultFilterIDs: defaultFilterIDs }
		*/
	}
	initFilterPropertyPanel(filterPropertyPanelParams)
	
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
					setElemObjectRef(updatedSummaryTable.summaryTableID,updatedObjRef)
			})
		}
	}
	initDashboardComponentSummaryColsPropertyPanel(columnsPropertyPanelParams)
	
	// Toggle to the summary properties, hiding the other property panels
	hideSiblingsShowOne('#summaryTableProps')
	
}