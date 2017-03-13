function initThresholdValuesPropertyPanel(panelParams) {
		
		
	var $addThresholdButton = $(createPrefixedSelector(panelParams.elemPrefix,'ValueThresholdAddThresholdButton'))
	var $thresholdList = $(createPrefixedSelector(panelParams.elemPrefix,'ThresholdValuesList'))
	
	function addThreshold() {
		var $thresholdItem = $('#thresholdValuesPanelListItem').clone()
		$thresholdItem.attr("id","")
		$thresholdList.append($thresholdItem)
		
		var $deleteButton = $thresholdItem.find(".thresholdValuesListItemDeleteThresholdButton")
		initButtonControlClickHandler($deleteButton,function() {
			console.log("delete threshold button clicked")
			$thresholdItem.remove()
		})
	}
	
	initButtonControlClickHandler($addThresholdButton,function() {
		console.log("add threshold button clicked")
		addThreshold()
	})	
		
	initThresholdValuesList(panelParams)	
	
}
