function gaugeViewDashboardConfig(gaugeRef) {
	
	var currFilterRules = gaugeRef.properties.defaultFilterRules
		
	
	function reloadGauge($gaugeContainer) {
	
		var getDataParams = {
			parentDashboardID:gaugeRef.parentDashboardID,
			gaugeID:gaugeRef.gaugeID,
			filterRules: currFilterRules
		}
		jsonAPIRequest("dashboardController/getGaugeData",getDataParams,function(gaugeData) {
			initGaugeData(gaugeRef.parentDashboardID,$gaugeContainer, gaugeData)
		})		
		
	}
	
	
	function selectGauge($container,selectedGaugeRef) {
		
		var gaugeElemPrefix = "gauge_"
		
		var filterPaneParams = {
			elemPrefix: gaugeElemPrefix,
			databaseID: viewDashboardContext.databaseID,
			defaultFilterRules: currFilterRules,
			initDone: function () {},
			updateFilterRules: function (updatedFilterRules) {
				currFilterRules = updatedFilterRules
				reloadGauge($container)
			},
			refilterWithCurrentFilterRules: function() {
				reloadGauge($container)
			}
		}

		initRecordFilterViewPanel(filterPaneParams)

		// Toggle to the summary properties, hiding the other property panels
		hideSiblingsShowOne('#dashboardGaugeViewProps')
	}
	
		
	var viewConfig = {
		selectionFunc: selectGauge
	}
	
	return viewConfig
}