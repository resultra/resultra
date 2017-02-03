function initAdminNewItemPresetSettings(databaseID) {
		
		
	function addPresetToAdminPresetList(presetInfo) {
 
		var $presetListItem = $('#adminNewItemPresetListItemTemplate').clone()
		$presetListItem.attr("id","")
	
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
		
	// Retrieve presets from the server, populate the list of presets.
	var presetParams = { parentDatabaseID: databaseID }
	jsonAPIRequest("newItem/getPresets",presetParams,function(presetList) {
		for(var presetIndex = 0; presetIndex < presetList.length; presetIndex++) {
			var currPreset = presetList[presetIndex]
			addPresetToAdminPresetList(currPreset)
		}
	})
	
	
	initButtonClickHandler('#adminNewNewItemPresetButton',function() {
		console.log("New field button clicked")
		openNewNewItemPresetDialog(databaseID)
	})
	
	
	
}