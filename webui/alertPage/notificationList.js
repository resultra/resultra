function initAlertNotificationList(databaseID) {
	
	var getAlertListParams = {
		parentDatabaseID: databaseID
	}
	
	jsonAPIRequest("alert/getNotificationList",getAlertListParams,function(notificationList) {
		
		var $notificationListTable = $('#notificationListTable')
		
		var timeDataCol = {
			data: 'timestamp',
			type:'date',
			render: function(data, type, row, meta) {
				if (type==='display') {
					return moment(row.timestamp()).format('MM/DD/YYYY')
				} else {
					return data
				}
			},			
			defaultContent: ''
		}
		
		var alertNameDataCol = {
			data: 'alertName', 
			defaultContent: '',
			type:'string'
		}
		
		var dataCols = [timeDataCol,alertNameDataCol]
		
		function AlertDisplayData(rawDataIndex) {
			
			this.rawDataIndex = rawDataIndex
			this.notification = notificationList.notifications[this.rawDataIndex]
			this.alert = notificationList.alertsByID[this.notification.alertID]
			
			this.alertName = function() {
				return this.alert.name
			}
			
			this.timestamp = function() {
				return moment(this.notification.timestamp).toDate()
			}
		}
		
		var displayData = []
		$.each(notificationList.notifications,function(index,notification) {
			displayData.push(new AlertDisplayData(index))
		})
		
		var dataTable = $notificationListTable.DataTable({
			destroy:true, // Destroy existing table before applying the options
			searching:false, // Hide the search box
			paging:true, // pagination must be enabled for pageResize plug-in
			pageResize:true, // enable plug-in for vertical page resizing
			lengthChange:true, // needed for pageResize plug-in
			deferRender:true, // only create elements when required (needed with paging)
			columns:dataCols,
			data: displayData,
			order: [[0,"desc"]] // initially sort by the notification time column
		})
		
	})
	
}