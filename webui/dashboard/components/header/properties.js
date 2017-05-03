

function loadDashboardHeaderProperties($header,headerRef) {

	var headerElemPrefix = "header_"
	
	function initHeaderSizeProperties() {
		var $sizeSelection = $('#adminDashboardHeaderSizeSelection')
		$sizeSelection.val(headerRef.properties.size)
		initSelectControlChangeHandler($sizeSelection,function(newSize) {
		
			var sizeParams = {
				parentDashboardID:headerRef.parentDashboardID,
				headerID: headerRef.headerID,
				size: newSize
			}
			console.log("Setting new header size: " + JSON.stringify(sizeParams))
		
			jsonAPIRequest("dashboard/header/setSize",sizeParams,function(updatedHeader) {
				setContainerComponentInfo($header,updatedHeader,updatedHeader.headerID)				
				setHeaderDashboardComponentLabel($header,updatedHeader)			
			})
		
		})
		
	}
	initHeaderSizeProperties()
	
	function initHeaderUnderlinedProperties() {
		initCheckboxChangeHandler('#adminDashboardHeaderUnderline', 
					headerRef.properties.underlined, function (newVal) {
				
			var underlinedParams = {
				parentDashboardID:headerRef.parentDashboardID,
				headerID: headerRef.headerID,
				underlined: newVal
			}

			jsonAPIRequest("dashboard/header/setUnderlined",underlinedParams,function(updatedHeader) {
				setContainerComponentInfo($header,updatedHeader,updatedHeader.checkBoxID)
				setHeaderDashboardComponentLabel($header,updatedHeader)			
			})
		})
		
	}
	initHeaderUnderlinedProperties()
	
	
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
