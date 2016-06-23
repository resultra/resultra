// Javascript for sign-in dialog

function openSigninDialog() {
	$('#signinDialogAlert').hide()
	$('#signInModalDialog').modal('show')
	
	$('#signInDialogSigninButton').unbind("click")
	$('#signInDialogSigninButton').click(function(e) {
	    console.log("Sign in dialog button clicked")
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality
		
		var loginParams = {
			emailAddr: $('#signinDialogEmailInput').val(),
			password: $('#signinDialogPasswordInput').val()
		}
		jsonRequest("/auth/login",loginParams,function(loginResp) {
			if(loginResp.success == true) {
				$('#signInModalDialog').modal('hide')
				
				// Reload the current page, while forcing the page to refresh
				// from the server, rather than the cache. This will cause
				// the server to authenticate the current user while serving
				// up the current page.
				location.reload(true)
			} else {
				$('#signinDialogErrorMsg').text(loginResp.msg)
				$('#signinDialogAlert').show()
			}
		})
	});
	
}