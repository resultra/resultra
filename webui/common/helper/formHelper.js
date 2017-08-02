

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

function validateNumberFormField(controlSelector) {
	var numberInput = $(controlSelector).val()
	var inputAsNum = Number(numberInput)
	if(isNaN(inputAsNum)) {
		return false
	} else {
		return true
	}
	
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


// A group of radio buttons can use a name prefix to identify what the radio buttons are for, then a
// unique ID value as the names' suffix. For the same commonNamePrefix, this function returns a list
// of IDs mapped to the radio button values.
// 
// This is applicable to "columns of radio buttons groups" used for groups of values, such as setting
// the priviliges to none,view, or edit for each of the different forms.
function getGroupedRadioButtonVals(commonNamePrefix) {
	var radioVals = {}
	
	$("input:radio").each (function() {
		var radioName = $(this).attr('name')
		if(radioName.indexOf(commonNamePrefix) == 0) {
			var idVal = radioName.replace(commonNamePrefix,'')
			var radioSelector = 'input[name="'+radioName+'"]:checked'
			radioVals[idVal] = $(radioSelector).val()			
		}
	})
	
	console.log("Radio vals: " + JSON.stringify(radioVals))
	
	return radioVals
	
}

function nonEmptyStringVal(val) {
	
	if (val === undefined) {
		return false
	}
	if (val === null) {
		return false
	}
	
	
	var valWithoutSpace = val.replace(/\s/g,'')
	if(valWithoutSpace.length > 0) {
		return true
	} else {
		return false
	}
	
}

function stripLeadingAndTrailingSpace(val) {
	var valWithoutLeadingSpace = val.replace(/^\s+/,'')
	var valWithoutTrailingSpace = valWithoutLeadingSpace.replace(/\s+$/,'')
	return valWithoutTrailingSpace
}




