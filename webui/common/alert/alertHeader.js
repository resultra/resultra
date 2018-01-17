function initAlertHeader(databaseID,seeAllAlertsCallback) {
	
	var $header = $('#alertPageHeaderMenu')
	var $alertCountBadge = $header.find(".badge")
	
	function clearAlertCount() {
		$alertCountBadge.text(0)
		$alertCountBadge.hide()
		
		var advanceAlertParams = { parentDatabaseID: databaseID }
		jsonAPIRequest("alert/advanceNotificationTime",advanceAlertParams,function(reply) {})
		
	}
	
	
	loadNotificationListInfo(databaseID, function(notificationList,formsByID,fieldsByID) {
		
		function initAlertCountBadge() {
		
			var alertCount = notificationList.numAlertsNotSeen
			if(alertCount > 0) {
				$alertCountBadge.show()
				$alertCountBadge.text(alertCount)
			} else {
				$alertCountBadge.hide()
			}
			
		}
		
		function createAlertListItem(rawDataIndex) {
			
			var currNotification = notificationList.notifications[rawDataIndex]
			var currAlert = notificationList.alertsByID[currNotification.alertID]
				
			var $alertListItem = $('<li><a class="alertFormLink notificationListItem"></a></li>')
							
			var $alertLink = $alertListItem.find("a")
			
			$alertLink.click(function(e) {
				clearAlertCount()
				e.preventDefault()
				var viewFormLink = '/viewItem/' + currAlert.properties.formID + '/' + currNotification.recordID
				 win = window.open(viewFormLink,"_blank")
				win.focus()
				
			})
			
			var $alertName = $('<div class="h5 alertHeader"></div>')
			$alertName.text(currAlert.name)	
			$alertLink.append($alertName)
			
			var $summary = $("<div></div>")
			var summaryText = currNotification.caption
			$summary.text(summaryText)
			$alertLink.append($summary)
			
			
			var alertMoment = moment(currNotification.timestamp)
			var alertTime = alertMoment.calendar() + " (" + alertMoment.fromNow() + ")"
			var $alertTime = "<div><small>" + alertTime + "</small></div>"
			$alertLink.append($alertTime)
					
			return $alertListItem
		}
		
		var $alertList = $header.find("ul")
		$alertList.empty()
		var maxNotificationDisplay = 5
		for(var notificationIndex = 0; 
			(notificationIndex < notificationList.notifications.length) && (notificationIndex < maxNotificationDisplay);
			notificationIndex++) {
			$alertList.append(createAlertListItem(notificationIndex))
		}
			
		var $sellAllListItem = $('<li><a>See all alerts</a></li>')
		var seeAllUrl = "/alerts/" + databaseID
		var $seeAllLink = $sellAllListItem.find("a")
		$alertList.append($sellAllListItem)
		
		$seeAllLink.click(function(e) {
			e.preventDefault()
			clearAlertCount()
			if(seeAllAlertsCallback !== undefined) {
				seeAllAlertsCallback()
			}
		})
		
		initAlertCountBadge()
		
	})
}