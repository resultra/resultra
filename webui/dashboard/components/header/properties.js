

function loadDashboardHeaderProperties($header,headerRef) {

	var headerElemPrefix = "header_"
	
	var titlePropertyPanelParams = {
		dashboardID: headerRef.parentDashboardID,
		title: headerRef.properties.title,
		setTitleFunc: function(newTitle) {

			var setTitleParams = {
				parentDashboardID:headerRef.parentDashboardID,
				headerID: headerRef.headerID,
				newTitle:newTitle
			}
			jsonAPIRequest("dashboard/header/setTitle",setTitleParams,function(updatedHeader) {
				setContainerComponentInfo($header,updatedHeader,updatedHeader.headerID)
				setHeaderDashboardComponentLabel($header,updatedHeader)
			})

		}
	}
	initDashboardComponentTitlePropertyPanel(headerElemPrefix,titlePropertyPanelParams)


	// Toggle to the bar chart properties, hiding the other property panels
	hideSiblingsShowOne('#dashboardHeaderProps')

}
