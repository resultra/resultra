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

		}) // set record's number field value
	
		
	
	
})