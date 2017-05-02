function headerFromDashboardHeaderContainer($header) {
	return 	$header.find(".dashboardHeader")
}


function dashboardHeaderContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class="layoutContainer dashboardHeaderFormContainer" id="'+elementID+'">' +
			'<span class="h3 dashboardHeader">' +
			'New Header' +
			'</span>' +
		'</div><';
						
	return containerHTML
}

function setHeaderDashboardComponentHeaderSize($container,headerSize) {
	
	var sizeSizeClassMap = {
		"xlarge":"h2",
		"large":"h3",
		"medium":"h4",
		"small":"h5",
		"xsmall":"h6"
	}
	var sizeClass = sizeSizeClassMap[headerSize]
	
	$container.find("span").removeClass("h1 h2 h3 h4 h5 h6")
	$container.find("span").addClass(sizeClass)
}

function setHeaderDashboardComponentUnderlined($container,isUnderlined) {
	
	var $header = $container.find("span")
	
	$header.removeClass("underlinedFormHeader")
	if(isUnderlined) {
		$header.addClass("underlinedFormHeader")
	}
}

function setHeaderDashboardComponentLabel($container,headerRef) {
	var $label = $container.find("span")
	$label.text(headerRef.properties.title)
}
