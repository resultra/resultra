function loadSummaryTableProperties(propArgs) {

	var $summaryTable = $('#'+propArgs.summaryTableID)


	var summaryTableRef = getContainerObjectRef(propArgs.$summaryTable)
	var summaryTableElemPrefix = "summaryTable_"
	
	
	function reloadSummaryTable(summaryTableRef) {
	
		var getDataParams = {
			parentDashboardID:summaryTableRef.parentDashboardID,
			summaryTableID:summaryTableRef.summaryTableID,
			filterRules: summaryTableRef.properties.defaultFilterRules
		}
		jsonAPIRequest("dashboardController/getSummaryTableData",getDataParams,function(updatedSummaryTableData) {
			console.log("Repopulating summary table after changing properties")
			initSummaryTableData(summaryTableRef.parentDashboardID,propArgs.$summaryTable,updatedSummaryTableData)
		})
		
	}
	


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
				reloadSummaryTable(updatedSummaryTable)
				setContainerComponentInfo(propArgs.$summaryTable,updatedSummaryTable,updatedSummaryTable.summaryTableID)
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
				reloadSummaryTable(updatedSummaryTable)
				setContainerComponentInfo(propArgs.$summaryTable,updatedSummaryTable,updatedSummaryTable.summaryTableID)
			})
		}

	}
	initDashboardValueGroupingPropertyPanel(rowGroupingPropertyPanelParams)

	var preFilterElemPrefix = "summaryTablePreFilter_"
	var preFilterPropertyPanelParams = {
		elemPrefix: preFilterElemPrefix,
		databaseID: propArgs.databaseID,
		defaultFilterRules: summaryTableRef.properties.preFilterRules,
		initDone: function () {},
		updateFilterRules: function (updatedFilterRules) {
			var setDefaultFiltersParams = {
				parentDashboardID:propArgs.dashboardID,
				summaryTableID: summaryTableRef.summaryTableID,
				preFilterRules: updatedFilterRules
			}
			jsonAPIRequest("dashboard/summaryTable/setPreFilterRules",setDefaultFiltersParams,function(updatedSummaryTable) {
				console.log(" Default filters updated")
				reloadSummaryTable(updatedSummaryTable)
				setContainerComponentInfo(propArgs.$summaryTable,updatedSummaryTable,updatedSummaryTable.summaryTableID)
			}) // set record's number field value

		}
	}
	initFilterPropertyPanel(preFilterPropertyPanelParams)

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
				reloadSummaryTable(updatedSummaryTable)
				setContainerComponentInfo(propArgs.$summaryTable,updatedSummaryTable,updatedSummaryTable.summaryTableID)
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
				reloadSummaryTable(updatedSummaryTable)
				setContainerComponentInfo(propArgs.$summaryTable,updatedSummaryTable,updatedSummaryTable.summaryTableID)
			})
		}
	}
	initDashboardComponentSummaryColsPropertyPanel(columnsPropertyPanelParams)
	
	
	var helpPopupParams = {
		initialMsg: summaryTableRef.properties.helpPopupMsg,
		elemPrefix: summaryTableElemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentDashboardID:propArgs.dashboardID,
				summaryTableID: summaryTableRef.summaryTableID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("dashboard/summaryTable/setHelpPopupMsg",params,function(updatedSummaryTable) {
				setContainerComponentInfo(propArgs.$summaryTable,updatedSummaryTable,
								updatedSummaryTable.summaryTableID)
				updateComponentHelpPopupMsg(propArgs.$summaryTable, updatedSummaryTable)
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	

	// Toggle to the summary properties, hiding the other property panels
	hideSiblingsShowOne('#summaryTableProps')

}
