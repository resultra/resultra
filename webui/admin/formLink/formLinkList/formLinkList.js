function initAdminFormLinkSettings(pageContext) {
		
	var $formLinkList = $('#adminFormLinkList')
		
	function addPresetToAdminPresetList(presetInfo) {
 
		var $presetListItem = $('#formLinkItemTemplate').clone()
		$presetListItem.attr("id","")
	
		$presetListItem.attr("data-linkID",presetInfo.linkID)
	
		var $editPropsButton = $presetListItem.find(".editFormLinkPropsButton")
		$editPropsButton.click(function(e) {
			e.preventDefault()
			$editPropsButton.blur()			
			var editPropsContentURL = '/admin/formLink/' + presetInfo.linkID
			setSettingsPageContent(editPropsContentURL,function() {
				initNewItemLinkPropsSettingsPageContent(pageContext,presetInfo)
			})
		})
		
		var $nameLabel = $presetListItem.find(".adminFormLinkItemLabel")
		$nameLabel.text(presetInfo.name)
		 	
		$formLinkList.append($presetListItem)		
	}
		
    $formLinkList .sortable({
		placeholder: "ui-state-highlight",
		cursor:"move",
		update: function( event, ui ) {
			
			var linkOrder = []
			$formLinkList.find(".formLinkListItem").each( function() {
				var linkID = $(this).attr('data-linkID')
				linkOrder.push(linkID)
			})
			var setOrderParams = {
				databaseID:pageContext.databaseID,
				formLinkOrder: linkOrder
			}
			console.log("New form link sort order:" + JSON.stringify(linkOrder))
			jsonAPIRequest("database/setFormLinkOrder",setOrderParams,function(dbInfo) {
				console.log("Done changing database form link order")
			})
			
		}
    });

	// Retrieve presets from the server, populate the list of presets.
	var presetParams = { parentDatabaseID: pageContext.databaseID }
	jsonAPIRequest("formLink/getList",presetParams,function(presetList) {
		for(var presetIndex = 0; presetIndex < presetList.length; presetIndex++) {
			var currPreset = presetList[presetIndex]
			addPresetToAdminPresetList(currPreset)
		}
	})
	
	
	initButtonClickHandler('#adminNewFormLinkButton',function() {
		console.log("New field button clicked")
		openNewNewItemPresetDialog(pageContext)
	})
	
	
	
}