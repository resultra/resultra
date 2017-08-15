$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#alertPropsPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(alertPropsContext.databaseID)
	
	function initAlertFormProperties(alertInfo) {
		var selectFormParams = {
			menuSelector: "#alertPropFormSelection",
			parentDatabaseID: alertPropsContext.databaseID,
			initialFormID: alertInfo.properties.formID
		}	
		populateFormSelectionMenu(selectFormParams)
		var $formSelection = $("#alertPropFormSelection")
		initSelectControlChangeHandler($formSelection, function(selectedFormID) {

			var setFormParams = {
				alertID: alertPropsContext.alertID,
				formID: selectedFormID
			}	

			jsonAPIRequest("alert/setForm",setFormParams,function(setFormParams) {
				console.log("Done setting form for alert")
			}) 
			
		})
		
	} // initItemListFormProperties
	
	var getAlertParams = { 
		alertID: alertPropsContext.alertID
	}
	jsonAPIRequest("alert/get",getAlertParams,function(alertInfo) {
		initAlertFormProperties(alertInfo)
	}) 
	
	var conditionPropsParams = {
		databaseID: alertPropsContext.databaseID,
		alertID: alertPropsContext.alertID
	}
	initAlertConditionProps(conditionPropsParams)
})