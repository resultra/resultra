// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function userTagControlFromUserTagComponentContainer($userSelContainer) {
	return $userSelContainer.find(".userTagCompSelectionControl")
}


function userTagInputButtonsContainerHTML() {
		
	var userTagDropdownHTML = '' +
		'<div class="userTagDropdown">'+
			'<button type="button" class="btn btn-default btn-sm dropdown-toggle clearButton" ' + 
					'data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">' +
					'<span class="caret"></span></button>'+
			'<ul class="dropdown-menu valueDropdownMenu">' +
			'</ul>'+
		'</div>'
	
	var  clearValButtonHTML = '<button type="button" class="btn btn-default btn-sm clearButton userTagComponentClearValueButton" >' + 
				'<small><i class="glyphicon glyphicon-remove"></i></small>' +
		'</button>'
	
	return userTagDropdownHTML + clearValButtonHTML
}

function userTagControlContainerHTML() {
	
	return   '<div class="formUserTagControl">' + 
					'<select class="form-control userTagCompSelectionControl" multiple></select>' +
				'</div>'
}

function userTagContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class=" layoutContainer userTagFormContainer">' +
			'<div class="row">' +
				'<div class="col-sm-8">' +
					'<label>New User Selection</label>' + 
				'</div>' +
				'<div class="col-sm-4 userTagButtons">' +
					userTagInputButtonsContainerHTML() + 
					componentHelpPopupButtonHTML() +
				'</div>' +
			'</div>' +
			'<div class="form-group marginBottom0">'+
				userTagControlContainerHTML() +
			'</div>'+
		'</div>';
										
	return containerHTML
}

function userTagTablePopupEditorContainerHTML() {
	var containerHTML = ''+
		'<div class=" layoutContainer userTagTableCellContainer userTagPopupEditorContainer">' +
			'<div class="userEditorHeader">' +
				'<button type="button" class="close closeEditorPopup" data-dismiss="modal" aria-hidden="true">x</button>' +
				userTagInputButtonsContainerHTML() +
			'</div>' +
			'<div class="marginTop5">' +
				userTagControlContainerHTML() +
			'</div>' +
		'</div>';									
	return containerHTML
	
}


function userTagTableCellContainerHTML() {
	return '<div class="layoutContainer userTagEditTableCell">' +
			'<div>' +
				'<a class="btn userTagEditPopop"></a>'+
			'</div>' +
		'</div>'
}

function setUserTagComponentLabel($userTag,userTag) {
	var $label = $userTag.find('label')
	
	setFormComponentLabel($label,userTag.properties.fieldID,
			userTag.properties.labelFormat)	
	
}

function configureUserTagDropdown(componentContext,$userTag,userTag,setValueCallback) {
	
	var $userTagDropdown = $userTag.find(".userTagDropdown")
	var userSelFieldID = userTag.properties.fieldID
	
	function hideDropdownControls() {
		$userTagDropdown.css("display","none")
	}
	
	function showDropdownControls() {
		// The jQuery show() method will set the display to "block", which causes the controls to display on a
		// new line.
		$userTagDropdown.css("display","")	
	}
	
	var fieldRef = getFieldRef(userSelFieldID)
	if(fieldRef.isCalcField) {
		hideDropdownControls()
		return
	}
	
	if(formComponentIsReadOnly(userTag.properties.permissions)) {
		hideDropdownControls()
		return		
	}
	
	function createDropdownMenuItem(currUserInfo,userInfo) {
		var $menuItem = $('<li><a href="#"></a></li>')
		var $menuLink = $menuItem.find('a')
		$menuLink.click(function(e) {
			console.log("User selected: " + JSON.stringify(userInfo))
			setValueCallback(userInfo)
			e.preventDefault()
		})
		var userNameDisplay = '@' + userInfo.userName
		if (currUserInfo.userID === userInfo.userID) {
			userNameDisplay = userNameDisplay + ' (me)'
		}
		$menuItem.find('a').text(userNameDisplay)
		return $menuItem
	}
	
	function getDropdownMenuInfo(infoCallback) {
		var callsRemaining = 2
		
		var usersByRole
		var currUserInfo
		
		function processInfo() {
			callsRemaining--
			if(callsRemaining <= 0) {
				infoCallback(usersByRole,currUserInfo)
			}
		}
		
		var getUsersByRoleParams = { databaseID: componentContext.databaseID }
		jsonAPIRequest("userRole/getUsersByRole",getUsersByRoleParams,function(usersByRoleResp) {
			usersByRole = usersByRoleResp
			processInfo()
		})
		
		var getUserInfoParams = {}
		jsonRequest("/auth/getCurrentUserInfo",getUserInfoParams,function(userInfoResp) {
			currUserInfo = userInfoResp
			processInfo()
		})
		
	}
	
	var selectableRoles = userTag.properties.selectableRoles
	if (userTag.properties.currUserSelectable || (selectableRoles !== undefined && selectableRoles.length > 0)) {
		showDropdownControls()
		
		getDropdownMenuInfo(function(usersByRole,currUserInfo) {
			var $valDropdownMenu = $userTag.find('.valueDropdownMenu')
			$valDropdownMenu.empty()
			
			if(userTag.properties.currUserSelectable) {
				$valDropdownMenu.append(createDropdownMenuItem(currUserInfo,currUserInfo))
			}
			
			var firstRoleUserAppended = false
			
			for (var roleIndex = 0; roleIndex < selectableRoles.length ; roleIndex++) {
				var currSelectableRole = selectableRoles[roleIndex]
				var roleUserInfo = usersByRole[currSelectableRole]
				
	
				if ((roleUserInfo !== null) && (roleUserInfo.roleUsers.length >0)) {
					
					if(userTag.properties.currUserSelectable && (firstRoleUserAppended===false)) {
						$valDropdownMenu.append('<li role="separator" class="divider"></li>')						
					}
					firstRoleUserAppended = true
					
					
					var $header = $('<li class="dropdown-header"></li>')
					$header.text(roleUserInfo.roleName)
					$valDropdownMenu.append($header)						

					for(var userInfoIndex=0; userInfoIndex<roleUserInfo.roleUsers.length;userInfoIndex++) {
						var currRoleUser = roleUserInfo.roleUsers[userInfoIndex]
						$valDropdownMenu.append(createDropdownMenuItem(currUserInfo,currRoleUser))
					}		
					
				}
			}
			
		})
		
		var getUsersByRoleParams = { databaseID: componentContext.databaseID }
		jsonAPIRequest("userRole/getUsersByRole",getUsersByRoleParams,function(usersByRole) {
			
		})
	} else {
		hideDropdownControls()
		return		
	}
	
	showDropdownControls()
	
	
}

function initUserTagClearValueButton($userTag,userTag) {
	
	var $clearValueButton = $userTag.find(".userTagComponentClearValueButton")
	
	var fieldID = userTag.properties.fieldID
	
	function hideClearValueButton() {
		$clearValueButton.css("display","none")
	}
	
	function showClearValueButton() {
		$clearValueButton.css("display","")
	}
	
	
	var fieldRef = getFieldRef(fieldID)
	if(fieldRef.isCalcField) {
		hideClearValueButton()
		return
	}
	
	if(formComponentIsReadOnly(userTag.properties.permissions)) {
		hideClearValueButton()
	} else {
		if(userTag.properties.clearValueSupported) {
			showClearValueButton()
		} else {
			hideClearValueButton()
		}
	}
	
}

function initDummiedUpUserTagControl($container) {
	// When dragging and dropping the control, a dummied-up select2 control needs to be initialized.
	// This ensures the geometry of the control is similar to what it will look like when fully functional.
	var $userTagControl = userTagControlFromUserTagComponentContainer($container)
	$userTagControl.select2({
		placeholder: "Select a collaborator",
		width:'100%'
	})
}



function initUserTagFormComponentContainer(componentContext,$container,userTag) {
	setUserTagComponentLabel($container,userTag)
	initUserTagClearValueButton($container,userTag)
	
	
	function dummySetVal(userID) {}
	configureUserTagDropdown(componentContext,$container,userTag,dummySetVal)
	
	var $userTagControl = userTagControlFromUserTagComponentContainer($container)
	
	var userTagParams = {
		$selectionInput: $userTagControl,
		databaseID: componentContext.databaseID,
	}
	initCollaboratorUserSelection(userTagParams)
	
	
	
	initComponentHelpPopupButton($container, userTag)	
}