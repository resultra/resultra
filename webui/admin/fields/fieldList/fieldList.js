function initAdminFieldSettings(databaseID) {
		
		
	function addFieldToAdminFieldList(fieldInfo) {
 
		var $fieldListItem = $('#adminFieldItemTemplate').clone()
		$fieldListItem.attr("id","")
	
		$fieldListItem.attr("data-fieldID",fieldInfo.fieldID)
	
		var $editFieldPropsButton = $fieldListItem.find(".editFieldPropsButton")
		var editPropsLink = '/admin/field/' + fieldInfo.fieldID
		$editFieldPropsButton.attr("href",editPropsLink)

		var $nameLabel = $fieldListItem.find(".fieldNameLabel")
		$nameLabel.text(fieldInfo.name)
		
	
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
		openNewFieldDialog(databaseID)
	})
	
	
	
}