// Setup Boostrap to attach dropdown menus to the body instead of the local context.
// See - http://stackoverflow.com/questions/14473769/bootstrap-drop-down-cutting-off
(function () {
    // hold onto the drop down menu                                             
    var dropdownMenu;
	var isBodyAttachDropdownMenu;
	

    // and when you show it, move it to the body                                     
    $(window).on('show.bs.dropdown', function (e) {

        // grab the menu        
        dropdownMenu = $(e.target).find('.dropdown-menu');
	
		// Only attach the dropdown to the body if the option is explicitely
		// turned on for the given dropdown menu.
		// Using this option introduces conflicts with other dropdown menu
		// options, notably dropdown-menu-right.
		isBodyAttachDropdownMenu = dropdownMenu.hasClass("dropdown-menu-attach-body")
		
		if (isBodyAttachDropdownMenu) {
	        // detach it and append it to the body
	        $('body').append(dropdownMenu.detach());

	        // grab the new offset position
	        var eOffset = $(e.target).offset();

	        // make sure to place it where it would normally go (this could be improved)
	        dropdownMenu.css({
	            'display': 'block',
	                'top': eOffset.top + $(e.target).outerHeight(),
	                'left': eOffset.left
	        });
		}

    });
	

    // and when you hide it, reattach the drop down, and hide it normally                                                   
    $(window).on('hide.bs.dropdown', function (e) {
		if (isBodyAttachDropdownMenu) {
	        $(e.target).append(dropdownMenu.detach());
	        dropdownMenu.hide();
		}
		
    });
	
	
})();