

function addDashboardLinkToTOCList(dashboardInfo) {
	// TODO - Link to "dashboard view" page instead of dashboard design page (view page isn't implemented yet)
	var dashboardListItemHTML = '<a href="/viewDashboard/' + dashboardInfo.dashboardID 
			+ '" class="list-group-item">' + dashboardInfo.name + '</a>'
	$('#tocDashboardList').append(dashboardListItemHTML)		
}

function addItemListLinkToTOCList(listInfo) {
	var itemListItemHTML = '<a href="/viewList/' + listInfo.listID 
			+ '" class="list-group-item">' + listInfo.name + '</a>'
	$('#tocListList').append(itemListItemHTML)		
	
}

function addFormLinkToTOCList(tocConfig, formInfo) {
	var formLinkHTML = '<div class="row marginTop5">' + 
			'<div class="col-md-12">' +
				'<button type="button" class="btn btn-default btn-sm formLinkButton">' + formInfo.name + '</button>' +
			'</div>' +
		'</div>';
	var $formLinkButton = $(formLinkHTML)
	
	initButtonControlClickHandler($formLinkButton,function() {
		console.log("Form button clicked: " + JSON.stringify(formInfo))
		
		var viewFormContext = {
			databaseID: tocConfig.databaseID,
			formID: formInfo.formID
		}
		
		tocConfig.newItemFormButtonFunc(viewFormContext)
	})
	
	$('#tocFormList').append($formLinkButton)
}


function initDatabaseTOC(tocConfig) {
	
	var getDBInfoParams = { databaseID: tocConfig.databaseID }
	jsonAPIRequest("database/getInfo",getDBInfoParams,function(dbInfo) {
		console.log("Got database info: " + JSON.stringify(dbInfo))
		
		$('#tocFormList').empty()
		for(var formInfoIndex = 0; formInfoIndex < dbInfo.formsInfo.length; formInfoIndex++) {
			var formInfo = dbInfo.formsInfo[formInfoIndex]
			if(formInfo.properties.addNewItemFromForm) {
				addFormLinkToTOCList(tocConfig,formInfo)		
			}
		}
		
		
		$('#tocListList').empty()
		for(var listInfoIndex = 0; listInfoIndex < dbInfo.listsInfo.length; listInfoIndex++) {
			var listInfo = dbInfo.listsInfo[listInfoIndex]
			addItemListLinkToTOCList(listInfo)
		}

		$('#tocDashboardList').empty()
		for (var dashboardInfoIndex = 0; dashboardInfoIndex < dbInfo.dashboardsInfo.length; dashboardInfoIndex++) {
			var dashboardInfo = dbInfo.dashboardsInfo[dashboardInfoIndex]
			addDashboardLinkToTOCList(dashboardInfo)
		}

		
	}) // getRecord
	
}