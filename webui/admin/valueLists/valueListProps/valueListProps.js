function initValueListValueListProperties(valueListInfo) {
	
	var $valueList = $('#valueListValuesList')
	
	function updateValueListValues() {
		console.log("Updating value list values ...")
		var values = []
		$valueList.find('.valueListValueItem').each(function() {
			var $valueInput = $(this).find('.valueListValue')
			var value = $valueInput.val()
			if(value != null && value.length > 0) {
				console.log("Got a value: " + value)			
				var valueParam = { textValue: value }
				values.push(valueParam)
			}
		})
		var setValuesParams = {
			valueListID:valueListInfo.valueListID,
			values:values
		}
		jsonAPIRequest("valueList/setValues",setValuesParams,function(updatedLinkInfo) {
			console.log("Done changing value list")
		})
	}
	
	function addValueListValue(initialVal) {
		var $listItem = $('#valueListValueItem').clone()
		$listItem.attr("id","")
		
		var $valueInput = $listItem.find('.valueListValue')
		if (initialVal != null) {
			$valueInput.val(initialVal)
		}
		$valueInput.blur(function() {
			updateValueListValues()
		})
		
		var $deleteButton = $listItem.find('.valueListDeleteValButton')
		initButtonControlClickHandler($deleteButton,function() {
			$listItem.remove()
			updateValueListValues()
		})
		
		$valueList.append($listItem)
	}
	
	// Initialize the list with any existing values
	$valueList.empty()
	for(var currValIndex=0;currValIndex<valueListInfo.properties.values.length;currValIndex++) {
		var currVal = valueListInfo.properties.values[currValIndex]
		addValueListValue(currVal.textValue)
	}
	
    $valueList.sortable({
		placeholder: "ui-state-highlight",
		cursor:"move",
		update: function( event, ui ) {
			console.log("Value list order changed")	
			updateValueListValues()		
		}
    });
	
	initButtonClickHandler('#adminValueListAddValueButton',function() {
		addValueListValue()
	})
	
	
}