function initAlertNotificationList(pageLayout, databaseID) {
	
//	pageLayout.disablePropertySidebar()
	pageLayout.setCenterContentHeader("Alerts")		
	
	var $alertLayoutCanvas = $('#alertPageLayoutCanvas')
	$alertLayoutCanvas.empty()
	
	loadNotificationListInfo(databaseID, function(notificationList,formsByID,fieldsByID) {
			
		var $notificationListTable = $('#notificationListTableTemplate').clone()
		$notificationListTable.attr("id","notificationListTable")
		$alertLayoutCanvas.append($notificationListTable)
		
		var timeDataCol = {
			data: 'timestamp',
			type:'date',
			render: function(data, type, row, meta) {
				if (type==='display') {
					var alertMoment = moment(row.timestamp())
					var alertTime = alertMoment.calendar() + 
							" (" + alertMoment.fromNow() + ")"
					return alertTime
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
				
				
				$formLink.click(function(e) {
					e.preventDefault()
					var viewFormLink = '/viewItem/' + rowData.formID() + '/' + rowData.recordID()
					 win = window.open(viewFormLink,"_blank")
					win.focus()
					
				})
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
		
		var dataCols = [alertNameDataCol,notificationSummaryCol,timeDataCol,formLinkCol]
		
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