function headerFromHeaderContainer($header) {
	return 	$header.find(".formHeader")
}


function formHeaderContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class="layoutContainer headerFormContainer" id="'+elementID+'">' +
			'<span class="h3 formHeader">' +
			'New Header' +
			'</span>' +
		'</div><';
						
	return containerHTML
}

function setHeaderFormComponentHeaderSize($container,headerSize) {
	
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

function setHeaderFormComponentUnderlined($container,isUnderlined) {
	
	var $header = $container.find("span")
	
	$header.removeClass("underlinedFormHeader")
	if(isUnderlined) {
		$header.addClass("underlinedFormHeader")
	}
}