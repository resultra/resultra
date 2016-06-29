function openRegisterUserDialog() {
	$('#registerDialogAlert').hide()
	
	$('#registerDialogRegisterButton').unbind("click")
	$('#registerDialogRegisterButton').click(function(e) {
	    console.log("Register button clicked")
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality
		
		if($('#registerDialogPasswordInput').val() != $('#registerDialogRepeatPasswordInput').val()) {
			$('#registerDialogAlert').text("Passwords don't match. Please re-enter a password.")
			$('#registerDialogRepeatPasswordInput').val("")
			$('#registerDialogPasswordInput').val("")
			$('#registerDialogAlert').show()
			return
		}
		
		var loginParams = {
			firstName: $('#registerDialogFirstNameInput').val(),
			lastName: $('#registerDialogLastNameInput').val(),
			userName: $('#registerDialogUserNameInput').val(),
			emailAddr: $('#registerDialogEmailInput').val(),
			password: $('#registerDialogPasswordInput').val(),
		}
		jsonRequest("/auth/register",loginParams,function(registerResp) {
			if(registerResp.success == true) {
				$('#registerModalDialog').modal('hide')
				
			} else {
				$('#registerDialogAlert').text(registerResp.msg)
				$('#registerDialogAlert').show()
			}
		})
	});
	
	$('#registerModalDialog').modal('show')
}