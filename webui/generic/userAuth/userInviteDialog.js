function openUserInviteDialog() {
	var $inviteDialog = $('#userInviteDialog')
	var $inviteForm = $("#inviteUsersForm")
	
	var $emailInput1 = $('#userInviteEmail1')
	var $emailInput2 = $('#userInviteEmail2')
	var $emailInput3 = $('#userInviteEmail3')
	
	var $inviteAlert = $('#userInviteAlert')
	var $inviteAlertMsg = $('#userInviteAlertMsg')
	var $inviteInfo = $('#userInviteInfo')
	var $inviteConfirm = $('#userInviteConfirm')
	var $inviteButton = $('#sendUserInviteButton')
	
	$inviteAlert.hide()
	$inviteInfo.show()
	$inviteConfirm.hide()
	$inviteButton.prop('disabled',false)
	
	$emailInput1.val("")
	$emailInput2.val("")
	$emailInput3.val("")
	resetFormValidationFeedback($inviteForm)
	
	var validator = $inviteForm.validate({
		rules: {
			userInviteEmail1: {
				minlength: 3,
				required: true,
				email: true,
				remote: {
					url: '/auth/validateNewUserEmail',
					data: {
						emailAddr: function() { return  $emailInput1.val(); }
					}
				} // remote
			},
			userInviteEmail2: {
				email: true,
				remote: {
					url: '/auth/validateNewUserEmail',
					data: {
						emailAddr: function() { return  $emailInput2.val(); }
					}
				} // remote
			},
			userInviteEmail3: {
				email: true,
				remote: {
					url: '/auth/validateNewUserEmail',
					data: {
						emailAddr: function() { return  $emailInput3.val(); }
					}
				} // remote
			},

		},
		messages: {
			userInviteEmail1: {
				required: "Email address is required",
				email: "Please enter a valid email address",
				remote:"Email address already registered."
			},
			userInviteEmail2: {
				email: "Please enter a valid email address",
				remote:"Email address already registered."
			},
			userInviteEmail3: {
				email: "Please enter a valid email address",
				remote:"Email address already registered."
			},
		}
	})
	
	initButtonControlClickHandler($inviteButton,function() {
		
		if($inviteForm.valid()) {
			
			var inviteEmailAddrs = []
			inviteEmailAddrs.push($emailInput1.val())
			if ($emailInput2.val() !== null && $emailInput2.val().length > 0) {
				inviteEmailAddrs.push($emailInput2.val())
			}
			if ($emailInput3.val() !== null && $emailInput3.val().length > 0) {
				inviteEmailAddrs.push($emailInput3.val())
			}
				
			var inviteParams = { emailAddrs: inviteEmailAddrs }
			$inviteButton.prop('disabled',true)
			jsonRequest("/auth/sendUserInvites",inviteParams,function(inviteResp) {
				if(inviteResp.success == true) {
					$inviteConfirm.show()
					$inviteInfo.hide()
					setTimeout(function() {
						$inviteDialog.modal('hide')
					}, 3000);
				} else {
					$inviteAlertMsg.text(inviteResp.msg)
					$inviteAlert.show()
				}
			})
		}
		
	})
	
	$inviteDialog.modal('show')
}