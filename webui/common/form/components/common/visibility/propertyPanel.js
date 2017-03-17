function initFormComponentVisibilityPropertyPanel(params) {
	
	var visibilityFilterConditionPropertyPanelParams = {
		elemPrefix: params.elemPrefix,
		databaseID: params.databaseID,
		defaultFilterRules: params.initialConditions,
		initDone: function () {},
		updateFilterRules: function (updatedFilterRules) {
			
			console.log("Updating form component filter conditions: " + JSON.stringify(updatedFilterRules))
			/*
			var setPreFiltersParams = {
				listID: listInfo.listID,
				filterRules: updatedFilterRules
			}
			jsonAPIRequest("itemList/setPreFilterRules",setPreFiltersParams,function(updatedList) {
				console.log(" Pre filters updated")
			}) // set record's number field value
			*/
		
		}
	
	}
	initFilterPropertyPanel(visibilityFilterConditionPropertyPanelParams)
	
}