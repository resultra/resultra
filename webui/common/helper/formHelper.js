

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

// resetFormControlError is the same as removeFormControlError, except the
// help block is immediately removed. This is intended for initialization
// functions, where a transition out of an error state is not wanted.
function resetFormControl(controlParentSelector) {
	$(controlParentSelector).removeClass("has-error")
	$(controlParentSelector).find(".help-block").hide()
	$(controlParentSelector).find(".form-control").val("")
}


function radioButtonIsChecked(radioButtonSelector) {
	return $(radioButtonSelector).prop('checked')
}

function formFieldValueIsNonEmpty(controlSelector) {
	var selectedVal = $(controlSelector).val()
	if((selectedVal != null) && (selectedVal.length > 0))
	{
		return true
	} else {
		return false
	}
	
}


function formFieldValueIsEmpty(controlSelector) {
	var selectedVal = $(controlSelector).val()
	if((selectedVal != null) && (selectedVal.length > 0))
	{
		return false
	} else {
		return true
	}
	
}

function revalidateNonEmptyFormFieldOnChange(controlSelector) {
	$(controlSelector).change(function() { 
		validateNonEmptyFormField(controlSelector) 
	})
}

function validateNonEmptyFormField(controlSelector) {
	if(formFieldValueIsEmpty(controlSelector)) {
		$(controlSelector).parent().addClass("has-error")
		$(controlSelector).siblings(".help-block").slideDown()
		return false
	} else {
		$(controlSelector).parent().removeClass("has-error")
		$(controlSelector).siblings(".help-block").slideUp()
		return true
	}
}


function emptyOptionHTML(prompt) {
	return '<option value="">' + prompt + '</option>'	
}

function selectOptionHTML(selItemVal, selItemText) {
	var selOptionHTML = '<option value="' + selItemVal + '">' + selItemText + '</option>'
	return selOptionHTML
}

function defaultSelectOptionPromptHTML(selItemPrompt) {
	var selOptionHTML = '<option disabled selected value="">' + selItemPrompt + '</option>'
	return selOptionHTML
}
