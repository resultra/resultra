function jsonRequest(requestUrl, requestData,replyFunc) {
	var jsonReqData = JSON.stringify(requestData)
	
	function formatUserErrorMsg(jqXHR, exception) {
	    if (jqXHR.status === 0) {
	         return ('Not connected.\nPlease verify your network connection.');
	     } else if (jqXHR.status == 404) {
	         return ('The requested page was not found. [404]');
	     } else if (jqXHR.status == 500) {
	         return ('Internal Server Error [500].');
	     } else if (exception === 'parsererror') {
	         return ('Requested JSON parse failed.');
	     } else if (exception === 'timeout') {
	         return ('Time out error.');
	     } else if (exception === 'abort') {
	         return ('request aborted.');
	     } else {
	         return ('Uncaught Error.\n' + jqXHR.responseText);
	     }
	}
			
	// TODO - In debug builds, the API logging could be enabled, but disabled in production
	console.log("JSON Request: Sending Request: url = " + requestUrl + " requestData =" + jsonReqData)
	
    $.ajax({
       url: requestUrl,
		contentType : 'application/json',
       data: jsonReqData,
       error: function(xhr, status, error) {
		  var userErrMsg = formatUserErrorMsg(xhr,status)

		  alert(userErrMsg)

		  var logMsg = "ERROR: Request failed: url = " + requestUrl + " requestData =" + jsonReqData +
		  		" user error msg = " + userErrMsg
		  console.log(logMsg)
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
	
	this.removeID = function(id) {
		delete this.idLookup[id]
	}
	
	this.getIDList = function() {
		var idList = []
		for(var currID in this.idLookup) {
			idList.push(currID) 
		}
		return idList
	}
	
}

function navigateToURL(url) {
	window.location.href = url
}


function convertStringToNumber(numberStr) {
	var strForConv = numberStr + '' // make sure value passed in is already a string
	if (/^\s*$/.test(strForConv)) {
		return null
	}
	else {
		var numberVal = Number(strForConv)
		if (!isNaN(numberVal)) {
			return numberVal
		} else {
			return null
		}
	} 
}
