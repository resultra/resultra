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
  			'<button class="btn btn-xs btn-danger deleteFormButton">' + 
				// padding-bottom: 2px makes the button image vertically line up better.
				'<span class="glyphicon glyphicon-remove" style="padding-bottom:2px;"></span>' +
			'</button>';

			'</div>'
}


function addFormToAdminFormList(formInfo) {
	 
	var formListFormID = adminFormListElemPrefix + formInfo.formID
	
	var formListItemHTML = '<li class="list-group-item" id="' + formListFormID + '">' + 
		formInfo.name +
		adminFormListButtonsHTML(formInfo) +
	 '</li>'
	
	$('#adminFormList').append(formListItemHTML)		
}

function initAdminFormSettings(databaseID) {
	
    $("#adminFormList").sortable({
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
	
	
	var getDBInfoParams = { databaseID: databaseID }
	jsonAPIRequest("database/getInfo",getDBInfoParams,function(dbInfo) {
		console.log("Got database info: " + JSON.stringify(dbInfo))
		
		$('#adminFormList').empty()
		for (var formInfoIndex = 0; formInfoIndex < dbInfo.formsInfo.length; formInfoIndex++) {
			var formInfo = dbInfo.formsInfo[formInfoIndex]
			addFormToAdminFormList(formInfo)
		}
		
	})
	
	
	initButtonClickHandler('#adminNewFormButton',function() {
		console.log("New form button clicked")
		openNewFormDialog(databaseID)
	})
	
	
	
}