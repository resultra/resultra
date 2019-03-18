// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
	
	
	function formatCommentBoxVal() {
		switch(format){
		case "general":
			return "Added comment: " + rawVal + "."
		default:
			return "Added comment:  " + rawVal + "."
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
	case "commentBox":
		return formatCommentBoxVal()
	default:
		return "Changed value to " + rawVal + 
					" -- " + JSON.stringify(fieldValChange.valueFormat)
	}
	
}