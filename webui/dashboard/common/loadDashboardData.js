function loadDashboardData(loadDashboardConfig)
{
		
	var dashboardID = loadDashboardConfig.dashboardContext.dashboardID
	
	var dashboardLayoutSelector = '#dashboardCanvas'
	
	function initBarChartLayout($componentRow,barChartData) {
		
		
		var barChartHTML = barChartContainerHTML();
		var $barChartElem = $(barChartHTML)

		setContainerComponentInfo($barChartElem,barChartData.barChart,barChartData.barChartID)
		
		$componentRow.append($barChartElem)
		setElemFixedWidthFlexibleHeight($barChartElem,barChartData.barChart.properties.geometry.sizeWidth)
		
		initBarChartData(dashboardID,$barChartElem, barChartData);	
		
		loadDashboardConfig.initBarChartComponent($barChartElem,barChartData.barChart)	
	}

	function initSummaryTableLayout($componentRow,summaryTableData) {
		
		var summaryTableHTML = summaryTableComponentHTML(summaryTableData.summaryTableID);
		var $summaryTableElem = $(summaryTableHTML)
	
		setContainerComponentInfo($summaryTableElem,summaryTableData.summaryTable,summaryTableData.summaryTableID)
		
		$componentRow.append($summaryTableElem)
		setElemDimensions($summaryTableElem,summaryTableData.summaryTable.properties.geometry)
		
		initSummaryTableData(dashboardID,$summaryTableElem,summaryTableData)
		
		loadDashboardConfig.initSummaryTableComponent($summaryTableElem,summaryTableData.summaryTable)
	}


	function initHeaderLayout($componentRow,header) {
		
		var headerHTML = dashboardHeaderContainerHTML(header.headerID);
		var $header = $(headerHTML)
		
		setHeaderDashboardComponentLabel($header,header)
	
		setContainerComponentInfo($header,header,header.headerID)
		
		$componentRow.append($header)
		setElemFixedWidthFlexibleHeight($header,header.properties.geometry.sizeWidth)
				
		loadDashboardConfig.initHeaderComponent($header,header)
	}


	function initGaugeLayout($componentRow,gaugeData) {
		
		var gaugeHTML = dashboardGaugeContainerHTML(gaugeData.gaugeID);
		var $gauge = $(gaugeHTML)
		
		var gaugeRef = gaugeData.gauge
		
		setGaugeDashboardComponentLabel($gauge,gaugeRef)
	
		setContainerComponentInfo($gauge,gaugeRef,gaugeRef.gaugeID)
		
		initGaugeData(dashboardID,$gauge,gaugeData)
		
		$componentRow.append($gauge)
		setElemFixedWidthFlexibleHeight($gauge,gaugeRef.properties.geometry.sizeWidth)
				
		loadDashboardConfig.initGaugeComponent($gauge,gaugeRef)
	}

	function initSummaryValLayout($componentRow,summaryValData) {
		
		var summaryValHTML = dashboardSummaryValContainerHTML(summaryValData.summaryValID);
		var $summaryVal = $(summaryValHTML)
		
		var summaryValRef = summaryValData.summaryVal
		
		setGaugeDashboardComponentLabel($summaryVal,summaryValRef)
	
		setContainerComponentInfo($summaryVal,summaryValRef,summaryValRef.summaryValID)
		
		initSummaryValData(dashboardID,$summaryVal,summaryValData)
		setSummaryValDashboardComponentLabel($summaryVal,summaryValRef)
		
		$componentRow.append($summaryVal)
		setElemFixedWidthFlexibleHeight($summaryVal,summaryValRef.properties.geometry.sizeWidth)
				
		loadDashboardConfig.initSummaryValComponent($summaryVal,summaryValRef)
	}
	
	// Load the dashboard data
	var loadBarChartDataParams = { dashboardID: dashboardID }
	jsonAPIRequest("dashboardController/getDefaultData",loadBarChartDataParams,function(dashboardData) {
		
		var compenentIDComponentMap = {}
			
		for (var barChartDataIndex in dashboardData.barChartsData) {
			var barChartData = dashboardData.barChartsData[barChartDataIndex]
			console.log ("Loading bar chart: id = " + barChartData.barChartID)
			compenentIDComponentMap[barChartData.barChartID] = {
				componentInfo: barChartData,
				initFunc: initBarChartLayout
			}		
		}
		for (var summaryTableDataIndex in dashboardData.summaryTablesData) {
			var summaryTableData = dashboardData.summaryTablesData[summaryTableDataIndex]
			console.log ("Loading summary table: id = " + summaryTableData.summaryTableID)
			compenentIDComponentMap[summaryTableData.summaryTableID] = {
				componentInfo: summaryTableData,
				initFunc: initSummaryTableLayout
			}		
		}


		for (var gaugeDataIndex in dashboardData.gaugesData) {
			var gaugeData = dashboardData.gaugesData[gaugeDataIndex]
			console.log ("Loading summary table: id = " + gaugeData.gaugeID)
			compenentIDComponentMap[gaugeData.gaugeID] = {
				componentInfo: gaugeData,
				initFunc: initGaugeLayout
			}		
		}

		for (var summaryValDataIndex in dashboardData.summaryValsData) {
			var summaryValData = dashboardData.summaryValsData[summaryValDataIndex]
			console.log ("Loading summary table: id = " + gaugeData.gaugeID)
			compenentIDComponentMap[summaryValData.summaryValID] = {
				componentInfo: summaryValData,
				initFunc: initSummaryValLayout
			}		
		}


		for (var headerIndex in dashboardData.headers) {
			var header = dashboardData.headers[headerIndex]
			console.log ("Loading header: id = " + header.headerID)
			compenentIDComponentMap[header.headerID] = {
				componentInfo: header,
				initFunc: initHeaderLayout
			}		
		}
		
		
		var dashboardLayout = dashboardData.dashboard.properties.layout
		var $parentLayout = $(dashboardLayoutSelector)
		
		populateComponentLayout(dashboardLayout,$parentLayout,compenentIDComponentMap)
		
		loadDashboardConfig.doneLoadingDashboardDataFunc()
						
	}) // getData
	
}
