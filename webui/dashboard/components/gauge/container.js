function gaugeFromDashboardGaugeContainer($gauge) {
	return 	$gauge.find(".dashboardGauge")
}


function dashboardGaugeContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class="layoutContainer dashboardGaugeFormContainer" id="'+elementID+'">' +
			'<div class="row">' +
				'<div class="col-sm-10">' +
					'<label class="gaugeLabel">' + 'New Gauge' + '</label>' +
				'</div>' +
				'<div class="col-sm-2 componentHeaderButtons">' +
					componentHelpPopupButtonHTML() +
				'</div>' +
			'</div>' +
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
	var numberFormat = gaugeRef.properties.valSummary.numberFormat
	
	function formatGaugeVal(val) {
		var formattedVal = formatNumberValue(numberFormat,val)
		return formattedVal
	}
			
	var gaugeConfig = 
	{
		size: gaugeRef.properties.geometry.sizeWidth,
		min: gaugeRef.properties.minVal,
		max: gaugeRef.properties.maxVal,
		minorTicks: 5,
		valueFormatter: formatGaugeVal,
		thresholdVals: gaugeRef.properties.thresholdVals
	}	
	gaugeConfig.yellowZones = [];
	gaugeConfig.redZones = [];
	gaugeConfig.greenZones = [];
	
	initGaugeDashboardComponentControl($gauge,gaugeConfig,gaugeVal)
	
}
