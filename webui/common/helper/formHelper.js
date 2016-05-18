

function enableFormControl(controlSelector) {
	$(controlSelector).removeAttr('disabled');
}

function disableFormControl(controlSelector) {
	$(controlSelector).attr('disabled','disabled');
}

function addFormControlError(controlParentSelector) {
	$(controlParentSelector).addClass("has-error")
	$(controlParentSelector).find(".help-block").slideDown()
}

function removeFormControlError(controlParentSelector) {
	$(controlParentSelector).removeClass("has-error")
	$(controlParentSelector).find(".help-block").slideUp()
}

function radioButtonIsChecked(radioButtonSelector) {
	$(radioButtonSelector).prop('checked')
}