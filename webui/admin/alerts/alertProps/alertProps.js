$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#alertPropsPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(alertPropsContext.databaseID)
	
	var conditionPropsParams = {
		databaseID: alertPropsContext.databaseID,
		alertID: alertPropsContext.alertID
	}
	initAlertConditionProps(conditionPropsParams)
})