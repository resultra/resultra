// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function loadDashboardGaugeProperties(gaugePropsArgs) {

	var gaugeElemPrefix = "gauge_"	
	
	var gaugeRef = getContainerObjectRef(gaugePropsArgs.$gauge)
	var $gauge = gaugePropsArgs.$gauge
	
	function reloadGauge(gaugeRef) {
		var gaugeDataParams = { 
			parentDashboardID: gaugeRef.parentDashboardID,
			gaugeID: gaugeRef.gaugeID,
			filterRules: gaugeRef.properties.defaultFilterRules
		}
		jsonAPIRequest("dashboardController/getGaugeData",gaugeDataParams,function(gaugeData) {
			initGaugeData(gaugeRef.parentDashboardID,$gauge, gaugeData)
		})		
	}
	
	
	var titlePropertyPanelParams = {
		dashboardID: gaugeRef.parentDashboardID,
		title: gaugeRef.properties.title,
		setTitleFunc: function(newTitle) {

			var setTitleParams = {
				parentDashboardID:gaugeRef.parentDashboardID,
				gaugeID: gaugeRef.gaugeID,
				newTitle:newTitle
			}
			jsonAPIRequest("dashboard/gauge/setTitle",setTitleParams,function(updatedGauge) {
				setContainerComponentInfo($gauge,updatedGauge,updatedGauge.gaugeID)
				setGaugeDashboardComponentLabel($gauge,updatedGauge)
			})

		}
	}
	initDashboardComponentTitlePropertyPanel(gaugeElemPrefix,titlePropertyPanelParams)
	
	
	function setGaugeRange(minVal,maxVal) {
		var setRangeParams = {
			parentDashboardID:gaugeRef.parentDashboardID,
			gaugeID: gaugeRef.gaugeID,
			minVal: minVal,
			maxVal: maxVal
		}
		console.log("Setting gauge range: " + JSON.stringify(setRangeParams))
		jsonAPIRequest("dashboard/gauge/setRange", setRangeParams, function(updatedGauge) {
			reloadGauge(updatedGauge)
			setContainerComponentInfo($gauge,updatedGauge,updatedGauge.gaugeID)
		})
	}
	var gaugeRangeParams = {
		defaultMinVal: gaugeRef.properties.minVal,
		defaultMaxVal: gaugeRef.properties.maxVal,
		setRangeCallback: setGaugeRange
	}
	initGaugeRangeProperties(gaugeRangeParams)
	
	function saveGaugeThresholds(newThresholdVals) {
		var setThresholdParams = {
			parentDashboardID:gaugeRef.parentDashboardID,
			gaugeID: gaugeRef.gaugeID,
			thresholdVals: newThresholdVals
		}
		jsonAPIRequest("dashboard/gauge/setThresholds", setThresholdParams, function(updatedGauge) {
			reloadGauge(updatedGauge)
			setContainerComponentInfo($gauge,updatedGauge,updatedGauge.gaugeID)
		})	
	}
	var thresholdParams = {
		elemPrefix: gaugeElemPrefix,
		saveThresholdsCallback: saveGaugeThresholds,
		initialThresholdVals: gaugeRef.properties.thresholdVals
	}
	initThresholdValuesPropertyPanel(thresholdParams)
	
	
		

	var preFilterGaugeElemPrefix = "gaugePreFilter_"
	var preFilterPropertyPanelParams = {
		elemPrefix: preFilterGaugeElemPrefix,
		databaseID: gaugePropsArgs.databaseID,
		defaultFilterRules: gaugeRef.properties.preFilterRules,
		initDone: function () {},
		updateFilterRules: function (updatedFilterRules) {
			var setPreFiltersParams = {
				parentDashboardID:gaugePropsArgs.dashboardID,
				gaugeID: gaugeRef.gaugeID,
				preFilterRules: updatedFilterRules
			}
			jsonAPIRequest("dashboard/gauge/setPreFilterRules",setPreFiltersParams,function(updatedGauge) {
				console.log(" Pre filters updated")
				reloadGauge(updatedGauge)
				setContainerComponentInfo($gauge,updatedGauge,updatedGauge.gaugeID)
			}) 
		}
	}
	initFilterPropertyPanel(preFilterPropertyPanelParams)
	

	var filterPropertyPanelParams = {
		elemPrefix: gaugeElemPrefix,
		databaseID: gaugePropsArgs.databaseID,
		defaultFilterRules: gaugeRef.properties.defaultFilterRules,
		initDone: function () {},
		updateFilterRules: function (updatedFilterRules) {
			var setDefaultFiltersParams = {
				parentDashboardID:barChartPropsArgs.dashboardID,
				gaugeID: gaugeRef.gaugeID,
				defaultFilterRules: updatedFilterRules
			}
			jsonAPIRequest("dashboard/gauge/setDefaultFilterRules",setDefaultFiltersParams,function(updatedGauge) {
				console.log(" Default filters updated")
				reloadGauge(updatedGauge)
				setContainerComponentInfo($gauge,updatedGauge,updatedGauge.gaugeID)
			}) // set record's number field value
		}
	}
	initFilterPropertyPanel(filterPropertyPanelParams)
	
	
	var valSummaryPropertyPanelParams = {
		elemPrefix: gaugeElemPrefix,
		databaseID: gaugePropsArgs.databaseID,
		valSummaryProps: gaugeRef.properties.valSummary,
		saveValueSummaryFunc: function(newValSummaryParams) {
			var setValSummaryParams = {
				parentDashboardID:gaugePropsArgs.dashboardID,
				gaugeID: gaugeRef.gaugeID,
				valSummary:newValSummaryParams
			}
			jsonAPIRequest("dashboard/gauge/setValSummary",setValSummaryParams,function(updatedGauge) {
				reloadGauge(updatedGauge)
				setContainerComponentInfo($gauge,updatedGauge,updatedGauge.gaugeID)
			})

		}
	}
	initDashboardValueSummaryPropertyPanel(valSummaryPropertyPanelParams)
	
	var helpPopupParams = {
		initialMsg: gaugeRef.properties.helpPopupMsg,
		elemPrefix: gaugeElemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentDashboardID:gaugePropsArgs.dashboardID,
				gaugeID: gaugeRef.gaugeID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("dashboard/gauge/setHelpPopupMsg",params,function(updatedGauge) {
				setContainerComponentInfo($gauge,updatedGauge,updatedGauge.gaugeID)
				updateComponentHelpPopupMsg($gauge, updatedGauge)
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	
	var deleteParams = {
		elemPrefix: gaugeElemPrefix,
		parentDashboardID: gaugePropsArgs.dashboardID,
		componentID: gaugeRef.gaugeID,
		componentLabel: 'gauge',
		$componentContainer: $gauge
	}
	initDeleteDashboardComponentPropertyPanel(deleteParams)
	
	

	// Toggle to the bar chart properties, hiding the other property panels
	hideSiblingsShowOne('#dashboardGaugeProps')

}
