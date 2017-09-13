function initAlertNotificationList(databaseID) {
	
	loadNotificationListInfo(databaseID, function(notificationList,formsByID,fieldsByID) {
			
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
		
		var itemSummaryCol = {
			data: 'itemSummary', 
			defaultContent: '',
			type:'string'
		}
		
		var notificationSummaryCol = {
			data: 'notificationSummary', 
			defaultContent: '',
			type:'string'
		}
		
		var formLinkCol = {
			data: 'formName',
			type:'date',
			createdCell: function( cell, cellData, rowData, rowIndex, colIndex ) {
				var $formLink = $(cell).find('a')
				$formLink.text(rowData.formName())
				
				var viewFormLink = '/viewItem/' + rowData.formID() + '/' + rowData.recordID()
				$formLink .attr("href",viewFormLink)
			},
			render: function(data, type, row, meta) {
				if (type==='display') {
					return '<a href="">Link goes here</a>'
				} else {
					return data
				}
			},			
			defaultContent: ''
			
		}
		
		var dataCols = [timeDataCol,alertNameDataCol,itemSummaryCol,notificationSummaryCol, formLinkCol]
		
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
			
			this.formName = function() {
				var formID = this.alert.properties.formID
				if (formID.length > 0) {
					return formsByID[formID].name
				}
				return ""
			}
			this.formID = function() {
				return this.alert.properties.formID
			}
			
			this.recordID = function() {
				return this.notification.recordID
			}
			
			this.itemSummary = function() {
				return this.notification.itemSummary
			}
			
			this.notificationSummary = function() {
				return this.notification.caption			
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