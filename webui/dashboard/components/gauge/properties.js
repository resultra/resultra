

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
	
	
	var formatSelectionParams = {
		elemPrefix: gaugeElemPrefix,
		initialFormat: gaugeRef.properties.valueFormat.format,
		formatChangedCallback: function (newFormat) {
			console.log("Format changed for gauge: " + newFormat)
			
			var newValueFormat = {
				format: newFormat
			}
			var formatParams = {
				parentDashboardID:gaugeRef.parentDashboardID,
				gaugeID: gaugeRef.gaugeID,
				valueFormat: newValueFormat
			}
			jsonAPIRequest("dashboard/gauge/setValueFormat", formatParams, function(updatedGauge) {
				reloadGauge(updatedGauge)
				setContainerComponentInfo($gauge,updatedGauge,updatedGauge.gaugeID)
			})	
			
		}
	}
	initNumberFormatSelection(formatSelectionParams)
	

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
	
	
	

	// Toggle to the bar chart properties, hiding the other property panels
	hideSiblingsShowOne('#dashboardGaugeProps')

}
