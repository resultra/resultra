function initAlertRecipientProps(alertInfo) {
	
	function roleAlertPrivilegeListItem(rolePriv) {
		
		var privCheckboxHTML = '' +
			'<div class="list-group-item">' +
				'<div class="checkbox">' +
					'<input type="checkbox" id="' + rolePriv.roleID + '"></input>'+
					'<label for="' + rolePriv.roleID +  '"><span class="noselect roleNameLabel"></span></label>' +
				'</div>' + 
			'</div>'
		
		var $privCheckbox = $(privCheckboxHTML)
		$privCheckbox.find('.roleNameLabel').text(rolePriv.roleName)
		
		var $checkboxInput = $privCheckbox.find("input")
		
		initCheckboxControlChangeHandler($checkboxInput,rolePriv.alertEnabled,function(alertEnabled) {
			
			var getPrivParams = {
				roleID: rolePriv.roleID,
				alertID: alertInfo.alertID,
				alertEnabled: alertEnabled
			}			
			jsonAPIRequest("userRole/setAlertRolePrivs",getPrivParams,function(setPrivsStatus) {
			})
		})
		
		return $privCheckbox
		
	}
	
	jsonAPIRequest("userRole/getAlertRolePrivs", { alertID: alertInfo.alertID }, function(roleAlertPrivs) {
		
		var $privList = $('#adminAlertRecipientList')
		
		$privList.empty()
		
		for(var privIndex=0; privIndex < roleAlertPrivs.length; privIndex++) {
			var rolePriv = roleAlertPrivs[privIndex]
			$privList.append(roleAlertPrivilegeListItem(rolePriv))
		}
	})
	
}