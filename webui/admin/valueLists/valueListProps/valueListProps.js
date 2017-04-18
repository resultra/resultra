function initValueListValueListProperties(valueListInfo) {
	
	var $valueList = $('#valueListValuesList')
	
	function addValueListValue() {
		var $listItem = $('#valueListValueItem').clone()
		$listItem.attr("id","")
		
		var $deleteButton = $listItem.find('.valueListDeleteValButton')
		initButtonControlClickHandler($deleteButton,function() {
			$listItem.remove()
		})
		
		$valueList.append($listItem)
	}
	
    $valueList.sortable({
		placeholder: "ui-state-highlight",
		cursor:"move",
		update: function( event, ui ) {
			console.log("Value list order changed")			
		}
    });
	
	
		
	initButtonClickHandler('#adminValueListAddValueButton',function() {
		addValueListValue()	
	})
	
	
}