function ItemListContentLayout() {
	
	var contentLayout = $('#listViewContentLayout').layout({
		onresize_end: function(pane, $pane, paneState, paneOptions) {
			if(pane === 'center'){
				// only propagate the resize event for the center/content pane
				console.log("resize triggered")
			//	resizeMainWindow()
			}
		},
		north: fixedUILayoutPaneAutoSizeToFitContentsParams(),
		south: {
			size: 40,
			resizable:false,
			slidable: false,
			spacing_open:0,
			spacing_closed:0,
			initClosed:false,
			fxName:"none"	
		},
		north__showOverflowOnHover:	true,
		south__showOverflowOnHover:	true
	})
	
	function hideFooterLayout() {
		contentLayout.hide("south")
	}
	
	function showFooterLayout() {
		contentLayout.show("south")
	}
	
	
	var $refreshButton = $('#listViewRefreshButton')
	
	function disableRefreshButton() {
		$refreshButton.hide()
	}
	
	function enableRefreshButton(refreshCallback) {
		initButtonControlClickHandler($refreshButton,refreshCallback)
		$refreshButton.show()
	}
	
	function setCenterContentHeader(header) {
		var $header = $('#itemListCenterContentHeader')
		$header.text(header)
	}
	
	initButtonClickHandler("#viewTableOfContentsMenuButton", function() {
		console.log("TOC button clicked")
		theMainWindowLayout.toggleLHSSidebar()
	})
	
	
	
	initButtonClickHandler("#viewListOptionsButton", function() {
		var $iconSpan = $('#viewListOptionsButton > span')
		
		if(theMainWindowLayout.rhsSidebarIsOpen()) {
			$iconSpan.removeClass("fa-toggle-right")
			$iconSpan.addClass("fa-toggle-left")
			theMainWindowLayout.closeRHSSidebar()
		} else {
			$iconSpan.removeClass("fa-toggle-left")
			$iconSpan.addClass("fa-toggle-right")
			theMainWindowLayout.openRHSSidebar()	
		}
	})
	
	// This method will refresh the icon on the sidebar toggle button
	// to match the current state of the RHS sidebar. This is needed
	// when switching between lists and the icon may need initializing.
	function refreshToggleButtonIcon() {
		var $iconSpan = $('#viewListOptionsButton > span')
		
		if(theMainWindowLayout.rhsSidebarIsOpen()) {
			$iconSpan.removeClass("fa-toggle-left")
			$iconSpan.addClass("fa-toggle-right")
		} else {
			$iconSpan.removeClass("fa-toggle-right")
			$iconSpan.addClass("fa-toggle-left")
		}
	}
	
	
	
	
	this.showFooterLayout = showFooterLayout
	this.hideFooterLayout = hideFooterLayout
	this.enableRefreshButton = enableRefreshButton
	this.setCenterContentHeader = setCenterContentHeader
	this.refreshToggleButtonIcon = refreshToggleButtonIcon
	
}