function openNewUserDialog(databaseID) {
	
	var $newUserForm = $('#adminNewUserForm')
	var $newUserDialog = $('#adminNewUserDialog')
	
	$newUserDialog.modal('show')

	initButtonClickHandler('#newUserDialogSaveUserButton',function() {
		console.log("Add new user save button clicked")
		$newUserDialog.modal('hide')
//		jsonAPIRequest("global/addUser",newGlobalParams,function(newUserInfo) {
//			console.log("Add new user: " + JSON.stringify(newUserInfo))
//		})
	})
	
}