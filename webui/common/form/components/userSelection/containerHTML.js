
function userSelectionControlFromUserSelectionComponentContainer($userSelContainer) {
	return $userSelContainer.find(".userSelectionCompSelectionControl")
}


function userSelectionInputButtonsContainerHTML() {
		
	var userSelectionDropdownHTML = '' +
		'<div class="userSelectionDropdown">'+
			'<button type="button" class="btn btn-default btn-sm dropdown-toggle clearButton" ' + 
					'data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">' +
					'<span class="caret"></span></button>'+
			'<ul class="dropdown-menu valueDropdownMenu">' +
			'</ul>'+
		'</div>'
	
	var  clearValButtonHTML = '<button type="button" class="btn btn-default btn-sm clearButton userSelectionComponentClearValueButton" >' + 
				'<small><i class="glyphicon glyphicon-remove"></i></small>' +
		'</button>'
	
	return userSelectionDropdownHTML + clearValButtonHTML
}

function userSelectionControlContainerHTML() {
	
	return   '<div class="formUserSelectionControl">' + 
					'<select class="form-control userSelectionCompSelectionControl"></select>' +
				'</div>'
}

function userSelectionContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class=" layoutContainer userSelectionFormContainer">' +
			'<div class="row">' +
				'<div class="col-sm-8">' +
					'<label>New User Selection</label>' + 
				'</div>' +
				'<div class="col-sm-4 userSelectionButtons">' +
					userSelectionInputButtonsContainerHTML() + 
					componentHelpPopupButtonHTML() +
				'</div>' +
			'</div>' +
			'<div class="form-group marginBottom0">'+
				userSelectionControlContainerHTML() +
			'</div>'+
		'</div>';
										
	return containerHTML
}

function userSelectionTablePopupEditorContainerHTML() {
	var containerHTML = ''+
		'<div class=" layoutContainer userSelectionTableCellContainer userSelectionPopupEditorContainer">' +
			'<div class="userEditorHeader">' +
				'<button type="button" class="close closeEditorPopup" data-dismiss="modal" aria-hidden="true">x</button>' +
				userSelectionInputButtonsContainerHTML() +
			'</div>' +
			'<div class="marginTop5">' +
				userSelectionControlContainerHTML() +
			'</div>' +
		'</div>';									
	return containerHTML
	
}


function userSelectionTableCellContainerHTML() {
	return '<div class="layoutContainer userSelectionEditTableCell">' +
			'<div>' +
				'<a class="btn userSelectionEditPopop"></a>'+
			'</div>' +
		'</div>'
}

function setUserSelectionComponentLabel($userSelection,userSelection) {
	var $label = $userSelection.find('label')
	
	setFormComponentLabel($label,userSelection.properties.fieldID,
			userSelection.properties.labelFormat)	
	
}

function configureUserSelectionDropdown(componentContext,$userSelection,userSelection,setValueCallback) {
	
	var $userSelectionDropdown = $userSelection.find(".userSelectionDropdown")
	var userSelFieldID = userSelection.properties.fieldID
	
	function hideDropdownControls() {
		$userSelectionDropdown.css("display","none")
	}
	
	function showDropdownControls() {
		// The jQuery show() method will set the display to "block", which causes the controls to display on a
		// new line.
		$userSelectionDropdown.css("display","")	
	}
	
	var fieldRef = getFieldRef(userSelFieldID)
	if(fieldRef.isCalcField) {
		hideDropdownControls()
		return
	}
	
	if(formComponentIsReadOnly(userSelection.properties.permissions)) {
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
	
	var selectableRoles = userSelection.properties.selectableRoles
	if (userSelection.properties.currUserSelectable || (selectableRoles !== undefined && selectableRoles.length > 0)) {
		showDropdownControls()
		
		getDropdownMenuInfo(function(usersByRole,currUserInfo) {
			var $valDropdownMenu = $userSelection.find('.valueDropdownMenu')
			$valDropdownMenu.empty()
			
			if(userSelection.properties.currUserSelectable) {
				$valDropdownMenu.append(createDropdownMenuItem(currUserInfo,currUserInfo))
			}
			
			var firstRoleUserAppended = false
			
			for (var roleIndex = 0; roleIndex < selectableRoles.length ; roleIndex++) {
				var currSelectableRole = selectableRoles[roleIndex]
				var roleUserInfo = usersByRole[currSelectableRole]
				
	
				if ((roleUserInfo !== null) && (roleUserInfo.roleUsers.length >0)) {
					
					if(userSelection.properties.currUserSelectable && (firstRoleUserAppended===false)) {
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

function initUserSelectionClearValueButton($userSelection,userSelection) {
	
	var $clearValueButton = $userSelection.find(".userSelectionComponentClearValueButton")
	
	var fieldID = userSelection.properties.fieldID
	
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
	
	if(formComponentIsReadOnly(userSelection.properties.permissions)) {
		hideClearValueButton()
	} else {
		if(userSelection.properties.clearValueSupported) {
			showClearValueButton()
		} else {
			hideClearValueButton()
		}
	}
	
}

function initDummiedUpUserSelectionControl($container) {
	// When dragging and dropping the control, a dummied-up select2 control needs to be initialized.
	// This ensures the geometry of the control is similar to what it will look like when fully functional.
	var $userSelectionControl = userSelectionControlFromUserSelectionComponentContainer($container)
	$userSelectionControl.select2({
		placeholder: "Select a collaborator",
		width:'100%'
	})
}



function initUserSelectionFormComponentContainer(componentContext,$container,userSelection) {
	setUserSelectionComponentLabel($container,userSelection)
	initUserSelectionClearValueButton($container,userSelection)
	
	
	function dummySetVal(userID) {}
	configureUserSelectionDropdown(componentContext,$container,userSelection,dummySetVal)
	
	var $userSelectionControl = userSelectionControlFromUserSelectionComponentContainer($container)
	
	var userSelectionParams = {
		$selectionInput: $userSelectionControl,
		databaseID: componentContext.databaseID,
	}
	initCollaboratorUserSelection(userSelectionParams)
	
	
	
	initComponentHelpPopupButton($container, userSelection)	
}