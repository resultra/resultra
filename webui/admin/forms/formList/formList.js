var adminFormListElemPrefix = "adminFormList_"


function designFormPageHTMLLink(formID) {
	return '/admin/frm/' + formID
}

function adminFormListButtonsHTML(formInfo) {
return '' +
			'<div class="pull-right adminFormListButtons">' + 
				'<a class="btn btn-xs btn-default" href="' + designFormPageHTMLLink(formInfo.formID) + '" role="button">' + 
					'<span class="glyphicon glyphicon-pencil" style="padding-bottom:2px;"></span>' +
				'</a>' +
	'</div>';
}


function addFormToAdminFormList(pageContext,formInfo) {
	 
	var formListFormID = adminFormListElemPrefix + formInfo.formID
	
	var $formListItem = $('#adminFormListItemTemplate').clone()
	$formListItem.attr("id",formListFormID)
	
	$formListItem.find(".adminFormListFormName").text(formInfo.name)
	
	var $editFormButton = $formListItem.find(".adminFormListEditFormButton")
	$editFormButton.click(function(e) {
		e.preventDefault()
		$editFormButton.blur()
		navigateToFormDesignerPageContent(pageContext,formInfo)
	})
	// TODO - initialize button
	
	$('#adminFormList').append($formListItem)		
}

function initAdminFormSettings(pageContext) {
	
	var $adminFormList = $("#adminFormList")
	
    $adminFormList.sortable({
		placeholder: "ui-state-highlight",
		cursor:"move",
		update: function( event, ui ) {
			// Get the new sorted list of form IDs. The prefix needs to be stripped from the ID.
			var prefixRegexp = new RegExp('^' + adminFormListElemPrefix)
			var sortedIDs =  $("#adminFormList").sortable("toArray").map(function(elem) {
				return elem.replace(prefixRegexp,'')
			})
			console.log("New sort order:" + JSON.stringify(sortedIDs))
		}
    });
	
	
	var getDBInfoParams = { databaseID: pageContext.databaseID }
	jsonAPIRequest("database/getInfo",getDBInfoParams,function(dbInfo) {
		console.log("Got database info: " + JSON.stringify(dbInfo))
		
		$adminFormList.empty()
		for (var formInfoIndex = 0; formInfoIndex < dbInfo.formsInfo.length; formInfoIndex++) {
			var formInfo = dbInfo.formsInfo[formInfoIndex]
			addFormToAdminFormList(pageContext,formInfo)
		}
		
	})
	
	
	initButtonClickHandler('#adminNewFormButton',function() {
		console.log("New form button clicked")
		openNewFormDialog(pageContext.databaseID)
	})
	
	
	
}