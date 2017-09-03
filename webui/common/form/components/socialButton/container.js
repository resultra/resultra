
function getSocialButtonControlFromContainer($socialButtonContainer) {
	var $socialButtonControl = $socialButtonContainer.find(".socialButtonFormComponentControl")
	assert($socialButtonControl !== undefined, "getSocialButtonControlFromContainer: Can't get control")
	return $socialButtonControl
}

function socialButtonFormComponentButtonControlHTML() {
	
	var buttonHTML = '<button type="button" class="socialButtonFormComponentControl btn btn-default">' +
  			'<span class="controlIcon glyphicon glyphicon-star" aria-hidden="true"></span>&nbsp;' +
  			'<span class="controlLabel" aria-hidden="true"></span>' +
		'</button>'
	
	return buttonHTML
}

function socialButtonContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class=" layoutContainer socialButtonFormContainer">' +
			componentHelpPopupButtonHTML() +
			'<div class="socialButtonControl">' +
				socialButtonFormComponentButtonControlHTML() + // Rating control from Bootstrap Rating plugin
			'</div>' +
		'</div><';
										
	return containerHTML
}

function socialButtonTableCellContainerHTML() {
	return '' +
		'<div class=" layoutContainer socialButtonTableCellContainer">' +
					'<div class="socialButtonControl">' +
						socialButtonFormComponentButtonControlHTML() + // Rating control from Bootstrap Rating plugin
					'</div>' +
		'</div>';
}


function getSocialButtonIconClasses(socialButtonObjectRef) {
	var socialButtonIconNameClassesMap = {
		"heart": {
			filled: 'glyphicon glyphicon-heart socialButtonColorFireRed',
			empty: 'glyphicon glyphicon-heart socialButtonIconEmptyBackground'
		},
		"star": {
			filled: 'glyphicon glyphicon-star socialButtonColorStarYellow fa-lg',
			empty: 'glyphicon glyphicon-star socialButtonIconEmptyBackground fa-lg'
		},
		"eyeball": {
			filled: 'glyphicon glyphicon-eye-open',
			empty: 'glyphicon glyphicon-eye-open socialButtonIconEmptyBackground'
		},
		"warning": {
			filled: 'glyphicon glyphicon-warning-sign',
			empty: 'glyphicon glyphicon-warning-sign socialButtonIconEmptyBackground'
		},
		"fire": {
			filled: 'glyphicon glyphicon-fire socialButtonColorFireRed',
			empty: 'glyphicon glyphicon-fire socialButtonIconEmptyBackground'
		},
		"redFlag": {
			filled: 'glyphicon glyphicon-flag socialButtonColorFireRed',
			empty: 'glyphicon glyphicon-flag socialButtonIconEmptyBackground'
		},
		"blackFlag": {
			filled: 'glyphicon glyphicon-flag socialButtonColorBlack',
			empty: 'glyphicon glyphicon-flag socialButtonIconEmptyBackground'
		},
		"yellowFlag": {
			filled: 'glyphicon glyphicon-flag socialButtonColorStarYellow',
			empty: 'glyphicon glyphicon-flag socialButtonIconEmptyBackground'
		},
		"trash": {
			filled: 'glyphicon glyphicon-trash',
			empty: 'glyphicon glyphicon-trash socialButtonIconEmptyBackground'
		},
		"time": {
			filled: 'glyphicon glyphicon-time',
			empty: 'glyphicon glyphicon-time socialButtonIconEmptyBackground'
		},
		"thumbsUp": {
			filled: 'glyphicon glyphicon-thumbs-up  fa-lg',
			empty: 'glyphicon glyphicon-thumbs-up socialButtonIconEmptyBackground  fa-lg'
		},
		"greenCheck": {
			filled: 'glyphicon glyphicon-ok socialButtonColorGreen fa-lg',
			empty: 'glyphicon glyphicon-ok socialButtonIconEmptyBackground  fa-lg'
		},
		"greenForward": {
			filled: 'glyphicon glyphicon-circle-arrow-right socialButtonColorGreen fa-lg',
			empty: 'glyphicon glyphicon-circle-arrow-right socialButtonIconEmptyBackground  fa-lg'
		},
	}
	
	// Other possible icons: people, happy face, sad face, graducation cap, stop hand
	// thumbs up, pig, money, dollar, bug, check mark, certificate, exclamation,
	// diamond, cog, fill (circle), arrow?, book, bell, lock (privacy), lightening bolt,
	// calculator, apple (rate a teacher), magnifying class (depth), stopwatch(urgency)
	
	var socialButtonIconClasses = socialButtonIconNameClassesMap[socialButtonObjectRef.properties.icon]
	if (socialButtonIconClasses === undefined) {
		socialButtonIconClasses = socialButtonIconNameClassesMap["star"]
	}
	return socialButtonIconClasses
	
}

function initSocialButtonFormComponentControl($container,socialButtonObjectRef) {
	
	var $socialButtonControl = getSocialButtonControlFromContainer($container)
	
	
	var socialButtonIconClasses = getSocialButtonIconClasses(socialButtonObjectRef)
	
	var $iconSpan = $container.find(".controlIcon")
	$iconSpan.addClass(socialButtonIconClasses.filled)
	$iconSpan.addClass("controlIcon")
	
}

function setSocialButtonButtonIcon(iconIsSet,$container,socialButtonObjectRef) {
	
	var $iconSpan = $container.find(".controlIcon")
	$iconSpan.removeClass()
	$iconSpan.addClass("controlIcon")
	
	var socialButtonIconClasses = getSocialButtonIconClasses(socialButtonObjectRef)
	if(iconIsSet) {
		$iconSpan.addClass(socialButtonIconClasses.filled)
		
	} else {
		$iconSpan.addClass(socialButtonIconClasses.empty)
		
	}
}


/* There isn't a method (that I know of) to re-initialize a rating container. So, to re-initialize,
   the rating control, the DOM elements need to be cleared out and re-initialized. */
function reInitSocialButtonFormComponentControl($container,socialButtonObjectRef) {
	var $socialButtonControlContainer = $container.find(".socialButtonControl")
	$socialButtonControlContainer.empty()
	$socialButtonControlContainer.append(socialButtonFormComponentButtonControlHTML())
	initSocialButtonFormComponentControl($container,socialButtonObjectRef)
}

function setSocialButtonComponentLabel($socialButton,socialButtonRef) {
	var $label = $socialButton.find('.controlLabel')
	
	setFormComponentLabel($label,socialButtonRef.properties.fieldID,
			socialButtonRef.properties.labelFormat)	
	
}

function initSocialButtonFormComponentContainer($container,socialButtonObjectRef) {
	setSocialButtonComponentLabel($container,socialButtonObjectRef)
	initComponentHelpPopupButton($container, socialButtonObjectRef)
	
	initSocialButtonFormComponentControl($container,socialButtonObjectRef)
		
	setElemFixedWidthFlexibleHeight($container,
				socialButtonObjectRef.properties.geometry.sizeWidth)
}

