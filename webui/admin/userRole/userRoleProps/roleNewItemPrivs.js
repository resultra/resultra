function initRoleNewItemPrivs(roleID) {
	
	function roleNewItemPrivilegeListItem(newItemPriv) {
		var privCheckboxHTML = '' +
			'<div class="list-group-item">' +
				'<div class="checkbox">' +
					'<input type="checkbox" id="' + newItemPriv.linkID + '"></input>'+
					'<label for="' + newItemPriv.linkID +  '"><span class="noselect linkNameLabel"></span></label>' +
				'</div>' + 
			'</div>'
		
		var $privCheckbox = $(privCheckboxHTML)
		$privCheckbox.find('.linkNameLabel').text(newItemPriv.linkName)
		
		var $checkboxInput = $privCheckbox.find("input")
		
		initCheckboxControlChangeHandler($checkboxInput,newItemPriv.linkEnabled,function(linkEnabled) {
			
			var params = {
				roleID: roleID,
				linkID: newItemPriv.linkID,
				linkEnabled: linkEnabled
			}			
			jsonAPIRequest("userRole/setNewItemRolePrivs",params,function(setPrivsStatus) {
			})
		})
		
		return $privCheckbox
		
	}
	
	jsonAPIRequest("userRole/getRoleNewItemPrivs", { roleID: roleID }, function(roleNewItemPrivs) {
		
		var $privList = $('#adminNewItemLinkRolesPrivilegesList')
		
		$privList.empty()
		
		for(var privIndex=0; privIndex < roleNewItemPrivs.length; privIndex++) {
			var newItemPriv = roleNewItemPrivs[privIndex]
			$privList.append(roleNewItemPrivilegeListItem(newItemPriv))
		}
	})
	
}