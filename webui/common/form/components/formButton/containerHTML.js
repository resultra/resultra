function buttonFromFormButtonContainer($formButton) {
	return 	$formButton.find(".formButton")
}


function formButtonContainerHTML()
{	
	var containerHTML = ''+
		'<div class=" layoutContainer buttonFormContainer">' +
			'<button type="button" class="btn btn-primary formButton">' + 
			'Open Form' +
			'</button>' +
		'</div>';
						
	return containerHTML
}

function formButtonTableCellContainerHTML()
{	
	var containerHTML = ''+
		'<div class=" layoutContainer buttonTableCellContainer">' +
			'<button type="button" class="btn btn-primary formButton">' + 
			'Open Form' +
			'</button>' +
		'</div>';
						
	return containerHTML
}


function setFormButtonSize($container,newSize) {
	
	var $button = $container.find('button')
	$button.removeClass("btn-lg btn-sm btn-xs")
	
	var sizeBtnClassMap = {
		"large": "btn-lg",
		"medium":"",
		"small":"btn-sm",
		"xsmall":"btn-xs"
	}
	
	var btnSizeClass = sizeBtnClassMap[newSize]
	
	if (btnSizeClass !== undefined && btnSizeClass.length > 0) {
			$button.addClass(btnSizeClass)
	}

}

function setFormButtonColorScheme($container,newScheme) {
	var $button = $container.find('button')
	$button.removeClass("btn-primary btn-default btn-success btn-info btn-warning btn-danger btn-link")
	
	var schemeBtnClassMap = {
		"default": "btn-default",
		"primary":"btn-primary",
		"success":"btn-success",
		"info":"btn-info",
		"warning":"btn-warning",
		"danger":"btn-danger",
		"link":"btn-link"
	}
		
	var schemeClass = schemeBtnClassMap[newScheme]
	
	if (schemeClass !== undefined && schemeClass.length > 0) {
			$button.addClass(schemeClass)
	}
	
}

function setFormButtonLabel($container,buttonRef) {
	
	var iconNameClassMap = {
		"none":undefined,
		"check":"glyphicon glyphicon-check",
		"option":"glyphicon glyphicon-option-horizontal",
		"rchevron":"glyphicon glyphicon-chevron-right",
		"enter":"glyphicon glyphicon-log-in",
		"exit":"glyphicon glyphicon-log-out",
		"comment":"glyphicon glyphicon-comment",
		"zoom":"glyphicon glyphicon-zoom-in",
		"time":"glyphicon glyphicon-time",
		"cog":"glyphicon glyphicon-cog",
		"calculator":"fa fa-calculator"
	}
	
	function setButtonLabel(buttonLabel) {
		var $button = $container.find(".formButton")
		
		var iconClass = iconNameClassMap[buttonRef.properties.icon]
		$button.empty()
		
		var $nameSpan = $('<span></span>')
		$nameSpan.text(buttonLabel)
		if(iconClass !== undefined) {
			var $iconSpan = $('<span aria-hidden="true"></span>')
			$iconSpan.addClass(iconClass)
			$button.append($iconSpan)
			$nameSpan.addClass("marginLeft5")
			$button.append($nameSpan)
		} else {
			$button.append($nameSpan)			
		}	
		
	}
	
	var labelFormat = buttonRef.properties.buttonLabelFormat
	
    switch (labelFormat.labelType) {
    	case "none":
			setButtonLabel("")
    		break;
    	case "custom":
			setButtonLabel(labelFormat.customLabel)
    		break;
    	default:
			jsonAPIRequest("frm/getFormInfo", { formID: buttonRef.properties.linkedFormID }, function(formInfo) {
				setButtonLabel(formInfo.form.name)
			})
    		break;
    }
	
}

function calcFormButtonMinTableCellWidth(buttonRef, formsByID) {
	var labelFormat = buttonRef.properties.buttonLabelFormat
	
	var label = ""
    switch (labelFormat.labelType) {
    	case "none":
    		break;
    	case "custom":
			label = labelFormat.customLabel
    		break;
    	default:
			var formInfo = formsByID[buttonRef.properties.linkedFormID]
			label = formInfo.name
    		break;
    }
	
	var overallWidth = calcTextWidth(label)
	
	if (buttonRef.properties.icon !== undefined && buttonRef.properties.icon !== 'none') {
		overallWidth += 30
	}
	if (overallWidth < 75) {
		overallWidth = 75
	}
	
	return overallWidth
	
	
}

function setFormButtonHeader($container,buttonRef) {
	jsonAPIRequest("frm/getFormInfo", { formID: buttonRef.properties.linkedFormID }, function(formInfo) {
		var $nameSpan = $('<span></span>')
		$nameSpan.text(formInfo.form.name)
		$container.append($nameSpan)
	})
	
}