function summaryValFromDashboardSummaryValContainer($summaryVal) {
	return 	$summaryVal.find(".dashboardSummaryVal")
}


function dashboardSummaryValContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class="layoutContainer dashboardSummaryValContainer" id="'+elementID+'">' +
			'<label class="summaryValLabel">' + 'New Summary Val' + '</label>' +
			'<div class="dashboardSummaryVal">' +
				'<span class="summaryValControl"></span>'+
  			'</div>' +
		'</div><';
						
	return containerHTML
}

function setSummaryValDashboardComponentLabel($container,summaryValRef) {
	var $label = $container.find(".summaryValLabel")
	$label.text(summaryValRef.properties.title)	
}

function initSummaryValDashboardComponentControl($summaryVal,summaryValConfig,summaryVal) {
	
	var $summaryValControlContainer = $summaryVal.find(".summaryValControl")
	
	var summaryValControl = new GaugeUIControl($summaryValControlContainer, summaryValConfig);
	summaryValControl.render()
	summaryValControl.redraw(0.0)
	$summaryVal.data("summaryValControl",summaryValControl) 
	
	summaryValControl.redraw(summaryVal)
	
}


function initSummaryValData(dashboardID,$summaryVal, summaryValData) {
	
	var summaryValRef = summaryValData.summaryVal	
	var summaryValVal = summaryValData.groupedSummarizedVals.overallDataRow.summaryVals[0]
	var numberFormat = summaryValRef.properties.valSummary.numberFormat
	
	function formatSummaryValVal(val) {
		var formattedVal = formatNumberValue(numberFormat,val)
		return formattedVal
	}
			
	var summaryValConfig = 
	{
		size: summaryValRef.properties.geometry.sizeWidth,
		min: 0.0,
		max: 100.0,
		minorTicks: 5,
		valueFormatter: formatSummaryValVal,
		thresholdVals: summaryValRef.properties.thresholdVals
	}	
	summaryValConfig.yellowZones = [];
	summaryValConfig.redZones = [];
	summaryValConfig.greenZones = [];
	
	initSummaryValDashboardComponentControl($summaryVal,summaryValConfig,summaryValVal)
	
}
