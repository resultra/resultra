function setUserSelectionControlVal($userSelectionControl,userIDs) {
	
	// Setting the value for a select2 selection menu involves putting an
	// option inside the select element then setting the value to the value
	// of this option.
	var getUserInfoParams = { userIDs: userIDs }
	jsonRequest("/auth/getUsersInfo",getUserInfoParams,function(usersInfo) {
		
		$userSelectionControl.empty()
		for (var userIndex = 0; userIndex < usersInfo.length; userIndex++) {
			var userInfo = usersInfo[userIndex]
			var userLabel = "@" + userInfo.userName
			$userSelectionControl.append('<option value="'+userInfo.userID+'">'+userLabel+'</option')
		}
		
		$userSelectionControl.val(userIDs)
	})
	
}

function clearUserSelectionControlVal($userSelectionControl) {
	$userSelectionControl.empty()
	$userSelectionControl.val("")
}


function initUserSelection(selectionParams) {
	
	var configParams = {
		width: '250px'
	}
	$.extend(configParams,selectionParams)
	
	configParams.selectionInput.select2({
		placeholder: "", // TODO - Allow a property to configure the placeholder.
		dropdownParent: configParams.dropdownParent,
		minimumInputLength: 2,
		width: configParams.width,
		ajax: {
			dataType: 'json',
			url: '/auth/searchUsers',
			delay: 250,
			data: function (params) {
				var queryParams = {
				  searchTerm: params.term, // search term
				  page: params.page
				}
	      	  return queryParams
		  	},
			processResults: function (data, params) {
			      // parse the results into the format expected by Select2
			      // since we are using custom formatting functions we do not need to
			      // alter the remote JSON data, except to indicate that infinite
			      // scrolling can be used
			      params.page = params.page || 1;
				  
				  var select2results = []
				  for(var matchIndex = 0; matchIndex < data.matchedUserInfo.length; matchIndex++) {
					  var currMatch = data.matchedUserInfo[matchIndex]
					  var select2result = {
						  id:currMatch.userID,
						  text:'@'+currMatch.userName
					  }
					  select2results.push(select2result)
				  }

			      return {
			        results: select2results,
			        pagination: {
			          more: (params.page * 30) < data.matchedUserInfo.length
			        }
			      };
			  },
			cache: true	
		}
	});
	
	
	
}


function initCollaboratorUserSelection(params) {
	
	var configParams = {
		width: '250px'
	}
	$.extend(configParams,params)
	
	
	var getCollaboratorsParams = {
		databaseID: params.databaseID
	}
	
	jsonAPIRequest("admin/getAllCollaboratorInfo",getCollaboratorsParams,function(collabUserInfo) {
		
		params.$selectionInput.empty()
		
		$.each(collabUserInfo,function(index,userInfo) {
			var newOption = new Option('@'+userInfo.userID, userInfo.userID);
			params.$selectionInput.append(newOption)
		})
		
		params.$selectionInput.select2({
			placeholder: "Select a collaborator", // TODO - Allow a property to configure the placeholder.
			width: params.width,
		});
		
	})
	
	
	
	
}


