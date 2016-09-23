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
					setElemObjectRef(updatedSummaryTable.summaryTableID,updatedObjRef)
			})
			
		}
	}
	initDashboardComponentTitlePropertyPanel(summaryTableElemPrefix,titlePropertyPanelParams)
	
	
	var filterPropertyPanelParams = {
		elemPrefix: summaryTableElemPrefix,
		tableID: summaryTableRef.properties.dataSrcTableID,
		defaultFilterIDs: summaryTableRef.properties.defaultFilterIDs,
		setDefaultFilterFunc: function(defaultFilterIDs) {
			var params = {
				parentDashboardID: propArgs.dashboardID,
				summaryTableID: summaryTableRef.summaryTableID,
				defaultFilterIDs: defaultFilterIDs }
			jsonAPIRequest("dashboard/summaryTable/setDefaultFilters",params,function(updatedSummaryTable) {
				setElemObjectRef(updatedSummaryTable.summaryTableID,updatedSummaryTable)
				console.log("Default filters updated")
			}) // set record's number field value
			
		},
		availableFilterIDs: summaryTableRef.properties.availableFilterIDs,
		setAvailableFilterFunc: function(availFilterIDs) {
			var params = {
				parentDashboardID: propArgs.dashboardID,
				summaryTableID: summaryTableRef.summaryTableID,
				availableFilterIDs: availFilterIDs }
			jsonAPIRequest("dashboard/summaryTable/setAvailableFilters",params,function(updatedSummaryTable) {
				setElemObjectRef(updatedSummaryTable.summaryTableID,updatedSummaryTable)
				console.log("Available filters updated")
			}) // set record's number field value
			
		}
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