function initAdminFieldSettings(databaseID) {
		
		
	function addFieldToAdminFieldList(fieldInfo) {
 
		var $fieldListItem = $('#adminFieldItemTemplate').clone()
		$fieldListItem.attr("id","")
	
		$fieldListItem.attr("data-fieldID",fieldInfo.fieldID)
	
		var $editFieldPropsButton = $fieldListItem.find(".editFieldPropsButton")

		var $nameLabel = $fieldListItem.find(".fieldNameLabel")
		$nameLabel.text(fieldInfo.name)
		
		initButtonControlClickHandler($editFieldPropsButton,function() {
			console.log("Edit field properties: " + JSON.stringify(fieldInfo))
			openFieldPropertiesDialog(fieldInfo)
		}) 
 	
		$('#adminFieldList').append($fieldListItem)		
	}
		
	
	loadFieldInfo(databaseID, [fieldTypeAll],function(fieldsByID) {
		for (var fieldID in fieldsByID) {
	
			var fieldInfo = fieldsByID[fieldID]	
			
			addFieldToAdminFieldList(fieldInfo)	
	
		} // for each  field
	})
	
	initButtonClickHandler('#adminNewFieldButton',function() {
		console.log("New field button clicked")
		// TODO - Open new field dialog
	})
	
	
	
}