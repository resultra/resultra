
function initSocialButtonIconProps(params) {
	var $iconSelection = $('#adminSocialButtonComponentIconSelection')
	$iconSelection.val(params.initialIcon)
	initSelectControlChangeHandler($iconSelection,function(newIcon) {
		params.setIcon(newIcon)
	})
	
}
