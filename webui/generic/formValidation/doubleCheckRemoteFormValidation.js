// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


// There is a limitation with the jQuery Validation plugin that remote validations are not completed when
// validation is triggered programmatically (e.g., after a blur event). One possible workaround is to use
// the 'async:false' option on the remote validation request, but this triggeres a warning from jQuery
// about synchronous instead of asynchronous AJAX calls.
//
// A number of workarounds are discussed on Stack Exchange and the support comments for the plug-in.
// However, a simple work-around which doesn't involve patching the plugin is to simply do an asynchronous
// remote validation at the time the field's value is ready to be saved.
//
// This function is expected to take all the same agruments as the ones given to the "remote" option
// in the jQuery Validation plugin rules.
function doubleCheckRemoteFormValidation(requestUrl, requestData,replyFunc) {
	
    $.ajax({
       url: requestUrl,
       data: requestData,
       error: function() {
		  replyFunc(false)
       },
       success: function(replyData) {	   
		   if (replyData == true) {
		   		replyFunc(true)
		   } else {
		   		replyFunc(false)
		   }		   
       },
       type: 'POST'
    });
	
}
