function initFormComponentVisibilityPropertyPanel(params) {
	
	var visibilityFilterConditionPropertyPanelParams = {
		elemPrefix: params.elemPrefix,
		databaseID: params.databaseID,
		defaultFilterRules: params.initialConditions,
		initDone: function () {},
		updateFilterRules: function (updatedFilterRules) {
			
			console.log("Updating form component filter conditions: " + JSON.stringify(updatedFilterRules))
			params.saveVisibilityConditionsCallback(updatedFilterRules)		
		}
	
	}
	initFilterPropertyPanel(visibilityFilterConditionPropertyPanelParams)
	
}