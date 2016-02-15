function jsonAPIRequest(apiName,requestData, successFunc)
{
	var jsonReqData = JSON.stringify(requestData)
			
	// TODO - In debug builds, the API logging could be enabled, but disabled in production
	console.log("JSON API Request: api name = " + apiName + " requestData =" + jsonReqData)
	
    $.ajax({
       url: '/api/'+ apiName,
		contentType : 'application/json',
       data: jsonReqData,
       error: function() {
		  var errMsg = "ERROR: API Request failed: api name = " + apiName + " requestData =" + jsonReqData
		  console.log(errMsg)
          alert(errMsg)
       },
       dataType: 'json',
       success: function(replyData) {
		  console.log("JSON API Request succeeded: api name = " + apiName + " replyData =" + JSON.stringify(replyData))
		  successFunc(replyData)
       },
       type: 'POST'
    });
	
}


// A placeholderID is a temporary ID to assign to a div. After saving a 
// new object via JSON call, it is replaced with a unique ID created by the server.
var placeholderNum = 1
function allocNextPlaceholderID()
{
	placeholderID = "placeholderContainerID" + placeholderNum.toString()
	placeholderNum = placeholderNum + 1
	return placeholderID
}

// Parameters to setup a jQuery UI Layout pane without a resize bar and of a fixed size.
function fixedUILayoutPaneParams(paneSize) {
	return { 
		size: paneSize,
		resizable:false,
		spacing_open:0,
		spacing_closed:0
	}
}

function fixedUILayoutPaneAutoSizeToFitContentsParams() {
	return { 
		//Note the size for fixed size panes defaults to "auto", which causes
		// the pane to fit the content
		resizable:false,
		spacing_open:0,
		spacing_closed:0
	}
}
