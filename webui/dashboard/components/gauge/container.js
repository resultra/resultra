function gaugeFromDashboardGaugeContainer($gauge) {
	return 	$gauge.find(".dashboardGauge")
}


function dashboardGaugeContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class="layoutContainer dashboardGaugeFormContainer" id="'+elementID+'">' +
			'<span class="h3 dashboardGauge">' +
			'New Gauge' +
			'</span>' +
		'</div><';
						
	return containerHTML
}

function setGaugeDashboardComponentLabel($container,gaugeRef) {
	var $label = $container.find("span")
	$label.text(gaugeRef.properties.title)	
}
