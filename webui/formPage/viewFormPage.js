$(document).ready(function() {	
	 
	function getPageConfig(pageConfigDoneCallback) {
		
		var pageConfig = {}
		
		pageConfig.databaseID = viewFormPageContext.databaseID
		pageConfig.recordID =  viewFormPageContext.recordID
		pageConfig.formID = viewFormPageContext.formID
		
		if(viewFormPageContext.srcColumnID.length > 0) {
			// Get the default values from the column used to open the form. 
			var getButtonParams = {
				buttonID: viewFormPageContext.srcColumnID
			}
			jsonAPIRequest("tableView/formButton/getFromButtonID",getButtonParams,function(buttonRef) {
				pageConfig.defaultVals = buttonRef.properties.defaultValues
				pageConfig.saveMode = buttonRef.properties.popupBehavior.popupMode
				pageConfigDoneCallback(pageConfig)	
			})
		} else if(viewFormPageContext.srcFrmButtonID.length > 0) {
			var getButtonParams = {
				buttonID: viewFormPageContext.srcFrmButtonID
			}
			jsonAPIRequest("frm/formButton/get",getButtonParams,function(buttonRef) {
				pageConfig.defaultVals = buttonRef.properties.defaultValues
				pageConfig.saveMode = buttonRef.properties.popupBehavior.popupMode
				pageConfigDoneCallback(pageConfig)	
			})
				
		} else {
			// Load without default values.
			pageConfig.defaultVals = []
			pageConfig.saveMode = FormViewModeSave
			pageConfigDoneCallback(pageConfig)	
		}
		
	}
	
	getPageConfig(function(pageConfig) {
		getRecordRefAndChangeSetID(pageConfig,initRecordFormView)
	})
	
					
}); // document ready