// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initThresholdValuesPropertyPanel(panelParams) {
		
		
	var $addThresholdButton = $(createPrefixedSelector(panelParams.elemPrefix,'ValueThresholdAddThresholdButton'))
	
	var $thresholdList = $(createPrefixedSelector(panelParams.elemPrefix,'ThresholdValuesList'))
	$thresholdList.empty()
	
	function saveUpdatedThresholdValues() {
		
		var thresholds = []
		$thresholdList.find(".thresholdValuesPanelListItem").each(function () {
			var $thresholdItem = $(this)
			
			var colorScheme = $thresholdItem.find(".thresholdColorSchemeSelection").val()
			var startingValInput = $thresholdItem.find(".thresholdStartingValInput").val()
			var startingVal = Number(startingValInput)
			
			if( (startingValInput.length > 0) && (!isNaN(startingVal)) && 
				(colorScheme != null) && (colorScheme.length>0)) {
				var thresholdVal = {
					startingVal: startingVal,
					colorScheme: colorScheme
				}
				thresholds.push(thresholdVal)
			}
			
		}) // Each threshold item
		thresholds.sort(function(a,b) { return a.startingVal-b.startingVal })
		console.log("Saving thresholds: " + JSON.stringify(thresholds))
		panelParams.saveThresholdsCallback(thresholds)
	}
	
	function addThreshold(initialVal) {
		var $thresholdItem = $('#thresholdValuesPanelListItem').clone()
		$thresholdItem.attr("id","")
		$thresholdList.append($thresholdItem)
		
		var $deleteButton = $thresholdItem.find(".thresholdValuesListItemDeleteThresholdButton")
		initButtonControlClickHandler($deleteButton,function() {
			console.log("delete threshold button clicked")
			$thresholdItem.remove()
			saveUpdatedThresholdValues()
		})
		
		var $schemeSelection = $thresholdItem.find(".thresholdColorSchemeSelection")
		if (initialVal != null) {
			$schemeSelection.val(initialVal.colorScheme)
		}
		initSelectControlChangeHandler($schemeSelection,function(newScheme) {
			console.log("new color scheme selected: " + newScheme)
			saveUpdatedThresholdValues()
		})
		
		var $startingVal = $thresholdItem.find(".thresholdStartingValInput")
		if(initialVal != null) {
			$startingVal.val(initialVal.startingVal)
		}
		$startingVal.unbind("blur")
		$startingVal.blur(function() { saveUpdatedThresholdValues() })
	}
	
	initButtonControlClickHandler($addThresholdButton,function() {
		console.log("add threshold button clicked")
		addThreshold(null)
	})
	
	for(var initialValIndex = 0; initialValIndex < panelParams.initialThresholdVals.length; initialValIndex++) {
		var currThreshVal = panelParams.initialThresholdVals[initialValIndex]
		addThreshold(currThreshVal)
	}
	
}
