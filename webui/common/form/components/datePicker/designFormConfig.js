
// Definition of parameters and callbacks for a date picker to be editable within the form editor.
// this javascript file needs to included after the other check box related files, so all the functions
// are already defined.

function initDesignFormDatePicker() {
	console.log("Init checkbox design form behavior")
	initNewDatePickerDialog()
}

function selectFormDatePicker(datePickerObjRef) {
	console.log("Selected date picker: " + JSON.stringify(datePickerObjRef))
	loadDatePickerProperties(datePickerObjRef)
}


function resizeDatePicker(datePickerID,geometry) {
	var resizeParams = {
		parentFormID: designFormContext.formID,
		datePickerID: datePickerID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/datePicker/resize", resizeParams, function(updatedObjRef) {
		setElemObjectRef(datePickerID,updatedObjRef)
	})	
}


var datePickerDesignFormConfig = {
	draggableHTMLFunc:	datePickerContainerHTML,
	startPaletteDrag: function(placeholderID,$paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewDatePickerDialog,
	resizeConstraints: elemResizeConstraints(100,640,30,30),
	resizeFunc: resizeDatePicker,
	initFunc: initDesignFormDatePicker,
	selectionFunc: selectFormDatePicker
}
