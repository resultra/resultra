// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function openNewUserDialog(pageContext) {
	
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
			'<div class="list-group-item">' +
				'<div class="checkbox">' +
					'<input type="checkbox" id="' + roleCheckbox.id + '"></input>'+
					'<label for="'+ roleCheckbox.id +'"><span class="noselect"></span></label>' +
				'</div>' +
			'</div>'
		
		var $roleCheckbox = $(roleCheckboxHTML)
		$roleCheckbox.find("span").text(roleInfo.roleName)
		
	
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
	
	var userSelectionParams = {
		selectionInput: $userNameInput,
		dropdownParent: $newUserDialog
	}
	
	initUserSelection(userSelectionParams)
			
	initButtonClickHandler('#newUserDialogSaveUserButton',function() {
		console.log("Add new user save button clicked")
		if($newUserForm.valid()) {
			var selectedUserID = $userNameInput.val()
			var selectedRoleIDs = getRoleListSelectedRoleIDs()
			
			var addCollabParams = {
				databaseID: pageContext.databaseID,
				userID: selectedUserID,
				roleIDs: selectedRoleIDs
			}
			console.log("Adding new collaborator: " + JSON.stringify(addCollabParams))
			jsonAPIRequest("admin/addCollaborator",addCollabParams,function(collabUserRoleInfo) {
					console.log("Added new collaborator: " + JSON.stringify(collabUserRoleInfo))
				$('#userListTableBody').append(userListTableRow(pageContext,collabUserRoleInfo))
				$newUserDialog.modal('hide')				
				
			})
			$newUserDialog.modal('hide')
		}
	})
	
	
	var dbRolesParams = {
		databaseID: pageContext.databaseID
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