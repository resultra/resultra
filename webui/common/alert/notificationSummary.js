
function formatNotificationSummary(fieldsByID,notification,alert) {
	
	
	function formatTimeNotification(fieldInfo,triggerCond) {
		
		switch (triggerCond.conditionID) {
		case "changed":
			var dateBefore = moment(notification.dateBefore).format('MM/DD/YYYY')
			var dateAfter = moment(notification.dateAfter).format('MM/DD/YYYY')
			return "'" + fieldInfo.name + "' changed from " + dateBefore + " to " + dateAfter
		default:
			return ""
		}
		
	}
	
	
	var triggerCond = notification.triggerCondition
	var triggerFieldID = triggerCond.fieldID
	var triggerFieldInfo = fieldsByID[triggerFieldID]
	
	switch (triggerFieldInfo.type) {
	case fieldTypeTime:
		return formatTimeNotification(triggerFieldInfo,triggerCond)
	default:
		return ""
	}
	
	return ""
}