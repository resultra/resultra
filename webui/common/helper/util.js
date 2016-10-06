function jsonRequest(requestUrl, requestData,replyFunc) {
	var jsonReqData = JSON.stringify(requestData)
			
	// TODO - In debug builds, the API logging could be enabled, but disabled in production
	console.log("JSON Request: Sending Request: url = " + requestUrl + " requestData =" + jsonReqData)
	
    $.ajax({
       url: requestUrl,
		contentType : 'application/json',
       data: jsonReqData,
       error: function() {
		  var errMsg = "ERROR: Request failed: url = " + requestUrl + " requestData =" + jsonReqData
		  console.log(errMsg)
          alert(errMsg)
       },
       dataType: 'json',
       success: function(replyData) {
		  console.log("JSON Received Reply: url = " + requestUrl + " replyData =" + JSON.stringify(replyData))
		  replyFunc(replyData)
       },
       type: 'POST'
    });
	
}

function jsonAPIRequest(apiName,requestData, successFunc)
{
	var jsonReqData = JSON.stringify(requestData)
	
	var requestUrl = '/api/' + apiName
	
	jsonRequest(requestUrl,requestData,successFunc)
				
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

function fixedInitiallyHiddenUILayoutPaneAutoSizeToFitContentsParams() {
	return { 
		//Note the size for fixed size panes defaults to "auto", which causes
		// the pane to fit the content
		resizable:false,
		spacing_open:0,
		spacing_closed:0,
		initClosed:true
	}
}

function hidableUILayoutPaneAutoSizeToFitContentsParams() {
	return { 
		// The pane's size will scale to the contents. It can open and close,
		// but is not resizable.
		resizable:false,
	}
}


function assert(condition,message) {
	if(!condition) {
		throw message || "Assertion failed"
	}
}

function IDLookupTable(idList) {
	
	function createIDLookupTable(idList) {	
		var idLookup = {}
		for(var idIndex = 0; idIndex < idList.length; idIndex++) {
			var currID = idList[idIndex]
			idLookup[currID] = true
		}
		return idLookup
	}
	
	this.idLookup = createIDLookupTable(idList)
	this.hasID = function hasID(id) {
		if(this.idLookup.hasOwnProperty(id)) {
			return true
		} else {
			return false
		}
	}
}
