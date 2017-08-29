

function loadDashboardSummaryValProperties(summaryValPropsArgs) {

	var elemPrefix = "summaryVal_"	
	
	var summaryValRef = getContainerObjectRef(summaryValPropsArgs.$summaryVal)
	var $summaryVal = summaryValPropsArgs.$summaryVal
	
	function reloadSummaryVal(summaryValRef) {
		var gaugeDataParams = { 
			parentDashboardID: summaryValRef.parentDashboardID,
			summaryValID: summaryValRef.summaryValID,
			filterRules: summaryValRef.properties.defaultFilterRules
		}
		jsonAPIRequest("dashboardController/getSummaryValData",gaugeDataParams,function(summaryValData) {
			initSummaryValData(summaryValRef.parentDashboardID,$summaryVal, summaryValData)
		})		
	}
	
	
	var titlePropertyPanelParams = {
		dashboardID: summaryValRef.parentDashboardID,
		title: summaryValRef.properties.title,
		setTitleFunc: function(newTitle) {

			var setTitleParams = {
				parentDashboardID:summaryValRef.parentDashboardID,
				summaryValID: summaryValRef.summaryValID,
				newTitle:newTitle
			}
			jsonAPIRequest("dashboard/summaryVal/setTitle",setTitleParams,function(updatedSummaryVal) {
				setContainerComponentInfo($summaryVal,updatedSummaryVal,updatedSummaryVal.summaryValID)
				setGaugeDashboardComponentLabel($summaryVal,updatedSummaryVal)
			})

		}
	}
	initDashboardComponentTitlePropertyPanel(elemPrefix,titlePropertyPanelParams)
	
		
	function saveThresholds(newThresholdVals) {
		var setThresholdParams = {
			parentDashboardID:summaryValRef.parentDashboardID,
			summaryValID: summaryValRef.summaryValID,
			thresholdVals: newThresholdVals
		}
		jsonAPIRequest("dashboard/summaryVal/setThresholds", setThresholdParams, function(updatedSummaryVal) {
			reloadGauge(updatedSummaryVal)
			setContainerComponentInfo($summaryVal,updatedSummaryVal,updatedSummaryVal.summaryValID)
		})	
	}
	var thresholdParams = {
		elemPrefix: elemPrefix,
		saveThresholdsCallback: saveThresholds,
		initialThresholdVals: summaryValRef.properties.thresholdVals
	}
	initThresholdValuesPropertyPanel(thresholdParams)
	
	
	var preFilterElemPrefix = "summaryValPreFilter_"
	var preFilterPropertyPanelParams = {
		elemPrefix: preFilterElemPrefix,
		databaseID: summaryValPropsArgs.databaseID,
		defaultFilterRules: summaryValRef.properties.preFilterRules,
		initDone: function () {},
		updateFilterRules: function (updatedFilterRules) {
			var setPreFiltersParams = {
				parentDashboardID:summaryValPropsArgs.dashboardID,
				summaryValID: summaryValRef.summaryValID,
				preFilterRules: updatedFilterRules
			}
			jsonAPIRequest("dashboard/summaryVal/setPreFilterRules",setPreFiltersParams,function(updatedSummaryVal) {
				console.log(" Pre filters updated")
				reloadGauge(updatedSummaryVal)
				setContainerComponentInfo($summaryVal,updatedSummaryVal,updatedSummaryVal.summaryValID)
			}) 
		}
	}
	initFilterPropertyPanel(preFilterPropertyPanelParams)
	

	var filterPropertyPanelParams = {
		elemPrefix: elemPrefix,
		databaseID: summaryValPropsArgs.databaseID,
		defaultFilterRules: summaryValRef.properties.defaultFilterRules,
		initDone: function () {},
		updateFilterRules: function (updatedFilterRules) {
			var setDefaultFiltersParams = {
				parentDashboardID:summaryValPropsArgs.dashboardID,
				summaryValID: summaryValRef.summaryValID,
				defaultFilterRules: updatedFilterRules
			}
			jsonAPIRequest("dashboard/summaryVal/setDefaultFilterRules",setDefaultFiltersParams,function(updatedSummaryVal) {
				console.log(" Default filters updated")
				reloadSummaryVal(updatedSummaryVal)
				setContainerComponentInfo($summaryVal,updatedSummaryVal,updatedSummaryVal.summaryValID)
			}) // set record's number field value
		}
	}
	initFilterPropertyPanel(filterPropertyPanelParams)
	
	
	var valSummaryPropertyPanelParams = {
		elemPrefix: elemPrefix,
		databaseID: summaryValPropsArgs.databaseID,
		valSummaryProps: summaryValRef.properties.valSummary,
		saveValueSummaryFunc: function(newValSummaryParams) {
			var setValSummaryParams = {
				parentDashboardID:summaryValPropsArgs.dashboardID,
				summaryValID: summaryValRef.summaryValID,
				valSummary:newValSummaryParams
			}
			jsonAPIRequest("dashboard/summaryVal/setValSummary",setValSummaryParams,function(updatedSummaryVal) {
				reloadSummaryVal(updatedSummaryVal)
				setContainerComponentInfo($summaryVal,updatedSummaryVal,updatedSummaryVal.summaryValID)
			})

		}
	}
	initDashboardValueSummaryPropertyPanel(valSummaryPropertyPanelParams)
	
	// Toggle to the bar chart properties, hiding the other property panels
	hideSiblingsShowOne('#dashboardSummaryValProps')

}
