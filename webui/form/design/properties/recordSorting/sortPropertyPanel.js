function initFormSortPropertyPanel(formInfo) {
	
	var defaultSortRules = []
	
	function saveDefaultFormSortRules(sortRules) {
		console.log("Saving default sort rules for form: " + JSON.stringify(sortRules))
/*		var saveSortRulesParams = {
			parentFormID: viewFormContext.formID,
			sortRules: sortRules
		}
		jsonAPIRequest("recordSort/saveFormSortRules",saveSortRulesParams,function(saveReply) {}) // getRecord			
*/
	}
	
	
	var sortPaneParams = {
		defaultSortRules: defaultSortRules,
		resortFunc: function() {}, // no-op
		initDoneFunc: function() {}, // no-op
		saveUpdatedSortRulesFunc: saveDefaultFormSortRules}
	
	
	initSortRecordsPane(sortPaneParams)
}