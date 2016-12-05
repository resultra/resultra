$(document).ready(function() {	
	
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }


	$('#editItemListPropsPage').layout({
			inset: zeroPaddingInset,
			north: fixedUILayoutPaneParams(40),
			west: {
				size: 250,
				resizable:false,
				slidable: false,
				spacing_open:4,
				spacing_closed:4,
				initClosed:false // panel is initially open	
			}
		})
		
	initDatabaseTOC(itemListPropsContext.databaseID)
		
	initUserDropdownMenu()
		
		var listElemPrefix = "itemList_"
		
		var getItemListParams = { listID: itemListPropsContext.listID }
		jsonAPIRequest("itemList/get",getItemListParams,function(listInfo) {
			var filterPropertyPanelParams = {
				elemPrefix: listElemPrefix,
				tableID: listInfo.parentTableID,
				defaultFilterRules: listInfo.properties.defaultFilterRules,
				initDone: function () {},
				updateFilterRules: function (updatedFilterRules) {
					var setDefaultFiltersParams = {
						listID: listInfo.listID,
						filterRules: updatedFilterRules
					}
					jsonAPIRequest("itemList/setDefaultFilterRules",setDefaultFiltersParams,function(updatedList) {
						console.log(" Default filters updated")
					}) // set record's number field value
				
				}
			
			}
			initFilterPropertyPanel(filterPropertyPanelParams)
			
			
			function saveDefaultListSortRules(sortRules) {
				console.log("Saving default sort rules for list: " + JSON.stringify(sortRules))
				var saveSortRulesParams = {
					listID:listInfo.listID,
					sortRules: sortRules
				}
				jsonAPIRequest("itemList/setDefaultSortRules",saveSortRulesParams,function(saveReply) {
					console.log("Done saving default sort rules")
				})			

			}
	
	
			var sortPaneParams = {
				defaultSortRules: listInfo.properties.defaultRecordSortRules,
				tableID: listInfo.parentTableID,
				resortFunc: function() {}, // no-op
				initDoneFunc: function() {}, // no-op
				saveUpdatedSortRulesFunc: saveDefaultListSortRules}
	
	
			initSortRecordsPane(sortPaneParams)
			

		}) // set record's number field value
	
		
	
	
})