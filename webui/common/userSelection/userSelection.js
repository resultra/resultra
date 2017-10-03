
function userSelectionNameDisplay(userInfo) {
	var userNameDisplay = userInfo.firstName + ' ' + userInfo.lastName + ' (@' + userInfo.userName + ')'
	return userNameDisplay
}

function createUserSelectionOption(userInfo) {
	var newOption = new Option(userSelectionNameDisplay(userInfo), userInfo.userID,false,false);
	return newOption
}

function addUserInfoSelectionOptionIfNotExists($userSelectionControl,userInfo) {	
	function optionExists() {
		
		var optionExists = false
		$userSelectionControl.find('option').each(function() {
			var optionUserVal = $(this).val()
			if (optionUserVal === userInfo.userID) {
				optionExists = true
			}
		})
		return optionExists
	}
	
	if(!optionExists()) {
		$userSelectionControl.append(createUserSelectionOption(userInfo))
	}
	
}


function setMultipleUserSelectionControlVal($userSelectionControl,userIDs) {
	
	// Setting the value for a select2 selection menu involves putting an
	// option inside the select element then setting the value to the value
	// of this option.
	var getUserInfoParams = { userIDs: userIDs }
	jsonRequest("/auth/getUsersInfo",getUserInfoParams,function(usersInfo) {
		
		$userSelectionControl.empty()
		for (var userIndex = 0; userIndex < usersInfo.length; userIndex++) {
			var userInfo = usersInfo[userIndex]
			addUserInfoSelectionOptionIfNotExists($userSelectionControl,userInfo)
		}	
		$userSelectionControl.val(userIDs).trigger("change")
	})
	
}

function setSingleUserSelectionControlVal($userSelectionControl,userID) {
	
	// Setting the value for a select2 selection menu involves putting an
	// option inside the select element then setting the value to the value
	// of this option.
	var getUserInfoParams = { userID: userID }
	jsonRequest("/auth/getUserInfo",getUserInfoParams,function(userInfo) {
		addUserInfoSelectionOptionIfNotExists($userSelectionControl,userInfo)
		$userSelectionControl.val(userID).trigger("change")
	})
	
}


function clearUserSelectionControlVal($userSelectionControl) {
	$userSelectionControl.val('').trigger("change")
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
		width: '100%'
	}
	$.extend(configParams,params)
	
	
	var getCollaboratorsParams = {
		databaseID: params.databaseID
	}
	
	params.$selectionInput.select2({
		placeholder: "Select a collaborator", // TODO - Allow a property to configure the placeholder.
		width: configParams.width,
//		data:selectionOptions
	});
	
	
	jsonAPIRequest("admin/getAllCollaboratorInfo",getCollaboratorsParams,function(collabUserInfo) {
					
		var selectionOptions = []
		
		var emptyOption = { id:'', text:'' }
		selectionOptions.push(emptyOption)
		
		$.each(collabUserInfo,function(index,userInfo) {
			// TODO - The following code initializes the control asynchronously with the value
			// being set by loading the record. This causes an incorrect value to be displayed.
			// Either the values need to be retrieved dynamically, or somehow populated in a way
			// which doesn't change the current value.
			var currOption = {
				id: userInfo.userID,
				text: userSelectionNameDisplay(userInfo)
			}
			selectionOptions.push(currOption) 
	//		addUserInfoSelectionOptionIfNotExists(params.$selectionInput,userInfo)
		})
		
//		params.$selectionInput.select2('data',selectionOptions)
				
		
	})
	
	
	
	
}


