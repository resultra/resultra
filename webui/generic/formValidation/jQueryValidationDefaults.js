function resetFormValidationFeedback($form) {
	$form.find(".form-group").removeClass("has-feedback has-success has-error has-success-feedback ")
	$form.find(".form-control-feedback").remove()
}

$(document).ready(function() {	
	$.validator.setDefaults({
        highlight: function(element) {
            $(element).closest('.form-group').removeClass('has-success has-feedback').addClass('has-error has-feedback');
			$(element).closest('.form-group').find('i.glyphicon').remove();
			$(element).closest('.form-group').append('<i class="form-control-feedback glyphicon glyphicon-exclamation-sign"></i>');
        },
        unhighlight: function(element) {
            $(element).closest('.form-group').removeClass('has-error  has-feedback').addClass('has-success  has-feedback');
			$(element).closest('.form-group').find('i.glyphicon').remove();
			$(element).closest('.form-group').append('<i class="form-control-feedback glyphicon glyphicon-ok-sign"></i>');
        },
        errorElement: 'span',
        errorClass: 'help-block',
        errorPlacement: function(error, element) {
            if(element.parent('.input-group').length) {
                error.insertAfter(element.parent());
            } else {
                error.insertAfter(element);
            }
		}
		
	})
	
	jQuery.validator.addMethod("itemName", function(value, element) {
	  // allow any non-whitespace characters as the host part
	  var itemName = /^[a-zA-Z0-9][a-zA-Z0-9 \'\.\-]{2,24}$/.test(value)	
	  var allPunc = /^[ \'\.\-]+$/.test(value)
		
	  return itemName && (!allPunc)
	}, 'Please enter a valid name');
	
   // add the rule here
    $.validator.addMethod("optionSelectionRequired", function(value, element, arg){
		console.log("validating option: " + value)
		return (value != null) && (value.length > 0);
    }, "Please select a {0}");

	$.validator.addMethod('positiveNumber',function (value) { 
		console.log("validating positive number: " + value)
        return Number(value) > 0;
    }, 'Enter a positive number.');	
	
	$.validator.addMethod('greaterThan', function(value, element, otherValSelector) { 
		console.log("validating positive number: " + value)
		var $otherVal = $(otherValSelector)
		
		var currVal = Number(value)
		var otherVal = Number($otherVal.val())
		
		if (isNaN(currVal) || isNaN(otherVal)) { 
			return false
		} else if(currVal <= otherVal) {
			return false
		} else {
			return true
		}
    }, 'Value must be greater.');	
	
	
	$.validator.addMethod('floatNumber',function (value) { 
		console.log("validating floating point number: " + value)
		if(isNaN(value)) {
			console.log("validating floating point number: false")
			return false
		} else {
			console.log("validating floating point number: true")
			return true
		}
    }, 'Enter a number.');	
	
	
})


