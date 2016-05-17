

function enableFormControl(controlSelector) {
	$(controlSelector).removeAttr('disabled');
}

function disableFormControl(controlSelector) {
	$(controlSelector).attr('disabled','disabled');
}

function addFormControlError(controlParentSelector) {
	$(controlParentSelector).addClass("has-error")
}

function removeFormControlError(controlParentSelector) {
	$(controlParentSelector).removeClass("has-error")
}

function radioButtonIsChecked(radioButtonSelector) {
	$(radioButtonSelector).prop('checked')
}