function formatTimelineValue(fieldValChange) {
	
	var context = fieldValChange.valueFormat.context
	var format = fieldValChange.valueFormat.format
	var rawVal = fieldValChange.updatedValue
	
	function formatRatingVal() {
		switch(format){
		case "star":
			return "Rating changed to " + rawVal + " stars."
		default:
			return "Rating changed to " + rawVal
		}
	}

	function formatDatePickerVal() {
		switch(format){
		case "date":
			return "Set to " + moment(rawVal).format('MMMM Do, YYYY') + "."
		default:
			return "Set to " + moment(rawVal).calendar()
		}
	}


	function formatImageVal() {
		switch(format){
		case "general":
			var fileName = escapeHTML(rawVal.origName)
			return "Uploaded a new image: " + '<a href="'+rawVal.url+'" target="new">' +fileName +'</a>'
		default:
			return "Uploaded a new image."
		}
	}


	function formatTextBoxVal() {
		switch(format){
		case "general":
			return "Entered " + rawVal + "."
		default:
			return "Entered " + rawVal + "."
		}
	}

	function formatSelectVal() {
		switch(format){
		case "general":
			return "Selected " + rawVal + "."
		default:
			return "Selected " + rawVal + "."
		}
	}


	function formatSelectUserVal() {
		switch(format){
		case "general":
			return "Selected @" + rawVal.userName + "."
		default:
			return "Selected @" + rawVal.userName + "."
		}
	}

	
	function formatCheckboxVal() {
		return "Set to " + rawVal + "."
	}
	
	switch (context) {
	case "rating":
		return formatRatingVal()
	case "checkbox":
		return formatCheckboxVal()
	case "datePicker":
		return formatDatePickerVal()
	case "textBox":
		return formatTextBoxVal()
	case "image":
		return formatImageVal()
	case "select":
		return formatSelectVal()
	case "selectUser":
		return formatSelectUserVal()
	default:
		return "Changed value to " + rawVal + 
					" -- " + JSON.stringify(fieldValChange.valueFormat)
	}
	
}