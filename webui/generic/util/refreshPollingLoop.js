$(document).ready(function() {
	// Special jQuery function to detect when an element is removed from the DOM.
	// Based upon the following: https://stackoverflow.com/questions/2200494/jquery-trigger-event-when-an-element-is-removed-from-the-dom
	(function($){
	  $.event.special.destroyed = {
	    remove: function(o) {
	      if (o.handler) {
	        o.handler()
	      }
	    }
	  }
	})(jQuery)
	
})

function initRefreshPollingLoop($parentContainer, refreshFrequencySecs, refreshCallback) {
    var userActivityTimer;
	
	var userIsActive = true
	var parentContainerAlive = true
	
	$parentContainer.bind("destroyed",function() {
		parentContainerAlive = false
		$(window).off("mousemove",resetUserActivityTimer)
		$(window).off("click",resetUserActivityTimer)
		$(window).off("mousedown",resetUserActivityTimer)
		$(window).off("keypress",resetUserActivityTimer)
		$(window).off("scroll",resetUserActivityTimer)
	})
	
	$(window).mousemove(resetUserActivityTimer)
	$(window).click(resetUserActivityTimer)
	$(window).mousedown(resetUserActivityTimer)
	$(window).keypress(resetUserActivityTimer)
	$(window).scroll(resetUserActivityTimer)

    function setUserInactive() {
		userIsActive = false
    }

    function resetUserActivityTimer() {
        clearTimeout(userActivityTimer);
		userIsActive = true
		
		// If the timer completes before user activity is seen, then 
		// disable the inactivity timer.
        userActivityTimer = setTimeout(setUserInactive, 10000);  // time is in milliseconds
    }
	
	function refresh() {
		if(userIsActive && parentContainerAlive) {
			refreshCallback()
		} 
		if(parentContainerAlive) {
			setTimeout(refresh,refreshFrequencySecs * 1000)	
		}
	}
	refresh()
}
