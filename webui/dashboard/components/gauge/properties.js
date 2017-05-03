

function loadDashboardGaugeProperties($gauge,gaugeRef) {

	var gaugeElemPrefix = "gauge_"	
	
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


	// Toggle to the bar chart properties, hiding the other property panels
	hideSiblingsShowOne('#dashboardGaugeProps')

}
