function initFormSortPropertyPanel(formInfo) {
		
	function saveDefaultFormSortRules(sortRules) {
		console.log("Saving default sort rules for form: " + JSON.stringify(sortRules))
		var saveSortRulesParams = {
			formID:formInfo.formID,
			sortRules: sortRules
		}
		jsonAPIRequest("frm/setDefaultSortRules",saveSortRulesParams,function(saveReply) {
			console.log("Done saving default sort rules")
		})			

	}
	
	
	var sortPaneParams = {
		defaultSortRules: formInfo.properties.defaultRecordSortRules,
		resortFunc: function() {}, // no-op
		initDoneFunc: function() {}, // no-op
		saveUpdatedSortRulesFunc: saveDefaultFormSortRules}
	
	
	initSortRecordsPane(sortPaneParams)
}