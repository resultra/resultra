// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function loadNotificationListInfo(databaseID, loadDoneCallback) {
	var loadsRemaining = 3
	
	var notificationList
	var formsByID
	var fieldsByID
	
	function oneLoadComplete() {
		loadsRemaining--
		if (loadsRemaining <= 0) {
			loadDoneCallback(notificationList,formsByID,fieldsByID)
		}
	}
	
	var getAlertListParams = { parentDatabaseID: databaseID }
	jsonAPIRequest("alert/getNotificationList",getAlertListParams,function(notListReply) {
		notificationList = notListReply
		oneLoadComplete()
	})

	var getFormsParams = { parentDatabaseID: databaseID }
	jsonAPIRequest("frm/formsByID",getFormsParams,function(formsByIDReply) {
		formsByID = formsByIDReply
		oneLoadComplete()
	})
	
	loadFieldInfo(databaseID,[fieldTypeAll],function(fieldsByIDReply) {
		fieldsByID = fieldsByIDReply
		oneLoadComplete()			
	})		
	
}
