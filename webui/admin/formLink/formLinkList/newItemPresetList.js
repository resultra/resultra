function initAdminNewItemPresetSettings(databaseID) {
		
		
	function addPresetToAdminPresetList(presetInfo) {
 
		var $presetListItem = $('#adminNewItemPresetListItemTemplate').clone()
		$presetListItem.attr("id","")
	
		$presetListItem.attr("data-presetID",presetInfo.presetID)
	
		var $editPropsButton = $presetListItem.find(".editNewItemPresetPropsButton")
		var editPropsLink = '/admin/formLink/' + presetInfo.presetID
		$editPropsButton.attr('href',editPropsLink)
		
		var $nameLabel = $presetListItem.find(".newItemPresetLabel")
		$nameLabel.text(presetInfo.name)
		 	
		$('#adminNewItemPresetList').append($presetListItem)		
	}
		
	// Retrieve presets from the server, populate the list of presets.
	var presetParams = { parentDatabaseID: databaseID }
	jsonAPIRequest("formLink/getPresets",presetParams,function(presetList) {
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