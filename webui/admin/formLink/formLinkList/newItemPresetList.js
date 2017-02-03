function initAdminNewItemPresetSettings(databaseID) {
		
		
	function addPresetToAdminPresetList(presetInfo) {
 
		var $presetListItem = $('#adminNewItemPresetListItemTemplate').clone()
		$presetListItem.attr("id","")
	
		$presetListItem.attr("data-linkID",presetInfo.linkID)
	
		var $editPropsButton = $presetListItem.find(".editNewItemPresetPropsButton")
		var editPropsLink = '/admin/formLink/' + presetInfo.linkID
		$editPropsButton.attr('href',editPropsLink)
		
		var $nameLabel = $presetListItem.find(".newItemPresetLabel")
		$nameLabel.text(presetInfo.name)
		 	
		$('#adminNewItemPresetList').append($presetListItem)		
	}
		
	// Retrieve presets from the server, populate the list of presets.
	var presetParams = { parentDatabaseID: databaseID }
	jsonAPIRequest("formLink/getList",presetParams,function(presetList) {
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