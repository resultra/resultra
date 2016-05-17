

function enableFormControl(controlSelector) {
	$(controlSelector).removeAttr('disabled');
}

function disableFormControl(controlSelector) {
	$(controlSelector).attr('disabled','disabled');
	
}