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
		databaseID: propArgs.databaseID,
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
		databaseID: propArgs.databaseID,
		defaultFilterRules: summaryTableRef.properties.defaultFilterRules,
		initDone: function () {},
		updateFilterRules: function (updatedFilterRules) {
			var setDefaultFiltersParams = {
				parentDashboardID:propArgs.dashboardID,
				summaryTableID: summaryTableRef.summaryTableID,
				defaultFilterRules: updatedFilterRules
			}
			jsonAPIRequest("dashboard/summaryTable/setDefaultFilterRules",setDefaultFiltersParams,function(updatedSummaryTable) {
				console.log(" Default filters updated")
				setElemObjectRef(updatedSummaryTable.summaryTableID,updatedSummaryTable)
			}) // set record's number field value
			
		}
	}
	initFilterPropertyPanel(filterPropertyPanelParams)
	
	var columnsPropertyPanelParams = {
		listElemPrefix: summaryTableElemPrefix,
		databaseID: propArgs.databaseID,
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