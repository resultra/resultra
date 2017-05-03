function gaugeFromDashboardGaugeContainer($gauge) {
	return 	$gauge.find(".dashboardGauge")
}


function dashboardGaugeContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class="layoutContainer dashboardGaugeFormContainer" id="'+elementID+'">' +
			'<label class="gaugeLabel">' + 'New Gauge' + '</label>' +
			'<div class="dashboardGauge">' +
				'<span class="gaugeControl"></span>'+
  			'</div>' +
	
		'</div><';
						
	return containerHTML
}

function setGaugeDashboardComponentLabel($container,gaugeRef) {
	var $label = $container.find(".gaugeLabel")
	$label.text(gaugeRef.properties.title)	
}

function initGaugeDashboardComponentControl($gauge,gaugeConfig,gaugeVal) {
	var $gaugeControlContainer = $gauge.find(".gaugeControl")
	var gaugeControl = new GaugeUIControl($gaugeControlContainer, gaugeConfig);
	gaugeControl.render()
	gaugeControl.redraw(gaugeConfig.min)
	$gauge.data("gaugeControl",gaugeControl) 
	
	gaugeControl.redraw(gaugeVal)
	
}


function initGaugeData(dashboardID,$gauge, gaugeData) {
	
	var gaugeRef = gaugeData.gauge	
	var gaugeVal = gaugeData.groupedSummarizedVals.overallDataRow.summaryVals[0]
	
	function formatGaugeVal(val) {
		var numberFormat = "general"
//		var numberFormat = gaugeRef.properties.valueFormat.format
		var formattedVal = formatNumberValue(numberFormat,val)
		return formattedVal
	}
	
//	var minVal = gaugeRef.properties.minVal
	var minVal = 0
// var maxVal = gaugeRef.properties.maxVal
	var maxVal = 100
		
	var gaugeConfig = 
	{
		size: gaugeRef.properties.geometry.sizeWidth,
		min: minVal,
		max: maxVal,
		minorTicks: 5,
		valueFormatter: formatGaugeVal
	}	
	gaugeConfig.yellowZones = [];
	gaugeConfig.redZones = [];
	gaugeConfig.greenZones = [];
	
	initGaugeDashboardComponentControl($gauge,gaugeConfig,gaugeVal)
	
}
