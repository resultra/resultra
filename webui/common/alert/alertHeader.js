// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initAlertHeader(databaseID,seeAllAlertsCallback) {
	
	var $header = $('#alertPageHeaderMenu')
	var $alertCountBadge = $header.find(".badge")
	
	function clearAlertCount() {
		$alertCountBadge.text(0)
		$alertCountBadge.hide()
		
		var advanceAlertParams = { parentDatabaseID: databaseID }
		jsonAPIRequest("alert/advanceNotificationTime",advanceAlertParams,function(reply) {})
		
	}
	
	function reloadAlerts() {
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
					openNewWindowWithElectronOptions(viewFormLink)
				
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
			
			if(notificationList.notifications !== undefined && notificationList.notifications.length > 0) {
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
				
			} else {
				const $noAlertsListItem = $('<li><a>No alerts</a></li>')
				$alertList.append($noAlertsListItem)
			}
			
		
			initAlertCountBadge()
		
		})		
	}
	
	initRefreshPollingLoop($header,10,reloadAlerts)
	
	

}