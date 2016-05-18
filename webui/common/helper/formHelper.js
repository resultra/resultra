

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
	return $(radioButtonSelector).prop('checked')
}

function formFieldValueIsNonEmpty(controlSelector) {
	var selectedVal = $(controlSelector).val()
	if(selectedVal.length > 0)
	{
		return true
	} else {
		return false
	}
	
}

function formFieldValueIsEmpty(controlSelector) {
	var selectedVal = $(controlSelector).val()
	if(selectedVal.length > 0)
	{
		return false
	} else {
		return true
	}
	
}