function initRoleAlertPrivs(roleID) {
	
	function roleAlertPrivilegeListItem(alertPriv) {
		var privCheckboxHTML = '' +
			'<div class="list-group-item">' +
				'<div class="checkbox">' +
					'<input type="checkbox" id="' + alertPriv.alertID + '"></input>'+
					'<label for="' + alertPriv.alertID +  '"><span class="noselect alertNameLabel"></span></label>' +
				'</div>' + 
			'</div>'
		
		var $privCheckbox = $(privCheckboxHTML)
		$privCheckbox.find('.alertNameLabel').text(alertPriv.alertName)
		
		var $checkboxInput = $privCheckbox.find("input")
		
		initCheckboxControlChangeHandler($checkboxInput,alertPriv.alertEnabled,function(alertEnabled) {
			
			var params = {
				roleID: roleID,
				alertID: alertPriv.alertID,
				alertEnabled: alertEnabled
			}			
			jsonAPIRequest("userRole/setAlertRolePrivs",params,function(setPrivsStatus) {
			})
		})
		
		return $privCheckbox
		
	}
	
	jsonAPIRequest("userRole/getAlertRolePrivs", { roleID: roleID }, function(roleAlertPrivs) {
		
		var $privList = $('#adminAlertRolesPrivilegesList')
		
		$privList.empty()
		
		for(var privIndex=0; privIndex < roleAlertPrivs.length; privIndex++) {
			var alertPriv = roleAlertPrivs[privIndex]
			$privList.append(roleAlertPrivilegeListItem(alertPriv))
		}
	})
	
}