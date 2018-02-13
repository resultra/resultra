function initRefreshPollingLoop(refreshFrequencySecs, refreshCallback) {
    var userActivityTimer;
	
	var userIsActive = true
	
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
		if(userIsActive) {
			refreshCallback()
		} 
		setTimeout(refresh,refreshFrequencySecs * 1000)
	}
	refresh()
}
