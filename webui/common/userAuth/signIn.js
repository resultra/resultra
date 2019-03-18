// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initUserSigninComponents() {
	
	$('#signinDialogAlert').hide()
	
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