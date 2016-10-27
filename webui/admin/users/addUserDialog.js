function openNewUserDialog(databaseID) {
	
	var $newUserForm = $('#adminNewUserForm')
	var $newUserDialog = $('#adminNewUserDialog')
	var elemPrefix = "addUser_"
	
	var validator = $newUserForm.validate({
		rules: {
			adminNewUserNameInput: {
				required: true			
			}, // adminNewUserNameInput
		},
		messages: {
			adminNewUserNameInput: {
				required: "User name is required"
			}
		}
	})

	validator.resetForm()
	
	
	function addRoleToRoleCheckboxList(roleInfo) {
		
		var roleCheckbox = createIDWithSelector(elemPrefix + roleInfo.roleID)
	
		var roleCheckboxHTML = '' +
			'<div class="checkbox list-group-item addRoleCheckboxListItem">' +
				'<label>' +
					'<input type="checkbox" id="' + roleCheckbox.id + '"></input>'+
					'<span class="noselect">' + roleInfo.roleName + '</span>' +
				'</label>' +
			'</div>'
		
		var $roleCheckbox = $(roleCheckboxHTML)
		
	
		$('#adminNewUserRoleList').append($roleCheckbox)	
		
		$(roleCheckbox.selector).data("roleInfo",roleInfo)
	
	}
	
	function getRoleListSelectedRoleIDs() {
	
		var selectedRoleIDs = []
	
		// TODO - Is this selector too generic?
		var checkboxSelector = "#adminNewUserRoleList input[type=checkbox]:checked"
	
		$(checkboxSelector).each(function() {
			var roleInfo = $(this).data("roleInfo")
			console.log("Selected role: " + JSON.stringify(roleInfo))
			selectedRoleIDs.push(roleInfo.roleID)
		});
	
		return selectedRoleIDs
	
	}
	
	var $userNameInput = $('#adminNewUserNameInput')
		
	$userNameInput.select2({
		placeholder: "Select a user",
		dropdownParent: $newUserDialog,
		minimumInputLength: 2,
		width:'250px',
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
	
	

	initButtonClickHandler('#newUserDialogSaveUserButton',function() {
		console.log("Add new user save button clicked")
		if($newUserForm.valid()) {
			var selectedUserID = $userNameInput.val()
			var selectedRoleIDs = getRoleListSelectedRoleIDs()
			
			var addUserParams = {
				userID: selectedUserID,
				roles: selectedRoleIDs
			}
			console.log("Adding new user: " + JSON.stringify(addUserParams))
			$newUserDialog.modal('hide')
		}
//		jsonAPIRequest("global/addUser",newGlobalParams,function(newUserInfo) {
//			console.log("Add new user: " + JSON.stringify(newUserInfo))
//		})
	})
	
	
	var dbRolesParams = {
		databaseID: databaseID
	}
	jsonAPIRequest("userRole/getDatabaseRoles",dbRolesParams,function(rolesInfo) {
		
		console.log("Got roles info: " + JSON.stringify(rolesInfo))
		$('#adminNewUserRoleList').empty()
		for(var roleInfoIndex = 0; roleInfoIndex<rolesInfo.length; roleInfoIndex++) {
			var currRoleInfo = rolesInfo[roleInfoIndex]
			addRoleToRoleCheckboxList(currRoleInfo)
		}
		
		$newUserDialog.modal('show')
	})
	
	
}