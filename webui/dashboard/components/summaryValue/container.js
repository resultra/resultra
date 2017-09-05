function summaryValFromDashboardSummaryValContainer($summaryVal) {
	return 	$summaryVal.find(".dashboardSummaryVal")
}


function dashboardSummaryValContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class="layoutContainer dashboardSummaryValContainer" id="'+elementID+'">' +
			'<div class="row">' +
				'<div class="col-sm-8">' +
					'<label class="summaryValLabel">' + 'New Summary Val' + '</label>' +
				'</div>' +
				'<div class="col-sm-4 componentHeaderButtons">' +
					componentHelpPopupButtonHTML() +
				'</div>' +
			'</div>' +
			'<div class="dashboardSummaryVal alert defaultAlertBackground" role="alert">' +
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
	
	var formattedVal = summaryValConfig.valueFormatter(summaryVal)
	$summaryValControlContainer.text(formattedVal)
	
	var colorScheme = getThresholdColorScheme(summaryValConfig.thresholdVals,summaryVal)
	var colorClass = "defaultAlertBackground"
	if (colorScheme !== colorThresholdSchemeDefault) {
		colorClass = "alert-" + colorScheme
	}
		
	var $summaryValValContainer = $summaryVal.find(".dashboardSummaryVal")
	$summaryValValContainer.removeClass("defaultAlertBackground alert-success alert-danger alert-info alert-warning")
	$summaryValValContainer.addClass(colorClass)
	
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
		valueFormatter: formatSummaryValVal,
		thresholdVals: summaryValRef.properties.thresholdVals
	}	
	
	initSummaryValDashboardComponentControl($summaryVal,summaryValConfig,summaryValVal)
	
}
