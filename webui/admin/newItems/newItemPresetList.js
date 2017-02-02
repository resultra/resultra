function initAdminNewItemPresetSettings(databaseID) {
		
		
	function addPresetToAdminPresetList(presetInfo) {
 
		var $presetListItem = $('#adminNewItemPresetListItemTemplate').clone()
		$fieldListItem.attr("id","")
	
		$presetListItem.attr("data-presetID",presetInfo.presetID)
	
		var $editPropsButton = $presetListItem.find(".editNewItemPresetPropsButton")

		var $nameLabel = $presetListItem.find(".newItemPresetLabel")
		$nameLabel.text(presetInfo.name)
		
		initButtonControlClickHandler($editPropsButton,function() {
			console.log("Edit preset properties: " + JSON.stringify(presetInfo))
//			openFieldPropertiesDialog(fieldInfo)
		}) 
 	
		$('#adminNewItemPresetList').append($presetListItem)		
	}
		
	// TODO - Retrieve presets from the server, populate the list of presets.
	// For each presetInfo ... addPresetToAdminPresetList(presetInfo)
	
	
	initButtonClickHandler('#adminNewNewItemPresetButton',function() {
		console.log("New field button clicked")
		openNewNewItemPresetDialog(databaseID)
	})
	
	
	
}