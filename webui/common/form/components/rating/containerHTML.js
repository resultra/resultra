
function getRatingControlFromRatingContainer($ratingContainer) {
	var $ratingControl = $ratingContainer.find(".ratingFormComponentControl")
	assert($ratingControl !== undefined, "getRatingControlFromRatingContainer: Can't get control")
	return $ratingControl
}

function ratingFormComponentRatingControlHTML() {
	return '<input type="hidden" class="ratingFormComponentControl"/>' // Rating control from Bootstrap Rating plugin
}

function ratingContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class=" layoutContainer ratingFormContainer">' +
			'<label class="marginBottom0">Rating</label>' +
			'<div class="formRatingControl">' +
				ratingFormComponentRatingControlHTML() + // Rating control from Bootstrap Rating plugin
			'</div>' +
		'</div><';
										
	return containerHTML
}



function initRatingFormComponentControl($container,ratingObjectRef) {
	
	var $ratingControl = getRatingControlFromRatingContainer($container)
	
	function getRatingIconClasses() {
		var ratingIconNameClassesMap = {
			"heart": {
				filled: 'glyphicon glyphicon-heart ratingColorFireRed',
				empty: 'glyphicon glyphicon-heart ratingIconEmptyBackground'
			},
			"star": {
				filled: 'glyphicon glyphicon-star ratingColorStarYellow',
				empty: 'glyphicon glyphicon-star ratingIconEmptyBackground'
			},
			"eyeball": {
				filled: 'glyphicon glyphicon-eye-open',
				empty: 'glyphicon glyphicon-eye-open ratingIconEmptyBackground'
			},
			"warning": {
				filled: 'glyphicon glyphicon-warning-sign',
				empty: 'glyphicon glyphicon-warning-sign ratingIconEmptyBackground'
			},
			"fire": {
				filled: 'glyphicon glyphicon-fire ratingColorFireRed',
				empty: 'glyphicon glyphicon-fire ratingIconEmptyBackground'
			},
			"redFlag": {
				filled: 'glyphicon glyphicon-flag ratingColorFireRed',
				empty: 'glyphicon glyphicon-flag ratingIconEmptyBackground'
			},
			"blackFlag": {
				filled: 'glyphicon glyphicon-flag ratingColorBlack',
				empty: 'glyphicon glyphicon-flag ratingIconEmptyBackground'
			},
			"yellowFlag": {
				filled: 'glyphicon glyphicon-flag ratingColorStarYellow',
				empty: 'glyphicon glyphicon-flag ratingIconEmptyBackground'
			},
			"trash": {
				filled: 'glyphicon glyphicon-trash',
				empty: 'glyphicon glyphicon-trash ratingIconEmptyBackground'
			},
			"time": {
				filled: 'glyphicon glyphicon-time',
				empty: 'glyphicon glyphicon-time ratingIconEmptyBackground'
			},
		}
		
		// Other possible icons: people, happy face, sad face, graducation cap, stop hand
		// thumbs up, pig, money, dollar, bug, check mark, certificate, exclamation,
		// diamond, cog, fill (circle), arrow?, book, bell, lock (privacy), lightening bolt,
		// calculator, apple (rate a teacher), magnifying class (depth), stopwatch(urgency)
		
		var ratingIconClasses = ratingIconNameClassesMap[ratingObjectRef.properties.icon]
		if (ratingIconClasses === undefined) {
			ratingIconClasses = ratingIconNameClassesMap["star"]
		}
		return ratingIconClasses
	}
	
	var ratingIconClasses = getRatingIconClasses()
	
	$ratingControl.rating({
		extendSymbol: function(rating) {
			var ratingIndex = rating-1 // 0 based index
			if(ratingObjectRef.properties.tooltips[ratingIndex] !== undefined) {
				var tooltipText = ratingObjectRef.properties.tooltips[ratingIndex]
				if(tooltipText.length > 0) {
					var tooltipHTML = '<p class="ratingTooltip">' + escapeHTML(tooltipText) + '</p>'
					$(this).tooltip({
						container: 'body',
						placement: 'bottom',
						title: tooltipHTML,
						html: true 
					});
					
				}
			}
			
		},
		filled: ratingIconClasses.filled,
  	  	empty: ratingIconClasses.empty,
		fractions: 2
	})
	
}


/* There isn't a method (that I know of) to re-initialize a rating container. So, to re-initialize,
   the rating control, the DOM elements need to be cleared out and re-initialized. */
function reInitRatingFormComponentControl($container,ratingObjectRef) {
	var $ratingControlContainer = $container.find(".formRatingControl")
	$ratingControlContainer.empty()
	$ratingControlContainer.append(ratingFormComponentRatingControlHTML())
	initRatingFormComponentControl($container,ratingObjectRef)
}
